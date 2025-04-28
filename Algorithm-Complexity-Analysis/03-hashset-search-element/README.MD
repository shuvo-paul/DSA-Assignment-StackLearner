## Input - Output
| Hash Set Contains  | Search Element | Output |
|------------ | ------- | --------- |
|["banana", "papaya", "gelato", "poopaye"] | "gelato" | true
|["banana", "papaya", "gelato", "poopaye"] | "Bello" | false

## Step by Step Cost
1. Get the index of the value from hash function. // Time: O(1), Space: O(1)
2. Check if a bucket exist at the index of the hash table. // Time: O(1), Space: O(1)
3. If the bucket exists, traverse through the nodes and look for the element. // Time: O(n), Space: O(1)

## Final asymptotic complexities
| Case | Time | Space | Reason |
| ----- | ---- | ----- | ---- |
| Best | O(1) | O(1) | No collision when hashing.
| Average | O(1) | O(1) | Few collisions
| Worst | O(n) | O(1) | All elements end up in the same bucket.