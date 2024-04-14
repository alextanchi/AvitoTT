package bootstrap

import (
	"AvitoTestTask/internal/controller"
	"AvitoTestTask/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	cnt        controller.Controller
	mdw        middleware.IBannerMiddleware
}

// NewServer объединили  пакет http , контроллер и мидлвар
func NewServer(cnt controller.Controller, mdw middleware.IBannerMiddleware) Server {
	return Server{
		httpServer: &http.Server{
			Addr:           ":8080",
			MaxHeaderBytes: 1 << 20,          //1MB
			ReadTimeout:    10 * time.Second, //10 сек
			WriteTimeout:   10 * time.Second,
		},
		cnt: cnt,
		mdw: mdw,
	}
}

// InitRoutes инициализируем все наши эндпоинты
func (s Server) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/user_banner",
		//s.mdw.Authorization(), s.mdw.CheckRole(false),
		s.cnt.GetBanner)
	router.GET("/banner",
		//s.mdw.Authorization(), s.mdw.CheckRole(true),
		s.cnt.GetBannersByFeatureAndTag)
	router.POST("/banner",
		//s.mdw.Authorization(), s.mdw.CheckRole(true),
		s.cnt.CreateBanner)
	router.PATCH("/banner/:id",
		//s.mdw.Authorization(), s.mdw.CheckRole(true),
		s.cnt.UpdateBannerById)
	router.DELETE("/banner/:id",
		//	s.mdw.Authorization(), s.mdw.CheckRole(true),
		s.cnt.DeleteBannerById)

	return router

}
