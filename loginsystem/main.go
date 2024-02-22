package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// ログイン機構の仮データ(実運用時はSQLなどを使用)
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func main() {
	router := gin.Default()
	//ログインページの表示
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	//ログイン処理
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		//ユーザーが存在しパスワードが一致するかどうか検証
		if storedPassword, ok := users[username]; ok {
			if password == storedPassword {
				c.HTML(http.StatusOK, "loggedin.html", gin.H{"username": username})
				return
			}
		}
		//ログイン失敗時の処理
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "Invalid credentials"})
	})

	router.LoadHTMLGlob("templates/*")

	//サーバー起動
	router.Run(":8080")
}
