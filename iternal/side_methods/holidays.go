package side_methods

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

//TODO Отредактировать

func GetListOfHolidays() (string, error) {
	response, err := http.Get("https://calend.online/holiday/")
	if err != nil {
		return "", err
	}

	readerWeather, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	table := readerWeather.Find(".holidays-list")
	header, err := table.Prev().Html()
	if err != nil {
		return "", err
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

	return fmt.Sprintf(header + "\n\n" + holidays), nil
}
