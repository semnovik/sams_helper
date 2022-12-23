package side_methods

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

//TODO Объединить в один метод + отредактировать

func GetCurrentCurrencyUSD() string {
	responseDollar, err := http.Get("https://quote.rbc.ru/ticker/59111")
	if err != nil {
		log.Fatal(err)
	}

	BodyDollar := responseDollar.Body

	docDollar, err := goquery.NewDocumentFromReader(BodyDollar)
	if err != nil {
		log.Fatal(err)
	}
	currentDollarPrice := docDollar.Find(".chart__info__sum").Text()

	responseEuro, err := http.Get("https://quote.rbc.ru/ticker/59090")
	if err != nil {
		log.Fatal(err)
	}

	BodyEuro := responseEuro.Body

	docEuro, err := goquery.NewDocumentFromReader(BodyEuro)
	if err != nil {
		log.Fatal(err)
	}
	euroPrice := docEuro.Find(".chart__info__sum").Text()

	responseKTG, err := http.Get("https://www.google.com/finance/quote/RUB-KZT")
	if err != nil {
		log.Fatal(err)
	}

	BodyKTG := responseKTG.Body
	docKTG, err := goquery.NewDocumentFromReader(BodyKTG)
	if err != nil {
		log.Fatal(err)
	}
	tengePrice := docKTG.Find(".YMlKec.fxKbKc").Text()

	return "Текущий курс доллара: " + strings.TrimSpace(currentDollarPrice) + "\n" + "Текущий курс евро: " + strings.TrimSpace(euroPrice) + "\n" + "Текущий курс тенге: 1₽ = " + strings.TrimSpace(tengePrice) + " T"
}

func GetAliCurrency() string {
	response, err := http.Get("https://alicoup.ru/currency/")
	if err != nil {
		log.Fatal(err)
	}

	Body := response.Body

	doc, err := goquery.NewDocumentFromReader(Body)
	if err != nil {
		log.Fatal(err)
	}
	dollarAliTable, _ := doc.Find(".font-weight-light").Children().Next().Next().Children().Next().Html()
	aliCurrencyMessage := "Текущий курс USD на Алиэкспресс: " + strings.TrimSpace(dollarAliTable)
	return aliCurrencyMessage
}
