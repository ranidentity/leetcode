import heapq

def find_kth_largest(nums, k):
    """
    Finds the k-th largest element in an unsorted array using a min-heap.

    Args:
        nums: The input list of numbers.
        k: The desired k-th largest position.

    Returns:
        The k-th largest element.
    """
    if not nums or k <= 0 or k > len(nums):
        raise ValueError("Invalid input")
    
    min_heap = []

    for num in nums:
        if len(min_heap)<k:
            heapq.heappush(min_heap,num)
        else:
            if num> min_heap[0]:
                heapq.heapreplace(min_heap, num)
    
    return min_heap[0]