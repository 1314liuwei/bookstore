package cmd

import (
	"back/internal/controller"
	"back/internal/service"
	"context"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of bookstore",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(service.Middleware().CORS)

			pwd, _ := os.Getwd()
			s.BindHandler("/test/:file", func(r *ghttp.Request) {
				file := r.GetRouter("file")
				r.Response.ServeFileDownload(pwd + "/ebooks/" + file.String())
			})

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareHandlerResponse,
				)
				group.Bind(
					controller.User,
					controller.Book,
				)

				group.Middleware(
					service.Middleware().Auth,
				)

				group.GET("/user/profile", controller.User.Profile)
				group.Bind(
					controller.Order,
					controller.ShoppingCart,
				)
			})
			s.Run()
			return nil
		},
	}
)
