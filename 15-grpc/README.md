# gRPC - High-Performance RPC Framework

## Learning Objectives
- Understand gRPC and Protocol Buffers
- Learn service definition and code generation
- Master streaming and error handling
- Implement production-ready gRPC services

## Topics to Practice
1. **Protocol Buffers** - Schema definition language
2. **Service Definition** - RPC method declarations
3. **Code Generation** - protoc compiler usage
4. **Streaming** - Unary, server, client, bidirectional
5. **Interceptors** - Middleware for gRPC

## Files to Create
- `proto/user.proto` - Protocol buffer definitions
- `server.go` - gRPC server implementation
- `client.go` - gRPC client implementation
- `streaming.go` - Streaming examples
- `interceptors.go` - Authentication, logging middleware
- `gateway.go` - gRPC-Gateway for REST compatibility

## gRPC vs REST
- **Performance** - Binary protocol, HTTP/2
- **Type Safety** - Strong typing with protobuf
- **Streaming** - Real-time bidirectional communication
- **Code Generation** - Auto-generated client/server code
- **Language Support** - Multi-language compatibility

## Big Tech gRPC Usage
- **Google** - Internal service communication
- **Netflix** - Microservice communication
- **Uber** - Real-time location services
- **Square** - Payment processing
- **Dropbox** - File synchronization
- **CoreOS** - etcd distributed key-value store

## Production Features
- **Load Balancing** - Client-side and server-side
- **Health Checking** - Service health monitoring
- **Reflection** - Dynamic service discovery
- **Compression** - Gzip compression support
- **TLS** - Secure communication
- **Deadlines** - Request timeout handling

## Tools and Libraries
- **protoc** - Protocol buffer compiler
- **grpc-go** - Go gRPC implementation
- **grpc-gateway** - REST to gRPC proxy
- **buf** - Modern protobuf toolchain
