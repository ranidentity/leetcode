1. key consideration
- Huge data streaming in for processing
- conversion rate = click / impression. how many times showed 

2. functional requirements
- aggregrate metric of a given ads ID in X minutes
- Top conversion rate in X duration

3. non-functional requirements
- scalability
- high throughput
- data integrity
- fault tolerent
- latency

4. Diagram
                                            query
                                            |
                                            |
 data collection - > realtime processing -> data storage

what data collected?
1. ads ID
2. user ID
3. timestamp
4. ip_address


[data collection] Consider the storage size
cannot use rdms - too high QPS? query per seconds
Use NoSQL - high throughput. But complex
redis - good but need in-disc, asynchronous go into disc
message queue - kafka store data longer than memory. Introduce latency because queu
direct log file - since message queue need to introduce queue.

[realtime processing]
1. way
batch processing - latency
mini batch - 
streaming - reduce latency, increase complexity
2. need checkpoint to catch breakpoint/downpoint
kafka easier
3. hot spot consideration

[data storage]
1. aggregated data storage. ads id, impression count, demography...
(consider each row how big data)
2. rdbms - partitioning
3. + cache