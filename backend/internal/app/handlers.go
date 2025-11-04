package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Here we implement basic handler stubs. Later these should interact with services/repositories.

// ListPatients returns a list of patients (stub).
func (s *Server) ListPatients(c *gin.Context) {
	// TODO: replace with real logic
	c.JSON(http.StatusOK, gin.H{"data": []string{}, "message": "list patients - stub"})
}

// CreatePatient creates a new patient (stub).
func (s *Server) CreatePatient(c *gin.Context) {
	// TODO: parse body, validate, persist
	c.JSON(http.StatusCreated, gin.H{"message": "create patient - stub"})
}

// GetPatient returns a single patient by id (stub).
func (s *Server) GetPatient(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "get patient - stub"})
}

// UpdatePatient updates a patient by id (stub).
func (s *Server) UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "update patient - stub"})
}

// DeletePatient deletes a patient by id (stub).
func (s *Server) DeletePatient(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "delete patient - stub"})
}
