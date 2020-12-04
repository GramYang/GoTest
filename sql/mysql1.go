package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	name string
	age  string
)

func main() {
	//连接DB
	db, err := sql.Open("mysql", "gram:yangshu88@tcp(127.0.0.1:3306)/test")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	//验证连接
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	//查询
	multiQuery(db)
	//单行查询
	singleQuery(db)
	//预编译，用于非查询操作
	ps(db)
	//预编译+事务
	pstx(db)

}

func multiQuery(db *sql.DB) {
	//如果你不知道这个column字段的type，就需要使用sql.RawBytes
	useRawBytes(db)
	//如果某个字段可能为空，就要使用sql.NullString来处理了
	useNullString(db)
}

func useRawBytes(db *sql.DB) {
	//查询，这里传入的参数可以是: "1"、01、1
	rows, err := db.Query("select * from student where SId = ?", "1")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	cols, err := rows.Columns()
	vals := make([]sql.RawBytes, len(cols))
	ages := make([]interface{}, len(cols))
	for i := range ages {
		ages[i] = &vals[i]
	}
	for rows.Next() {
		//Scan只接受指针
		err = rows.Scan(ages...)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range ages {
		s := v.(*sql.RawBytes)
		fmt.Printf("useRawBytes: %s\n", *s)
		//输出：
		//赵雷
		//1990-01-01 00:00:00
		//男
	}
}

func useNullString(db *sql.DB) {
	//查询
	rows, err := db.Query("select Sname, Sage from student where Sname = ?", "赵雷")
	defer rows.Close() //根据Close的注释，这里可以有也可以没有，因为Next在返回false后会自动调用Close
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&name, &age)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(name, age) //赵雷 1990-01-01 00:00:00
		//如果某个字段可能为空，就要使用sql.NullString来处理了
		var s1, s2 sql.NullString
		err = rows.Scan(&s1, &s2)
		if err != nil {
			fmt.Println(err)
		}
		if s1.Valid && s2.Valid {
			fmt.Println("useNullString: ", s1.String, s2.String) //赵雷 1990-01-01 00:00:00
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
}

func singleQuery(db *sql.DB) {
	err := db.QueryRow("select Sage from student where Sname = ?", "赵雷").Scan(&age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("singleQuery: ", age) //1990-01-01 00:00:00
}

func ps(db *sql.DB) {
	stmt, err := db.Prepare("select Sage from student where Sname = ?")
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}
	rows, err := stmt.Query("赵雷")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ps: ", rows) //因为我不想修改数据库表，所以用的查询，打印的是Rows
}

func pstx(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("update sc set sc.score = sc.score + ? where sc.SId = ? and sc.CId = ?")
	if err != nil {
		fmt.Println(err)
	}
	//这里传入的涉及到加减运算的参数（第一个）可以是正负整数，正负整数的字符串
	rows, err := stmt.Exec("-1", 6, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("pstx: ", rows) //打印的是driverResult
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
	_ = stmt.Close()
}
