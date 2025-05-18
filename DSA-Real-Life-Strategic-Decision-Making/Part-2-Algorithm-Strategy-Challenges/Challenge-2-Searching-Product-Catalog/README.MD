# Searching in Product Catalog
Binary Search is better suited than an Interpolation Search here.

## Why Binary Search over Interpolation Search?
- Interpolation Search performs better for uniformly distributed data. The data here is non-uniformly distributed.
- For sorted data, Binary Search runs in O(log n) time in the worst case, whereas Interpolation Search may degenerate to something similar to Linear Search, if the data is not uniformly distributed.