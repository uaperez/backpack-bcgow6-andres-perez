package main

import "fmt"

var (
	SALARY_CATEGORY_A = 1_000
	SALARY_CATEGORY_B = 1_500
	SALARY_CATEGORY_C = 3_000
)

func calculateSalary(workedMinutesPerMonth int, category string) float32 {
	switch category {
	case "A":
		return float32(SALARY_CATEGORY_A) * float32(workedMinutesPerMonth)
	case "B":
		monthlySalary := float32(SALARY_CATEGORY_B) * float32(workedMinutesPerMonth)
		return (0.20 * monthlySalary) + monthlySalary
	case "C":
		monthlySalary := float32(SALARY_CATEGORY_B) * float32(workedMinutesPerMonth)
		return (0.50 * monthlySalary) + monthlySalary
	default:
		return 0
	}
}

func main() {
	fmt.Println(calculateSalary(128, "C"))
}
