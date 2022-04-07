package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"time"
)

var listenAddr = flag.String("l", ":8999", "listen address")

func init() {
	flag.Parse()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Any("/app/:rname", showCurrentMachine)
	router.GET("/healthz", showHealthz)
	router.Run(*listenAddr)
}

func showCurrentMachine(ctx *gin.Context) {
	hName, _ := os.Hostname()
	rname := ctx.Param("rname")
	ctx.JSON(200, gin.H{
		"message":   "Hello - " + rname,
		"hostname":  hName,
		"timestamp": time.Now().UTC().Unix(),
		"ipAddrs":   getIPAddrs(),
	})
	return
}

func getIPAddrs() []string {
	addrs, _ := net.InterfaceAddrs()
	var ipAddrs []string
	for _, addr := range addrs {
		ipAddrs = append(ipAddrs, addr.String())
	}
	return ipAddrs
}

func showHealthz(ctx *gin.Context) {
	ctx.String(200, "OK")
	return
}
