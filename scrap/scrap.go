package scrap

import (
	"fmt"
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

func StealListOfHolidays(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	readerHolidays, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	headerLocator := readerHolidays.Find(".holidays-list")
	header, err := headerLocator.Prev().Html()
	if err != nil {
		return "", err
	}

	rawHolidays := strings.Split(headerLocator.Children().Text(), "\n")

	var listFinal []string
	for _, v := range rawHolidays {
		v = strings.TrimSpace(v)
		if v != "" {
			listFinal = append(listFinal, v)
		}
	}
	holidays := strings.Join(listFinal, "\n")

	return fmt.Sprintf("🎊🎊🎊" + "\n" + header + "\n\n" + holidays), nil
}
