package main

func kLengthApart(nums []int, k int) bool {
	prev1 := 0
	if nums[0] == 0 {
		prev1 = -k
	}
	for i:=1; i<len(nums); i++ {
		if nums[i] == 1 {
			if i-prev1 <= k {
				return false
			} 
			prev1 = i
		}
	}
    return true
}
