package app

import (
	"github.com/gin-gonic/gin"
	"github.com/voicurobert/bookstore_oauth-api/src/clients/cassandra"
	"github.com/voicurobert/bookstore_oauth-api/src/domain/access_token"
	"github.com/voicurobert/bookstore_oauth-api/src/http"
	"github.com/voicurobert/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	service := access_token.NewService(db.NewRepository())
	atHandler := http.NewHandler(service)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
