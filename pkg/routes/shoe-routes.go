package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/heavenmise/vicco-shoes-api/pkg/controllers"
)

func RegisterNewRouter() *gin.Engine {

	r := gin.Default()

	v0 := r.Group("/")
	{
		v0.GET("shop", controllers.GetShoes)
		v0.GET("product", controllers.GetShoe)
		v0.POST("product", controllers.CreateShoe)
		v0.PUT("product", controllers.UpdateShoe)
		v0.DELETE("product", controllers.DeleteShoe)
	}

	v1 := r.Group("/product-category")
	{
		v1.GET("pervyj-shag-19-21", controllers.GetShoesPervyjShag)
		v1.GET("malyshi-22-25", controllers.GetShoesMalyshi)
		v1.GET("doshkolniki-26-30", controllers.GetShoesDoshkolniki)
		v1.GET("shkolniki-31-36", controllers.GetShoesShkolniki)
		v1.GET("podrostki-37-39", controllers.GetShoesPodrostki)
	}

	return r
}