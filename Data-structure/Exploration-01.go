package main

import (
	"fmt"
)

type Course struct {
	Name                  string
	MentorSatisfaction    float64
	LearningSatisfaction  float64
}

func main() {
	data := []Course{
		{"Course A", 4.5, 4.0},
		{"Course B", 4.7, 4.1},
		{"Course C", 4.6, 4.3},
		{"Course A", 4.8, 4.4},
		{"Course B", 4.3, 4.2},
		{"Course C", 4.7, 4.5},
	}

	avgMentorSatisfaction := averageMentorSatisfaction(data)
	avgLearningSatisfaction := averageLearningSatisfaction(data)
	highestMentorSatisfactionCourse := courseWithHighestMentorSatisfaction(data)

	fmt.Printf("Average Mentor Satisfaction: %.2f\n", avgMentorSatisfaction)
	fmt.Printf("Average Learning Satisfaction: %.2f\n", avgLearningSatisfaction)
	fmt.Printf("Course with Highest Mentor Satisfaction: %s\n", highestMentorSatisfactionCourse)
}

func averageMentorSatisfaction(courses []Course) float64 {
	total := 0.0
	for _, course := range courses {
		total += course.MentorSatisfaction
	}
	return total / float64(len(courses))
}

func averageLearningSatisfaction(courses []Course) float64 {
	total := 0.0
	for _, course := range courses {
		total += course.LearningSatisfaction
	}
	return total / float64(len(courses))
}

func courseWithHighestMentorSatisfaction(courses []Course) string {
	courseTotals := make(map[string]float64)
	courseCounts := make(map[string]int)

	for _, course := range courses {
		courseTotals[course.Name] += course.MentorSatisfaction
		courseCounts[course.Name]++
	}

	highestAverage := 0.0
	highestCourse := ""

	for name, total := range courseTotals {
		average := total / float64(courseCounts[name])
		if average > highestAverage {
			highestAverage = average
			highestCourse = name
		}
	}

	return highestCourse
}
