package controller

import (
	"ms/dao"
	"ms/models"
	"ms/setting"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AdaptCatReceive(c *gin.Context) {
	var adapts []models.AdaptInfo
	result := dao.DB.Find(&adapts)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": adapts,
	})
}

func AdaptCatRespond(c *gin.Context) {
	var adapt models.AdaptInfo
	if err := c.ShouldBindJSON(&adapt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cat models.CatInfo
	var adapt_find models.AdaptInfo
	dao.DB.Where("id = ?", adapt.CatId).First(&cat)
	dao.DB.Where("id = ?", adapt.ID).First(&adapt_find)

	if adapt.Status == "已同意" {
		if cat.Adapt {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"data": gin.H{
					"msg": "这只猫猫已经被领养了！",
				},
			})
			return
		}
		cat.Owner = adapt.OwnerId
		cat.Adapt = true
		adapt_find.Status = adapt.Status
		dao.DB.Save(&cat)
		dao.DB.Save(&adapt_find)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"msg": "已经同意该申请！",
			},
		})
	} else if adapt.Status == "已拒绝" {
		adapt_find.Status = adapt.Status
		dao.DB.Save(&adapt_find)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"msg": "已经拒绝该申请！",
			},
		})
	} else if adapt.Status == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "未获取到修改操作！",
			},
		})
	}
}

func  UploadUserFile(c *gin.Context) {
	// 处理头像文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"msg":   "获取出错",
		})
		return
	}

	fileName := generateUniqueFileName(file.Filename)
	// 保存文件到指定目录
	if err := c.SaveUploadedFile(file, "./uploads/userAvatars/"+fileName); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
			"msg":   "保存出错",
		})
		return
	}
	fileURL := "http://localhost:" + strconv.Itoa(setting.Conf.Port) + "/uploads/useravatars/" + fileName
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"url": fileURL,
		},
	})
}

func GetPersonInfo(c *gin.Context) {
	Method := c.Query("method")
	Content := c.Query("content")
	var users []models.UserInfo
	var result *gorm.DB
	if Content == "" {
		result = dao.DB.Find(&users)
	} else if Method == "" {
		result = dao.DB.Where("name like ?", "%"+Content+"%").Find(&users)
	} else if Method != "" {
		result = dao.DB.Where(Method+" like ?", "%"+Content+"%").Find(&users)
	}

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
func CreateAPerson(c *gin.Context) {
	var user models.UserInfo

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"msg": "上传成功",
		},
	})
	if err := dao.DB.Create(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"msg": "创建成功！",
			},
		})
		return
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "创建失败！",
			},
		})
		return
	}
}

func UpdateAPerson(c *gin.Context) {
	var userNew models.UserInfo

	if err := c.ShouldBindJSON(&userNew); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userOrigin models.UserInfo
	dao.DB.Where("id = ?", userNew.ID).First(&userOrigin)
	if userOrigin.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "找不到该用户",
			},
		})
		return
	}
	if userOrigin.Name != userNew.Name {
		userOrigin.Name = userNew.Name
	}
	if userOrigin.Password != userNew.Password {
		userOrigin.Password = userNew.Password
	}
	if userOrigin.Gender != userNew.Gender {
		userOrigin.Gender = userNew.Gender
	}
	if userOrigin.Phone != userNew.Phone {
		userOrigin.Phone = userNew.Phone
	}
	if userOrigin.Age != userNew.Age {
		userOrigin.Age = userNew.Age
	}
	if userOrigin.Avatar != userNew.Avatar {
		userOrigin.Avatar = userNew.Avatar
	}
	dao.DB.Save(&userOrigin)
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": gin.H{
			"msg": "修改成功！",
		},
	})
}
func DeleteAPerson(c *gin.Context) {
	var user models.UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user_find models.UserInfo
	dao.DB.Where("id = ?", user.ID).First(&user_find)
	if user_find.ID != 0 {
		if user_find.Avatar != "" {
			len := len("http://localhost:" + strconv.Itoa(setting.Conf.Port) + "/uploads/useravatars")
			filePath := "uploads/userAvatars" + user_find.Avatar[len:]
			if err := os.Remove(filePath); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		dao.DB.Delete(user_find)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"msg": "删除成功！",
			},
		})
		return
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "删除失败，未找到该猫猫！",
			},
		})
		return
	}
}
