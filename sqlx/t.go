package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var dbname = "mysql"
var datasource = "gram:yangshu##8867@tcp(112.74.205.92:3306)/test"
var dscqdq = "cqdq:cqdq12345@tcp(106.54.87.204:3306)/iot_admin"
var db *sqlx.DB

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

type Time struct {
	Ts string
	D  string
	Dt string
	Y  string
	T  string
}

func main() {
	//sqlx使用测试，用的是以前的阿里云服务器，可惜现在已经过期了
	//t1()
	//cqdq的云服务器
	//t2()
}

func t1() {
	//就是sql的Open+Ping
	db = sqlx.MustConnect(dbname, datasource)
	//tx:=db.MustBegin()
	//tx.MustExec("insert into person(first_name,last_name,email) values(?,?,?)","Jason","Moiron","jmoiron@jmoiron.url")
	//tx.MustExec("insert into person(first_name,last_name,email) values(?,?,?)","John", "Doe", "johndoeDNE@gmail.url")
	//tx.MustExec("insert into place (country, city, telcode) values (?,?,?)","United States", "New York", "1")
	//tx.MustExec("insert into place (country, telcode) values (?,?)","Hong Kong", "852")
	//tx.MustExec("insert into place(country, telcode) values(?,?)","Singapore", "65")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	//tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)",
	//	&Person{"Jane", "Citizen", "jane.citzen@example.com"})
	//tx.Commit()

	//db.MustExec("insert into time(ts,d,dt,y,t) values(?,?,?,?,?)",time.Now(),time.Now(),time.Now(),2020,time.Now())
	//插入结果
	//id|ts                 |d         |dt                 |y   |t       |
	//--|-------------------|----------|-------------------|----|--------|
	//1|2020-07-05 09:38:26|2020-07-05|2020-07-05 09:38:26|2020|09:38:26|
	//秒数time.Now().Unix()就不行，会报错
	//t := time.Unix(1593942253,0)//将秒数转变为Time实例
	//db.MustExec("insert into time(ts,d,dt,y,t) values(?,?,?,?,?)",t,t,t,2020,t)
	//插入成功
	timeMap := make(map[string]interface{})
	times, _ := db.Queryx("select ts from time where id=1")
	if times != nil {
		for times.Next() {
			_ = times.MapScan(timeMap)
		}
	}
	ts, _ := time.ParseInLocation("2006-01-02 15:04:05", string(timeMap["ts"].([]uint8)), time.Local)
	fmt.Println("timestamp:", ts.Unix()) //timestamp: 1593913453，除了ts，其他的域都不能转成time.Time实例
	var t Time
	db.Get(&t, "select ts,d,dt,y,t from `time` where id=1")
	fmt.Println(t) //{2020-07-05 09:38:26 2020-07-05 2020-07-05 09:38:26 2020 09:38:26}

	var people []Person
	//Select用于抓多行数据
	db.Select(&people, "select first_name,last_name from person order by first_name asc")
	jason, john := people[0], people[1]
	fmt.Printf("%#v\n%#v", jason, john)
	//main.Person{FirstName:"Bin", LastName:"Smuth", Email:""}
	//main.Person{FirstName:"Jane", LastName:"Citizen", Email:""}
	count := make(map[string]interface{})
	rows, err := db.Queryx("select count(first_name) as num from person")
	if rows != nil {
		for rows.Next() {
			err = rows.MapScan(count)
		}
	}
	fmt.Println("count: ", string(count["num"].([]uint8)))

	jason1 := Person{}
	//Get用于抓单行数据
	err = db.Get(&jason1, "select * from person where first_name=?", "Jason")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", jason1)
	//main.Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.url"}

	var places []Place
	err = db.Select(&places, "select * from place order by telcode asc")
	if err != nil {
		fmt.Println(err)
	}
	usa, singsing, honkers := places[0], places[1], places[2]
	fmt.Printf("%#v\n%#v\n%#v\n", usa, singsing, honkers)
	//main.Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	//main.Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}
	//main.Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}

	place := Place{}
	rows, err = db.Queryx("select * from place")
	if rows != nil {
		for rows.Next() {
			err := rows.StructScan(&place)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%#v\n", place)
		}
	}

	//main.Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	//main.Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}
	//main.Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}
	_ = db.QueryRowx("select * from place where telcode=?", 1).StructScan(&place)
	fmt.Println(place)
	//{United States {New York true} 1}

	// Named queries, using `:name` as the bindvar.  Automatic bindvar support
	// which takes into account the dbtype based on the driverName on sqlx.Open/Connect
	//_, err=db.NamedExec(`insert into person(first_name,last_name,email) values(:first,:last,:email)`, map[string]interface{}{
	//	"first":"Bin",
	//	"last":"Smuth",
	//	"email":"bensmith@allblacks.nz",
	//})
	//
	//rows,err=db.NamedQuery(`select * from person where first_name=:fn`,map[string]interface{}{"fn":"Bin"})
	//
	//rows,err=db.NamedQuery(`select * from person where first_name=:first_name`,jason)
}

type timebag struct {
	Id   int    `json:"id" db:"id"`
	Time string `json:"time" db:"time"`
}

type timebag1 struct {
	Id   int   `json:"id" db:"id"`
	Time int64 `json:"time" db:"time"`
}

func t2() {
	db = sqlx.MustConnect(dbname, dscqdq)
	//db.MustExec("insert into test.test(time) values(from_unixtime(?))",1593942253)
	//插入timestamp成功，这说明毫秒数是可以转换成timestamp的，但是需要加上mysql的内置函数from_unixtime

	//报错：mysql的timestamp不能直接映射成time.Time
	//timestamp只能映射成string，然后由string转换成time.Time
	var tb timebag
	err := db.Get(&tb, "select * from test.test where id=?", 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tb) //{1 2020-07-05 17:44:13}，这里可以看出timestamp映射成string的输出格式
	timeTemplate := "2006-01-02 15:04:05"
	//为什么解析的只能是string？我猜应该是方便格式多元化的原因吧。
	stamp, _ := time.ParseInLocation(timeTemplate, tb.Time, time.Local)
	fmt.Println(stamp) //2020-07-05 17:44:13 +0800 CST，成功

	//将timestamp映射成int64
	var tb1 timebag1
	err = db.Get(&tb1, "select id,unix_timestamp(`time`) as `time` from test.test where id=?", 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tb1) //{1 1593942253}
	db = sqlx.MustConnect(dbname, dscqdq)
	resMap := map[string]interface{}{}
	rows, err := db.Queryx("select secret_key as ProductSecret,product_id as ProductId,"+
		"device_name as DeviceName,device_id as DeviceId from device_detail where device_name=?", "测试119")
	if rows != nil {
		for rows.Next() {
			err = rows.MapScan(resMap)
		}
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(resMap["ProductId"].([]uint8)), string(resMap["DeviceId"].([]uint8)))
	//10079997 cb30899f92dc410aa826dd2881351130
	//注意，用mapscan，map里的所有域都会变成string
}
