package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 75, Assignment)
	gradeCalculator.AddGrade("exam", 70, Exam)
	gradeCalculator.AddGrade("essay", 72, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected '%s'; got '%s'", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam", 60, Exam)
	gradeCalculator.AddGrade("essay", 62, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected '%s'; got '%s'", expected_value, actual_value)
	}
}
func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 10, Assignment)
	gradeCalculator.AddGrade("exam 1", 20, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 15, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestNewGradeCalculator(t *testing.T) {
	gc := NewGradeCalculator()
	if gc == nil {
		t.Errorf("Expected non-nil GradeCalculator")
	}
	if len(gc.grades) != 0 {
		t.Errorf("Expected empty grades slice in new GradeCalculator")
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected 'assignment'; got '%s'", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected 'exam'; got '%s'", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected 'essay'; got '%s'", Essay.String())
	}
}

func TestComputeAverageEmpty(t *testing.T) {
	avg := computeAverageByType([]Grade{}, Assignment)
	if avg != 0 {
		t.Errorf("Expected average of empty slice to be 0; got %d", avg)
	}
}
