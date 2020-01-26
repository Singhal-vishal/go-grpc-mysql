package main

import (
	"grpc-go-demo/proto"
	"context"
	"google.golang.org/grpc"
	"time"
	"fmt"
)
func main(){
	conn, err := grpc.Dial("localhost:4040")
	if err!=nil{
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client := proto.NewAddServiceClient(conn)
	req := &proto.Request{
		A: 1,
		B: 2,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	//if response, err := client.GetUser(ctx,req); err==nil{
	//	fmt.Println(response.Email,response.Password)
	//}else{
	//	fmt.Println("Error occured : ", err.Error())
	//}
	//if response, err := client.DeleteUser(ctx,req); err==nil{
	//	fmt.Println(response.Response)
	//}else{
	//	fmt.Println("Error occured : ", err.Error())
	//}
	if response, err := client.Add(ctx,req); err==nil{
		fmt.Println(response.Result)
	}else{
		fmt.Println("Error occured : ", err.Error())
	}
	//if response, err := client.GetUsers(ctx,req); err==nil{
	//	fmt.Println(response.Users)
	//}else{
	//	fmt.Println("Error occured : ", err.Error())
	//}
	/*
	hardcode a and b
	 */
	//a := 10
	//b := 20
	//req := &proto.Request{A:int64(a),B:int64(b)}
	//
	//if response, err := client.Add(ctx,req); err==nil{
	//	fmt.Println(response.Result)
	//}else{
	//	fmt.Println("Error occured : ",err.Error())
	//}
	//if response, err := client.Multiply(ctx,req); err==nil{
	//	fmt.Println(response.Result)
	//}else{
	//	fmt.Println("Error occured : ",err.Error())
	//}
}
