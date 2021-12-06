package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func readDb(){
	time.Sleep(10*time.Second) // 模拟超时情况

}

func helloWorld(c *gin.Context) {
	res := gin.H{
		"message": "hello world!",
	}
	time.Sleep(10*time.Second) // 模拟超时情况
	readDb()
	c.JSON(http.StatusOK, res)
}

func main() {
	r := gin.Default()
	r.GET("/", helloWorld)

	srv := http.Server{Addr: "127.0.0.1:8000", Handler: r}

	go func() {
		if err:= srv.ListenAndServe(); err!=nil && err != http.ErrServerClosed{
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<- quit

	log.Println("Shut down server")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 收到quit信号后，两秒未正常关闭，则抛出异常
	if err:= srv.Shutdown(ctx); err != nil{
		log.Fatalf("Shut down err: %s\n", err)
	}

	log.Println("Server closed!")
}
