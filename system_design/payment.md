Component
1. Client
2. Api Gateway
3. Payment Services
4. Database
5. Notification Services
6. Banking network
7. Reporting services

                        3rd Party Payment validation
                            ||
                        (Payment)
client -> API Gateway -> Services -> database
                            ||
                        Notification ---> reporting

Detail breakdown
1. Client: we use SDK for easier integration
2. API Gateway: handle authentication and authorization, order validation, request and response transformation 
3. Services: Processess payment requests. 
    Fraud detection:
    Process Payment: 
        integrate third party payment gateway
        connections to card networks (Visa, Master)
            settlement, refund policy
    Notification: send notification to merchants, users
4. Database: CRUD. Create payment records. High availability with replication. 
    INNO db when down, has a log for double records so asssure not losing data when down

NFC
1. End to end encryption
2. Tokenize sensitive data
3. Microservices - JWT
4. Reliability
    Redis 
    - store session - maintain payment session state
    - Distributed Locking - prevent duplicate processing of the same transcation
5. Performance
    Kafka 
    - decouple services through publish-subscribe model. payment created, transaction processed,...
6. Scalability
    Kafka
    - scale automatically according to traffic, managing resources through partitioning- consumer and producer can keep write events
7. Compliance
