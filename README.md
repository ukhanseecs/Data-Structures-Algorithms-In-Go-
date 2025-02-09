# Data Structures & Algorithms in Go
This repository provides solutions for the [Data Structures & Algorithms In Go](https://www.educative.io/courses/data-structures-and-algorithms-go) course.

## Table of Contents
- [Arrays](#arrays)
- [Recursive Functions](#recursive-functions)
- [Sorting](#sorting)
- [Stacks](#stacks)
- [Queues](#queues) 
- [Trees](#trees)
- [Binary Search Trees](#binary-search-trees)

## [Arrays](./arrays)
- [**Sum Array**](./arrays/sum.go): Calculate sum of array elements
- [**Sequential Search**](./arrays/search.go): Linear search implementation
- [**Binary Search**](./arrays/binary_search.go): Search in sorted array
- [**Largest Sum Subarray**](./arrays/max_sum_sub_array.go): Find subarray with maximum sum (Kadane's Algorithm)
- [**Rotating an Array**](./arrays/rotate.go): Rotate array by k positions
- [**Array Waveform**](./arrays/swap.go): Rearrange array in wave pattern
- [**Index Array**](./arrays/index.go): Map elements to corresponding indices
- [**Sort 1 to N**](./arrays/sort1toN.go): Sort array containing numbers from 1 to N
- [**Smallest Positive Missing**](./arrays/smallest_missing_number.go): Find smallest missing positive integer
- [**Max Min Array**](./arrays/max_min.go): Rearrange array in max/min pattern
- [**Array Index Max Difference**](./arrays/index_max_diff.go): Find maximum index difference satisfying given conditions
- [**Tower of Hanoi**](./arrays/tower_hanoi.go): Solve Tower of Hanoi puzzle

## [Recursive Functions](./recursive.go)
All recursive functions are implemented in a single file:
- **Factorial**: Calculate factorial recursively
- **Print Base 16**: Convert and print integers in hexadecimal
- **GCD**: Greatest Common Divisor using Euclidean algorithm
- **Fibonacci**: Generate Fibonacci numbers
- **Permutation**: Generate all permutations

## [Sorting](./sorting.go)
All sorting algorithms are implemented in a single file:
- **Partition 0/1**: Segregate 0s and 1s
- **Partition 0/1/2**: Dutch National Flag problem

## [Stacks](./stack)
### [Array Implementation](./stack/stack_using_array.go)
- **Basic Operations**: Push, Pop, Top, IsEmpty, Length
### Stack Problems:
- [**Sorted Insert**](./stack/sorted_insert.go)
- [**Sort Stack**](./stack/sort.go)
- [**Bottom Insert**](./stack/bottom_insert.go)
- [**Reverse Stack**](./stack/reverse.go)

### [Linked List Implementation](./stack/stack_using_ll.go)
- **Basic Operations**: Push, Pop, Peek, IsEmpty, Size
- **Print**: Display stack contents

## [Queues](./queue)
- [**Queue using Stacks**](./queue/queue_using_stack.go): Implementation using two stacks
  - Add, Remove operations
  - Length, IsEmpty checks

## [Trees](./tree.go)
All tree operations are implemented in a single file:
- **Binary Tree Construction**: Level order construction
- **Tree Properties and Operations**
- **Various Traversals**
- **Special Operations**

## [Binary Search Trees](./binary_search_trees)
- [**Tree Structure**](./binary_search_trees/tree.go): Basic BST structure and creation
- [**Find Value**](./binary_search_trees/find_value.go): Search for a value
- [**Insertion**](./binary_search_trees/insertion.go): Add new values
- [**Delete Node**](./binary_search_trees/delete_node.go): Remove nodes
- [**Find Min/Max**](./binary_search_trees/find_min_max.go): Find extreme values
- [**Floor**](./binary_search_trees/floor.go): Find floor value
- [**Ceil**](./binary_search_trees/ceil.go): Find ceiling value
- [**Trim Tree**](./binary_search_trees/trim_tree.go): Remove nodes outside range
- [**Print Range**](./binary_search_trees/print_range.go): Print values within range
- [**LCA BST**](./binary_search_trees/LCA_BST.go): Find Lowest Common Ancestor
- [**Is BST**](./binary_search_trees/is_bst.go): Validate BST property