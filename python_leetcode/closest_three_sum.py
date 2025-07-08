def closestThreeSum(nums: list[int], target: int) -> int:
    nums.sort()
    n = len(nums)
    closest_sum = float('inf')

    for i in range(n - 2):
        left, right = i + 1, n - 1
        while left < right:
            current_sum = nums[i] + nums[left] + nums[right]
            # Update closest_sum if current_sum is closer to target
            if abs(current_sum - target) < abs(closest_sum - target):
                closest_sum = current_sum
            # Early exit if exact match found
            if current_sum == target:
                return target
            # Adjust pointers
            if current_sum < target:
                left += 1
            else:
                right -= 1
    return closest_sum
