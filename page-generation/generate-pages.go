package pagegeneration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
)

type classModulesData []map[string]string
type pageData struct {
	pageName string
	pageHtml string
}

func readJson(jsonFile string) classModulesData {
	c, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(bytes.NewReader(c))
	var d classModulesData
	err = dec.Decode(&d)
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func sortModulesByTitle(modules classModulesData) classModulesData {
	sort.Slice(modules, func(i int, j int) bool {
		return modules[i]["name"] < modules[j]["name"]
	})
	return modules
}

func generateIndexAndPages(jsonFile string) (indexHtml string, pagesHtml []pageData) {
	classModules := readJson(jsonFile)
	classModulesOrdered := sortModulesByTitle(classModules)
	pagesHtml = make([]pageData, len(classModulesOrdered))
	indexPageListHtml := ""
	for i, classModule := range classModules {
		pageName := classModule["name"]
		pageId := classModule["id"]
		indexPageListHtml = indexPageListHtml + fmt.Sprintf(pageListItemTemplate, pageName, pageName)
		pageHtml := fmt.Sprintf(pageTemplate, pageName, pageName, pageId)
		pagesHtml[i] = pageData{pageName, pageHtml}
	}
	indexHtml = fmt.Sprintf(indexTemplate, indexPageListHtml)
	return
}

func removeFiles(dir string) {
	dirInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, d := range dirInfo {
		os.RemoveAll(path.Join(dir, d.Name()))
	}
}

func WriteIndexAndPages(modulesJson string, indexHtmlOutDir string, pagesOutDir string, removeOldPages bool) {
	if removeOldPages {
		removeFiles(pagesOutDir)
	}

	indexHtml, pages := generateIndexAndPages(modulesJson)
	indexHtmlFilePath := path.Join(indexHtmlOutDir, "index.html")
	indexHtmlFile, err := os.Create(indexHtmlFilePath)
	if err != nil {
		panic(err)
	}
	defer indexHtmlFile.Close()
	_, err = indexHtmlFile.Write([]byte(indexHtml))
	if err != nil {
		panic(err)
	}
	for _, page := range pages {
		pageHtmlFilePath := path.Join(pagesOutDir, fmt.Sprintf("%v.html", page.pageName))
		pageHtmlFile, err := os.Create(pageHtmlFilePath)
		if err != nil {
			panic(err)
		}
		defer pageHtmlFile.Close()
		_, err = pageHtmlFile.Write([]byte(page.pageHtml))
		if err != nil {
			panic(err)
		}
	}
	log.Printf("index.html generatated in directoy \"%v\"", indexHtmlOutDir)
	log.Printf("Pages generatated in directoy \"%v\"", pagesOutDir)
}
