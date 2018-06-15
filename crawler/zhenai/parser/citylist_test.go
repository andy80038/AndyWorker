package parser

import (
	"io/ioutil"
	"testing"
)

func aa() {

}
func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("citylist_text_data.html")
	if err != nil {

	}
	result := ParseCityList(content, "")
	const resultSize = 470
	// expctedUrls := []string{
	// 	"",
	// 	"",
	// 	"",
	// }
	// expctedCitys := []string{
	// 	"",
	// 	"",
	// 	"",
	// }
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests   but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests   but had %d", resultSize, len(result.Items))
	}
}
