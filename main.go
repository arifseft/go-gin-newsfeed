package main

import (
	"github.com/arifseft/go-gin-crud/handler"
	"github.com/arifseft/go-gin-crud/platform/newsfeed"
	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run()

	// feed := newsfeed.New()
	//
	// fmt.Println(feed)
	//
	// feed.Add(newsfeed.Item("Hello", "How are you doing?"))
}
