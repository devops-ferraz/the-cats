package reitsHelpers

import (
	"math"

	"github.com/devops-ferraz/the-cats/api/structs/reits"
)

func CalculateEvolution(fund reits.RealEstateFund) ([]reits.Evolution, error) {
	currentValue, err := GetCurrentValue(fund.Name)
	if err != nil {
		return nil, err
	}
	lastDividend, err := GetLastDividend(fund.Name)
	if err != nil {
		return nil, err
	}

	var evolution []reits.Evolution
	quantityShares := fund.Quantity
	change := 0.0
	shareValue := currentValue
	dividend := lastDividend

	for month := 1; month <= fund.Months; month++ {
		dividendsReceived := float64(quantityShares) * dividend
		totalAvailable := dividendsReceived + change + fund.MonthlyDeposit
		newShares := int(totalAvailable / shareValue)
		change = totalAvailable - float64(newShares)*shareValue
		quantityShares += newShares

		evolution = append(evolution, reits.Evolution{
			Month:          month,
			QuantityShares: quantityShares,
			Dividends:      math.Round(dividendsReceived*100) / 100,
			Change:         math.Round(change*100) / 100,
			ShareValue:     shareValue,
			CurrentYield:   dividend,
		})
	}
	return evolution, nil
}
