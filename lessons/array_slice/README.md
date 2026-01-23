# Array vs Slice
## The Core Differences (Staff Level Depth)

### Value Types vs. Reference Descriptors

* **Arrays are Values:** Assigning an array to a new variable or passing it to a function creates a full copy of the array's memory.
    * **Implication:** Passing large arrays (e.g., [100000]int) to functions is expensive due to memory copying. You almost always want to pass a pointer to the array or use a slice.

* **Slices are "Reference-like":** A slice variable is actually a struct (header) passed by value, but it contains a pointer to the underlying data.
    * **Implication:** Passing a slice is cheap (constant cost) regardless of the data size, as you only copy the small struct header.

## The Slice Header (Internals)

To truly master slices, you must know what a slice looks like in the Go runtime (reflect.SliceHeader):
```Go
type SliceHeader struct {
    Data uintptr // Pointer to the underlying array
    Len  int     // Number of elements in the slice
    Cap  int     // Total slots available in the underlying array starting from 'Data'
}
```

When you pass a slice to a function, you are copying these three fields. This explains why:
1. Modifying elements works: The copy points to the same underlying array.
2. Append might not update the original: If append causes a re-allocation (exceeding Cap), the function's local slice header points to a new array, but the caller's slice header still points to the old array.

## Slice Growth & Allocation Strategy

When you append to a slice that is full (len == cap), Go allocates a new, larger array and copies the data over.
    * **Heuristic:** Traditionally, it doubled capacity for small slices and grew by ~1.25x for large slices to avoid memory fragmentation. (Note: The exact formula evolves with Go versions, but understanding the O(N) copy cost during growth is key).

    * **Optimization:** If you know the size beforehand, always pre-allocate using make([]T, len, cap). This prevents repeated allocations and copies.

## The "Memory Leak" (Keeping References)

This is a classic discussion point. If you load a large file into memory (e.g., a 100MB array) and return a small slice of it (e.g., data[:10]), the entire 100MB backing array remains in memory because the small slice holds a pointer to it. The Garbage Collector (GC) cannot reclaim the large array.

* **Solution:** Copy the small chunk you need into a fresh slice so the massive backing array can be garbage collected.
```Go

    // Bad
    return hugeArray[:4]

    // Good
    small := make([]byte, 4)
    copy(small, hugeArray[:4])
    return small
```

## Nil Slice vs. Empty Slice

They behave the same structurally in many contexts (both have len 0, cap 0), but they are initialized differently.

`var s []int` (Nil slice, no underlying array allocated). `s == nil` is true.

`s := []int{}` or `make([]int, 0)` (Empty slice, usually holds a pointer to a specialized zero-size variable). `s == nil` is false.

**Interview note:** JSON marshaling treats them differently. A nil slice marshals to null, whereas an empty slice marshals to []. This matters for API contracts.