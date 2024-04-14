package app

import (
	"AvitoTestTask/internal/bootstrap"
	"AvitoTestTask/internal/cache"
	"AvitoTestTask/internal/controller"
	"AvitoTestTask/internal/middleware"
	"AvitoTestTask/internal/repository"
	"AvitoTestTask/internal/service"
	"context"
	"log"
	"time"
)

func Run() error {
	ctx := context.Background()
	db, err := repository.ConnectDb()
	if err != nil {
		log.Println(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	store := repository.NewBanner(db)
	mdw := middleware.NewBannerMiddleware(store)
	temp := cache.NewBannerCache(store)
	go func() {
		err = temp.Refresh(ctx)
		if err != nil {
			log.Println(err)
		}
		timer := time.NewTimer(5 * time.Minute)
		for {
			select {
			case <-timer.C:
				err = temp.Refresh(ctx)
				if err != nil {
					log.Println(err)
				}
			}
		}

	}()
	srv := service.NewService(store, temp)
	cnt := controller.NewController(srv)
	serv := bootstrap.NewServer(cnt, mdw)
	router := serv.InitRoutes()
	router.Run(":8080")
	return nil
}
