package controller

import (
	// "fmt"
	"fmt"
	"hello-world/db"
	"hello-world/models"
	"net/http"

	// "strconv"

	// "strings"

	"github.com/codehand/echo-restful-crud-api-example/types"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

// Getusers is func get all user
func GetAllCompany(c *gin.Context) {
	data, err := db.GetAllCompany()
	if err != nil {
		c.JSON(http.StatusNotFound, types.ParseStatus("NOT_FOUND", err.Error()))
		return
	}

	resp := data
	c.JSON(http.StatusOK, resp)
}

// Createuser is func create new user
func CreateCompany(c *gin.Context) {

	var objRequest models.Company

	if err := c.Bind(&objRequest); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_ERR", "Please Enter Data"))
		return
	}

	name, _ := db.ValidateCompany(objRequest.Company_Code)
	if name != nil {
		c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_ERR", "company_code Dupplicate"))
		return
	}

	_, err := db.CreateCompany(&objRequest)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, types.ParseStatus("NOT_ACCEPTED", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, &objRequest)

}

// Getuser is func get one user
func GetCompanyById(c *gin.Context) {
	fmt.Println("get : ", c.Param("id"))
	id := c.Param("id")
	fmt.Println(id)
	data, err := db.GetCompanyById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, types.ParseStatus("NOT_FOUND", err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

// Updateuser is func update one user
func UpdateCompany(c *gin.Context) {
	id := c.Param("id")

	var objRequest models.Company

	if err := c.Bind(&objRequest); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_ERR", "Please enter role id"))
		return
	}

	data, err := db.UpdateCompany(id, &objRequest)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, types.ParseStatus("NOT_ACCEPTED", err.Error()))
		return
	}
	c.JSON(http.StatusOK, data)
}

// Deleteuser is func delete one user
// func DeleteExample(c echo.Context) error {
// 	id := c.Param("id")

// 	data, err := db.DeleteAtUser(id)
// 	if err != nil {
// 		return c.JSON(http.StatusNotAcceptable, types.ParseStatus("NOT_ACCEPTED", err.Error()))
// 	}
// 	return c.JSON(http.StatusOK, data)
// }
