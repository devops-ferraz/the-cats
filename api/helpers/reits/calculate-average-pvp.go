package reitsHelpers

import (
	"fmt"
	"strconv"
	"strings"

	reitsDetails "github.com/devops-ferraz/the-cats/api/helpers/reits-details"
)

func CalculatePVP(tickerSymbol string) (map[string]interface{}, error) {
	seriesPatrimLiq, seriesValorMercado, _, err := initializeReitDetails(tickerSymbol)
	if err != nil {
		return nil, err
	}

	if len(seriesPatrimLiq) == 0 {
		return nil, fmt.Errorf("SeriesPatrimLiq is empty")
	}

	var pvps []float64
	for i := range seriesPatrimLiq {
		pvps = append(pvps, float64(seriesValorMercado[i])/float64(seriesPatrimLiq[i]))
	}

	var sum float64
	for _, pvp := range pvps {
		sum += pvp
	}

	currentPvp := pvps[len(pvps)-1]
	averagePVP := sum / float64(len(pvps))
	discount := (1 - currentPvp/averagePVP) * 100

	result := map[string]interface{}{
		"AveragePVP": averagePVP,
		"CurrentPvp": currentPvp,
		"Discount":   strconv.FormatFloat(discount, 'f', 2, 64) + "%",
	}

	return result, nil
}

func initializeReitDetails(tickerSymbol string) ([]int64, []int64, []int64, error) {
	var seriesPatrimLiq []int64
	var seriesValorMercado []int64
	var labels []int64

	switch strings.ToUpper(tickerSymbol) {
	case "XPML11":
		reitsDetails.XpmlDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "HGLG11":
		reitsDetails.HglgDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "ALZR11":
		reitsDetails.AlzrDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "LVBI11":
		reitsDetails.LvbiDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "RZTR11":
		reitsDetails.RztrDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "CPTS11":
		reitsDetails.CptsDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "NEWL11":
		reitsDetails.NewlDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "VISC11":
		reitsDetails.ViscDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "XPLG11":
		reitsDetails.XplgDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "BTLG11":
		reitsDetails.BtlgDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "KNCR11":
		reitsDetails.KncrDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "MXRF11":
		reitsDetails.MxrfDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "RBRF11":
		reitsDetails.RbrfDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "MALL11":
		reitsDetails.MallDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "PVBI11":
		reitsDetails.PvbiDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "KFOF11":
		reitsDetails.KfofDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "KNCA11":
		reitsDetails.KncaDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "RBRY11":
		reitsDetails.RbryDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "KCRE11":
		reitsDetails.KcreDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "KNSC11":
		reitsDetails.KnscDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	case "BBGO11":
		reitsDetails.BbgoDetailsInit(&seriesPatrimLiq, &seriesValorMercado, &labels)
	default:
		return nil, nil, nil, fmt.Errorf("Ticker symbol not found")
	}

	return seriesPatrimLiq, seriesValorMercado, labels, nil
}
