package expenses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintEmptyReport(t *testing.T) {
	expenses := []Expense{{}}
	printReport(expenses)
}

func TestPrintReport(t *testing.T) {
	expenses := []Expense{{1, 10}, {2, 20}}
	printReport(expenses)
}

func TestPrintReportLunch(t *testing.T) {
	expenses := []Expense{{4, 30}}
	printReport(expenses)
}

func TestPrintReportOverExpenses(t *testing.T) {
	expenses := []Expense{{4, 3000}}
	printReport(expenses)
}

func TestExpenseValid(t *testing.T) {
	expense := Expense{1, 5000}
	got := expenseValid(expense)
	assert.Equal(t, " ", got)

	expense = Expense{1, 5001}
	got = expenseValid(expense)
	assert.Equal(t, "X", got)

	expense = Expense{2, 1001}
	got = expenseValid(expense)
	assert.Equal(t, "X", got)

}

func TestGetExpenseName(t *testing.T) {
	expense := Expense{1, 10}
	got := getExpenseName(expense)
	assert.Equal(t, "Dinner", got)

	expense = Expense{2, 10}
	got = getExpenseName(expense)
	assert.Equal(t, "Breakfast", got)

	expense = Expense{3, 10}
	got = getExpenseName(expense)
	assert.Equal(t, "Car Rental", got)

}

func TestGetExpenses(t *testing.T) {
	expenses := []Expense{{1, 10}}
	meals, total := getExpenses(expenses)
	assert.Equal(t, 10, meals)
	assert.Equal(t, 10, total)

    expenses = []Expense{{1, 10},{2,50},{3,150}}
	meals, total = getExpenses(expenses)
	assert.Equal(t, 60, meals)
	assert.Equal(t, 210, total)

}
