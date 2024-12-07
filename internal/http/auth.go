package http

import (
	"encoding/json"
	"log"
	"net/http"

	"example.com/expenses-tracker/internal/handlers"
	"example.com/expenses-tracker/internal/requests"
	"example.com/expenses-tracker/internal/responses"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	internalHandler handlers.AuthHandler
}

// NewAuthHandler creates a new auth http handler
func NewAuthHandler(internalHandler handlers.AuthHandler) *AuthHandler {
	return &AuthHandler{internalHandler: internalHandler}
}

func (h *AuthHandler) RegisterRoutes(g *gin.Engine) {
	g.POST("/login", h.loginUser)
	g.POST("/forgot-password", h.initiateForgottenPassword)
	g.POST("/reset-password", h.updatePasswordFromResetRequest)
	g.POST("/logout", h.logoutUser)
}

func (h *AuthHandler) loginUser(c *gin.Context) {
	var loginRequest requests.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewErrorJsonHttpResponse(http.StatusBadRequest, loginRequest, err))
		return
	}

	response, err := h.internalHandler.HandleLoginRequest(c, &loginRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.NewErrorJsonHttpResponse(http.StatusUnauthorized, loginRequest, err))
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, responses.NewErrorJsonHttpResponse(http.StatusInternalServerError, loginRequest, err))
		return
	}

	c.JSON(http.StatusOK, jsonResponse)
}

func (h *AuthHandler) initiateForgottenPassword(c *gin.Context) {

}

func (h *AuthHandler) updatePasswordFromResetRequest(c *gin.Context) {

}

func (h *AuthHandler) logoutUser(c *gin.Context) {
	token, exists := c.Get("user_token")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are not logged in"})
		return
	}

	success, err := h.internalHandler.HandleLogout(c, token.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "logout could not be completed"})
		return
	}

	if success {
		c.JSON(http.StatusOK, gin.H{"success": true})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "logout could not be completed"})
}
