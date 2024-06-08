package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	//create database connection
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create table
	_, err = db.Exec("CREATE TABLE my_table (id int)")
	if err != nil {
		log.Fatal(err)
	}

	//update table
	_, err = db.Exec("UPDATE my_table SET id = 2")
	if err != nil {
		log.Fatal(err)
	}

	//delete table
	_, err = db.Exec("DELETE FROM my_table")
	if err != nil {
		log.Fatal(err)
	}

	//drop table
	_, err = db.Exec("DROP TABLE my_table")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")

	//select table
	rows, err := db.Query("SELECT * FROM my_table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		// 定义变量接收查询结果
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

	// begin transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// execute a series of transactions
	_, err = tx.Exec("INSERT INTO my_table (name) VALUES (?)", "John Doe")
	if err != nil {
		tx.Rollback() // 回滚事务
		log.Fatal(err)
	}

	err = tx.Commit() // 提交事务
	if err != nil {
		log.Fatal(err)
	}

}
