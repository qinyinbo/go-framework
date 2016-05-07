package handler

import (
	"github.com/fuxiaohei/GoInk"
	"server/model"
)

func Logout(context *GoInk.Context) {

	msg, data := model.Test("test mysql")
	//context.Body = []byte(msg)
	context.Layout("test/test")
	context.Render("test/test", map[string]interface{}{
		"Msg":  msg,
		"Data": data,
	})

}
