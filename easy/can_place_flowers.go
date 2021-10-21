package easy

// You have a long flowerbed in which some of the plots are planted, and some are not. However,
// flowers cannot be planted in adjacent plots.
//
// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty,
// and an integer n, return if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule.
//
//
//
// Example 1:
// Input: flowerbed = [1,0,0,0,1], n = 1
// Output: true
//
// Example 2:
// Input: flowerbed = [1,0,0,0,1], n = 2
// Output: false
//
// Constraints:
//
// 1 <= flowerbed.length <= 2 * 104
// flowerbed[i] is 0 or 1.
// There are no two adjacent flowers in flowerbed.
// 0 <= n <= flowerbed.length
func canPlaceFlowers(flowerbed []int, n int) bool {
	counter := 0

	for i := 0; i < len(flowerbed); i++ {
		back := i - 1
		front := i + 1

		if back < 0 {
			back = 0
		}

		if front > len(flowerbed)-1 {
			front = len(flowerbed) - 1
		}

		if flowerbed[back] == 0 && flowerbed[front] == 0 && flowerbed[i] != 1 {
			flowerbed[i] = 1
			counter++
		}

		if counter >= n {
			return true
		}
	}

	return counter >= n
}
