package main

import (
	"fmt"
)

// This function calculares the salary having two parameters: amount of minutes per month and employee category.
func salaryCalculation(minutesWorkedByMonth int, category string) float64 {
	var salaryPerHour float64
	switch category {
	case "C":
		salaryPerHour = 1000.0
	case "B":
		salaryPerHour = 1500.0
	case "A":
		salaryPerHour = 3500.0
	default:
		return 0.0 //invalid category, return salary 0
	}

	hoursWorkedPerMonth := float64(minutesWorkedByMonth) / 60.0
	monthlySalary := salaryPerHour * hoursWorkedPerMonth

	//apply the aditional porcentage in base of the category

	switch category {
	case "B":
		monthlySalary += monthlySalary * 0.2
	case "A":
		monthlySalary += monthlySalary * 0.5
	}
	return monthlySalary
}

func main() {
	// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
	// Categoría C: su salario es de $1.000 por hora.
	// Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
	// Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.
	// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

	workingMinutes := 160 * 80
	categoryEmployee := "A"

	salary := salaryCalculation(workingMinutes, categoryEmployee)

	fmt.Printf("The employee salary is: $%.2f\n", salary)
}
