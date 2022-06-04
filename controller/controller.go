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
// GetBalance godoc
// @Summary      Show balance.
// @ID get-all-Balance
// @Description  get all balance from the database.
// @Tags         Accounts
// @Accept       json
// @Produce      json
// @Success      200  {struct}  models.Accounts
// @Failure      404  {object}  models.Accounts
// @Router       /balance [post]
func GetBalance() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var balance []models.Accounts
		defer cancel()
		if err := database.Db.Find(&balance).Error; err != nil {
			ctx.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			ctx.JSON(200, balance)
		}
	}

}
// TopUp godoc
// @Summary      Top up the Balance.
// @ID top-up
// @Description  Top-up e-wallet account.
// @Tags         TopUps
// @Accept       json
// @Produce      json
// @Success      200  {struct}  models.TopUps
// @Failure      404  {object}  models.TopUps
// @Router       /top_up [post]
func TopUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var topUp models.TopUps
		err := ctx.BindJSON(&topUp)
		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		topUp.Created_at = time.Now().Local()
		database.Db.Create(&topUp)
		ctx.JSON(http.StatusOK, topUp)

	}
}
// CheckAcount godoc
// @Summary      Check user.
// @ID check-acount
// @Description  Check user if it is exist or not.
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param  	     id path int true "id"
// @Success      200  {object}  models.Users
// @Failure      404  {object}  models.Users
// @Router       /user/{id} [post]
func CheckAcount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		id := ctx.Params.ByName("id")
		log.Println("ID", id)
		var user models.Users
		defer cancel()
		if err := database.Db.Where("id=?", id).First(&user).Error; err != nil {
			ctx.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			ctx.JSON(http.StatusOK, user)
		}

	}
}
// GetTotalTopUps godoc
// @Summary      Get total balance.
// @ID get-all-Students
// @Description  Get total count and total amount of top-up operations for current month.
// @Tags         TopUps
// @Accept       json
// @Produce      json
// @Success      200  {struct}  models.TopUps
// @Failure      404  {object}  models.TopUps
// @Router       /total [post]
func GetTotalTopUps() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var amount models.Accounts
		var total models.TopUps
		if total.Created_at == time.Now().Add(time.Duration(time.Now().Month())) {
			if err := database.Db.Find(&total).Error; err != nil {
				ctx.AbortWithStatus(http.StatusNotFound)
				fmt.Println(err)
			} else {
				ctx.JSON(http.StatusOK, total.Amount)
			}
			if err := database.Db.Find(&amount).Error; err != nil {
				ctx.AbortWithStatus(http.StatusNotFound)
				fmt.Println(err)
			} else {
				ctx.JSON(http.StatusOK, amount.Balance)
			}
		}

	}
}
