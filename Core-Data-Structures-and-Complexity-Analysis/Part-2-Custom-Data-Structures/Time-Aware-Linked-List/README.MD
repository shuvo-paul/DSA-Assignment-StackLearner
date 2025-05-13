# Time aware Linked List:
A linkedlist whose nodes store the time they were inserted. Each new node is added at the head. This makes it easy to retrieve nodes added in the last n time by traversing from the most recently inserted nodes.

## Complexity
| Operation | Time | Space |
|-----------|------|-------|
|Insert     | O(1) | O(1)  |
|Remove     | O(n) | O(1)  |
|RetriveRecent | O(n) | O(1)  |
