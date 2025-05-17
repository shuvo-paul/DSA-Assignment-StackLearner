# Array lookup

## Input - Output
| Initial State | Input |  Output |
| ------- | --------- |--------- |
| ["banana", "papaya", "gelato", "poopaye"] | "poopaye" | true |
| ["banana", "papaya", "gelato", "poopaye"] | "potato" | false |

## Step by Step Cost
1. Iterate over each elements in the array. // Time: O(n), Space: O(1)
2. If the current element matches the input, return true. If no match found return false. // Time: O(1), Space: O(1)

## Final asymptotic complexities
| Case | Time | Space |
| ----- | ---- | ----- |
| Best | O(1) | O(1) |
| Average | O(n) | O(1) |
| Worst | O(n) | O(1) |

# Hash Set lookup

## Input - Output
| Initial State | Input |  Output |
| ------- | --------- |--------- |
| {"banana", "papaya", "gelato", "poopaye"} | "poopaye" | true |
| {"banana", "papaya", "gelato", "poopaye"} | "potato" | false |

## Step by Step Cost
1. Get the index of the value from hash method. // Time: O(1), Space: O(1)
2. Return ture if the value exist in the indexed bucket. // Time: O(1), Space: O(1)

## Final asymptotic complexities
| Case | Time | Space |
| ----- | ---- | ----- |
| Best | O(1) | O(1) |
| Average | O(1) | O(1) |
| Worst | O(1) | O(1) |