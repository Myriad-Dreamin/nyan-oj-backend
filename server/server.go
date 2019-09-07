package main

import (
	"fmt"

	server "github.com/Myriad-Dreamin/core-oj/gin-server"

	"context"

	config "github.com/Myriad-Dreamin/core-oj/config"
	"github.com/Myriad-Dreamin/core-oj/log"
	"github.com/gin-gonic/gin"

	// import driver
	_ "github.com/go-sql-driver/mysql"

	jwt "github.com/Myriad-Dreamin/gin-middleware/auth/jwt"
	privileger "github.com/Myriad-Dreamin/gin-middleware/auth/privileger"
	morm "github.com/Myriad-Dreamin/nyan-oj-backend/types/orm"
	rbac "github.com/Myriad-Dreamin/nyan-oj-backend/types/rbac"
	"github.com/go-xorm/xorm"
)

func main() {
	userx, err := morm.NewUserX()
	if err != nil {
		fmt.Println(err)
		return
	}

	logger, err := log.NewZapColorfulDevelopmentSugarLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	engine, err := xorm.NewEngine(config.Config().DriverName, config.Config().MasterDataSourceName)
	if err != nil {
		logger.Error("prepare failed", "error", err)
		return
	}

	err = rbac.Init(engine)
	if err != nil {
		logger.Error("prepare failed", "error", err)
		return
	}

	morm.RegisterEngine(engine)

	engine.ShowSQL(true)

	x := rbac.GetEnforcer()

	srv, err := server.NewServer()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = srv.DefaultPrepare(config.Config())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer srv.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	codeRouter := r.Group("/code")
	srv.DefaultCodeRouter(codeRouter)
	problemRouter := r.Group("/problem")
	srv.DefaultProblemRouter(problemRouter)
	problemFSRouter := r.Group("/problemfs")
	srv.DefaultProblemFSRouter(problemFSRouter)

	userRouter := r.Group("/user")
	{
		var userService = NewUserService(userx, logger)
		// userRouter.GET("/:id", userService.Get)
		// userRouter.GET("/:id/content", userService.GetContent)
		// userRouter.GET("/:id/result", userService.GetResult)
		userRouter.POST("/register", userService.Register)
		userRouter.POST("/login", userService.Login)
		// userRouter.PUT("/:id/updateform-runtimeid", userService.UpdateRuntimeID)
		// userRouter.DELETE("/:id", userService.Delete)
	}

	authmw := privileger.NewMiddleWare(&x, "user:")
	r.Use(authmw.Build())
	// _ = authmw
	apiRouter := r.Group("/api")
	{
		apiRouter.Use(jwt.NewMiddleWare().Build())
		apiRouter.GET("/authv2", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "orzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz",
			})
		})

		// apiRouter

		apiRouter.Group("/")
		apiRouter.GET("/authv3", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "orzzzzzzzzzzzzzzzzzzzzzzzzz",
			})
		})

		authRouter := apiRouter.Group("/auth")
		{
			var authService = NewAuthService(logger)
			authRouter.GET("/policy", authService.GetPolicy)
			authRouter.PUT("/policy", authService.AddPolicy)
			authRouter.GET("/group/policy", authService.GetGroupingPolicy)
			authRouter.PUT("/group/policy", authService.AddGroupingPolicy)

		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	go srv.Serve(ctx)
	defer cancel()
	r.Run(":23336")
}
