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
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func UploadCatFile(c *gin.Context) {
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
	if err := c.SaveUploadedFile(file, "./uploads/catAvatars/"+fileName); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
			"msg":   "保存出错",
		})
		return
	}
	fileURL := "http://localhost:" + strconv.Itoa(setting.Conf.Port) + "/uploads/catavatars/" + fileName
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"url": fileURL,
		},
	})
}

func GetCatInfo(c *gin.Context) {
	Method := c.Query("method")
	Content := c.Query("content")
	var cats []models.CatInfo
	var result *gorm.DB
	if Content == "" {
		result = dao.DB.Find(&cats)
	} else if Method == "" {
		result = dao.DB.Where("name like ?", "%"+Content+"%").Find(&cats)
	} else if Method != "" {
		result = dao.DB.Where(Method+" like ?", "%"+Content+"%").Find(&cats)
	}

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}
func CreateACat(c *gin.Context) {
	var cat models.CatInfo

	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"msg": "上传成功",
		},
	})
	if err := dao.DB.Create(&cat); err != nil {
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

func UpdateACat(c *gin.Context) {
	var catNew models.CatInfo

	if err := c.ShouldBindJSON(&catNew); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var catOrigin models.CatInfo
	dao.DB.Where("id = ?", catNew.ID).First(&catOrigin)
	if catNew.Name != catOrigin.Name {
		catOrigin.Name = catNew.Name
	}
	if catNew.Type != catOrigin.Type {
		catOrigin.Type = catNew.Type
	}
	if catNew.Gender != catOrigin.Gender {
		catOrigin.Gender = catNew.Gender
	}
	if catNew.Sterilization != catOrigin.Sterilization {
		catOrigin.Sterilization = catNew.Sterilization
	}
	if catNew.Adapt != catOrigin.Adapt {
		catOrigin.Adapt = catNew.Adapt 
	}
	if catNew.Avatar != catOrigin.Avatar && catNew.Avatar != "" {
		catOrigin.Avatar = catNew.Avatar
	}
	dao.DB.Save(&catOrigin)
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": gin.H{
			"msg": "修改成功！",
		},
	})
}
func DeleteACat(c *gin.Context) {
	var cat models.CatInfo
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cat_find models.CatInfo
	dao.DB.Where("id = ?", cat.ID).First(&cat_find)
	if cat_find.ID != 0 {
		if cat_find.Avatar != "" {
			len := len("http://localhost:" + strconv.Itoa(setting.Conf.Port) + "/uploads/catavatars")
			filePath := "uploads/catAvatars" + cat_find.Avatar[len:]
			if err := os.Remove(filePath); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		dao.DB.Delete(cat_find)
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
