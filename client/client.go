package client

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"gitlab.com/ro-tex/grpc/proto"
)

func addHandler(addService proto.AddServiceClient) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		a, err := strconv.ParseInt(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})
			return
		}

		b, err := strconv.ParseInt(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})
			return
		}

		req := &proto.Request{A: a, B: b}

		if res, err := addService.Add(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(res.Result)})
		}
	}
}

func multHandler(addService proto.AddServiceClient) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		a, err := strconv.ParseInt(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})
			return
		}

		b, err := strconv.ParseInt(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameter"})
			return
		}

		req := &proto.Request{A: a, B: b}
		if res, err := addService.Multiply(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(res.Result)})
		}
	}
}

func Run() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	addService := proto.NewAddServiceClient(conn)

	g := gin.Default()

	/**
	One way to define handlers is directly here. It's a bit messy but it's fine for simple handlers.
	*/
	g.GET("/about", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"about": "Send a GET request to `/add/:a/:b` and you'll get their sum."})
	})

	/**
	Another way is to have a generator function that returns a func(ctx *gin.Context).
	This allows us to move the handler implementation somewhere else and have this part nice and clean.
	*/
	g.GET("/add/:a/:b", addHandler(addService))
	g.GET("/mult/:a/:b", multHandler(addService))

	if err := g.Run(":4041"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
