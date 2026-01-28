package main

func generatePossibleNextMoves(currentState string) []string {
	possibleStates := []string{}

	for i := 0; i < len(currentState)-1; i++ {
		if currentState[i] == '+' && currentState[i+1] == '+' {
			possibleStates = append(possibleStates, currentState[:i]+"--"+currentState[i+2:])
		}
	}
	return possibleStates
}
