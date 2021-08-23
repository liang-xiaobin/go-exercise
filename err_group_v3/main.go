package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello world")
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	g.Go(func() error {
		fmt.Println("http")
		go func() {
			r := <-ctx.Done()
			fmt.Println(ctx.Err().Error())
			fmt.Printf("http ctx done (%v)\n", r)
			err := srv.Shutdown(context.TODO())
			if err != nil {
				fmt.Printf("srv shutdown err is (%v)", err)
			}
		}()
		return srv.ListenAndServe()
	})
	g.Go(func() error {
		exitSingnals := []os.Signal{os.Interrupt, syscall.SIGALRM, syscall.SIGQUIT, syscall.SIGINT}
		sig := make(chan os.Signal, len(exitSingnals))
		signal.Notify(sig, exitSingnals...)
		for {
			fmt.Println("signal")
			select {
			case <-ctx.Done():
				fmt.Println("signal ctx done")
				fmt.Println(ctx.Err().Error())
				return ctx.Err()
			case <-sig:
				return nil
			}
		}
	})
	g.Go(func() error {
		fmt.Println("inject")
		time.Sleep(time.Second)
		fmt.Println("inject finish")
		return errors.New("sdfasadf")
	})
	err := g.Wait()
	fmt.Println(err)
}
