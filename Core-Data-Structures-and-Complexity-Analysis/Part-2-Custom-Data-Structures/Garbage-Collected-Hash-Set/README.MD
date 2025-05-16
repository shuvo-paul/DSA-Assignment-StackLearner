# Garbage-Collected Hash Set:
The HashSet implemented using an array of linked list. Each value is associated with a TTL expiration. The `CleanUpWorker` periodically removes expired entries. The TTL is updated on access.

## Complexity
| Operation   | Time | Space |
|-------------|------|-------|
|Insert       | O(1) | O(1)  |
|Remove       | O(1) | O(1)  |
|Find         | O(1) | O(1)  |
|Cleanup      | O(n) | O(1)  |
|CleanUpWorker| O(n) | O(1)  |
