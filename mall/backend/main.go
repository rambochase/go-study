package main

//go get -u github.com/kataras/iris
import (
	"github.com/kataras/iris"
)


func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	//注册模板
	tmplate := iris.HTML("./web/views","html").Layout(
		"shared/layout.html").Reload(true)
	app.RegisterView(tmplate)

	//设置模板目标
	app.StaticWeb("/assets",
		"web/assets")

	//出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",ctx.Values().GetStringDafault("message","访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	//注册控制器

	//监听对应端口
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
)

}