package parserinfocolor

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	goinfocolor "github.com/RB-PRO/infocolor/pkg/go-infocolor"
	minfocolor "github.com/RB-PRO/infocolor/pkg/go-infocolor/m-infocolor"
	"github.com/cheggaaa/pb"
)

func MobileStart() {
	ColorCodes, ErrColorFactory := LoadConfig("colorfactory.json") // colorfactory
	if ErrColorFactory != nil {
		panic(ErrColorFactory)
	}

	Brands, ErrBrands := minfocolor.GetCompanyList()
	if ErrBrands != nil {
		panic(ErrBrands)
	}
	fmt.Printf("Всего цветов в colorfactory - %d, а брендов в infocolor - %d\n",
		len(ColorCodes.Brands), len(Brands))

	iStart := 0
	if len(os.Args) == 2 {
		iStart, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Printf("Начинаем с %d. Первый бренд на рассмотрении - %s\n", iStart, Brands[iStart])
	for iBrand, Brand := range Brands {
		if iBrand < iStart {
			continue
		}
		for _, ColorCode := range ColorCodes.Brands {

			// Если бренды совпали
			Brand = goinfocolor.EditStr(Brand)
			ColorCode.Brand = goinfocolor.EditStr(ColorCode.Brand)
			if (strings.Contains(Brand, ColorCode.Brand) ||
				strings.Contains(ColorCode.Brand, Brand)) &&
				(len(Brand)/2 < len(ColorCode.Brand)) &&
				(len(ColorCode.Brand)/2 < len(Brand)) {
				fmt.Println("Начинаю обработку:", Brand, ColorCode.Brand)

				var FormulasBrand []minfocolor.Formulass
				Bar := pb.StartNew(len(ColorCode.ColorsCode))
				Bar.Prefix(Brand)
				for _, paintcode := range ColorCode.ColorsCode {
					for {
						// formuls, ErrGetFormulas := minfocolor.GetFormulas(Brand, paintcode)
						// if ErrGetFormulas != nil {
						// 	log.Printf("minfocolor.GetFormulas: %v", ErrGetFormulas)
						// 	continue
						// }

						formulIDs := []int{2, 3, 4}
						// for _, formul := range formuls {
						for _, formul := range formulIDs {
							formula, ErrFormula := minfocolor.GetFormulasGroupByPaintCode(formul, Brand, paintcode)
							if ErrFormula != nil {
								log.Printf("minfocolor.GetFormulasGroupByPaintCode: %v", ErrFormula)
							}

							FormulasBrand = append(FormulasBrand, formula.Formulas...)
							time.Sleep(time.Millisecond * 100)
						}
						break
					}
					Bar.Increment()
				}
				Bar.Finish()
				Brand = strings.ReplaceAll(Brand, "/", "")
				if len(FormulasBrand) != 0 {
					FileName := fmt.Sprintf("json/%s.json", Brand)
					fmt.Println("Save file:", FileName)
					minfocolor.SaveJson(FileName, FormulasBrand)
				}
				break
			}
		}
	}

	fmt.Println("Press 'q' to quit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			fmt.Println("Press 'q' to quit")
		}
	}
}

func MobileStart2() {
	Brands, ErrBrands := minfocolor.GetCompanyList()
	if ErrBrands != nil {
		panic(ErrBrands)
	}
	fmt.Printf("брендов в infocolor - %d\n", len(Brands))

	ColorTypes := []int{2, 3, 4}
	iStart := 0
	if len(os.Args) == 2 {
		iStart, _ = strconv.Atoi(os.Args[1])
	}
	Bar := pb.StartNew(len(Brands) - iStart)
	for iBrand, Brand := range Brands {
		if iBrand < iStart {
			continue
		}
		Bar.Prefix(Brand)
		var FormulasBrand []minfocolor.Formulass
		for _, ColorType := range ColorTypes {

			var formula minfocolor.FormulasGroupByPaintCode
			for {
				var ErrFormula error
				formula, ErrFormula = minfocolor.GetFormulasGroupByPaintCode(ColorType, Brand, "")
				if ErrFormula != nil {
					log.Printf("minfocolor.GetFormulasGroupByPaintCode: %v", ErrFormula)
				} else {
					break
				}
			}

			FormulasBrand = append(FormulasBrand, formula.Formulas...)
		}
		Brand = strings.ReplaceAll(Brand, "/", "")
		if len(FormulasBrand) != 0 {
			FileName := fmt.Sprintf("json/%s.json", Brand)
			fmt.Println("Save file:", FileName)
			minfocolor.SaveJson(FileName, FormulasBrand)
		}
		Bar.Increment()
	}
	Bar.Finish()

	fmt.Println("Press 'q' to quit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			fmt.Println("Press 'q' to quit")
		}
	}
}
