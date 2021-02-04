package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

type Config struct {
	AdminHost     string `json:"admin_host"`
	AdminPort     string `json:"admin_port"`
	RunMode       string `json:"run_mode"`
	MysqlAddr     string `json:"mysql_addr"`
	MysqlPort     string `json:"mysql_port"`
	MysqlUserName string `json:"mysql_username"`
	MysqlPassword string `json:"mysql_password"`
	MysqlDatabase string `json:"mysql_database"`
	LocalDebug    bool   `json:"local_debug"`
}

//该文件专门用来测试命令行，因为gomod开启后，你就必须在go.mod文件处执行go run
func main() {
	var p string
	c := &Config{}
	flag.StringVar(&p, "c", "", "配置文件路径")
	flag.Parse()
	if p != "" {
		file, err := ioutil.ReadFile(p)
		if err != nil {
			fmt.Println("读取配置文件", err)
		}
		fmt.Println(string(file))
		if err = json.Unmarshal(file, c); err != nil {
			fmt.Println("解析配置json", err)
		}
	}
	fmt.Printf("%+v\n", c)
}
