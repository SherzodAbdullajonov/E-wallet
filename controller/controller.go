package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SherzodAbdullajonov/e-wallet/database"
	"github.com/SherzodAbdullajonov/e-wallet/models"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var balance []models.Accounts
	defer cancel()
	if err := database.Db.Find(&balance).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, balance)
	}
}
func TopUp() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var topUp models.TopUps
		err := ctx.BindJSON(&topUp)
		defer cancel()
		if err != nil{
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		topUp.Created_at = time.Now().Local()
		database.Db.Create(&topUp)
		ctx.JSON(http.StatusOK, topUp)


	}
}

func CheckAcount() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		id := ctx.Params.ByName("id")
		log.Println("ID", id)
		var user models.Users
		defer cancel()
		if err := database.Db.Where("id=?", id).First(&user).Error; err!=nil{
			ctx.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			ctx.JSON(http.StatusOK, user)
		}

		
	}
}

func GetTotalTopUps() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var amount models.Accounts
		var total models.TopUps
		if total.Created_at == time.Now().Add(time.Duration(time.Now().Month())) {
			if err:= database.Db.Find(&total).Error; err!= nil{
				ctx.AbortWithStatus(http.StatusNotFound)
				fmt.Println(err)
			} else {
				ctx.JSON(http.StatusOK, total.Amount)
			}
			if err:= database.Db.Find(&amount).Error; err!= nil{
				ctx.AbortWithStatus(http.StatusNotFound)
				fmt.Println(err)
			} else {
				ctx.JSON(http.StatusOK, amount.Balance)
			}
		}

		
	}
}
