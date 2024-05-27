package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	db "sample_banking/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Balance  int64  `json:"balance" binding:"required,min=0"`
	Currency string `json:"currency" binding:"required, oneof=USD EUR BDT"`
}

func (server *Server) home(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"hello": "world from Gin!",
	})

}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}
type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type listAccountResponse struct {
	Accounts []db.Account `json:"accounts"`
}

func (server *Server) getAccount(context *gin.Context) {
	var req getAccountRequest
	if err := context.ShouldBindUri(&req); err != nil {
		context.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(context, req.ID)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		context.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	context.JSON(http.StatusOK, account)
}

func (server *Server) listAccounts(context *gin.Context) {
	var req listAccountRequest
	if err := context.ShouldBindQuery(&req); err != nil {
		context.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(context, arg)
	if err != nil {
		context.JSON(http.StatusInternalServerError, errorResponse(err))
		log.Fatal("cannot list accounts", err)
		return
	}
	context.JSON(http.StatusOK, accounts)
}

func (server *Server) createTransfer(context *gin.Context) {

}

func (server *Server) getTransfer(context *gin.Context) {

}

func (server *Server) createAccount(ctx *gin.Context) {

	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		log.Fatal("cannot bind JSON", err)
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  req.Balance,
		Currency: req.Currency,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		log.Fatal("cannot create account", err)
		return
	}

	ctx.JSON(http.StatusOK, account)

}

// Start  runs the server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
