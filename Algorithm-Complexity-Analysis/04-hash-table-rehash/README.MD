## Input - Output
| Operation  | Input | Output |
|------------ | ------- | --------- |
| Initial State | table with 4 items | Capacity: 11 (default) |
| Resize to 5 | resize(5) | Capacity: 5 |
| Insert two more items | Set() | Capacity: 10 |

## Step by Step Cost
1. Store all elements to a variable. // Time: O(n), Space: O(1)
2. Allocate the hash table with the new capacity. // Time: O(1), Space: O(n)
3. Go through all elements in hash table. // Time: O(1), Space: O(n)
4. Rehash each element. // Time: O(1), Space: O(1)
5. Insert each element into the new hash table. // Time: O(1), Space: O(1)

## Final asymptotic complexities
| Case | Time | Space |
| ----- | ---- | ----- |
| Best | O(n) | O(n) |
| Average | O(n) | O(n) |
| Worst | O(n) | O(n) |