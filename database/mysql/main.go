package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/**
	* type Driver interface {
	*	Open(name string) (Conn, error)
	}*/
	dsn := "root:fit4hiii@tcp(localhost:3308)/test?charset=utf8mb4"
	db, err := sql.Open("mysql", dsn)
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("insert into userinfo set username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("ric", "研发部", "2020-04-22")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Printf("insert id is %d\n", id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("rambochase", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)

	//res, err = stmt.Exec(id)
	//checkErr(err)

	//affect, err = res.RowsAffected()
	//checkErr(err)

	//fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
