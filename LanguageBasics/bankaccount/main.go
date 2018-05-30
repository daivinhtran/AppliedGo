package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/appliedgocourses/bank"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	// Restore the bank data.
	// Call bank.Load() here and handle any error.
	err := bank.Load()
	if err != nil {
		fmt.Println("Unable to load the bank database")
		return
	}

	// Save the bank data.
	// Use a deferred function for calling bank.Save().
	// (See the lecture on function behavior.)
	// If Save() returns an error, print it.
	defer func() {
		err := bank.Save()
		if err != nil {
			fmt.Println("Unable to save the bank database")
		}
	}()

	// Perform the action.
	// os.Args[0] is the path to the executable.
	// os.Args[1] is the first parameter - the action we want to perform:
	// create, list, update, transfer, or history.

	switch os.Args[1] {
	case "create":
		if len(os.Args) < 3 {
			usage()
			return
		}
		bank.NewAccount(string(os.Args[2]))
	case "list":
		fmt.Printf(bank.ListAccounts())
	case "update":
		if len(os.Args) < 4 {
			usage()
			return
		}
		amount, err := strconv.Atoi(os.Args[3])
		if err != nil {
			usage()
			return
		}
		if _, err := update(os.Args[2], amount); err != nil {
			fmt.Println(err)
			return
		}
	case "transfer":
		if len(os.Args) < 5 {
			usage()
			return
		}
		amount, err := strconv.Atoi(os.Args[4])
		if err != nil {
			usage()
			return
		}
		if _, _, err := transfer(os.Args[2], os.Args[3], amount); err != nil {
			fmt.Println(err)
			return
		}
	case "history":

	default:
		usage()
		return
	}
}

func usage() {
	fmt.Println(`Usage:

bank create <name>                     Create an account.
bank list                              List all accounts.
bank update <name> <amount>            Deposit or withdraw money.
bank transfer <name> <name> <amount>   Transfer money between two accounts.
bank history <name>                    Show an account's transaction history.
`)
}

// update takes a name and an amount, deposits the amount if it
// is greater than zero, or withdraws it if it is less than zero,
// and returns the new balance and any error that occurred.
func update(name string, amount int) (int, error) {
	// add error handling
	account, err := bank.GetAccount(name)
	if err != nil {
		return 0, err
	}
	if amount > 0 {
		return bank.Deposit(account, amount)
	}
	return bank.Withdraw(account, 0-amount)
}

// // transfer takes two names and an amount, transfers the amount from
// // the account belonging to name #1 to the account belonging to name #2,
// // and returns the new balances of both accounts and any error that occurred.
func transfer(name1, name2 string, amount int) (int, int, error) {
	if amount < 0 {
		return 0, 0, errors.New("<amount> must be positive")
	}
	balance1, err := update(name1, 0-amount)
	if err != nil {
		return 0, 0, err
	}
	balance2, err := update(name2, amount)
	if err != nil {
		return 0, 0, err
	}
	return balance1, balance2, nil
}

// // history takes an account name, retrieves the account, and calls bank.History()
// // to get the history closure function. Then it calls the closure in a loop,
// // formatting the return values and appending the result to the output string, until the boolean return parameter of the closure is `false`.
func history(name string) (string, error) {
	account, err := bank.GetAccount(name)
	if err != nil {
		return "", err
	}
	var ret string
	history := bank.History(account)
	for amt, bal, more := history(); more == true; amt, bal, more = history() {
		ret = ret + fmt.Sprintf("%s %d %d\n", name, amt, bal)
	}
	return ret, nil
}
