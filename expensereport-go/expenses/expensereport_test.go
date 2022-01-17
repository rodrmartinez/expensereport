package expenses

import (
	"testing"

	"github.com/stretchr/testify/assert"

	approvaltests "github.com/approvals/go-approval-tests"
)

func TestPrintEmptyReport(t *testing.T) {
	expenses := []Expense{{}}
	approvaltests.VerifyString(t, printReport(expenses))
}

func TestPrintReport(t *testing.T) {
	expenses := []Expense{{1, 10}, {2, 20},{3,40}}
	approvaltests.VerifyString(t, printReport(expenses))

}

func TestPrintReportLunch(t *testing.T) {
	expenses := []Expense{{4, 30}}
	approvaltests.VerifyString(t, printReport(expenses))
}

func TestPrintReportOverExpenses(t *testing.T) {
	expenses := []Expense{{4, 3000}}
	approvaltests.VerifyString(t, printReport(expenses))

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

	expenses = []Expense{{1, 10}, {2, 50}, {3, 150}}
	meals, total = getExpenses(expenses)
	assert.Equal(t, 60, meals)
	assert.Equal(t, 210, total)

}

func TestGetLunchName(t *testing.T) {
	expense := Expense{4, 10}
	got := getExpenseName(expense)
	assert.Equal(t, "Lunch", got)
}

func TestLunchValid(t *testing.T) {
	expense := Expense{4, 50}
	got := expenseValid(expense)
	assert.Equal(t, " ", got)

	expense = Expense{4, 2001}
	got = expenseValid(expense)
	assert.Equal(t, "X", got)

}

func TestGetExpenseWithLunch(t *testing.T) {
	expenses := []Expense{{4, 10}}
	meals, total := getExpenses(expenses)
	assert.Equal(t, 10, meals)
	assert.Equal(t, 10, total)

	expenses = []Expense{{1, 10}, {2, 50}, {3, 150}, {4, 30}}
	meals, total = getExpenses(expenses)
	assert.Equal(t, 90, meals)
	assert.Equal(t, 240, total)

}
