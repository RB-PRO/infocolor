package parserinfocolor

import (
	"fmt"

	goinfocolor "github.com/RB-PRO/infocolor/pkg/go-infocolor"
)

func Start() {
	bz, ErrBZ := goinfocolor.NewBrauzer()
	if ErrBZ != nil {
		panic(ErrBZ)
	}
	defer bz.Close()

	ErrAUF := bz.Authorization("Stepice", "Karen1986")
	if ErrAUF != nil {
		panic(ErrAUF)
	}

	Brands, ErrBrand := bz.BrandList("Официальные")
	if ErrBrand != nil {
		panic(ErrBrand)
	}
	fmt.Println("Всего брендов", len(Brands))

	for _, Brand := range Brands {
		cc, lc, ErrPage := bz.ParsePage(Brand, "Официальные", 1)
		if ErrPage != nil {
			panic(ErrPage)
		}
		
		fmt.Println(lc, "Всего данных на странице -", len(cc))
		break
	}
}
