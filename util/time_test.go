package util

import "testing"

func TestStringToDate(t *testing.T) {
	time := StringToTime("2021-11-27T10:10:10")
	if time.Year() != 2021 {
		t.Errorf("Ano diferente de 2021")
	}

	if time.Month() != 11 {
		t.Errorf("Mes diferente de 11")
	}

	if time.Hour() != 10 {
		t.Errorf("Hora diferente de 10")
	}

	if time.Minute() != 11 {
		t.Errorf("Minutos diferente de 10")
	}
}
