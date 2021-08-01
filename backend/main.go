package main

import (
	"Chise1/imooc-product/backend/web/controllers"
	"Chise1/imooc-product/db"
	"Chise1/imooc-product/repositories"
	"Chise1/imooc-product/services"
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	//初始化数据库
	db.InitDB("product.db")
	//1.创建iris 实例
	app := iris.New()
	//2.设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//3.注册模板
	tmplate := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	//4.设置模板目标
	app.HandleDir("/assets", iris.Dir("./backend/web/assets/"))
	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	//上下文控制器
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//5.注册控制器
	productRepository := repositories.NewProductManager()
	productSerivce := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productSerivce)
	product.Handle(new(controllers.ProductController))
	//注册order
	orderRepository := repositories.NewOrderManager()
	orderService := services.NewOrderService(orderRepository)
	fmt.Println(&orderService)
	orderParty := app.Party("/order")
	order := mvc.New(orderParty)
	order.Register(ctx, orderService)
	order.Handle(new(controllers.OrderController))

	//6.启动服务
	err := app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
	if err != nil {
		return
	}
}
