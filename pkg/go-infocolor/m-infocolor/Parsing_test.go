package minfocolor

import (
	"fmt"
	"testing"
)

func TestBrands(t *testing.T) {
	brands, err := GetCompanyList()
	if err != nil {
		t.Error(err)
	}
	if len(brands) == 0 {
		t.Error("len of brands = 0")
	}
}

func TestFormulas(t *testing.T) {
	Formulas, err := GetFormulas("LEXUS", "078")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", Formulas)
}
