要修改代码以便在数据库已经存在的情况下查询记录，我们可以在连接数据库后检查是否存在表。如果表存在，则直接查询记录；如果表不存在，则进行迁移。以下是更新后的代码示例：

### 更新后的代码示例

```go
package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 定义一个模型
type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	// 连接到 SQLite 数据库
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 检查表是否存在
	if !db.Migrator().HasTable(&User{}) {
		// 如果表不存在，则迁移模式
		db.AutoMigrate(&User{})
	}

	// 创建记录
	user := User{Name: "Alice", Age: 30}
	db.Create(&user)

	// 查询记录
	var users []User
	db.Find(&users)
	log.Println(users)
}
```

### 代码说明：
1. **检查表是否存在**：使用 `db.Migrator().HasTable(&User{})` 检查 `User` 表是否存在。
2. **条件迁移**：只有在表不存在的情况下才调用 `db.AutoMigrate(&User{})` 进行迁移。
3. **创建和查询记录**：无论表是否存在，都会创建记录并查询所有用户。

### 运行代码
确保在运行此程序之前，`test.db` 数据库文件已经存在于当前目录中。运行程序后，它将创建用户记录（如果表不存在），并查询所有用户记录并打印到日志中。
