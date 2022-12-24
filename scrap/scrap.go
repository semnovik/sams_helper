package scrap

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

func StealUSD(url string) (float64, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	readerUSA, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	USDstr := readerUSA.Find(".chart__info__sum").Text()[3:9]
	USDstr = strings.Replace(USDstr, ",", ".", 1)
	USDfloat, err := strconv.ParseFloat(USDstr, 64)
	if err != nil {
		return 0, err
	}

	return USDfloat, nil
}

func StealEuro(url string) (float64, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	readerEuro, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	EuroStr := readerEuro.Find(".chart__info__sum").Text()[3:9]
	EuroStr = strings.Replace(EuroStr, ",", ".", 1)
	EuroFloat, err := strconv.ParseFloat(EuroStr, 64)
	if err != nil {
		return 0, err
	}
	return EuroFloat, nil
}

func StealKZT(url string) (float64, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	readerKZT, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	kztString := readerKZT.Find(".YMlKec.fxKbKc").Text()
	kztFloat, err := strconv.ParseFloat(kztString, 64)
	if err != nil {
		return 0, err
	}

	return kztFloat, nil
}

func StealAli(url string) (float64, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	readerAli, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	rawAli, _ := readerAli.Find(".font-weight-light").Children().Next().Next().Children().Next().Html()
	strAli := strings.TrimSpace(rawAli)[:5]
	floatAli, err := strconv.ParseFloat(strAli, 64)
	if err != nil {
		return 0, err
	}

	return floatAli, err
}
