## Input - Output
| Input  | Output |
|------- | -------|
|["a", "b", "c", "d"] | true |
|["a", "b", "c", "d", "b"] | false |

## Step by Step Cost
1. Loop over each element. // Time: O(n), Space: O(1)
2. Create a hash set. // Time: O(1), Space: O(n)
3. Check if the value exists in the hash set. If not, add it to the hash set. // Time: O(1), Space: O(n)

## Final asymptotic complexities
Time: O(n)
Space: O(n)