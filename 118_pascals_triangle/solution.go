package main

/*func curString(prevString []int) []int {
	var curString []int
	curString = append(curString, 1)
	for i := 1; i < len(prevString); i++ {
		curString = append(curString, prevString[i]+prevString[i-1])
	}
	curString = append(curString, 1)
	return curString
}*/

func generate(numRows int) [][]int {
	var res [][]int
	res = append(res, []int{1})
	prevString := res[0]
	for i := 0; i < numRows-1; i++ {

		var curString []int
		curString = append(curString, 1)
		for i := 1; i < len(prevString); i++ {
			curString = append(curString, prevString[i]+prevString[i-1])
		}
		curString = append(curString, 1)

		res = append(res, curString)
		prevString = curString
	}

	return res
}
