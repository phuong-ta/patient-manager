package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// newRouter creates and returns a configured *gin.Engine with all routes registered.
func newRouter(s *Server) *gin.Engine {
	r := gin.Default()

	// health/home
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "This is my home page") })

	// Patient CRUD
	patients := r.Group("/patients")
	{
		patients.GET("", s.ListPatients)
		patients.POST("", s.CreatePatient)
		patients.GET(":id", s.GetPatient)
		patients.PUT(":id", s.UpdatePatient)
		patients.DELETE(":id", s.DeletePatient)
	}

	return r
}
