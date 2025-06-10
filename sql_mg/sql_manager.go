package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseManager 管理多个数据库连接
type DatabaseManager struct {
	Databases map[string]interface{} // 存储不同类型的数据库连接
	Config    *viper.Viper
}

// NewDatabaseManager 创建数据库管理器
func NewDatabaseManager(configPath string) (*DatabaseManager, error) {
	// 初始化Viper
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 读取配置
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 创建管理器
	manager := &DatabaseManager{
		Databases: make(map[string]interface{}),
		Config:    v,
	}

	return manager, nil
}

// ConnectAll 连接所有配置的数据库
func (dm *DatabaseManager) ConnectAll() error {
	// 检查数据库配置是否存在
	if !dm.Config.IsSet("databases") {
		return errors.New("配置中未找到数据库部分")
	}

	// 获取所有数据库配置
	dbConfigs := dm.Config.GetStringMap("databases")
	for dbName := range dbConfigs {
		dbType := dm.Config.GetString(fmt.Sprintf("databases.%s.driver", dbName))

		var err error
		switch dbType {
		case "postgres", "mysql", "sqlite3":
			err = dm.connectSQLDB(dbName)
		case "mongodb":
			err = dm.connectMongoDB(dbName)
		case "redis":
			// Redis连接逻辑（需要额外的Redis客户端库）
			// err = dm.connectRedis(dbName)
			log.Printf("Redis连接暂未实现: %s", dbName)
		default:
			log.Printf("不支持的数据库类型: %s", dbType)
		}

		if err != nil {
			return fmt.Errorf("连接数据库 %s 失败: %w", dbName, err)
		}
	}

	// 连接分片数据库
	if dm.Config.GetBool("sharding.enabled") {
		shardConfigs := dm.Config.GetStringMap("sharding.shards")
		for shardName := range shardConfigs {
			if err := dm.connectShardDB(shardName); err != nil {
				return fmt.Errorf("连接分片数据库 %s 失败: %w", shardName, err)
			}
		}
	}

	return nil
}

// connectSQLDB 连接SQL数据库（PostgreSQL, MySQL, SQLite）
func (dm *DatabaseManager) connectSQLDB(dbName string) error {
	dbType := dm.Config.GetString(fmt.Sprintf("databases.%s.driver", dbName))
	
	var dsn string
	switch dbType {
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			dm.Config.GetString(fmt.Sprintf("databases.%s.host", dbName)),
			dm.Config.GetInt(fmt.Sprintf("databases.%s.port", dbName)),
			dm.Config.GetString(fmt.Sprintf("databases.%s.user", dbName)),
			dm.Config.GetString(fmt.Sprintf("databases.%s.password", dbName)),
			dm.Config.GetString(fmt.Sprintf("databases.%s.dbname", dbName)),
			dm.Config.GetString(fmt.Sprintf("databases.%s.sslmode", dbName)),
		)
	case "mysql":
		params := dm.Config.GetString(fmt.Sprintf("databases.%s.params", dbName))
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?%s",
			dm.Config.GetString(fmt.Sprintf("databases.%s.user", dbName)),
			dm.Config.GetString(fmt.Sprintf("databases.%s.password", dbName)),
			dm.Config.GetString(fmt.Sprintf("databases.%s.host", dbName)),
			dm.Config.GetInt(fmt.Sprintf("databases.%s.port", dbName)),
			dm.Config.GetString(fmt.Sprintf("databases.%s.dbname", dbName)),
			params,
		)
	case "sqlite3":
		dsn = dm.Config.GetString(fmt.Sprintf("databases.%s.path", dbName))
	}

	// 打开数据库连接
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return fmt.Errorf("打开数据库连接失败: %w", err)
	}

	// 配置连接池
	db.SetMaxOpenConns(dm.Config.GetInt(fmt.Sprintf("databases.%s.max_open_conns", dbName)))
	db.SetMaxIdleConns(dm.Config.GetInt(fmt.Sprintf("databases.%s.max_idle_conns", dbName)))

	// 解析并设置连接生命周期
	if dm.Config.IsSet(fmt.Sprintf("databases.%s.conn_max_lifetime", dbName)) {
		lifetime, err := time.ParseDuration(dm.Config.GetString(fmt.Sprintf("databases.%s.conn_max_lifetime", dbName)))
		if err == nil {
			db.SetConnMaxLifetime(lifetime)
		}
	}

	// 解析并设置空闲连接超时
	if dm.Config.IsSet(fmt.Sprintf("databases.%s.conn_max_idle_time", dbName)) {
		idleTime, err := time.ParseDuration(dm.Config.GetString(fmt.Sprintf("databases.%s.conn_max_idle_time", dbName)))
		if err == nil {
			db.SetConnMaxIdleTime(idleTime)
		}
	}

	// 验证连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return fmt.Errorf("数据库连接验证失败: %w", err)
	}

	// 存储连接
	dm.Databases[dbName] = db
	log.Printf("成功连接到数据库: %s (%s)", dbName, dbType)

	return nil
}

