Requirement
- Functional
    - shorten url
    - redirect url
    - custom url
    - expiration
    - analytics
    - user account
- Non-Functional
    - High availability
    - latency
    - unpredictability
    - cost
- Special note:
    - capacity estimation
    - read to write ratio
        - high reading rate
        - hot data to cache
            - calculate their size
    - data retention
    - average record size:
        - short url ~14 bytes
        - long url (256) ~512 bytes
        - others data ~50 bytes
API - 
- Create short url
- Get short url

Component - 
1. Client
2. Api Gateway
3. Services
    - Shortener services
    - Key generation services
        - what logic can be used?
    - query
    - redirect services 
    - analytic services
        - kafka for buffer click events
4. Database
    -  RDBMS: ACID, consistency, simple
    - 
5. redis
    - policy recently used

Diagram

client --> api gateway --> services --> db

Improvement
- load balancer 
    - client to API gateway
        - can be internal load balancer
        - introduce latency because another layer
    - api gateway to services
        - suitable for microservices
        - internal load balancer
            - nginx
            - orchestration tool like k8s to scale
- dealing with hot data? redis
- db sharding: for scalability. Common key ID
- db replica for tolerance