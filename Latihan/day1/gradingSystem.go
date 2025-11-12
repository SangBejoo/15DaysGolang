package main

func gradingSystem(score int) string {
	switch {
	case score >= 90 && score <= 100:
		if score >= 95 {
			return "A+"
		} else {
			return "A"
		}
	case score >= 80 && score <= 90:
		if score >= 85 {
			return "B+"
		} else {
			return "B"
		}
	case score >= 70 && score <= 80:
		if score >= 75 {
			return "C+"
		} else {
			return "C"
		}
	case score >= 60 && score <= 70:
		if score >= 65 {
			return "D+"
		} else {
			return "D"
		}
	case score >= 0 && score < 60:
		return "F"
	default:
		return "Invalid score"
	}
}

func main() {
	// grading system
	scores := []int{98, 85, 76, 64, 50, 110, -5}
	for _, score := range scores {
		grade := gradingSystem(score)
		println("Score:", score, "Grade:", grade)
	}
}
