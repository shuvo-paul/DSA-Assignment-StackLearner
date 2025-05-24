## Explain one scenario where Selection Sort is a better choice over Insertion Sort.
Insertion sort has a time complexity of O(n^2) in the worst case, but can perform better on partially sorted arrays, with a best-case time complexity of O(n).
Selection sort has a time complexity of O(n^2) in all cases.
Even if the selection sort has complexity of O(n2) in all cases,  it takes a minimum swap unlike insertion sort. So, where memory is limited and writing memory is more expensive, selection sort is better instead of insertion sort.

## Why is Bubble Sort considered impractical for large datasets? Give a case where it may outperform Merge Sort.
Bubble Sort algorithm is not suitable for large data sets as its average and worst-case time complexity are quite high. Merge Sort is a divide-and-conquer algorithm and the time complexity is O(n log n) in all cases.


## Can you use Binary Search on an infinite data stream? Why or why not?
For an infinite datastream it’s not possible to divide by 2 and find the middle point. Without first and last indices it’s not possible to run binary search

## Would you choose a Stack (Array) or Stack (Linked List) for browser tab history? Justify.
I would choose Array. In a tab history a user might jump to the middle or their history. Array allows random access to any element via indexing. Where a linked list requires O(n) traversal time.
