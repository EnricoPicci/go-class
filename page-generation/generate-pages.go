package pagegeneration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type data []map[string]string

func readJson() data {
	c, err := os.ReadFile("./modules.json")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(bytes.NewReader(c))
	var d data
	err = dec.Decode(&d)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func generateIndexAndPages() (indexHtml string, pagesHtml []string) {
	pages := readJson()
	pagesHtml = make([]string, len(pages))
	indexPageListHtml := ""
	for i, pageData := range pages {
		pageName := pageData["name"]
		pageId := pageData["id"]
		indexPageListHtml = indexPageListHtml + fmt.Sprintf(pageListItemTemplate, pageId, pageName)
		pageHtml := fmt.Sprintf(pageTemplate, pageName, pageName, pageId)
		pagesHtml[i] = pageHtml
	}
	indexHtml = fmt.Sprintf(indexTemplate, indexPageListHtml)
	return
}
