package side_methods

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

//TODO Отредактировать

func GetListOfHolidays() string {
	response, err := http.Get("https://calend.online/holiday/")
	if err != nil {
		log.Fatal(err)
	}
	Body := response.Body

	doc, err := goquery.NewDocumentFromReader(Body)
	if err != nil {
		log.Fatal(err)
	}

	table := doc.Find(".holidays-list")
	header, err := table.Prev().Html()
	if err != nil {
		log.Fatal(err)
	}

	holiday := table.Children()
	listHolys := strings.Split(holiday.Text(), "\n")
	var listFinal []string
	for _, v := range listHolys {
		v = strings.TrimSpace(v)
		if v != "" {
			listFinal = append(listFinal, v)
		}
	}
	holidays := strings.Join(listFinal, "\n")

	return fmt.Sprintf(header + "\n\n" + holidays)
}
