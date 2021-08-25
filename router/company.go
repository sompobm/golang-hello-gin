package router

import (
	"hello-world/controller"

	"github.com/gin-gonic/gin"
)

func Company(api *gin.RouterGroup) {

	// ADD example
	api.GET("/company", controller.GetAllCompany)      // Returns all resources of this dialogue
	api.POST("/company", controller.CreateCompany)     // Creates a resource of this dialogue and stores the data you posted, then returns the ID
	api.GET("/company/:id", controller.GetCompanyById) // Returns the resource of this dialogue with that ID
	api.PUT("/company/:id", controller.UpdateCompany)  // Updates the resource of this dialogue with that ID
	// api.DELETE("/example/:id", controller.DeleteExample)  // Deletes the resource of this dialogue with that ID
	// END complaint

}
