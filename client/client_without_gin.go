package main

import (
	"fmt"
	"google.golang.org/grpc"
	"../proto"
	"time"
	"context"
)
func main(){
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client := proto.NewAddServiceClient(conn)

	/*
	hardcode a and b
	 */
	a := 10
	b := 20
	req := &proto.Request{A:int64(a),B:int64(b)}

	if response, err := client.Add(ctx,req); err==nil{
		fmt.Println(response.Result)
	}else{
		fmt.Println("Error occured : ",err.Error())
	}
	if response, err := client.Multiply(ctx,req); err==nil{
		fmt.Println(response.Result)
	}else{
		fmt.Println("Error occured : ",err.Error())
	}
}
