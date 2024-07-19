package reitsHelpers

import (
	"fmt"
	"log"
	"time"

	"github.com/xuri/excelize/v2"
)

type UpdateResponse struct {
	Success []string `json:"success"`
	Failed  []string `json:"failed"`
}

var knownStockTickers = map[string]bool{
	"BBAS3":  true,
	"BBSE3":  true,
	"ITUB4":  true,
	"AURE3":  true,
	"ITSA4":  true,
	"TAEE4":  true,
	"CXSE3":  true,
	"SAPR4":  true,
	"BITH11": true,
	"NASD11": true,
	"IVVB11": true,
	"TECK11": true,
	"NUBANK": true,
	"SOFISA": true,
}

func UpdateDataReits(filePath string, sheetName string) UpdateResponse {
	var response UpdateResponse

	tickers, err := getTickersFromSheet(sheetName, filePath)
	if err != nil {
		log.Fatalf("Error reading tickers: %v", err)
	}

	for _, ticker := range tickers {
		if ticker == "NUBANK" {
			break
		}
		quote, err := GetCurrentValue(ticker)
		if err != nil {
			log.Printf("Erro ao obter o valor atual para %s: %v", ticker, err)
			response.Failed = append(response.Failed, ticker)
			continue
		}

		var dividend float64
		// Verifica se o ticker é uma ação conhecida, se não for, obtém o dividendo
		if !knownStockTickers[ticker] {
			dividend, err = GetLastDividend(ticker)
			if err != nil {
				log.Printf("Erro ao obter o último dividendo para %s: %v", ticker, err)
				response.Failed = append(response.Failed, ticker)
				continue
			}
		}

		tickersToUpdate := map[string][2]float64{
			ticker: {quote, dividend},
		}
		err = updateProventosAndCotacoes(filePath, tickersToUpdate, sheetName, knownStockTickers[ticker])
		if err != nil {
			log.Printf("Erro ao atualizar proventos e cotações para %s: %v", ticker, err)
			response.Failed = append(response.Failed, ticker)
			continue
		}

		response.Success = append(response.Success, ticker)
		fmt.Printf("Ticker: %s, Quote: %.2f, Dividend: %.2f\n", ticker, quote, dividend)
	}

	return response
}

func updateProventosAndCotacoes(filePath string, tickers map[string][2]float64, sheetName string, isStockSection bool) error {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return fmt.Errorf("falha ao abrir o arquivo Excel: %w", err)
	}
	defer f.Close()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("falha ao obter as linhas da planilha: %w", err)
	}

	for ticker, values := range tickers {
		quote := values[0]
		dividend := values[1]
		tickerFound := false
		for i, row := range rows {
			if len(row) >= 4 && row[3] == ticker {
				cellQuote := fmt.Sprintf("E%d", i+1)
				err = f.SetCellValue(sheetName, cellQuote, quote)
				if err != nil {
					return fmt.Errorf("falha ao definir o novo valor de cotação: %w", err)
				}
				if !isStockSection {
					cellDividend := fmt.Sprintf("F%d", i+1)
					err = f.SetCellValue(sheetName, cellDividend, dividend)
					if err != nil {
						return fmt.Errorf("falha ao definir o novo valor de provento: %w", err)
					}
				}
				tickerFound = true
				break
			}
		}
		if !tickerFound {
			fmt.Printf("Ticker %s não encontrado na planilha %s\n", ticker, sheetName)
		}
	}

	maxAttempts := 3
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err = f.SaveAs(filePath)
		if err == nil {
			break
		}
		fmt.Printf("Tentativa %d de salvar o arquivo falhou: %v\n", attempt, err)
		if attempt < maxAttempts {
			fmt.Println("Aguardando 1 segundo antes de tentar novamente...")
			time.Sleep(time.Second)
		}
	}

	if err != nil {
		return fmt.Errorf("falha ao salvar o arquivo Excel após %d tentativas: %w", maxAttempts, err)
	}

	return nil
}

func getTickersFromSheet(sheetName string, filePath string) ([]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	var tickers []string
	for _, row := range rows {
		if len(row) > 3 {
			ticker := row[3]
			if ticker != "" && ticker != "TICKER" {
				tickers = append(tickers, ticker)
			}
		}
	}
	return tickers, nil
}
