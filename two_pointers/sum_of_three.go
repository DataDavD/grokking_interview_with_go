package palindrome

import (
	"log"
	"os"
	"sort"
)

// threeSum determines if a triplet of distinct values in a provided slice of ints called nums sums to equal
// the target value. True if so, else false.
// The time complexity is O(n^2) because:
// Sorting nums takes O(nlog(n)) given the Go standard library's `sort.Ints` function uses a variation of QuickSort.
// The nested loop takes O(n^2) in order to find the triplet.
// So, the total time complexity is O(nlog(n) + n^2) which simplifies to O(n^2).
// The space complexity is O(log(n)) given the `sort.Ints` function uses a variation of QuickSort and thus
// has a space complexity of O(log(n)).
func threeSum(nums []int, target int) bool {
	log.SetOutput(os.Stdout)

	sort.Ints(nums) // sort in ascending order
	log.Printf("nums: %v\n", nums)

	for i := 0; i < len(nums)-2; i++ {
		// skip if at least 2 of the values are the same
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Initialize the low and high pointers
		low := i + 1
		high := len(nums) - 1
		log.Printf("iteration={%d}; low={%d}; high={%d}\n", i, low, high)

		// Traverse the slice to find the triplet whose sum equals the target
		for low < high {
			triple := nums[i] + nums[low] + nums[high]
			log.Printf("triple: %d\n", triple)

			if triple == target {
				return true
			} else if triple < target {
				low++ // The sum of the triplet is less than the target, so move low forward.
			} else {
				high-- // The sum of the triplet is higher than the target, so move high backward.
			}
		}
	}

	// No triplet found that sums to the target value
	return false
}
