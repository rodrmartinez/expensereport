package expenses

import (
	"fmt"
	"time"
)

type Type int

const (
	DINNER Type = iota + 1
	BREAKFAST
	CAR_RENTAL
	LUNCH
)

type Expense struct {
	Type   Type
	Amount int
}

func getExpenseName(expense Expense) string {
	switch expense.Type {
	case DINNER:
		return "Dinner"
	case BREAKFAST:
		return "Breakfast"
	case CAR_RENTAL:
		return "Car Rental"
	case LUNCH:
		return "Lunch"
	default:
		return "Expense not found"
	}
}

func expenseValid(expense Expense) string {
	if expense.Type == DINNER && expense.Amount > 5000 || expense.Type == BREAKFAST && expense.Amount > 1000 || expense.Type == LUNCH && expense.Amount > 2000 {
		return "X"
	} else {
		return " "
	}
}

func getExpenses(expenses []Expense) (int, int) {
	total := 0
	mealExpenses := 0

	for _, expense := range expenses {
		if expense.Type != CAR_RENTAL {
			mealExpenses += expense.Amount
		}

		total += expense.Amount
	}
	return mealExpenses, total
}

func printReport(expenses []Expense) {
	fmt.Printf("Expenses %s\n", time.Now().Format("2006-01-02"))
	for _, expense := range expenses {
		fmt.Printf("%s\t%d\t%s\n", getExpenseName(expense), expense.Amount, expenseValid(expense))
	}
	mealExpenses, total := getExpenses(expenses)
	fmt.Printf("Meal expenses: %d\n", mealExpenses)
	fmt.Printf("Total expenses: %d\n", total)

}
