package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type book struct {
	title     string
	price     float64
	inventory int
}

var accountBalance float64 = 2000.00

var books = []book{
	{
		title:     "1984",
		price:     12.53,
		inventory: 11,
	},
	{
		title:     "1985",
		price:     9.81,
		inventory: 17,
	},
	{
		title:     "1986",
		price:     21.63,
		inventory: 30,
	},
}

func restock() {
	fmt.Println("-RESTOCK BOOK. ")
	fmt.Println("-Please enter book title: ")
	reader := bufio.NewReader(os.Stdin)
	bookTitle, _ := reader.ReadString('\n')
	bookTitle = strings.TrimSuffix(bookTitle, "\n")
	for key, book := range books {
		if book.title == bookTitle {
			fmt.Printf("Add how many inventory to <%v>? :", bookTitle)
			addInventory := 0
			fmt.Scanf("%d", &addInventory)
			if accountBalance-book.price*float64(addInventory) < 0 {
				fmt.Println("-You don't have enough money.\n ")
				return
			}
			balance("-", book.price*float64(addInventory))
			books[key].inventory = book.inventory + addInventory
			return
		}
	}
	fmt.Printf("-Add how many inventory to new book <%v>? :", bookTitle)
	addInventory := 0
	fmt.Scanf("%d", &addInventory)
	fmt.Printf("-What is new book %v's value? :", bookTitle)
	var bookValue float64 = 0.0
	fmt.Scanf("%f", &bookValue)
	if accountBalance-bookValue*float64(addInventory) < 0 {
		fmt.Println("-You don't have enough money.\n ")
		return
	}
	balance("-", bookValue*float64(addInventory))
	newBook := book{
		title:     bookTitle,
		price:     bookValue,
		inventory: addInventory,
	}
	books = append(books, newBook)
}

func sell() {
	fmt.Println("SELL BOOK\n ")
	list()
	fmt.Println("\n-Please enter a book title you want to sell: ")
	reader := bufio.NewReader(os.Stdin)
	bookTitle, _ := reader.ReadString('\n')
	bookTitle = strings.TrimSuffix(bookTitle, "\n")
	for key, book := range books {
		if book.title == bookTitle {
			fmt.Printf("-How many <%v> you want to sell? You have %v in total. :", bookTitle, book.inventory)
			sellNumber := 0
			fmt.Scanf("%d", &sellNumber)
			if sellNumber > book.inventory {
				fmt.Println("-You don't have enough inventory.\n ")
				sell()
			} else if sellNumber == book.inventory {
				balance("+", float64(book.inventory)*float64(sellNumber))

				books = append(books[:key], books[key+1:]...)

			} else {
				balance("+", float64(book.inventory)*float64(sellNumber))
				books[key].inventory = book.inventory - sellNumber
			}
			return
		}
	}
	fmt.Println("-The title you entered is not found in inventory.\n ")
}

func list() {
	fmt.Println("LIST OF BOOK")
	if len(books) < 1 {
		fmt.Println("You have no book in stock\n ")
	} else {
		fmt.Println("Title          \t Inventory\t Price ")
		for _, book := range books {
			fmt.Printf("%-15s\t %-10d\t $ %v\n", book.title, book.inventory, book.price)
		}
	}
}

func balance(operation string, value float64) {
	switch operation {
	case "display":
		var booksTotalValue float64 = 0.0
		for _, book := range books {
			booksTotalValue = booksTotalValue + book.price*float64(book.inventory)
		}
		fmt.Printf("-You have $ %.2f in account\n", accountBalance)
		fmt.Printf("-Your books inventory value is $ %.2f \n", booksTotalValue)
		fmt.Printf("-Your total asset value is $ %.2f \n", accountBalance+booksTotalValue)
	case "-":
		accountBalance = accountBalance - value
		fmt.Printf("-Book RESTOCK successful, you have $ %.2f left in account.\n ", accountBalance)
	case "+":
		accountBalance = accountBalance + value
		fmt.Printf("-Book SOLD successful, you have $ %.2f left in account.\n ", accountBalance)
	}
}

func menu() (string, error) {
	fmt.Println("\n-Please select your operation:\n r: Restock books \n s: Sell books\n l: List books\n b: Balance and Value\n q: Quit\n ")
	reader := bufio.NewReader(os.Stdin)
	operationCode, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("-Fatal error:", err)
		return "error", err
	}
	// remove \n from read string
	operationCode = strings.TrimSuffix(operationCode, "\n")
	if operationCode != "r" && operationCode != "s" && operationCode != "l" && operationCode != "b" && operationCode != "q" {
		fmt.Println("-Wrong input, please select again.\n ")
		//recursive
		return menu()
	}
	return operationCode, nil
}

func main() {
	// _ error return is not used
	opCode, _ := menu()
	for opCode != "q" {
		switch opCode {
		case "r":
			restock()
		case "s":
			sell()
		case "l":
			list()
		case "b":
			balance("display", 0)
		}
		opCode, _ = menu()
	}
}
