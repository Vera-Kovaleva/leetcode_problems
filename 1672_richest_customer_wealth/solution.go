package main

func maximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, curAccount := range accounts {
		curWealth :=0
		for _, curBank := range curAccount {
			curWealth += curBank
		}
		if curWealth > maxWealth {
			maxWealth = curWealth
		}
	}
    return maxWealth
}