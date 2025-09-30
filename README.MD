# ESEP Grade Calculator

This library is a simple utility written in Go to calculate final grades for a course.
It computes the average of assignments, exams, and essays, applies weightings, and returns
a final letter grade (A-F)

Runtime requirements:
Golang 1.24+

Setup
```
git clone https://github.com/YOUR-USERNAME/esep-grade-calculator.git
cd esep-grade-calculator/grade_calculator
```

Make sure Go is installed:
```
go version
```

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

The grade weighting is currently:
Assignments = 50%
Exams = 35%
Essays = 15%