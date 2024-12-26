package routers

import (
	"ms/controller"
	"ms/setting"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Static("/static", "./static")
	r.Static("/uploads/catavatars", "./uploads/catAvatars")
	r.Static("/uploads/useravatars", "./uploads/userAvatars")

	r.GET("/home", controller.Home)

	// 登录信息处理 √
	r.POST("/login", controller.Login)
	// 注册信息处理 √
	r.POST("/register", controller.Register)

	/*流浪猫管理*/
	// 上传流浪猫图片 √
	r.POST("/uploadcat", controller.UploadCatFile)
	// 查看流浪猫 √
	r.GET("/catm", controller.GetCatInfo)
	// 创建流浪猫 √
	r.POST("/catm", controller.CreateACat)
	// 修改流浪猫 √
	r.PUT("/catm", controller.UpdateACat)
	// 删除流浪猫 √
	r.DELETE("/catm", controller.DeleteACat)

	/*用户管理*/
	// 上传用户图片 √
	r.POST("/uploaduser", controller.UploadUserFile)
	// 查看用户档案 √
	r.GET("/userm", controller.GetPersonInfo)
	// 创建用户档案 √
	r.POST("/userm", controller.CreateAPerson)
	// 修改用户 √
	r.PUT("/userm", controller.UpdateAPerson)
	// 删除用户 √
	r.DELETE("/userm", controller.DeleteAPerson)

	// 用户进入收养界面
	// 用户查看最新收养的三条消息 √
	r.GET("/adaptu", controller.AdaptCatForm)
	// 提出收养请求 √
	r.POST("/adaptu", controller.AdaptCatRequest)

	//管理员进入收养界面
	// 管理员查看收养信息
	r.GET("/adaptm", controller.AdaptCatReceive)
	// 管理员对收养做出响应
	r.PUT("/adaptm", controller.AdaptCatRespond)

	// 用户进入个人档案界面，查看自己的个人档案 √
	r.GET("/personalfile", controller.GetPersonalProfile)
	// 修改个人的档案 √
	r.PUT("/personalfile", controller.UpdatePersonalProfile)

	return r

}
