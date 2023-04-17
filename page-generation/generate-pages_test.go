package pagegeneration

import "testing"

func TestReadJson(t *testing.T) {
	data := readJson()
	if data == nil {
		t.Errorf("expected %v, got %v", data, nil)
	}
}

func TestGeneratePages(t *testing.T) {
	index, pages := generateIndexAndPages()
	if index == "" {
		t.Errorf("expected %v, got %v", index, nil)
	}
	if pages == nil {
		t.Errorf("expected %v, got %v", pages, nil)
	}
}
