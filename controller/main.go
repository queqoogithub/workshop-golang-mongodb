package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gosoft.co.th/workshop-api/models"
	sv "gosoft.co.th/workshop-api/service"
	st "gosoft.co.th/workshop-api/structs"
)

func InsertEmployeeCtl(c *gin.Context) {
	req := models.EmployeeModel{}
	jsonErr := json.NewDecoder(c.Request.Body).Decode(&req)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(jsonErr))
		return
	}
	id, insErr := sv.InsertEmployeeSv(req)
	if insErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(insErr))
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
func UpdateEmployeeCtl(c *gin.Context) {
	req := st.EmployeeReq{}
	jsonErr := json.NewDecoder(c.Request.Body).Decode(&req)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(jsonErr))
	}
	n, updErr := sv.UpdateEmployeeSv(req)
	if updErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(updErr))
	}
	c.JSON(http.StatusCreated, gin.H{
		"rowAffect": n,
	})
}
func GetEmployeeCtl(c *gin.Context) {
	req := st.EmployeeReq{}
	jsonErr := json.NewDecoder(c.Request.Body).Decode(&req)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(jsonErr))
	}
	obj, objErr := sv.GetEmployeeSv(req)
	if objErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(objErr))
	}
	c.JSON(http.StatusOK, obj)
}
func GetEmployeesCtl(c *gin.Context) {
	req := st.EmployeeReq{}
	jsonErr := json.NewDecoder(c.Request.Body).Decode(&req)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(jsonErr))
	}
	listObj, totalResults, totalPages, objErr := sv.GetEmployeesSv(req)
	if objErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(objErr))
	}
	c.JSON(http.StatusOK, gin.H{
		"datas":        listObj,
		"totalResults": totalResults,
		"totalPages":   totalPages,
	})
}
func DeleteEmployeeCtl(c *gin.Context) {
	req := st.EmployeeReq{}
	jsonErr := json.NewDecoder(c.Request.Body).Decode(&req)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(jsonErr))
	}
	delNum, delErr := sv.DeleteEmployeeSv(req)
	if delErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(delErr))
	}
	c.JSON(http.StatusOK, gin.H{
		"rowAffect": delNum,
	})
}
