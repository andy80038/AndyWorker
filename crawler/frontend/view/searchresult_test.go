package view

import (
	"os"
	"testing"

	"github.com/andy80038/AndyWorker/crawler/frontend/model"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CteateSearchResultView("template.html")


	out,err:=os.Create("template_test.html")
	page := model.SearchResult{}
	page.Hits = 99999595
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}

}
