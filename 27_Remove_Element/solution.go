package main

func removeElement(nums []int, val int) int {
	var i int = 0

	for _, el := range nums {
		if el != val {
			nums[i] = el
			i++
		}
	}

	return i
}

/*func main() {
	nums := []int{3, 2, 2, 3} // Input array
	val := 2 // Value to remove
	expectedNums := []int {2, 2} // The expected answer with correct length.
                            // It is sorted with no values equaling val.

	k := removeElement(nums, val); // Calls your implementation

	assert k == expectedNums.length;
	sort(nums, 0, k); // Sort the first k elements of nums
	for (int i = 0; i < actualLength; i++) {
    	assert nums[i] == expectedNums[i];
	}
}*/
