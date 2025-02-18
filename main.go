package main

import (
	"gin-practice/controllers"
	"gin-practice/infra"

	//"gin-practice/models"
	"gin-practice/repositories"
	"gin-practice/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	// items := []models.Item{
	// 	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
	// 	{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
	// 	{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	// }

	//itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	ItemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", ItemController.FindAll)
	r.GET("/items/:id", ItemController.FindById)
	r.POST("/items", ItemController.Create)
	r.PUT("/items/:id", ItemController.Update)
	r.DELETE("/items/:id", ItemController.Delete)
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
