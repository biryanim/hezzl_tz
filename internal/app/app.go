package app

import (
	"context"
	"github.com/biryanim/hezzl_tz/internal/api/middleware"
	"github.com/biryanim/hezzl_tz/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	loggerCtx, loggerCancel := context.WithCancel(context.Background())
	defer loggerCancel()
	go a.serviceProvider.LoggerService(loggerCtx).Run(loggerCtx)
	err := a.runHTTPServer()
	if err != nil {
		log.Fatalf("failed to run http server: %v", err)
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load("local.env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	router := gin.Default()

	router.Use(middleware.ErrorHandler())

	public := router.Group("/good")
	{
		public.POST("/create", a.serviceProvider.GoodsImpl(ctx).Create)
		public.PATCH("/update", a.serviceProvider.GoodsImpl(ctx).Update)
		public.DELETE("/remove", a.serviceProvider.GoodsImpl(ctx).Delete)
		public.GET("/list", a.serviceProvider.GoodsImpl(ctx).List)
		public.PATCH("/reprioritize", a.serviceProvider.GoodsImpl(ctx).Reprioritize)
	}

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: router,
	}

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())
	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
