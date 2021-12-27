package controllers

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/heavenmise/vicco-shoes-api/pkg/config"
	"github.com/heavenmise/vicco-shoes-api/pkg/models"
	"github.com/heavenmise/vicco-shoes-api/pkg/utils"
)

func GetShoes(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var shoe models.Shoe
	shoeLists, err := models.GetShoes(&shoe, &pagination)
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

func GetShoe(c *gin.Context) {
	var name string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value) - 1]
		if key == "name" {
			name = queryValue
		}
	}
	// fmt.Println(name)
	shoe, err := models.GetShoe(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": shoe,
	})
}

func CreateShoe(c *gin.Context) {
	NewShoe := &models.Shoe{}
	utils.ParseBody(c, NewShoe)
	Shoe := NewShoe.CreateShoe()
	c.JSON(http.StatusOK, gin.H{
		"data": Shoe,
	})
}

func UpdateShoe(c *gin.Context) {
	var updateShoe = &models.Shoe{}
	utils.ParseBody(c, updateShoe)

	var name string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value) - 1]
		if key == "name" {
			name = queryValue
		}
	}
	
	shoe, _ := models.GetShoe(name)
	// if err != nil {
	// 	panic(err)
	// }

	if updateShoe.Name != "" {
		shoe.Name = updateShoe.Name
	}
	if updateShoe.Article != "" {
		shoe.Article = updateShoe.Article
	}
	if updateShoe.Color != "" {
		shoe.Color = updateShoe.Color
	}
	if updateShoe.Size != "" {
		shoe.Size = updateShoe.Size
	}
	if updateShoe.Gender != "" {
		shoe.Gender = updateShoe.Gender
	}
	if updateShoe.Material != "" {
		shoe.Material = updateShoe.Material
	}
	if updateShoe.Collection != "" {
		shoe.Collection = updateShoe.Collection
	}
	if updateShoe.Categories != "" {
		shoe.Categories = updateShoe.Categories
	}
	if updateShoe.Price != "" {
		shoe.Price = updateShoe.Price
	}
	
	config.DB.Save(&shoe)
	c.JSON(http.StatusOK, gin.H{
		"data": shoe,
	})
}

func DeleteShoe(c *gin.Context) {
	var name string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value) - 1]
		if key == "name" {
			name = queryValue
		}
	}
	// fmt.Println(name)
	shoe, err := models.DeleteShoe(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": shoe,
	})
}