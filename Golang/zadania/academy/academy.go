package academy

import (
	"math"
)

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	gradesLen := len(grades)
	if gradesLen == 0 {
		return 0
	}

	sum := 0
	for _, grade := range grades {
		sum += grade
	}

	avg := float64(sum) / float64(gradesLen)
	return int(math.Round(avg))
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {
	attendanceLen := len(attendance)
	if attendanceLen == 0 {
		return 0.0
	}

	presenceCount := 0
	for _, isPresent := range attendance {
		if isPresent {
			presenceCount++
		}
	}

	attendancePercentage := float64(presenceCount) / float64(attendanceLen)
	return attendancePercentage
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	averageAttendance := AttendancePercentage(s.Attendance)

	minimumFinalScore := 1
	insufficientProjectGrade := s.Project <= 1
	insufficientAttendance := averageAttendance < 0.6
	insufficientAverageGrade := AverageGrade(s.Grades) <= 1
	if insufficientProjectGrade || insufficientAttendance || insufficientAverageGrade {
		return minimumFinalScore
	}

	finalScoreGrades := []int{AverageGrade(s.Grades), s.Project}
	finalScore := AverageGrade(finalScoreGrades)
	if averageAttendance < 0.8 {
		return finalScore - 1
	}

	return finalScore
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	finalGradesMap := map[string]uint8{}
	if len(students) == 0 {
		return finalGradesMap
	}

	for _, student := range students {
		finalGradesMap[student.Name] = uint8(FinalGrade(student))
	}

	return finalGradesMap
}
