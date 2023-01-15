package main

import (
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	rt "gosoft.co.th/workshop-api/router"
)

func init() {
	godotenv.Load()
	gin.SetMode(gin.DebugMode)
}

func main() {
	r := gin.New()
	r.Use(rt.CORSMiddleware())
	rt.GetRoute(r)
	r.Run(":8080")
}
