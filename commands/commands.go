package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/renatobaez/simple_cli/expenses"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	fmt.Print("--> ")
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	str = strings.Replace(str, "\n", "", 1)
	return str, nil
}

func ShowInConsole(expensesList []float32) {
	fmt.Println(contentString(expensesList))
}

func contentString(expensesList []float32) string {
	builder := strings.Builder{}
	fmt.Println("")
	for i, expense := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))
		if i == (len(expensesList) - 1) {
			builder.WriteString("")
			builder.WriteString("==========================\n")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", expenses.Sum(expensesList...)))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", expenses.Max(expensesList...)))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", expenses.Min(expensesList...)))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", expenses.Average(expensesList...)))
			builder.WriteString("==========================")
		}
	}
	return (builder.String())
}
func Export(fileName string, list []float32) error {

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = w.WriteString(contentString(list))
	if err != nil {
		return err
	}
	return w.Flush()
}
