package main

import (
	"cs/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Print("hello world")
	engine := gin.Default()
	engine.Any("/", WebRoot)
	engine.Run(":8080")
}

func WebRoot(context *gin.Context) {
	user := model.UserInfo{}
	user.ID = 1
	user.Name = "你好"
	user.State = 12
	fmt.Print(user)
	context.String(http.StatusOK, "hello, world")
}
