package pagegeneration

import (
	"path"
	"testing"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

func TestReadJson(t *testing.T) {
	prjDir := testhelpers.ProjectDir()
	jsonFile := path.Join(prjDir, "page-generation", "modules.json")
	data := readJson(jsonFile)
	if data == nil {
		t.Errorf("expected %v, got %v", data, nil)
	}
}

func TestGeneratePages(t *testing.T) {
	prjDir := testhelpers.ProjectDir()
	jsonFile := path.Join(prjDir, "page-generation", "modules.json")
	index, pages := generateIndexAndPages(jsonFile)
	if index == "" {
		t.Errorf("expected %v, got %v", index, nil)
	}
	if pages == nil {
		t.Errorf("expected %v, got %v", pages, nil)
	}
}

func TestWriteIndexAndPages(t *testing.T) {
	prjDir := testhelpers.ProjectDir()
	jsonFile := path.Join(prjDir, "page-generation", "modules.json")
	outDir := path.Join(prjDir, "out")
	WriteIndexAndPages(jsonFile, outDir, outDir, true)
}
