// go build -o ./bin/generatePages ./page-generation/cmd
// ./bin/generatePages -outDir ./pages
// ./bin/generatePages -outDir ./pages -remove

package main

import (
	"flag"
	"log"

	pagegeneration "github.com/EnricoPicci/go-class/page-generation"
)

func main() {
	indexHtmlOutDir := "."
	pagesOutDir := flag.String("outDir", "", "pages output directory")
	modulesJson := flag.String("modulesJson", "./page-generation/modules.json", "json file containing modules data")
	removeOldPages := flag.Bool("remove", false, "remove old pages from pages output directory")
	flag.Parse()

	if *pagesOutDir == "" {
		log.Fatal("You must specify the directory where pages are published")
	}
	pagegeneration.WriteIndexAndPages(*modulesJson, indexHtmlOutDir, *pagesOutDir, *removeOldPages)
}
