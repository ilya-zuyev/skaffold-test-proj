package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/gin-gonic/gin"
)

var label = flag.String("l", "", "label")

func main() {
	flag.Parse()
	port, found := os.LookupEnv("PORT")
	if !found {
		port = "8080"
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"me":    os.Getuid(),
			"ver":   5,
			"label": *label,
		})

	})
	fmt.Println("zzzzm!!")

	r.Run(net.JoinHostPort("0.0.0.0", port))
}
