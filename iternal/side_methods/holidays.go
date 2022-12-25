package side_methods

import (
	"sams_helper/scrap"
)

func GetListOfHolidays() (string, error) {
	holidays, err := scrap.StealListOfHolidays("https://calend.online/holiday/")
	if err != nil {
		return "", err
	}

	return holidays, err
}
