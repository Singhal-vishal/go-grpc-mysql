package main

import (
	"fmt"
	"google.golang.org/grpc"
	"../proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func main()  {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)

	/*
	we want to provide end point to our api so that we can provide end points
	to it or we can use it through browser. So we are using "gin".
	 */

	/*
	client using gin_server
	 */
	gin_server := gin.Default()

	gin_server.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"),10,64)
		if err !=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{"error" : "Invalid Parameter a"})
		}

		b, err := strconv.ParseUint(ctx.Param("b"),10,64)
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid Parameter b"})
		}

		req := &proto.Request{A : int64(a), B: int64(b)}
		
		if response, err := client.Add(ctx, req); err == nil{
			ctx.JSON(http.StatusOK, gin.H{
				"result" : fmt.Sprint(response.Result),
			})
		} else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		}
	})

	gin_server.GET("/multi/:a/:b", func(ctx *gin.Context){
		a, err := strconv.ParseUint(ctx.Param("a"),10,64)
		if err!=nil {
			ctx.JSON(http.StatusBadRequest,gin.H{"error": "Invalid parameter a"})
		}

		b, err := strconv.ParseUint(ctx.Param("b"),10,64)
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid parameter b"})
		}

		req := &proto.Request{A : int64(a), B : int64(b)}

		if response, err := client.Multiply(ctx, req); err == nil{
			ctx.JSON(http.StatusOK, gin.H{
				"result" : fmt.Sprint(response.Result),
			})
		}else{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error": err.Error()})
		}
	})

	if err := gin_server.Run(":8080"); err!=nil {
		log.Fatal("Failed to run server: %v", err)
	}
}
