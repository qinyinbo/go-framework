//todo config
package server

import (
	"github.com/fuxiaohei/GoInk"
	"net/http"
	"os"
	"runtime/debug"
	"server/handler"
	"strings"
)

var (
	// Global GoInk application
	App *GoInk.App
)

// default handler, implement GoInk.Handler
func homeHandler(context *GoInk.Context) {
	context.Body = []byte("Hello GoInk !")
	context.Layout("test/test")
	context.Render("test/test", nil)
}

func Run() {
	// only bind GET handler by homeHandler.
	// if other http method, return 404.
	App = GoInk.New()

	App.Static(func(context *GoInk.Context) {
		url := strings.TrimPrefix(context.Url, "/")
		if strings.HasPrefix(url, "static") {
			http.ServeFile(context.Response, context.Request, url)
			context.IsEnd = true
		}
	})

	// set recover handler
	App.Recover(func(context *GoInk.Context) {
		context.Body = append([]byte("<pre>"), context.Body...)
		context.Body = append(context.Body, []byte("\n")...)
		context.Body = append(context.Body, debug.Stack()...)
		context.Body = append(context.Body, []byte("</pre>")...)
		context.End()
	})

	// set not found handler
	App.NotFound(func(context *GoInk.Context) {
		context.Status = 404
		context.Body = []byte("not found")
		context.End()
	})

	// add recover defer
	defer func() {
		e := recover()
		if e != nil {
			println("panic error, crash down")
			os.Exit(1)
		}
	}()

	// catch exit command

	// run application.
	// it listens localhost:9000 in pre-defined config.
	App.Get("/", homeHandler)
	App.Get("/logout/", handler.Logout)
	App.Run()
}
