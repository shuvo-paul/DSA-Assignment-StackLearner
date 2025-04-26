## Input - Output
Input: 9 -> 8 -> 7 -> 6  
Value to Delete: 8
Output: 9 -> 7 -> 6

## Step by Step Cost
1. Traverse the list until the value found. // Time: O(n), Space: O(1)
2. If the value found change the next node of previous node to the next node of the matched node. // Time: O(1), Space: O(1) 
3. If the matched Node is head change the head to the next node. //Time: O(1), Space: O(1)
4. If the matched Node is tail, change the tail to the previous node of matched node. //Time: O(1), Space: O(1)

## Final asymptotic complexities
Time: O(1)
Space: O(1)

## Diagram
![diagram](./linkedlist-delete.png)