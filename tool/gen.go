package main

import (
	"bytes"
	"festival/app/common/utils/file"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {

	mod := GenMod{}
	// 模块名称
	mod.Name = "兑换礼品表"
	// 简称
	mod.Abbr = "userGoods"
	// 文件名
	mod.FileName = "mod_user_goods"
	// 结构体名
	mod.ModName = "ModUserGoods"
	// 表名
	mod.TableName = "mod_user_goods"
	// 字段
	cols := make([]GenModCol, 0)
	// Id、CreatedAt、UpdatedAt、DeletedAt 不用重复创建

	col1 := GenModCol{Name: "userId", GoName: "UserId", Type: "uint", DbType: "int(11)", Mark: "用户ID"}
	col2 := GenModCol{Name: "goodsId", GoName: "GoodsId", Type: "uint", DbType: "int(11)", Mark: "礼品ID"}
	col3 := GenModCol{Name: "status", GoName: "Status", Type: "int", DbType: "tinyint(1)", Mark: "状态：0审核中1已发货-1驳回"}
	col4 := GenModCol{Name: "name", GoName: "Name", Type: "string", DbType: "varchar(280)", Mark: "用户名"}
	col5 := GenModCol{Name: "phone", GoName: "Phone", Type: "string", DbType: "varchar(11)", Mark: "电话号"}
	col6 := GenModCol{Name: "address", GoName: "Address", Type: "string", DbType: "varchar(300)", Mark: "联系地址"}
	cols = append(cols, col1, col2, col3, col4, col5, col6)

	// beginGen(mod, cols)
}

func beginGen(mod GenMod, cols []GenModCol) {
	//获取当前运行时目录
	curDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取路径失败")
		return
	}
	params := gin.H{"mod": mod, "cols": cols}
	//edit模板
	if tmp, err := loadTemplate("edit.txt", params); err == nil {
		fileName := strings.Join([]string{curDir, "/resource/views/backend/module/", mod.Abbr, "/edit.html"}, "")
		if !file.Exists(fileName) {
			f, err := file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}
	//list模板
	if tmp, err := loadTemplate("list.txt", params); err == nil {
		fileName := strings.Join([]string{curDir, "/resource/views/backend/module/", mod.Abbr, "/list.html"}, "")
		if !file.Exists(fileName) {
			f, err := file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}

	//Modle模板
	if tmp, err := loadTemplate("mod.txt", params); err == nil {
		fileName := strings.Join([]string{curDir, "/app/model/module/", mod.FileName, ".go"}, "")
		if !file.Exists(fileName) {
			f, err := file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}

	// 数据库文件模板
	if tmp, err := loadTemplate("migrate.txt", params); err == nil {
		fileName := strings.Join([]string{curDir, "/app/migrates/", mod.FileName, ".go"}, "")
		if !file.Exists(fileName) {
			f, err := file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}

	//Controller模板
	if tmp, err := loadTemplate("controller.txt", params); err == nil {
		fileName := strings.Join([]string{curDir, "/app/controller/module/", mod.Abbr, "_controller.go"}, "")
		if !file.Exists(fileName) {
			f, err := file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}

	//Service模板
	if tmp, err := loadTemplate("service.txt", params); err == nil {
		fileName := strings.Join([]string{curDir, "/app/service/module/", mod.Abbr, "_service.go"}, "")
		if !file.Exists(fileName) {
			f, err := file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}

	//Router模板
	if tmp, err := loadTemplate("router.txt", params); err == nil {
		fileName := strings.Join([]string{curDir, "/app/routers/", mod.Abbr, "_router.go"}, "")
		if !file.Exists(fileName) {
			f, err := file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}
	//Sql模板
	if tmp, err := loadTemplate("sql.txt", params); err == nil {
		fileName := strings.Join([]string{curDir, "/tool/", mod.Abbr, "_menus.sql"}, "")
		if !file.Exists(fileName) {
			f, err := file.Create(fileName)
			if err == nil {
				f.WriteString(tmp)
			}
			f.Close()
		}
	}
	log.Printf(`
	代码生成完成：
	/views/backend/module/%s
	/app/model/module/%s
	/app/migrates/%s
	/app/controller/module/%s
	/app/services/%s
	/app/routers/%s
	/tool/sql
	PowerBy：7be.cn
	`, mod.Abbr, mod.Abbr, mod.Abbr, mod.Abbr, mod.Abbr, mod.Abbr)
}

//读取模板
func loadTemplate(templateName string, data interface{}) (string, error) {
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadFile(cur + "/tool/tpl/" + templateName)
	if err != nil {
		return "", err
	}
	templateStr := string(b)

	tmpl, err := template.New(templateName).Parse(templateStr) //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		return "", nil
	}
	buffer := bytes.NewBufferString("")
	err = tmpl.Execute(buffer, data) //将string与模板合成，变量name的内容会替换掉{{.}}
	if err != nil {
		return "", nil
	}
	return buffer.String(), nil
}

type GenMod struct {
	Name      string // 模块名称
	Abbr      string // 简称
	ModName   string // 模块名
	FileName  string // 文件名
	TableName string // 表名
}

type GenModCol struct {
	Name   string //字段名
	GoName string //字段名
	Type   string // 类型
	DbType string // 数据库类型
	Mark   string // 备注
}
