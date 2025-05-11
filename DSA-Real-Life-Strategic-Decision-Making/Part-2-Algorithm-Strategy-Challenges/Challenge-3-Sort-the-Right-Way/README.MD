# Insertion Sort is best suited for nearly sorted arrays.

## Justification:
 - Insertion sort takes an element and inserts it into the correct position in the array. So, if the array is nearly sorted, only a few elements need to be shifted. In this case the Time complexity is O(n), and the Space complexity is O(1).
 - Selection sort and Bubble sort take O(n^2) time even for nearly sorted arrays. Although, Bubble sort can be optimized to O(n) time using and early exit condition.