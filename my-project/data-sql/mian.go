package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./myproject.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//创建表
	createtable(db)

	// 单条插入数据

	// 多条插入数据
	insertMultipeDataRowByRow(db)

	// 查询数据
	queryData(db)

}

// 创建数据表
func createtable(db *sql.DB) {
	//create users table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
}
func insertMultipeDataRowByRow(db *sql.DB) {
	//create users struct
	type User struct {
		Name     string
		Email    string
		Password string
	}
	//add multipe data by rows
	users := []User{
		{Name: "John Doe", Email: "john@example.com", Password: "123456"},
		{Name: "Jane Doe", Email: "jane@example.com", Password: "123456"},
	}

	// 为每行数据准备一个INSERT语句
	for _, row := range users {
		insertStmt := `INSERT INTO users (name, email, password) VALUES (?, ?, ?);`
		_, err := db.Exec(insertStmt, row.Name, row.Email, row.Password)
		if err != nil {
			log.Fatalf("Failed to insert data: %v", err)
		}
	}
	fmt.Println("Multiple data inserted successfully")
}

// add multipe data by column
//
//		_, err = db.Exec(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`, user.Name, user.Email, user.Password)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//
// 查询数据
func queryData(db *sql.DB) {
	//query data
	sqlStatement := `SELECT * FROM users`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var email string
		var password string
		err = rows.Scan(&id, &name, &email, &password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id = %d, name = %s, email = %s, password = %s\n", id, name, email, password)
	}
}
