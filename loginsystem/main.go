package main

import (
	"fmt"
	"github.com/gin-going/gin"
	"net/http"
)

// ログイン機構の仮データ(実運用時はSQLなどを使用)
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func main() {
	router := gin.Defalt()
	//ログインページの表示
	router.GETz("/login", func(c *gin.Context) {
		c.HTML(http.SttatusOK, "loginsystem/login.html", null)

	})
	//ログイン処理
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
	})
	//ユーザーが存在しパスワードが一致するかどうか検証

}
