package main

import (
	"fmt"

	pb "micro/micro" //import proto生成的类

	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

func main() {
	service := micro.NewService(micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{"type": "hello world"}))
	service.Init()
	greeter := pb.NewGreeterService("greeter", service.Client()) //调用proto生成的hello.micro.go中的NewGreeterService方法
	res := pb.HelloRequest{}
	res.Name = "lalalalalalalalalalalalalala"
	rsp, err := greeter.Hello(context.TODO(), &res) //Client API for Greeter service
	if err != nil {
		fmt.Println("请求服务出现了问题...", err)
		return
	}
	fmt.Println("服务返回的结果为：", rsp.Greeting)

}
