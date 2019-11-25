package route

import (
	"github.com/gin-gonic/gin"
)

/*
	@file	router.go
	@author	LoperLee
	@date	2019-11-20
*/

/*
	@author	LoperLee
	@brief	Start echo rest-api server
	@return	*Echo
*/
func StartAPIServer() *gin.Engine {
	engine := gin.Default()

	v1 := engine.Group("v1")
	{
		v1.POST("/calendar", func(con *gin.Context) {

		})
	}
	return engine
}
