package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
	"net/http"
	"time"
)

func main() {
	go server()

	cb := gobreaker.NewCircuitBreaker(
		gobreaker.Settings{
			Name:        "my-circuit-breaker",
			MaxRequests: 3,
			Timeout:     3 * time.Second,
			Interval:    1 * time.Second,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				return counts.ConsecutiveFailures > 6
			},
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				fmt.Printf("CircuitBreaker '%s' changed from '%s' to '%s'\n", name, from, to)
			},
		})
	fmt.Println("Call with circuit breaker")
	for i := 0; i < 100; i++ {
		_, err := cb.Execute(func() (interface{}, error) {
			err := DoReq()
			return nil, err
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

var startTime time.Time = time.Now()

func server() {
	e := gin.Default()
	e.GET("/ping", func(ctx *gin.Context) {
		if time.Since(startTime) < 3*time.Second {
			ctx.String(http.StatusInternalServerError, "pong")
			return
		}
		ctx.String(http.StatusOK, "pong")
	})

	fmt.Printf("Starting server at port 8081\n")
	e.Run(":8081")
}

func DoReq() error {
	resp, err := http.Get("http://localhost:8081/ping")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		return errors.New("bad response")
	}

	return nil
}
