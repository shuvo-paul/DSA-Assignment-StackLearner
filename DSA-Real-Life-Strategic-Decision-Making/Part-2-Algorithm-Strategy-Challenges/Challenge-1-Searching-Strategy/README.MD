# Searching Strategy
A hash map is better suited than Binary Search or Interpolation Search here.

## Why Hash Map over Binary Search, Interpolation Search?
 - Patients arrive randomly, so insertions are random and patient IDs are not sequential. A hash map can insert in O(1) time.
 - With a hash map, doctors can look up a patient in O(1) time.