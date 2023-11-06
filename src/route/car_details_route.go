package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	CarDetailsHandler "hvalfangst/imperative-golang-gin-api/src/handler"
)

func ConfigureRoute(r *gin.Engine, db *pg.DB) {
	r.GET("/car-details/:id", CarDetailsHandler.GetCarDetails(db))
}
