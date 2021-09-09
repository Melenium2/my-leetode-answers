package easy

// Given an integer array nums, handle multiple queries of the following type:
//
// Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
// Implement the numArray class:
//
// numArray(int[] nums) Initializes the object with the integer array nums.
// int sumRange(int left, int right) Returns the sum of the elements of nums between
// indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
//
//
// Example 1:
//
// Input
// ["numArray", "sumRange", "sumRange", "sumRange"]
// [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
// Output
// [null, 1, -1, -3]
//
// Explanation
// numArray numArray = new numArray([-2, 0, 3, -5, 2, -1]);
// numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
// numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
// numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3
type numArray struct {
	nums []int
	sums []int
}


func constructorNumArray(nums []int) numArray {
	sum := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i] // we use prefix sum in this task
	}

	return numArray{
		nums: nums,
		sums: sum,
	}
}


func (na *numArray) SumRange(left int, right int) int {
	return na.sums[right+1] - na.sums[left]
}


/**
 * Your numArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */