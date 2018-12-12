[![Build Status](https://travis-ci.com/unchartedsky/bodyguard.svg?token=265AAhxHviCMV4xC9qkK&branch=master)](https://travis-ci.com/unchartedsky/bodyguard)

# bodyguard

## How to test

[`grpc_cli`](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md) should be helpful

``` bash
brew install grpc
```

``` bash
grpc_cli ls localhost:50051 -l

grpc_cli call localhost:50051 SayHello "name: 'gRPC CLI'"
```

## See also

- [Getting started with Go and Protocol Buffers - Seven Story Rabbit Hole](https://tleyden.github.io/blog/2014/12/02/getting-started-with-go-and-protocol-buffers/)
