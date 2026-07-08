package delivery

import (
    "auth-service/domain"
	"auth-service/middleware"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func NewAuthHandler(r *gin.Engine, usecase domain.AuthUsecase) {
    handler := &AuthHandler{
        AuthUsecase: usecase,
    }

    api := r.Group("/api")
    {
        api.POST("/login", handler.Login)
        api.POST("/register", handler.Register)
        api.POST("/logout", handler.Logout)
        api.POST("/refresh", handler.RefreshToken) 
        
        protected := api.Group("")
        protected.Use(middleware.RequireAuth())
        protected.GET("/me", handler.GetMe)
    }
}

func (a *AuthHandler) Login(c *gin.Context) {
	var req domain.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := a.AuthUsecase.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", res.AccessToken, 900, "/", "localhost", false, true)
	c.SetCookie("refresh_token", res.RefreshToken, 604800, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": res.Message})
}

func (a *AuthHandler) Register(c *gin.Context) {
	var req domain.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := a.AuthUsecase.Register(&req)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", res.AccessToken, 900, "/", "localhost", false, true)
	c.SetCookie("refresh_token", res.RefreshToken, 604800, "/", "localhost", false, true)

	c.JSON(http.StatusCreated, gin.H{"message": res.Message})
}

func (a *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (a *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token missing"})
		return
	}

	newAccessToken, err := a.AuthUsecase.Refresh(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	c.SetCookie("access_token", newAccessToken, 900, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}

func (a *AuthHandler) GetMe(c *gin.Context) {
    userIDData, exists := c.Get("user_id")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    userID := uint(userIDData.(float64))

    user, err := a.AuthUsecase.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}