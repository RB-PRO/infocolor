package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	goinfocolor "github.com/RB-PRO/infocolor/pkg/go-infocolor"
	minfocolor "github.com/RB-PRO/infocolor/pkg/go-infocolor/m-infocolor"
)

// Переделать файлы из первых персингов в последние
func FilesEnded() {
	FolderInput := "json/old/"
	FolderOutput := "infocolor2/"

	files, err := ioutil.ReadDir(FolderInput)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filenameReplace := file.Name() // output file
		filenameReplace = strings.ReplaceAll(filenameReplace, "2.json2.json", "")
		filenameReplace = strings.ReplaceAll(filenameReplace, ".json2.json", "")
		filenameReplace = strings.ReplaceAll(filenameReplace, ".json", "")

		FilePatch := fmt.Sprintf(FolderOutput+"%s.json", filenameReplace)
		// fmt.Println(i, FilePatch)

		// read file
		fmt.Println(FolderInput + file.Name())
		data, err := os.ReadFile(FolderInput + file.Name())
		if err != nil {
			panic(err)
		}

		// var DatOut []minfocolor.Formulass

		var DataIn []goinfocolor.Color
		err = json.Unmarshal(data, &DataIn)
		if err != nil {
			panic(err)
		}

		DatOut := web2m(DataIn)

		minfocolor.SaveJson(FilePatch, DatOut)

	}

}

// Перевести веб формат в обычный формат
func web2m(DataIn []goinfocolor.Color) (DataOut []minfocolor.Formulass) {
	for _, in := range DataIn {
		for _, cf := range in.CF {
			for _, rec := range cf.Rec {
				var out minfocolor.Formulass

				out.Company = cf.Info.Brand
				out.PaintCode = cf.Info.Code
				out.Name = cf.Info.Name
				switch in.Type {
				case "Официальные":
					out.Type = 2
				case "Уч. центра":
					out.Type = 3
				case "Колористов":
					out.Type = 4
				}
				if cf.Seria != "" {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Серия",
						Value: cf.Seria,
					})
				}
				if cf.Coverage != "" {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Пократие",
						Value: cf.Coverage,
					})
				}
				if cf.Region != "" {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Регион",
						Value: cf.Region,
					})
				}
				if cf.Shade != "" {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Тень",
						Value: cf.Shade,
					})
				}
				if cf.STD != "" {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "STD",
						Value: cf.STD,
					})
				}
				if cf.Model != "" {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Модель",
						Value: cf.Model,
					})
				}
				if cf.Year != 0 {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Год",
						Value: strconv.Itoa(cf.Year),
					})
				}
				if cf.Manufacturer != "" {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Производитель",
						Value: cf.Manufacturer,
					})
				}
				if cf.Autor != "" {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Автор",
						Value: cf.Autor,
					})
				}
				out.Fields = append(out.Fields, minfocolor.Fields{
					Text:  "Номер слоя",
					Value: strconv.Itoa(rec.LayerNumber),
				})
				if rec.Note != "" {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Примечание",
						Value: rec.Note,
					})
				}
				if rec.Coast != 0.0 {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Примечание",
						Value: fmt.Sprintf("%f", rec.Coast),
					})
				}

				if !cf.Create.IsZero() {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Дата создания",
						Value: cf.Create.Format("02.01.2006"),
					})
				}
				if !cf.Add.IsZero() {
					out.Fields = append(out.Fields, minfocolor.Fields{
						Text:  "Дата создания",
						Value: cf.Add.Format("02.01.2006"),
					})
				}

				// комменты
				for _, comm := range rec.Comments {
					out.Commentaries = append(out.Commentaries, minfocolor.Commentaries{
						Text: comm.Message,
						User: comm.Autor,
						Date: comm.Data.Format("15:04 02.01.2006"),
					})
				}
				for _, f := range rec.Formuls {
					out.Components = append(out.Components, minfocolor.Components{
						// Text: comm.Message,
						// User: comm.Autor,
						// Date: comm.Data.Format("2 Jan 2006 15:04"),
						Code:              f.Code,
						CodeDisplay:       f.Code,
						WeightDisplay:     f.CapWeight,
						WeightDisplayInit: f.CapWeight,
						Weight:            fmt.Sprintf("%.2f", f.CapWeight),
					})

				}

				// save
				DataOut = append(DataOut, out)
			}
		}

	}
	return DataOut
}
