package easy

// We define a harmonious array as an array where the difference between its maximum value
// and its minimum value is exactly 1.
//
// Given an integer array nums, return the length of its longest harmonious subsequence
// among all its possible subsequences.
//
// A subsequence of array is a sequence that can be derived from the array by deleting
// some or no elements without changing the order of the remaining elements.
//
//
//
// Example 1:
// Input: nums = [1,3,2,2,5,2,3,7]
// Output: 5
// Explanation: The longest harmonious subsequence is [3,2,2,2,3].
//
// Example 2:
// Input: nums = [1,2,3,4]
// Output: 2
//
// Example 3:
// Input: nums = [1,1,1,1]
// Output: 0
func findLHS(nums []int) int {
	count := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	result := 0

	for num, length := range count {
		if l, ok := count[num+1]; ok {
			result = max(result, length+l)
		}
	}

	return result
}

//   int findLHS(vector<int>& nums) {
//        sort(nums.begin(),nums.end());
//        int len = 0;
//        for(int i = 1, start = 0, new_start = 0; i<nums.size(); i++)
//        {
//
//            if (nums[i] - nums[start] > 1)
//                start = new_start;
//            if (nums[i] != nums[i-1])
//                new_start = i;
//            if(nums[i] - nums[start] == 1)
//                len = max(len, i-start+1);
//        }
//   return len;
