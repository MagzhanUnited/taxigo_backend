package main

import (
	"avicena/controllers"

	"avicena/models"

	"avicena/scraping"

	"github.com/gin-gonic/gin"

	"avicena/middlewares"
)

func main() {
	models.ConnectDataBase()
	r := gin.Default()
	public := r.Group("/taxigo")
	public.POST("/book", models.Book)
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	public.GET("/data", scraping.ScrapedData)
	public.GET("/data/:id", scraping.GetInner)
	public.GET("/data/menu/:id", scraping.GetMenu)
	public.GET("/book/all", models.GetAllBookReqs)
	public.GET("/book/:number", models.GetBookReqsByNumber)
	public.POST("/book/update/:id", models.UpdateBookReq)
	public.POST("/book/food", models.PostBook)
	public.POST("/book/food/get", models.GetBook)
	public.POST("/order", models.PostOrderReq)
	public.GET("/order/all", models.GetAllOrderReqs)
	public.GET("/order/:number", models.GetOrderReqsByNumber)

	protected := r.Group("/taxigo/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	r.Run(":8000")
}
