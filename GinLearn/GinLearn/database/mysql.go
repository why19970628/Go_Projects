package database

import (
 "database/sql"
 "fmt"
 _ "github.com/go-sql-driver/mysql"
 "log"
)

var SqlDB *sql.DB

func init() {
 var err error
 SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
 if err != nil {
  log.Fatal(err.Error())
 }
 err = SqlDB.Ping()
 if err != nil {
  log.Fatal(err.Error())
 }
 rows, _ := SqlDB.Query("select id,first_name from person");
 //rows 查询 表里面所有的数据 结果应该是一个数组 方式db.Query
 fmt.Println(rows)
 id := 0;
 name := "";
 for rows.Next() {
  rows.Scan(&id, &name); //按顺序取值
  fmt.Println(id, name);
 }
}

