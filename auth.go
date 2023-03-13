package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var UserData map[string]string

func init() {
	UserData = map[string]string{
		"test": "test",
	}
}

func IsUserExists(username string) bool {
	_, exist := UserData[username]
	return exist
}

func IsCorrectPassword(p1, p2 string) error {
	if p1 == p2 {
		return nil
	}
	return errors.New("password incorrect")
}

func Auth(user, pass string) error {
	exists := IsUserExists(user)
	if exists {
		return IsCorrectPassword(UserData[user], pass)
	} else {
		return errors.New("user is not exists")
	}
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginAuth(c *gin.Context) {
	var (
		username string
		password string
	)
	if in, isExist := c.GetPostForm("username"); isExist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入使用者名稱"),
		})
		return
	}
	if in, isExist := c.GetPostForm("password"); isExist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": errors.New("必須輸入密碼名稱"),
		})
		return
	}
	if err := Auth(username, password); err == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"success": "登入成功",
		})
		return
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err,
		})
		return
	}
}
