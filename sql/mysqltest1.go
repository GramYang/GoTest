package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	test1()
	//测试了stmt.exec的变态用法，参数只能一个一个的传
	test2()
}

func test1() {
	db, err := sql.Open("mysql", "gram:yangshu88@tcp(127.0.0.1:3306)/gram_landlord")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("select * from player where player.name = ?", "yangping")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	var id, name, password, avatar, win, lose, money string
	for rows.Next() {
		err = rows.Scan(&id, &name, &password, &avatar, &win, &lose, &money)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(id, name, password, avatar, win, lose, money)
}

func test2() {
	db, err := sql.Open("mysql", "gram:yangshu88@tcp(127.0.0.1:3306)/gram_landlord")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	res := pstx1(db, "update player as p set p.win = p.win + 1, p.money = p.money + ? where p.name = ?",
		1000, "yangping")
	num, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	if num != 0 {
		fmt.Println("修改成功，行数：", num)
	} else {
		fmt.Println("修改失败")
	}
}

func pstx1(db *sql.DB, sql string, money int, name string) sql.Result {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(sql)
	if err != nil {
		fmt.Println(err)
	}
	res, err := stmt.Exec(money, name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
	}
	_ = stmt.Close()
	return res
}
