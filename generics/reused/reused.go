package reused

import "golang.org/x/exp/constraints"

// MaxWithoutGenerics return the max value
func MaxWithoutGenerics(num []int) (resMax int) {
	if len(num) == 0 {
		return
	}

	resMax = num[0]
	for _, v := range num {
		if v > resMax {
			resMax = v
		}
	}

	return resMax
}

// MaxWithGenerics return the max value
func MaxWithGenerics[T constraints.Ordered](nums []T) T {
	if len(nums) == 0 {
		return *new(T)
	}

	resMax := nums[0]
	for _, v := range nums[1:] {
		if v > resMax {
			resMax = v
		}
	}

	return resMax
}
