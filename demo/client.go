package main

import (
	"rpc-go/demo/common"
	"fmt"
	"net/rpc"
)

func main() {
	var args = common.Args{A: 32, B: 14}
	var result = common.Result{}

	var client, err = rpc.DialHTTP("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("connect rpc server failed, err:%v", err)
	}

	err = client.Call("MathService.Add", args, &result)
	if err != nil {
		fmt.Printf("call math service failed, err:%v", err)
	}
	fmt.Printf("call RPC server success, result:%f", result.Value)
}