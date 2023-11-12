package parsecolorfactory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	gocolorfactory "github.com/RB-PRO/infocolor/pkg/go-colorfactory"
)

func Parsing() {
	brands, ErrBrands := gocolorfactory.Brands()
	if ErrBrands != nil {
		panic(ErrBrands)
	}

	var CC ColorCoders
	for _, brand := range brands {
		brand = strings.ReplaceAll(brand, "avtokraski", "")
		brand = strings.ReplaceAll(brand, "/", "")
		fmt.Println(brand)
		ColorCodes, ErrPages := gocolorfactory.Pages(brand)
		if ErrPages != nil {
			panic(ErrPages)
		}
		CC.Brands = append(CC.Brands, Brand{Brand: brand, ColorsCode: ColorCodes})
	}

	SaveJson("colorfactory.json", CC)
}

type ColorCoders struct {
	Brands []Brand
}
type Brand struct {
	Brand      string
	ColorsCode []string
}

func SaveJson(filename string, DataColor ColorCoders) error {

	f, ErrCreateFile := os.Create(filename)
	if ErrCreateFile != nil {
		return ErrCreateFile
	}
	// as_json, ErrMarshalIndent := json.MarshalIndent(variety, "", "\t")
	as_json, ErrMarshalIndent := MarshalMy(DataColor)
	if ErrMarshalIndent != nil {
		return ErrMarshalIndent
	}
	f.Write(as_json)
	f.Close()
	return nil
}

func MarshalMy(i interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(i)
	return bytes.TrimRight(buffer.Bytes(), "\n"), err
}
