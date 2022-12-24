package side_methods

import (
	"fmt"
	"sams_helper/scrap"
)

func AllCurrencies() (string, error) {
	usd, err := scrap.StealUSD("https://quote.rbc.ru/ticker/59111")
	if err != nil {
		return "", err
	}
	euro, err := scrap.StealEuro("https://quote.rbc.ru/ticker/59090")
	if err != nil {
		return "", err
	}
	kzt, err := scrap.StealKZT("https://www.google.com/finance/quote/RUB-KZT")
	if err != nil {
		return "", err
	}
	ali, err := scrap.StealAli("https://alicoup.ru/currency/")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("🇺🇸Текущий курс доллара: %.3f₽\n🇪🇺Текущий курс евро: %.3f₽\n🇰🇿Текущий курс тенге: 1₽ = %.3fKZT\n🇨🇳Текущий курс на алиэкспресс: %.3f₽", usd, euro, kzt, ali), nil
}
