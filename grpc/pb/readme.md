## 安装依赖
```shell
# Go语言插件，用于生成.pb.go，其中包含proto文件中定义的类型以及序列化方法
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# Grpc插件，用于生成_grpc.pb.go，其中包含接口类型（供客户端调用的服务方法）以及服务器要实现的接口类型
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
注意需要将`$GOPATH/bin`加入环境变量中

## 编写pb文件
```proto
syntax = "proto3" // 版本声明

option go_package = "xx";  // 指定生成的Go代码在你项目中的导入路径

package pb; // 包名

// 定义服务
service Greeter {
    // SayHello 方法
    rpc SayHello (HelloRequest) returns (HelloResponse) {}
}

// 请求消息
message HelloRequest {
    string name = 1;
}

// 响应消息
message HelloResponse {
    string reply = 1;
}
```
