package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	CarDetailsService "hvalfangst/golang-gin-api-with-bun/src/service"
	"strconv"
)

func GetCarDetails(db *bun.DB) gin.HandlerFunc {
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
