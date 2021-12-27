package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/heavenmise/vicco-shoes-api/pkg/models"
	"github.com/heavenmise/vicco-shoes-api/pkg/utils"
)

func GetShoesPervyjShag(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var shoe models.Shoe
	shoeLists, err := models.GetShoesCategory(&shoe, &pagination, "pervyj-shag-19-21")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": shoeLists,
	})
}
func GetShoesMalyshi(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var shoe models.Shoe
	shoeLists, err := models.GetShoesCategory(&shoe, &pagination, "malyshi-22-25")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": shoeLists,
	})
}
func GetShoesDoshkolniki(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var shoe models.Shoe
	shoeLists, err := models.GetShoesCategory(&shoe, &pagination, "doskkolniki-26-30")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": shoeLists,
	})
}
func GetShoesShkolniki(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var shoe models.Shoe
	shoeLists, err := models.GetShoesCategory(&shoe, &pagination, "shkolniki-31-36")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": shoeLists,
	})
}
func GetShoesPodrostki(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var shoe models.Shoe
	shoeLists, err := models.GetShoesCategory(&shoe, &pagination, "podrostki-37-39")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": shoeLists,
	})
}