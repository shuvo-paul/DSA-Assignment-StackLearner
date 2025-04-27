# Part 1: Algorithm Complexity Analysis (30 Points)

Analyze the time and space complexity for each of the following algorithms. For every algorithm:

- Provide input-output mapping.
- Give step-by-step cost breakdown.
- Determine final asymptotic complexities (time and space).
- Suggest optimizations, if any.

## Algorithms:

1. Insert at the beginning of a dynamic array.
2. Insert at the end of a linked list.
3. Search for an element in a hash set.
4. Rehash a hash table after crossing load factor.
5. Delete a node from a singly linked list by value.
6. Check if an array contains all unique values.
7. Count common elements in two hash sets.
8. Convert an array into a linked list.
9. Clone a hash table with chaining.
10. Compare array vs. hash set lookup performance.

# Part 2: Custom Data Structures (30 Points)

Choose and implement any two of the following custom data structure options. Each implementation must include:

- Clean and reusable API.
- Design explanation with complexity.
- Real-world applicability and trade-offs.
- Inline documentation and README with diagram if applicable.

## Options:

1. Time-Aware Linked List:
   - Each node stores timestamp of insertion.
   - Add methods to retrieve nodes inserted within last n seconds.

2. Secure Hash Table:
   - Prime bucket sizing with auto-resizing.
   - Ordered key tracking.
   - Optional TTL (time-to-live) support for entries.

3. Garbage-Collected Hash Set:
   - Automatically remove entries not used for a configurable time.
   - O(1) average lookup and deletion.
   - Manual clean-up support.

4. Event-Driven Linked List:
   - Listens to insert/update/delete actions.
   - Supports registering external listeners with callbacks.

# Part 3: Real-World Problem Solving (20 Points)

Choose any two of the following problems and solve them using the best possible data structure(s). Your solution must:

- Include a code implementation.
- Explain your data structure decision.
- Provide time and space analysis.
- Optionally include supporting diagrams.

## Scenarios:

1. Inventory Lookup System:
   - Maintain a product catalog.
   - Allow quick availability checks and frequent insertions/removals.
   - Ensure scalability as inventory grows.

2. Recent Posts Feed:
   - Maintain the 10 most recent user posts.
   - Support constant-time insertion/removal.
   - Preserve chronological order.

3. Duplicate Detector in Stream:
   - Track stream of user actions.
   - Identify duplicates in O(1) average time.
   - Bonus: Support time-window filtering (e.g., ignore duplicates within 5 minutes).

# Part 4: Reflection & Documentation (10 Points)

Prepare a README.md file summarizing your assignment experience.

- List chosen tasks and their rationale.
- Reflect on challenges and lessons learned.
- Discuss trade-offs and complexity decisions.
- Include system or memory layout diagrams as needed.
- Optional: Add a short (max 2-minute) video walkthrough.
