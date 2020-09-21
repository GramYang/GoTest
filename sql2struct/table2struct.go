package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

var typeForMysqlToGo = map[string]string{
	"int":                "int32",
	"integer":            "int32",
	"tinyint":            "int8",
	"smallint":           "int16",
	"mediumint":          "int32", //go没有int24
	"bigint":             "int64",
	"int unsigned":       "int32", //unsigned类型和非unsigned类型大小一样
	"integer unsigned":   "int32",
	"tinyint unsigned":   "int8",
	"smallint unsigned":  "int16",
	"mediumint unsigned": "int32",
	"bigint unsigned":    "int64",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "string",
	"datetime":           "string",
	"timestamp":          "string",
	"time":               "string",
	"float":              "float32",
	"double":             "float64",
	"decimal":            "float64",
}

type Table2Struct struct {
	DataSource    string
	SavePath      string
	Table         string //扫描某个具体的表
	db            *sql.DB
	Prefix        string //这个前缀是表列名和表名共有的前缀
	PackageName   string //生成go文件的包名
	Err           error
	EnableJsonTag bool //启用json tag
	EnableDbTag   bool //启用db tag
}

type column struct {
	columnName    string
	dataType      string
	nullable      string
	tableName     string
	columnComment string
	tag           string
}

func (t *Table2Struct) Run() error {
	t.connectMysql()
	if t.Err != nil {
		return t.Err
	}
	tableColumns, err := t.getColumns()
	if err != nil {
		return err
	}
	var packageName string
	if t.PackageName == "" {
		packageName = "package model\n\n"
	} else {
		packageName = fmt.Sprintf("package %s\n\n", t.PackageName)
	}
	var structContent strings.Builder
	for tableName, item := range tableColumns {
		tableName = t.camelCase(tableName)
		structContent.WriteString("type " + tableName + " struct {\n")
		for _, v := range item {
			var columnComment string
			if v.columnComment != "" {
				columnComment = fmt.Sprintf(" // %s", v.columnComment)
			}
			structContent.WriteString(fmt.Sprintf("%s%s %s %s%s\n", tab(1),
				v.columnName, v.dataType, v.tag, columnComment))
		}
		structContent.WriteString("}\n\n")
	}
	savePath := t.SavePath
	if savePath == "" {
		savePath = "model.go"
	}
	f, err := os.Create(savePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	_, _ = f.WriteString(packageName + structContent.String())
	return nil
}

func (t *Table2Struct) connectMysql() {
	if t.db == nil {
		if t.DataSource == "" {
			t.Err = errors.New("dataSource缺失")
			return
		}
		t.db, t.Err = sql.Open("mysql", t.DataSource)
	}
}

// tableColumns是map，key是表名，[]column是列切片，且column中的信息已经适配了go
func (t *Table2Struct) getColumns() (tableColumns map[string][]column, err error) {
	tableColumns = make(map[string][]column)
	sqlStr := `SELECT COLUMN_NAME,DATA_TYPE,IS_NULLABLE,TABLE_NAME,COLUMN_COMMENT
		FROM information_schema.COLUMNS 
		WHERE table_schema = DATABASE()`
	if t.Table != "" {
		sqlStr += fmt.Sprintf(" AND TABLE_NAME = '%s'", t.Prefix+t.Table)
	}
	sqlStr += " order by TABLE_NAME asc, ORDINAL_POSITION asc"
	rows, err1 := t.db.Query(sqlStr)
	if err1 != nil {
		fmt.Println("Error reading table information: ", err1.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		col := column{}
		err = rows.Scan(&col.columnName, &col.dataType, &col.nullable, &col.tableName, &col.columnComment)
		if err != nil {
			fmt.Println(err)
			return
		}
		tmp := col.columnName                         //保存初始的域名
		col.columnName = t.camelCase(col.columnName)  //切换成go的域格式，驼峰且首字母大写
		col.dataType = typeForMysqlToGo[col.dataType] //切换成go的数据类型
		if t.EnableDbTag && t.EnableJsonTag {
			col.tag = fmt.Sprintf("`%s:\"%s\" %s:\"%s\"`", "db", tmp, "json",
				strings.ToLower(col.columnName[0:1])+col.columnName[1:])
		}
		if _, ok := tableColumns[col.tableName]; !ok {
			tableColumns[col.tableName] = []column{}
		}
		tableColumns[col.tableName] = append(tableColumns[col.tableName], col)
	}
	return
}

//处理域名。先去除前缀，再转换为大驼峰风格
func (t *Table2Struct) camelCase(name string) string {
	//去除表列名的前缀
	if t.Prefix != "" {
		name = strings.Replace(name, t.Prefix, "", 1)
	}
	var text string
	for _, p := range strings.Split(name, "_") {
		if len(p) == 1 {
			text += strings.ToUpper(p)
		} else if len(p) > 1 {
			text += strings.ToUpper(p[0:1]) + strings.ToLower(p[1:])
		}
	}
	//首字母大写
	return strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
}

func tab(depth int) string {
	return strings.Repeat("\t", depth)
}

func main() {
	t2t := &Table2Struct{
		DataSource:    "gram:yangshu##8867@tcp(112.74.205.92:3306)/blog",
		EnableDbTag:   true,
		EnableJsonTag: true,
	}
	err := t2t.Run()
	if err != nil {
		fmt.Println(err)
	}
}
