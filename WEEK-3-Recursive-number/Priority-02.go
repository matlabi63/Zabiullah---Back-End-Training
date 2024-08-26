
package main

import "fmt"

type Student struct {
        Name        string
        MathScore   int
        ScienceScore int
        EnglishScore int
}

func main() {
        students := []Student{
                {"jamie", 67, 88, 90},
                {"michael", 77, 77, 90},
                {"aziz", 100, 65, 88},
                {"jamal", 50, 80, 75},
                {"eric", 70, 80, 65},
                {"john", 80, 76, 68},
        }

        getInfo(students)
}

func getInfo(students []Student) {
        // Find best student in math
        bestMathStudent := findBestStudent(students, func(s Student) int { return s.MathScore })
        fmt.Printf("best student in math: %s with %d\n", bestMathStudent.Name, bestMathStudent.MathScore)

        // Find best student in science
        bestScienceStudent := findBestStudent(students, func(s Student) int { return s.ScienceScore })
        fmt.Printf("best student in science: %s with %d\n", bestScienceStudent.Name, bestScienceStudent.ScienceScore)

        // Find best student in english
        bestEnglishStudent := findBestStudent(students, func(s Student) int { return s.EnglishScore })
        fmt.Printf("best student in english: %s with %d\n", bestEnglishStudent.Name, bestEnglishStudent.EnglishScore)

        // Find best student in average
        bestAverageStudent := findBestStudent(students, func(s Student) int { return (s.MathScore + s.ScienceScore + s.EnglishScore) / 3 })
        fmt.Printf("best student in average: %s with %d\n", bestAverageStudent.Name, (bestAverageStudent.MathScore + bestAverageStudent.ScienceScore + bestAverageStudent.EnglishScore) / 3)
}

func findBestStudent(students []Student, scoreFunc func(Student) int) Student {
        bestStudent := students[0]
        bestScore := scoreFunc(students[0])

        for _, student := range students[1:] {
                score := scoreFunc(student)
                if score > bestScore {
                        bestStudent = student
                        bestScore = score
                }
        }

        return bestStudent
}