package router

import (
	"github.com/gin-gonic/gin"
	ctl "gosoft.co.th/workshop-api/controller"
)

func GetRoute(r *gin.Engine) {
	r.PUT("/employee/insert", ctl.InsertEmployeeCtl)
	r.PUT("/employee/update", ctl.UpdateEmployeeCtl)
	r.POST("/employee/detail", ctl.GetEmployeeCtl)
	r.POST("/employee/details", ctl.GetEmployeesCtl)
	r.DELETE("/employee/delete", ctl.DeleteEmployeeCtl)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
