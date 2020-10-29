package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"

	pb "micro/micro"
)

//import proto生成的类

type Greeter struct {
}

/*
实现proto生成的hello.micro.go中的
type GreeterHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error}
*/
func (g *Greeter) Hello(ctx context.Context, rep *pb.HelloRequest, rsp *pb.HelloResponse) error {
	rsp.Greeting = "Hello" + rep.Name
	//fmt.Printf()
	return nil
}

func main() {
	//新建一个服务
	service := micro.NewService(micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{"type": "hello world"}))
	service.Init()                                                   //初始化服务
	err := pb.RegisterGreeterHandler(service.Server(), new(Greeter)) //注册服务
	if err != nil {
		fmt.Println("注册服务出现了问题...", err)
		return
	}

	//运行服务
	if err := service.Run(); err != nil {
		fmt.Println("服务运行出现了错误：", err)
	}
}
