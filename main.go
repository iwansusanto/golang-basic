package main

import (
	"fmt"

	"golang-basic/helper"
)

func main() {
	// fmt.Println(helper.ExchangeToUSD(120000, 15000, 29, 1.5)) // contoh soal. output 27.333333
	// fmt.Println(helper.ExchangeToUSD(120000, 16000, 29, 1.5)) // output 26.833333
	// fmt.Println(helper.ExchangeToUSD(120000, 15000, 25, 1.5)) // output 24.666666

	// expenseList1 := []int{50000, 30000, 80000, 45000, 72000, 58000, 65000}
	// fmt.Println(helper.MinAndMaxTransaction(expenseList1)) // contoh soal. output 57142,857143 30000 80000

	expenseList2 := []int{80000, 60000, 30000, 45000, 72000, 58000, 20000, 15000}
	fmt.Println(helper.MinAndMaxTransaction(expenseList2)) // output 47500 15000 80000

	expenseList3 := []int{50000, 60000, 10000, 75000, 10000, 75000, 65000, 45000, 11000}
	fmt.Println(helper.MinAndMaxTransaction(expenseList3)) // output 44555,555555 10000 75000
}