// connectMongoDB 连接MongoDB数据库
func (dm *DatabaseManager) connectMongoDB(dbName string) error {
	uri := dm.Config.GetString(fmt.Sprintf("databases.%s.uri", dbName))
	
	// 创建MongoDB客户端选项
	clientOptions := options.Client().ApplyURI(uri)
	
	// 设置连接池选项
	if dm.Config.IsSet(fmt.Sprintf("databases.%s.max_pool_size", dbName)) {
		maxPoolSize := uint64(dm.Config.GetInt(fmt.Sprintf("databases.%s.max_pool_size", dbName)))
		clientOptions.SetMaxPoolSize(maxPoolSize)
	}
	
	if dm.Config.IsSet(fmt.Sprintf("databases.%s.min_pool_size", dbName)) {
		minPoolSize := uint64(dm.Config.GetInt(fmt.Sprintf("databases.%s.min_pool_size", dbName)))
		clientOptions.SetMinPoolSize(minPoolSize)
	}
	
	if dm.Config.IsSet(fmt.Sprintf("databases.%s.max_idle_time_ms", dbName)) {
		maxIdleTime := time.Duration(dm.Config.GetInt(fmt.Sprintf("databases.%s.max_idle_time_ms", dbName))) * time.Millisecond
		clientOptions.SetMaxConnIdleTime(maxIdleTime)
	}

	// 连接到MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("连接MongoDB失败: %w", err)
	}

	// 验证连接
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("MongoDB连接验证失败: %w", err)
	}

	// 存储连接
	dm.Databases[dbName] = client
	log.Printf("成功连接到MongoDB: %s", dbName)

	return nil
}

// connectShardDB 连接分片数据库
func (dm *DatabaseManager) connectShardDB(shardName string) error {
	dbType := dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.driver", shardName))
	
	var dsn string
	switch dbType {
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.host", shardName)),
			dm.Config.GetInt(fmt.Sprintf("sharding.shards.%s.port", shardName)),
			dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.user", shardName)),
			dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.password", shardName)),
			dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.dbname", shardName)),
			dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.sslmode", shardName)),
		)
	case "mysql":
		params := dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.params", shardName))
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?%s",
			dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.user", shardName)),
			dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.password", shardName)),
			dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.host", shardName)),
			dm.Config.GetInt(fmt.Sprintf("sharding.shards.%s.port", shardName)),
			dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.dbname", shardName)),
			params,
		)
	}

	// 打开数据库连接
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return fmt.Errorf("打开分片数据库连接失败: %w", err)
	}

	// 配置连接池
	db.SetMaxOpenConns(dm.Config.GetInt(fmt.Sprintf("sharding.shards.%s.max_open_conns", shardName)))
	db.SetMaxIdleConns(dm.Config.GetInt(fmt.Sprintf("sharding.shards.%s.max_idle_conns", shardName)))

	// 验证连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return fmt.Errorf("分片数据库连接验证失败: %w", err)
	}

	// 存储连接
	dm.Databases["shard_"+shardName] = db
	log.Printf("成功连接到分片数据库: %s (%s), ID范围: %s", 
		shardName, 
		dbType,
		dm.Config.GetString(fmt.Sprintf("sharding.shards.%s.id_range", shardName)),
	)

	return nil
}

// GetSQLDB 获取SQL数据库连接
func (dm *DatabaseManager) GetSQLDB(name string) (*sql.DB, error) {
	db, ok := dm.Databases[name]
	if !ok {
		return nil, fmt.Errorf("数据库 %s 未连接", name)
	}

	sqlDB, ok := db.(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("数据库 %s 不是SQL数据库", name)
	}

	return sqlDB, nil
}

// GetMongoDB 获取MongoDB客户端
func (dm *DatabaseManager) GetMongoDB(name string) (*mongo.Client, error) {
	db, ok := dm.Databases[name]
	if !ok {
		return nil, fmt.Errorf("MongoDB %s 未连接", name)
	}

	mongoClient, ok := db.(*mongo.Client)
	if !ok {
		return nil, fmt.Errorf("数据库 %s 不是MongoDB客户端", name)
	}

	return mongoClient, nil
}

// Close 关闭所有数据库连接
func (dm *DatabaseManager) Close() {
	for name, db := range dm.Databases {
		var err error
		switch v := db.(type) {
		case *sql.DB:
			err = v.Close()
		case *mongo.Client:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			err = v.Disconnect(ctx)
			cancel()
		}

		if err != nil {
			log.Printf("关闭数据库 %s 连接失败: %v", name, err)
		} else {
			log.Printf("成功关闭数据库 %s 连接", name)
		}
	}
}

func main() {
	// 创建数据库管理器
	manager, err := NewDatabaseManager("config.yaml")
	if err != nil {
		log.Fatalf("创建数据库管理器失败: %v", err)
	}

	// 连接所有数据库
	if err := manager.ConnectAll(); err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 确保在程序退出时关闭所有连接
	defer manager.Close()

	// 获取PostgreSQL主数据库
	pgDB, err := manager.GetSQLDB("postgres_master")
	if err != nil {
		log.Printf("获取PostgreSQL主数据库失败: %v", err)
	} else {
		// 使用PostgreSQL数据库
		var version string
		err := pgDB.QueryRow("SELECT version()").Scan(&version)
		if err != nil {
			log.Printf("查询PostgreSQL版本失败: %v", err)
		} else {
			log.Printf("PostgreSQL版本: %s", version)
		}
	}

	// 获取MongoDB
	mongoDB, err := manager.GetMongoDB("mongodb_primary")
	if err != nil {
		log.Printf("获取MongoDB失败: %v", err)
	} else {
		// 使用MongoDB
		dbNames, err := mongoDB.ListDatabaseNames(context.Background(), nil)
		if err != nil {
			log.Printf("列出MongoDB数据库失败: %v", err)
		} else {
			log.Printf("MongoDB数据库: %v", dbNames)
		}
	}

	// 应用主逻辑
	log.Println("数据库连接就绪，应用程序开始运行...")
	
	// 这里添加您的应用程序逻辑
	
	log.Println("应用程序结束")
}