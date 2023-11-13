package govseautoemali

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

// https://vseautoemali.ru/cars/chevrolet?sheet=2
//
/*
chevrolet 5
daewoo 3
uaz 2
*/

func TestParse(t *testing.T) {
	Colors, Err := Colors("uaz", 2)
	if Err != nil {
		t.Error(Err)
	}

	b := Brand{
		Brand:      "UAZ",
		ColorsCode: Colors,
	}
	// CHEVROLET

	SaveJson("uaz_1.json", ColorCoders{[]Brand{b}})
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
