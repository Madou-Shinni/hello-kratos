# proto文件输出swagger文档

官方文档：

- https://grpc-ecosystem.github.io/grpc-gateway/
- https://github.com/grpc-ecosystem/grpc-gateway

## 1. 安装工具

```shell
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## 2.复制目录

复制以下目录文件到项目中

[protoc-gen-openapiv2](..%2Fthird_party%2Fprotoc-gen-openapiv2)

## 3.添加proto选项

在proto文件中导入

```shell
import "protoc-gen-openapiv2/options/annotations.proto";
```

在需要生成swagger文档的服务接口上添加选项

```shell
// Sends a greeting
rpc SayHello (HelloRequest) returns (HelloReply) {
option (google.api.http) = {
  get: "/helloworld/{name}"
};

option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
  summary: "helloword";
  tags: ["hello-kratos"];
};
}
```


## 4.生成swagger文档

适当修改Makefile文件

在make api选项中添加以下命令

```shell
--openapiv2_out . \
--openapiv2_opt output_format=yaml \
--openapiv2_opt allow_merge=true,merge_file_name=openapiv2 \
```

```shell
在根目录下执行 make api
```

## 5.查看swagger文档

在根目录下生成openapiv2.swagger.yaml文件，将其复制到swagger编辑器中查看