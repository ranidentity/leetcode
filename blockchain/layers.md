- Cross layer interactions
    - simulating interactions between Layer 1 (L1) and Layer 2 (L2)—like bridge contracts or cross-chain messaging
    - Real bridges: complex deployment, expensive pay gas both side, waiting for cross-chain transaction (time)
    - if test with vm.Mockcall: instant test, avoid complexity

- Real-World Use Cases
    - Bridging Assets: mock a deposit from L1->L2 contract
    - Cross-chain messaging
    - Oracle update: mock a price update from L1 -> L2

- For L1 to read on L2 message
    - requires some form of finalization window, fraud proof, or validity proof before L1 can act on L2's data.
    - Write a bridge contract on L1 and L2, Emit events or send messages from L2, Prove and execute messages on L1 after finalization
    - slow and allow time for fraud detection

- When do you need L1 deployment:
    - Bridging, to receive message from L2
    - Finality, outcome need to be on L1 ( withdraw ETH, mint NFT)
    - Cross layer messaging
    - Shared security - You rely on L1 for dispute resolution or ZK proof verification


LAYER 2
Lightning Network, Optimistic Rollups, and ZK-Rollups are all layer-2 scaling