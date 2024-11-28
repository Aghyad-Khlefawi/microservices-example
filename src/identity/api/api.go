package api

import (
	"net/http"

	"github.com/aghyad-khlefawi/identity/pkg/servicecollection"
	"github.com/gin-gonic/gin"
)

type ApiContext struct{
  sc *servicecollection.ServiceCollection
}


func RegisterRoutes(router *gin.Engine,sc * servicecollection.ServiceCollection) {

  context:= &ApiContext{
		sc,
	}

 router.POST("/user",NewHandler(HandleCreateUser, context));
 router.POST("/auth",NewHandler(HandleAuthenticateUser, context));

  router.GET("/hc", func(c *gin.Context) {
		c.String(http.StatusOK,"Healthy")
	})


}


func NewHandler(handler (func(*gin.Context,*ApiContext)), apiContext *ApiContext) (func(*gin.Context)){

	return func(c *gin.Context) {
		handler(c,apiContext)
	}
}
