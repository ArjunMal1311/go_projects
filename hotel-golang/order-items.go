package main

import (
	"fmt"
	"strings"
)

type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

var menu = []Menu{
	{1, "Adrakh Chai", 20},
	{2, "Filter Coffee", 25},
	{3, "Chhas", 35.50},
	{4, "Lassi", 30},
	{5, "Mango Lassi", 60},
	{6, "Kadi Chaawal", 50},
	{7, "Chhole Bhature", 45},
	{8, "Khasta Kachori", 30},
	{9, "Raj Kachori", 30},
	{10, "Veg. Sandwich", 20},
	{11, "Veg. Masala Maggi", 60.00},
	{12, "Samosa", 20},
	{13, "Cream Roll", 15},
}

func printMenu() {
	fmt.Printf("%15s \n", menu)
	fmt.Printf("%s \n", strings.Repeat("-", 35))

	fmt.Printf("%-7s %6s    %12s\n", "S.No.", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, element := range menu {
		fmt.Printf("%-7d %.2f   %-4s \n", element.itemNo, element.itemNo, element.itemPrice, element.itemName)
	}

	fmt.Printf("%s", strings.Repeat("-", 35))

	fmt.Println()
}

func orderItems() {
	printMenu()
	var itemNumber uint
	var noOfPlates uint

	for {
		fmt.Println()
		fmt.Println("Enter '0' to exit.")
		fmt.Print("Enter the serial no. of the item to order: ")

		fmt.Scan(&itemNumber)
		if itemNumber == 0 {
			break
		}

		var choiceName string
		var itemPrice float64

		for index, item := range menu {
			if index+1 == int(itemNumber) {
				choiceName = item.itemName
				itemPrice = item.itemPrice
				break
			}
		}
		fmt.Printf("How many %v do you want: ", choiceName)
		fmt.Scan(&noOfPlates)
		if noOfPlates == 0 {
			continue
		} else {
			for index /*, item*/ := range menu {
				if index+1 == int(itemNumber) {
					customerOrder[choiceName] += noOfPlates
					subTotalBill += itemPrice * float64(noOfPlates)

					break
				}
			}
			fmt.Printf("\nYou just ordered %v %v which amounts to ₹%v.\n", noOfPlates, choiceName, itemPrice*float64(noOfPlates))
			orderTillNow()
		}
		fmt.Println()
	}
}

func orderTillNow() {
	fmt.Println("\nYour order till now: ")
	fmt.Printf("%s\n", strings.Repeat("-", 32))
	fmt.Printf(" %-12s %s\n", "Quantity", "Item")
	fmt.Printf("%s\n", strings.Repeat("-", 32))

	for i := range customerOrder {
		fmt.Printf(" %-12v %v\n", customerOrder[i], i)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 32))
}
