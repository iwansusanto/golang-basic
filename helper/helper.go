package helper

import "sort"

func sum(numbers []int) float64 {

	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}

	return float64(result)
}

// exchangeToUSD, output total in USD
func ExchangeToUSD(amountIDR float64, exchangeRateUSD_IDR float64, amountSGD float64, exchangeRateUSD_SGD float64) float64 {
	// write code here
	idrToUSD := amountIDR / exchangeRateUSD_IDR
	sgToUSD := amountSGD / exchangeRateUSD_SGD
	return idrToUSD + sgToUSD
}

// minAndMaxTransaction, output 3 values, mean, min & max transaction
func MinAndMaxTransaction(expenseList []int) (float64, int, int) {
	// write code here
	sort.Ints(expenseList)
	avgExpend := sum(expenseList) / float64(len(expenseList))

	min := expenseList[0]
	max := expenseList[len(expenseList)-1]
	return avgExpend, min, max
}
