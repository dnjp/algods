# Binary Search

If we are implementing a sequential search, there are at most n-1 more items to
look through if the first item is not what we're looking for.

Instead of searching in sequence, a binary search will start by examining the
middle item.

- If that item is the one we're searching for, we are done.
- Otherwise, we can use the ordered nature of the list to eliminate half of the
  remaining items.
- If the item we're searching for is greater than the middle
  item, we know that the entire lower half of the list as well as the middle
  item
  can be elimintated from further consideration.
- If the item is in the list, it must be in the upper half.
- We then repeat the process with the upper half.

Binary Search is a great example of a divide and conquer strategy - meaning we
divide the problem into smaller pieces and then reassemble the whole problem to
get the result.

Here's an implementation of binary search in go:

```go
func binarySearch(list []int, item int) bool {
	if len(list) == 0 {
		return false
	}

	mid := len(list) / 2

	if list[mid] == item {
		return true
	}

	if item < list[mid] {
		return binarySearch(list[:mid], item)
	}

	return binarySearch(list[mid+1:], item)
}
```

## Analysis

Binary Search elimintates around half of the remaining items from consideration. If we start with n items, about n/2 items will be left after the first comparison. After the second comparison there will be about n/4 items and on it goes. If we split the list enough times, we end up with a list that has just one item or the item is not in the list. The number of comparisons to get to this point is i where `n/2^i = 1`. Solving for i gives us i = log n. The maximum number of comparisons is logarithmic with respsect to the number of items in the list.

Therefore, Binary Search is O(log n).

