# Fixed-size Hash Map

This is a fixed-size hash map data structure implemented in Go. It uses **Linear Probe** to resolve collition. 

There are two common ways of resolving collitions in hash map: **Chaining** and **Linear Probe**. 

Below are the consideration of space and runtime efficiency when using these two different methods to handle collision:

- Space (Linear Probe wins): 
Since this is a fixed-size data structure that is not required to dynamically allocate more space, so it makes sense to fully utilize all space given. In the case of chaining method, when collition happens, all key/value pair that are hashed to the same bucket are chained in a linked list, so it is possible that some buckets are not used at all. Whereas for linear probe, whenever there is collition, the new key/value pair is placed in the next free bucket, so no pre-allocated space is wasted. Also, chaining method allocates new space every time there is collision,  so it will end up using far more space than linear probe if there are collisions.

- Runtime (Linear Probe wins): 
Due to the locality of reference, it is faster to access a series of elements in an array than to follow pointers in a linked list. So when looking up element in hash map that uses linked list chaining, the performance is worse than linear probe that looks up elements in a single array 

## Big O

Under Simple Uniform Hashing Assumption that states that a hypothetical hashing function will evenly distribute items into the slots of a hash table. The `Set` and `Get` function can be O(1)

- Set: O(1)
- Get: O(1)
- Delete : O(1)
- Load: O(1)

## Instruction

- Test
```
$go test *.go
```

To include the benchmark test
```
$go test -bench=.
```

### Usage
Example usage in /cmd folder
