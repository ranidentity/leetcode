process
Fan-Out on Write (for active users):
Post Service → Kafka → Feed Service → Redis/Cassandra.

Fan-Out on Read (for influencers/celebrities):
Feed Service queries Posts table, merges with cached results.

special user
influencers
active poster