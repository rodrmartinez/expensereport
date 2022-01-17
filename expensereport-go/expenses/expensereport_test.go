package expenses

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

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
