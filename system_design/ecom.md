Scope - storefront
Scale - how many users?
Geographic distribution - global / regional ? use CDN
read vs write heavy -
Latency Requirement
Consistency requirement
Feature to prioritize - search, product catalog, cart, checkout, order history, authentication

FR
- user management
- product catalog
- search
- cart
- order
- review/rating
- recommendation

NFR
- high availability
- Scalability
- Latency
- Security
- Consistency
- Fault tolerent
- Maintainability


Diagram
- Storefront:

        
Client -------> API Gateway ------> Services --> database


- Payment services [refer_to_payment.md]

Components:
- Data store
    - rdbms for structured data
        - users
        - inventory [strong consistency]
        - order
    - redis
        - user session data
        - cart if no need to be permenant
    - mongodb 
        - product catalog for flexible schema ( or elastic search)
    - Image storages - cloud storages
    - Search index: elastic search

Infrastructure
- Load balancer distribute traffic
- CDN
    - cache static contents
- Message Queue
    - asynchronous tasks: order+ inventory changes
- Caching
    - hot data
- logging

Scalability
- Horizontal scaling
    - load balancer
    - add new instance
    - db sharding, partitioning
- Vertical scaling  