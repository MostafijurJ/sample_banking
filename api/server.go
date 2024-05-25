package api

import (
	"github.com/gin-gonic/gin"
	db "sample_banking/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  *db.Queries
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store *db.Queries) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/home", server.home)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/transfers", server.createTransfer)
	router.GET("/transfers/:id", server.getTransfer)

	server.router = router
	return server
}
