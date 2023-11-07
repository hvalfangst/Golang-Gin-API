package route

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	CarDetailsHandler "hvalfangst/golang-gin-api-with-bun/src/handler"
)

func ConfigureRoute(r *gin.Engine, db *bun.DB) {
	r.GET("/car-details/:id", CarDetailsHandler.GetCarDetails(db))
	r.GET("/cars/:id", CarDetailsHandler.GetCarById(db))
}
