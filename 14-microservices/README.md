# Microservices Architecture - Scalable Systems

## Learning Objectives
- Understand microservices principles
- Learn service communication patterns
- Master service discovery and load balancing
- Implement distributed system patterns

## Topics to Practice
1. **Service Design** - Single responsibility, bounded contexts
2. **Inter-Service Communication** - HTTP, gRPC, messaging
3. **Service Discovery** - Registry patterns, health checks
4. **Load Balancing** - Distribution strategies
5. **Circuit Breaker** - Fault tolerance patterns

## Files to Create
- `user-service/` - User management microservice
- `order-service/` - Order processing microservice
- `api-gateway.go` - API gateway implementation
- `service-discovery.go` - Service registry
- `circuit-breaker.go` - Fault tolerance
- `load-balancer.go` - Load balancing logic

## Microservices Patterns
- **API Gateway** - Single entry point
- **Service Registry** - Service discovery
- **Circuit Breaker** - Prevent cascade failures
- **Bulkhead** - Isolate critical resources
- **Saga Pattern** - Distributed transactions
- **CQRS** - Command Query Responsibility Segregation

## Big Tech Microservices Stack
- **Service Mesh** - Istio, Linkerd
- **Container Orchestration** - Kubernetes, Docker Swarm
- **Message Brokers** - Kafka, RabbitMQ, NATS
- **Monitoring** - Prometheus, Grafana, Jaeger
- **Configuration** - Consul, etcd
- **Databases** - Database per service pattern

## Communication Patterns
- **Synchronous** - HTTP REST, gRPC
- **Asynchronous** - Message queues, event streaming
- **Request-Response** - Direct service calls
- **Publish-Subscribe** - Event-driven architecture
- **Event Sourcing** - Immutable event logs
