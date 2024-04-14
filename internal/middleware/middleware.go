package middleware

import (
	"AvitoTestTask/internal/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

const (
	AdminRole = "admin"
	UserRole  = "user"
)

type IBannerMiddleware interface {
	Authorization() gin.HandlerFunc
	CheckRole(isAdmin bool) gin.HandlerFunc
}

type BannerMiddleware struct {
	repo repository.Repository
}

func NewBannerMiddleware(repo repository.Repository) IBannerMiddleware {
	return BannerMiddleware{
		repo: repo,
	}

}

func (b BannerMiddleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		str := ctx.Request.Header.Get("Authorization")
		if !strings.Contains(str, "Bearer") {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("Пользователь не авторизован"))
			return
		}
		token := strings.Split(str, "Bearer ")
		t, _, err := new(jwt.Parser).ParseUnverified(token[1], jwt.MapClaims{})
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("Пользователь не авторизован"))
			return
		}

		var userId string
		if claims, ok := t.Claims.(jwt.MapClaims); ok {
			uu := claims["userId"]
			userId = uu.(string)
		}

		if userId == "" {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("Пользователь не авторизован"))
			return
		}

		ctx.Set("userId", userId)
	}
}

func (b BannerMiddleware) CheckRole(isAdmin bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, ok := ctx.Get("userId")
		if !ok {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("Пользователь не авторизован"))
			return
		}
		role, err := b.repo.GetRole(ctx, userId.(string))
		if isAdmin {
			if role != AdminRole {
				ctx.AbortWithError(http.StatusForbidden, errors.New("Пользователь не имеет доступа"))
				return
			}
		}
		if err != nil {
			ctx.AbortWithError(http.StatusForbidden, err)
			return
		}
		ctx.Set("role", role)
	}
}
