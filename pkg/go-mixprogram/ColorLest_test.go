package gomixprogram

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	Colors, Err := Colors("37", 2)
	if Err != nil {
		t.Error(Err)
	}

	b := Brand{
		Brand:      "LIFAN",
		ColorsCode: Colors,
	}

	SaveJson("lifan_2.json", ColorCoders{[]Brand{b}})
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
