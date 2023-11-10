package parserinfocolor

import (
	"fmt"
	"strings"

	goinfocolor "github.com/RB-PRO/infocolor/pkg/go-infocolor"
)

func Start() {
	//
	bz, ErrBZ := goinfocolor.NewBrauzer()
	if ErrBZ != nil {
		panic(ErrBZ)
	}
	defer bz.Close()

	//
	ErrAUF := bz.Authorization("Stepice", "Karen1986")
	if ErrAUF != nil {
		panic(ErrAUF)
	}

	//
	// Type := "Официальные" // "Уч. центра" "Колористов"
	Brands, ErrBrand := bz.BrandList("Официальные")
	if ErrBrand != nil {
		panic(ErrBrand)
	}
	fmt.Println("Всего брендов", len(Brands))

	Types := []string{"Официальные", "Уч. центра", "Колористов"}
	//
	var ColorsAll []goinfocolor.Color
	for _, Type := range Types {
		for iBrand, Brand := range Brands {
			fmt.Print(iBrand, Brand, "")

			var ColorsBrand []goinfocolor.Color //
			MaxList := 2
			for iList := 1; iList <= MaxList; iList++ {
				cc, MaxListPage, ErrPage := bz.ParsePage(Brand, Type, iList)
				if ErrPage != nil {
					panic(ErrPage)
				}
				MaxList = MaxListPage                    // максимальное к-во страниц
				ColorsBrand = append(ColorsBrand, cc...) // Сохраняем резы по бренду
			}
			fmt.Printf("У бренда '%s' для типа '%s' спарсили всего %d цветов\n", Brand, Type, len(ColorsBrand))

			// Парсинг каждой страницы цвета
			for iColorBrand := range ColorsBrand {
				var ErrParseColor error
				ColorsBrand[iColorBrand].CF, ErrParseColor = bz.ParseColor(ColorsBrand[iColorBrand].Link)
				if ErrParseColor != nil {
					panic(ErrParseColor)
				}
			}

			// SAVE
			FileNameBrand := strings.ReplaceAll(Brand, "/", "-")
			goinfocolor.SaveJson("json/"+FileNameBrand+".json", ColorsBrand)
			ColorsAll = append(ColorsAll, ColorsBrand...) // Сохраняем резы
		}
	}
	FileNameType := strings.ReplaceAll("Всё", "/", "-")
	goinfocolor.SaveJson("json/"+FileNameType+".json", ColorsAll)
}
