package avalara

import "testing"

func TestGetRates(t *testing.T) {
	client := New("ADD YOUR KEY HERE")
	rate, err := client.GetPostalCodeRate("USA", "10044")
	if err != nil {
		t.Error(err)
	}
	if rate.Amount <= 0 {
		t.Errorf("Invalid rate.")
	}
	return
}
