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

}
