package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

// func test(w http.ResponseWriter, r *http.Request) {
// 	tmpl := template.Must(template.ParseFiles("./index.html"))
// 	data := new(IndexData)
// 	data.Title = "Home Page"
// 	data.Content = "It's my first page 啊！！！！"
// 	tmpl.Execute(w, data)
// }

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "Check it by now"
	data.Content = "My first gin project !!"
	c.HTML(http.StatusOK, "index.html", data)
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")
	server.Static("/assets", "./template/assets")
	server.GET("/", test)
	server.GET("/login", LoginPage)
	server.POST("/login", LoginAuth)
	server.Run(":9999")
	// http.HandleFunc("/", test)
	// http.HandleFunc("/home", test)
	// err := http.ListenAndServe(":8888", nil)
	// if err != nil {
	// 	log.Fatal("Err msg:", err)
	// }
}
