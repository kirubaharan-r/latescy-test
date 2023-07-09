package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/ping/", test)
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "working")
	})
	r.Run(":8084")
}
func test(c *gin.Context) {
	a := time.Now()
	url := c.PostForm("url")
	htt := "https://" + url
	response, err := http.Get(htt)

	if err != nil {

		c.String(http.StatusOK, "site was down")
	}

	time.Sleep(time.Duration(100) * time.Millisecond)
	defer response.Body.Close()
	c.JSON(http.StatusOK, gin.H{
		"(may)IP address": c.RemoteIP(),
		"*IP address":     c.ClientIP(),
		"URL":             fmt.Sprintln(response.StatusCode, response.Request.URL),
		"time":            fmt.Sprintln(time.Since(a)),
	})

}
