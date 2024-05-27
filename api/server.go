package api

import (
	"github.com/gin-gonic/gin"
	db "sample_banking/db/sqlc"
	"sample_banking/db/utils"
	"sample_banking/token"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store      *db.Queries
	router     *gin.Engine
	tokenMaker token.Maker
	config     utils.Config
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config utils.Config, store *db.Queries) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}
	setupRouter(server)
	return server, nil
}

func setupRouter(server *Server) {
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/home", server.home)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.POST("/transfers", server.createTransfer)
	router.GET("/transfers/:id", server.getTransfer)

	//Add routes for users
	router.POST("/users", server.createUser)
	router.PUT("/users", server.updateUser)
	router.GET("/users/:username", server.getUserByUsername)
	router.POST("/users/login", server.loginUser)

	server.router = router
}
