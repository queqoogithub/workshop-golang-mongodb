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

func InsertCompanyCtl(c *gin.Context) {
	req := models.CompanyModel{}
	jsonErr := json.NewDecoder(c.Request.Body).Decode(&req)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(jsonErr))
		return
	}
	id, insErr := sv.InsertCompanySv(req)
	if insErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(insErr))
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
func UpdateCompanyCtl(c *gin.Context) {
	req := st.CompanyReq{}
	jsonErr := json.NewDecoder(c.Request.Body).Decode(&req)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(jsonErr))
	}
	n, updErr := sv.UpdateCompanySv(req)
	if updErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(updErr))
	}
	c.JSON(http.StatusCreated, gin.H{
		"rowAffect": n,
	})
}
func GetCompanyCtl(c *gin.Context) {
	req := st.CompanyReq{}
	jsonErr := json.NewDecoder(c.Request.Body).Decode(&req)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(jsonErr))
	}
	obj, objErr := sv.GetCompanySv(req)
	if objErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(objErr))
	}
	c.JSON(http.StatusOK, obj)
}
func GetCompaniesCtl(c *gin.Context) {
	req := st.CompanyReq{}
	jsonErr := json.NewDecoder(c.Request.Body).Decode(&req)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(jsonErr))
	}
	listObj, totalResults, totalPages, objErr := sv.GetCompaniesSv(req)
	if objErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(objErr))
	}
	c.JSON(http.StatusOK, gin.H{
		"datas":        listObj,
		"totalResults": totalResults,
		"totalPages":   totalPages,
	})
}
func DeleteCompanyCtl(c *gin.Context) {
	req := st.CompanyReq{}
	jsonErr := json.NewDecoder(c.Request.Body).Decode(&req)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(jsonErr))
	}
	delNum, delErr := sv.DeleteCompanySv(req)
	if delErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(delErr))
	}
	c.JSON(http.StatusOK, gin.H{
		"rowAffect": delNum,
	})
}
