1. FR
- Transaction services
- Account services

2. NFR
- high availability
    - active-passive db
    - async replica
- Strong consistency
    - ACID
    - idempotency key
- Scalability
    - Account sharding by geographic region
    - kafka partitioning by account ID prefix
    - cockroachDB automatic sharding
        # Capacity planning for 1M TPS:
        shards = 100  # Account shards
        tps_per_shard = 10_000  # PostgreSQL max sustained TPS
        total_capacity = shards * tps_per_shard  # 1M TPS

        # Resources:
        postgres_nodes = shards * 3 (for RF=3)  # 300 nodes
        kafka_brokers = 50  # For 1GB/s throughput
- Security
    - Data encryption
    - access control
- Durable
    - write-ahead logging
    - kafka down 
        - local disk buffering in producer
        - Dead-letter queue for undeliverable messages
- Compliance
    - Data retention
    - Audit trail
    - Transaction limit

3. Components
- Client app
- API Gateway
    - authentication and routings
- Services
    - transaction services
        - anti-fraud rules
        - logging
        - Parallel Writes with Consensus
    - account services
- Data store
    - mysql for ACID
        - write-ahead logging
        - Indexing
        - dynamic batching
    - Redis for user session

4. Design

client --> API --[kafka]--> Services --> Data store
                            [microservices]

5. Optimization
- Batching
    - pack a bundle of transactions
    - Kafka producer batching
- Caching
    - database optimization
        - redis/memcached for account balance (warm cache)
- Async Processing
    - Offload non-critical path (notifications, analytics)
    - Event-driven architecture