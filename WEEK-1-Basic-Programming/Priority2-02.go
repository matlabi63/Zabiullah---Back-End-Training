package main

import "fmt"

func main() {
	budget := 51
	duration := 10
	difficulty := 3

	// Calculate scores
	budgetScore := calculateBudgetScore(budget)
	durationScore := calculateDurationScore(duration)
	difficultyScore := calculateDifficultyScore(difficulty)

	// total score
	totalScore := budgetScore + durationScore + difficultyScore

	// Determine project priority
	var priority string
	switch {
	case totalScore >= 24:
		priority = "High"
	case totalScore >= 16:
		priority = "Medium"
	case totalScore >= 9:
		priority = "Low"
	default:
		priority = "Impossible"
	}

	// Output the result
	fmt.Println("Project Priority:", priority)
}

// Function to calculate budget
func calculateBudgetScore(budget int) int {
	switch {
	case budget <= 50:
		return 4
	case budget <= 80:
		return 3
	case budget <= 100:
		return 2
	default:
		return 1
	}
}

// Function to calculate duration
func calculateDurationScore(duration int) int {
	switch {
	case duration <= 20:
		return 10
	case duration <= 30:
		return 5
	case duration <= 50:
		return 2
	default:
		return 1
	}
}

// Function to calculate difficulty score
func calculateDifficultyScore(difficulty int) int {
	switch {
	case difficulty <= 3:
		return 10
	case difficulty <= 6:
		return 5
	case difficulty <= 10:
		return 1
	default:
		return 0
	}
}
