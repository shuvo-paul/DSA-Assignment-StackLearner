# Recent Posts Feed
Since the feed need to keep the 10 most recent ones, a fixed-size array is a good choice. Arrays let us insert, delete, and get data in O(1) time on average.

## Complexity
| Operation       | Time | Space |
|-----------------|------|-------|
|AddPost          | O(1) | O(1)  |
|GetRecentPosts   | O(1) | O(1)  |