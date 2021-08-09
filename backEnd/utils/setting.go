

package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var(
	AppMode string
	HttpPort string


	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	loadServer(file)
	LoadData(file)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")

}

func LoadData(file *ini.File) {
	Db = file.Section("server").Key("Db").MustString("mysql")
	DbHost = file.Section("server").Key("DbHost").MustString("localhost")
	DbPort = file.Section("server").Key("DbPort").MustString("3306")
	DbUser = file.Section("server").Key("DbUser").MustString("root")
	DbPassWord = file.Section("server").Key("DbPassWord").MustString("123456")
	DbName = file.Section("server").Key("DbName").MustString("library1")
}