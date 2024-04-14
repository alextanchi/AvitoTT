package controller

import (
	"AvitoTestTask/internal/convert"
	"AvitoTestTask/internal/middleware"
	"AvitoTestTask/internal/models"
	"AvitoTestTask/internal/repository"
	"AvitoTestTask/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Controller interface {
	CreateBanner(ctx *gin.Context)
	DeleteBannerById(ctx *gin.Context)
	GetBanner(ctx *gin.Context)
	GetBannersByFeatureAndTag(ctx *gin.Context)
	UpdateBannerById(ctx *gin.Context)
}
type BannerController struct {
	useCase    service.Service
	middleware middleware.IBannerMiddleware
}

func NewController(srv service.Service) Controller {
	return &BannerController{
		useCase: srv,
	}
}

func (c BannerController) CreateBanner(ctx *gin.Context) {
	var id int
	banner := models.CreateBannerRequest{}
	err := ctx.ShouldBind(&banner)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные",
		})
		return
	}
	id, err = c.useCase.CreateBanner(ctx, banner)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Внутренняя ошибка сервера",
		})
		return
	}
	response := models.CreateBannerResponse{
		BannerId: id,
	}
	ctx.JSON(http.StatusCreated, response)
	return
}

func (c BannerController) DeleteBannerById(ctx *gin.Context) {
	id := ctx.Param("id")
	idStr, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные",
		})
		return
	}
	err = c.useCase.DeleteBannerById(ctx, idStr)
	if errors.Is(err, repository.ErrBannerNotFound) {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Внутренняя ошибка сервера",
		})
		return
	}
	ctx.JSON(http.StatusNoContent, "Баннер успешно удален")
	return
}

func (c BannerController) GetBanner(ctx *gin.Context) {
	var err error
	banner := models.UserBannerFilter{}
	feature, ok := ctx.GetQuery("feature_id")
	if ok {
		f, _ := strconv.Atoi(feature)
		banner.FeatureId = f
	}
	tag, ok := ctx.GetQuery("tag_id")
	if ok {
		t, _ := strconv.Atoi(tag)
		banner.TagId = t
	}
	useLastRevision, ok := ctx.GetQuery("use_last_revision")
	if ok {
		u, _ := strconv.ParseBool(useLastRevision)
		banner.UseLastRevision = &u
	}
	data, err := c.useCase.GetBanner(ctx, banner)
	if errors.Is(err, repository.ErrBannerNotFound) {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Внутренняя ошибка сервера",
		})
		return
	}
	role, ok := ctx.Get("role")
	if ok {
		r := role.(string)
		if r == middleware.UserRole && !data.IsActive {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": repository.ErrBannerNotFound.Error(),
			})
		}
		return
	}
	response := convert.DomainBannerToContent(data)
	ctx.JSON(http.StatusOK, response)
	return
}

func (c BannerController) GetBannersByFeatureAndTag(ctx *gin.Context) {
	var err error
	banner := models.BannerListFilter{}
	feature, ok := ctx.GetQuery("feature_id")
	if ok {
		f, _ := strconv.Atoi(feature)
		banner.FeatureId = &f
	}
	tag, ok := ctx.GetQuery("tag_id")
	if ok {
		t, _ := strconv.Atoi(tag)
		banner.TagId = &t
	}
	limit, ok := ctx.GetQuery("limit")
	if ok {
		l, _ := strconv.Atoi(limit)
		banner.Limit = uint64(l)
	}
	offset, ok := ctx.GetQuery("offset")
	if ok {
		o, _ := strconv.Atoi(offset)
		banner.Offset = uint64(o)
	}
	data, err := c.useCase.GetBannersByFeatureAndTag(ctx, banner)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Внутренняя ошибка сервера",
		})
		return
	}
	response := convert.BannerByFeatureAndTagToBanner(data)
	ctx.JSON(http.StatusOK, response)
}

func (c BannerController) UpdateBannerById(ctx *gin.Context) {
	banner := models.BannerUpdateById{}
	id := ctx.Param("id")
	idStr, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные",
		})
		return
	}
	err = ctx.ShouldBind(&banner)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Некорректные данные",
		})
		return
	}
	banner.BannerId = idStr
	err = c.useCase.UpdateBannerById(ctx, banner)
	if errors.Is(err, repository.ErrBannerNotFound) {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Внутренняя ошибка сервера",
		})
		return
	}
	ctx.Status(http.StatusOK)
	return
}
