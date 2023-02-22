package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	proto "github.com/rodolfoviolla/go-rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getParam(ctx *gin.Context, param string) (value uint64) {
	value, err := strconv.ParseUint(ctx.Param(param), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter " + param})
	}
	return
}

func makeRequest(ctx *gin.Context, requestFn func(ctx context.Context, in *proto.Request, opts ...grpc.CallOption) (*proto.Response, error)) {
	a := getParam(ctx, "a")
	b := getParam(ctx, "b")
	req := &proto.Request{A: int64(a), B: int64(b)}
	if response, err := requestFn(ctx, req); err == nil {
		ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := proto.NewAddServiceClient(conn)
	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		makeRequest(ctx, client.Add)
	})
	g.GET("/mult/:a/:b", func(ctx *gin.Context) {
		makeRequest(ctx, client.Multiply)
	})
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}