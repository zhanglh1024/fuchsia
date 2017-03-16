package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Desire")
	if err != nil {
		log.Println(err)
	}
	err=db.Ping()
	if err!=nil{
		log.Println(err)
	}
	fmt.Println(db)
	var name []string
	rows,err:=db.Query("select name from student ")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next(){
		err:=rows.Scan(&name)
		if err!=nil{
			log.Println(err)
		}
	}
	err=rows.Err()
	if
	fmt.Println(name)
	stmt,err:=db.Prepare("insert into student(id,name,age)VALUES (?,?,?)")
	if err!=nil{
		log.Println(err)
	}
	re,err:=stmt.Exec(8,"nicework",17)
	if err!=nil{
		log.Panicln(err)
	}
	id ,err:=re.LastInsertId()
	affect,err:=re.RowsAffected()
	fmt.Println(id,affect)
	defer db.Close()
}
