package app

import (
	"github.com/gin-gonic/gin"
	"github.com/voicurobert/bookstore_oauth-api/src/http"
	"github.com/voicurobert/bookstore_oauth-api/src/repository/db"
	"github.com/voicurobert/bookstore_oauth-api/src/repository/users_rest"
	"github.com/voicurobert/bookstore_oauth-api/src/services"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewAccessTokenHandler(services.NewService(users_rest.NewRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
