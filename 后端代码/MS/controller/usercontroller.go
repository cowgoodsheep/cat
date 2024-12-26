package controller

import (
	"ms/dao"
	"ms/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfileInfo struct {
	ID       uint   `json:"user_id"`
	Name     string `json:"user_name"`
	Password string `json:"user_password"`
	Gender   string `json:"user_gender"`
	Phone    string `json:"user_phone"`
	Age      uint   `json:"user_age"`
	Avatar   string `json:"user_avatar"`
	Cat      string `json:"user_cat"`
}

func GetPersonalProfile(c *gin.Context) {
	id := c.Query("id")
	var user models.UserInfo
	if result := dao.DB.Where("id = ?", id).First(&user); result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	var adapts []models.AdaptInfo
	dao.DB.Where("owner_id = ?", id).Where("status = ?", "已同意").First(&adapts)
	var profile ProfileInfo
	var cats string
	for _, v := range adapts {
		cats += " " + strconv.Itoa(int(v.CatId))
	}
	profile.Name = user.Name
	profile.Password = user.Password
	profile.Gender = user.Gender
	profile.Phone = user.Phone
	profile.Age = user.Age
	profile.Avatar = user.Avatar
	profile.Cat = cats
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": profile,
	})
}
func UpdatePersonalProfile(c *gin.Context) {
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
	if userOrigin.Avatar != userNew.Avatar && userNew.Avatar != "" {
		userOrigin.Avatar = userNew.Avatar
	}

	var user_find models.UserInfo
	dao.DB.Where("Name = ?", userOrigin.Name).First(&user_find)
	if userOrigin.Name == "admin" || user_find.ID != userOrigin.ID {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "已经存在同名用户",
			},
		})
		return
	}
	if len(userOrigin.Password) < 6 || len(userOrigin.Password) > 20 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "密码长度应为6~20位",
			},
		})
		return
	}
	if len(userOrigin.Phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "手机号码长度不等于11位",
			},
		})
		return
	}
	for _, v := range userOrigin.Phone {
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
	dao.DB.Save(&userOrigin)
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": gin.H{
			"msg": "修改成功！",
		},
	})
}
func AdaptCatForm(c *gin.Context) {
	id := c.Query("id")
	var adapts []models.AdaptInfo
	if result := dao.DB.Where("owner_id = ?", id).Limit(3).Find(&adapts); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	for i := 0; len(adapts) < 3; i++ {
		adapts = append(adapts, models.AdaptInfo{
			ID:      0,
			OwnerId: 0,
			CatId:   0,
			Date:    "",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": adapts,
	})
}
func AdaptCatRequest(c *gin.Context) {
	var adapt models.AdaptInfo
	if err := c.ShouldBindJSON(&adapt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// c.JSON(http.StatusOK, adapt)
	var cat models.CatInfo
	result := dao.DB.Where("id = ?", adapt.CatId).First(&cat)
	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "没有这样的猫猫！",
			},
		})
	} else if cat.Adapt {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": gin.H{
				"msg": "猫猫已经被领养了！",
			},
		})
	} else {
		dao.DB.Create(&adapt)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"msg": "成功提交申请！",
			},
		})
	}
}
