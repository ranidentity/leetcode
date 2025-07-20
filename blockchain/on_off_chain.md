What’s the difference between on-chain and off-chain computation?

What are the trade-offs between putting data on-chain vs off-chain?

When would you choose off-chain over on-chain?

How would you design a decentralized application with off-chain components?

- Design Question
“You’re building a voting system for a DAO. Which parts should be on-chain? Which parts can be off-chain?”

“How do you store large files like images or documents in a blockchain app?”
- Metadata and images are often stored off-chain, specifically on IPFS(decentralized storage system that allows for the storage and retrieval of files), blockchain just store reference to the off-chain data   
- When NFT is created, smart contract store link to the image

“How would you verify off-chain data inside a smart contract?”
- Gas & optimization