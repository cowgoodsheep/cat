package controller

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"ms/dao"
	"ms/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func generateUniqueFileName(originalName string) string {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomInt := randomGenerator.Intn(1000)
	hash := md5.New()
	hash.Write([]byte(originalName + strconv.Itoa(randomInt)))
	md5Hash := hex.EncodeToString(hash.Sum(nil))

	// 使用原始文件名的后缀扩展名
	parts := strings.Split(originalName, ".")
	if len(parts) > 1 {
		return md5Hash + "." + parts[len(parts)-1]
	}

	return md5Hash
}

func Home(c *gin.Context) {
	var home models.HomeInfo
	var adapt models.AdaptInfo
	var cat models.CatInfo
	dao.DB.Where("status = ?", "已同意").First(&adapt)
	dao.DB.Where("id = ?", adapt.CatId).First(&cat)
	home.Cat_id = cat.ID
	home.Cat_name = cat.Name
	home.Cat_avatar = cat.Avatar
	dao.DB.Model(&models.CatInfo{}).Count(&home.Tot_amount)
	dao.DB.Model(&models.CatInfo{}).Where("gender = ?", "公").Count(&home.Cat_male)
	dao.DB.Model(&models.CatInfo{}).Where("gender = ?", "母").Count(&home.Cat_female)
	dao.DB.Model(&models.CatInfo{}).Where("adapt = ?", "1").Count(&home.Cat_adapt)
	dao.DB.Model(&models.CatInfo{}).Where("sterilization = ?", "1").Count(&home.Cat_sterilization)

	var firstHome models.HomeInfo
	dao.DB.Model(&models.HomeInfo{}).First(&firstHome)
	if firstHome.ID != 0 {
		var fl bool = false
		if firstHome.Cat_id != home.Cat_id {
			fl = true
		}
		if firstHome.Tot_amount != home.Tot_amount {
			fl = true
		}
		if firstHome.Cat_male != home.Cat_male {
			fl = true
		}
		if firstHome.Cat_female != home.Cat_female {
			fl = true
		}
		if firstHome.Cat_adapt != home.Cat_adapt {
			fl = true
		}
		if firstHome.Cat_adapt != home.Cat_adapt {
			fl = true
		}

		if fl {
			currentTime := time.Now()
			currentDateTime := currentTime.Format("2006-01-02 15:04:05")
			home.Date = currentDateTime
			dao.DB.Create(&home)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": home,
	})
}
func Login(c *gin.Context) {
	var user models.UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Name == "admin" && user.Password == "admin" {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"name": "管理员",
				"role": "admin",
				"id":   0,
			},
		})
		return
	}
	var user_find models.UserInfo
	dao.DB.Where("name = ?", user.Name).First(&user_find)
	if user_find.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "用户不存在",
			},
		})
		return
	}

	if user_find.Password != user.Password {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "密码错误",
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"name":   user_find.Name,
				"role":   "user",
				"id":     user_find.ID,
				"avatar": user_find.Avatar,
			},
		})
	}
}

func Register(c *gin.Context) {
	var user models.UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user_find models.UserInfo
	dao.DB.Where("Name = ?", user.Name).First(&user_find)
	if user.Name == "admin" || user_find.ID != 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "已经存在同名用户",
			},
		})
		return
	}
	if len(user.Password) < 6 || len(user.Password) > 20 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "密码长度应为6~20位",
			},
		})
		return
	}
	if len(user.Phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "手机号码长度不等于11位",
			},
		})
		return
	}
	for _, v := range user.Phone {
		if v < '0' || v > '9' {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"data": gin.H{
					"msg": "手机号码不能有非数字字符",
				},
			})
			return
		}
	}
	dao.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"msg": "注册成功！",
		},
	})
}
