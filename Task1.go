// Task: Student Grade Calculator
// Create a Go console application that allows students to calculate their average grade based on different subjects. The application should prompt the student to enter their name and the number of subjects they have taken. For each subject, the student should enter the subject name and the grade obtained (numeric value). After entering all subjects and grades, the application should display the student's name, individual subject grades, and the calculated average grade.
// Requirements:
// Use variables and data types to store student data.
// Use conditional statements to validate input (e.g., ensure grade values are within a valid range).
// Implement loops to handle multiple subjects and grades.
// Utilize collections (e.g., List, Dictionary) to store subject names and corresponding grades.
// Define a method to calculate the average grade based on the entered grades.
// Use string interpolation to display the results in a user-friendly format.

package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter Your Name: ")
	fmt.Scanln(&name)

	var numberOfSubjects int
	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&numberOfSubjects)

	if numberOfSubjects <= 0 {
		fmt.Println("Invalid number of subjects")
		return
	}

	subjectGrades := make(map[string]float64)

	for i := 0; i < numberOfSubjects; i++ {
		var subject string
		var grade float64

		fmt.Printf("Enter the name of subject %d: ", i+1)
		fmt.Scanln(&subject)

		fmt.Printf("Enter the grad for %d: ", i+1)
		fmt.Scanln(&grade)

		if grade < 0 || grade > 100 {
			fmt.Println("Invalid grade. (0-100)")
			i--
			continue
		}

		subjectGrades[subject] = grade
	}

	averageGrade := calculateAverage(subjectGrades)

	//Displaying the result
	fmt.Printf("\nStudent Name: %s\n", name)
	fmt.Println("Student Grades:")
	for subject, grade := range subjectGrades {
		fmt.Printf("	%s: %.2g\n", subject, grade)
	}
	fmt.Printf("Average Grade: %.2g", averageGrade)

}

func calculateAverage(grades map[string]float64) float64 {
	var total float64
	for _, grade := range grades {
		total += grade
	}
	return total / float64(len(grades))
}
