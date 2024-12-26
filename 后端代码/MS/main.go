package main

import (
	"fmt"
	"ms/dao"
	"ms/models"
	"ms/routers"
	"ms/setting"
	"os"
)

const defaultConfFile = "./conf/config.ini"

func main() {
	confFile := defaultConfFile
	if len(os.Args) > 2 {
		fmt.Println("use specified conf file: ", os.Args[1])
		confFile = os.Args[1]
	} else {
		fmt.Println("no configuration file was specified, use ./conf/config.ini")
	}
	// 加载配置文件
	if err := setting.Init(confFile); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	// 创建数据库
	// sql ：CREATE DATABASE catms;
	// 连接数据库

	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer dao.Close()

	// 模型绑定
	dao.DB.AutoMigrate(&models.UserInfo{})
	dao.DB.AutoMigrate(&models.CatInfo{})
	dao.DB.AutoMigrate(&models.AdaptInfo{})
	dao.DB.AutoMigrate(&models.HomeInfo{})
	// dao.DB.Model(&models.AdaptInfo{}).AddForeignKey("owner_id", "user_infos(id)", "CASCADE", "CASCADE")
	// dao.DB.Model(&models.AdaptInfo{}).AddForeignKey("cat_id", "cat_infos(id)", "CASCADE", "CASCADE")
	// dao.DB.Model(&models.CatInfo{}).AddForeignKey("owner", "user_infos(id)", "CASCADE", "CASCADE")

	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}

}
