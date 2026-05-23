package main

func separateDigits(nums []int) []int {
	res := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		for num > 0 {
			res = append(res, num%10)
			num /= 10
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

/*
func separateDigits(nums []int) []int {
	res := []int{}
	for _, num := range nums {
		cur_num := []int{}
		for num > 0 {
			cur_num = append(cur_num, num%10)
			num /= 10
		}
		for i := len(cur_num)-1; i >= 0; i-- {
			res = append(res, cur_num[i])
		}
	}
	return res
}
*/
