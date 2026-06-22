# Expense Tracker

A simple command-line expense tracker written in Go. Add, list, delete, and summarize expenses, with all data stored locally in a JSON file.

## Features

- Add an expense with a description and amount (the date is recorded automatically)
- List all expenses in a clean, boxed table
- Delete an expense by its ID
- Get a total summary, optionally filtered by month

## Requirements

- [Go](https://go.dev/dl/) 1.21 or newer (uses `slices.Delete`)

## Build

```bash
go build -o expense-tracker
```

This produces an `expense-tracker` binary in the current directory. You can also run commands directly with `go run .` during development.

## Usage

### Add an expense

```bash
./expense-tracker add --description "Lunch with client" --amount 45
```

### List all expenses

```bash
./expense-tracker list
```

```
┌──────┬────────────┬──────────────────────────────┬────────────┐
│ ID   │ Date       │ Description                  │ Amount     │
├──────┼────────────┼──────────────────────────────┼────────────┤
│ 1    │ 22/06/2025 │ Lunch with client            │        $45 │
└──────┴────────────┴──────────────────────────────┴────────────┘
```

### Delete an expense

```bash
./expense-tracker delete --id 1
```

### Summary

Total of all expenses:

```bash
./expense-tracker summary
```

Total for a specific month (1–12):

```bash
./expense-tracker summary --month 6
```

## Data Storage

Expenses are saved to `expenses.json` in the working directory. Each expense has the following shape:

```json
[
  {
    "id": 1,
    "description": "Lunch with client",
    "amount": 45,
    "date": "2025-06-22T14:33:09.847293-04:00"
  }
]
```

The file is created automatically on the first `add`. IDs are assigned as one higher than the current maximum, so they remain unique even after deletions.

## Project Status

- [x] Add expenses
- [x] List expenses
- [x] Delete expenses
- [x] Monthly summary
- [ ] Edit expenses

## Acknowledgements

This project is based on the [Expense Tracker challenge](https://roadmap.sh/projects/expense-tracker) from roadmap.sh.
