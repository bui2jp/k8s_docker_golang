package main

import (
	"net"
	//"os"
	"fmt"
	"github.com/gin-gonic/gin"
  )

func main() {

	addrs, _ := net.InterfaceAddrs()

	myip := "8.8.8.8"

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
		  if ipnet.IP.To4() != nil {
			fmt.Println(ipnet.IP.String())
			myip = ipnet.IP.String()
		  }
		}
	  }



	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {



		c.JSON(200, gin.H{
			"message": "my pong",
			"myip" : myip,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}