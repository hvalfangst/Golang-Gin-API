package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	CarDetailsService "hvalfangst/imperative-golang-gin-api/src/service"
	"strconv"
)

func GetCarDetails(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid Car ID"})
			return
		}

		// Query user by email
		carDetails, err := CarDetailsService.GetCarDetails(db, id)
		if err != nil {
			c.JSON(404, gin.H{"error": "CarDetails doesn't exist"})
			return
		}
		c.JSON(200, gin.H{"car_details": carDetails})
	}
}
