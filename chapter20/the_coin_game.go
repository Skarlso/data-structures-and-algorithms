package chapter20

func GameWinner(numberOfCoins int, currentPlayer string) string {
	var nextPlayer string
	if numberOfCoins <= 0 {
		return currentPlayer
	}
	if currentPlayer == "you" {
		nextPlayer = "them"
	} else {
		nextPlayer = "you"
	}

	if GameWinner(numberOfCoins-1, nextPlayer) == currentPlayer || GameWinner(numberOfCoins-2, nextPlayer) == currentPlayer {
		return currentPlayer
	}

	return nextPlayer
}

func GameWinnerWithMath(numberOfCoins int) string {
	if numberOfCoins%3 == 0 {
		return "them"
	}
	return "you"
}
