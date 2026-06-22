package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
}

func getExpenses() []Expense {
	var expenses []Expense
	data, err := os.ReadFile("expenses.json")
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading file: ", err)
		return nil
	}
	if len(data) > 0 {
		if err := json.Unmarshal(data, &expenses); err != nil {
			fmt.Println("Error parsing file: ", err)
			return nil
		}
	}
	return expenses
}

func saveExpenses(expenses []Expense) {
	out, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		fmt.Println("Error encoding: ", err)
		return
	}
	if err := os.WriteFile("expenses.json", out, 0o644); err != nil {
		fmt.Println("Error writing file: ", err)
	}
}

func usage() {
	fmt.Println(`Expense Tracker - track your expenses from the command line

Usage:
  expense-tracker <command> [flags]

Commands:
  add       Add an expense          --description <text> --amount <number>
  delete    Delete an expense       --id <id>
  list      List all expenses
  summary   Show total              [--month <1-12>]
  help      Show this help

Run "expense-tracker <command> --help" for flags on a specific command.`)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Exptected a subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help", "-h", "--help":
		usage()

	case "add":
		// NOTE: Add Logic
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := addCmd.String("description", "", "expense description")
		amount := addCmd.Float64("amount", 0, "expense amount")

		// parse everything after "add"
		addCmd.Parse(os.Args[2:])

		expenses := getExpenses()

		maxID := 0
		for _, e := range expenses {
			if e.ID > maxID {
				maxID = e.ID
			}
		}
		expenses = append(expenses, Expense{
			ID:          maxID + 1,
			Date:        time.Now(),
			Description: *description,
			Amount:      int(*amount),
		})
		saveExpenses(expenses)

	case "delete":
		// NOTE: Delete logic
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		targetID := deleteCmd.Int("id", 0, "which id to dlete")

		deleteCmd.Parse(os.Args[2:])

		expenses := getExpenses()

		idx := -1
		for i, e := range expenses {
			if e.ID == *targetID {
				idx = i
				break
			}
		}
		if idx == -1 {
			fmt.Println("No expense with id: ", *targetID)
			return
		}

		expenses = slices.Delete(expenses, idx, idx+1)

		saveExpenses(expenses)

	case "list":
		// NOTE: Consider using text/tabwriter package if description become too long
		expenses := getExpenses()
		fmt.Println("┌──────┬────────────┬──────────────────────────────┬────────────┐")
		fmt.Printf("│ %-4s │ %-10s │ %-28s │ %-10s │\n", "ID", "Date", "Description", "Amount")
		fmt.Println("├──────┼────────────┼──────────────────────────────┼────────────┤")
		for _, v := range expenses {
			amount := fmt.Sprintf("%d CHF", v.Amount)
			fmt.Printf("│ %-4d │ %-10s │ %-28s │ %10s │\n",
				v.ID, v.Date.Format("02/01/2006"), v.Description, amount)
		}
		fmt.Println("└──────┴────────────┴──────────────────────────────┴────────────┘")

	case "summary":
		summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		specificMonth := summaryCmd.Int("month", 0, "total for specific month")
		summaryCmd.Parse(os.Args[2:])
		expenses := getExpenses()

		sum := 0
		for _, v := range expenses {
			if *specificMonth == 0 || int(v.Date.Month()) == *specificMonth {
				sum += v.Amount
			}
		}

		if *specificMonth == 0 {
			fmt.Printf("Total Amount: %d\n", sum)
		} else {
			fmt.Printf("Total Amount in %s: %d\n", time.Month(*specificMonth).String(), sum)
		}

	case "edit":
		// TODO: Add edit logic

	default:
		fmt.Printf("unknown subcommand: %s\n", os.Args[1])
		os.Exit(1)
	}
}
