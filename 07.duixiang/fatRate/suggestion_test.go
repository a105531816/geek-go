package main

import "testing"

func Test(t *testing.T) {

}

func Test_fatRageSuggestion_GetSuggestion(t *testing.T) {
	sugg := GetFatRateSuggestion()
	tests := []Person{
		{
			sex:     "男",
			age:     25,
			fatRate: 0.24,
		},
	}

	if got := sugg.GetSuggestion(&tests[0]); got != "肥胖" {
		t.Fail()
	}

}
