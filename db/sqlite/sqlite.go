package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 数据库文件路径
	dbPath := "./mydatabase.db"

	// 打开或创建数据库
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// 验证数据库连接
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	// 创建表的SQL语句
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	);
	`

	// 执行SQL语句创建表
	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("Database and table created successfully.")

	// insert into multiple tables
	insertSQL := "INSERT INTO users (name, email) VALUES (?,?)"
	_, err = db.Exec(insertSQL, "Bob", "bob@localhost")
	if err != nil {
		fmt.Println("Error inserting into database:", err)
		return
	}
	_, err = db.Exec(insertSQL, "Charlie", "charlie@localhost")
	if err != nil {
		fmt.Println("Error inserting into database:", err)
		return
	}
	fmt.Println("Data inserted successfully.")

	// select from database
	selectSQL := "SELECT id, name, email FROM users"
	rows, err := db.Query(selectSQL)
	if err != nil {
		fmt.Println("Error selecting from database:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return
	}
	//close connection
	db.Close()
	fmt.Println("Connection closed.")
	fmt.Println("Done.")

}
