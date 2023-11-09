package goinfocolor

import (
	"fmt"
	"testing"
)

func TestPage(t *testing.T) {
	bz, ErrBZ := NewBrauzer()
	if ErrBZ != nil {
		t.Error(ErrBZ)
	}
	defer bz.Close()

	ErrAUF := bz.Authorization("Stepice", "Karen1986")
	if ErrAUF != nil {
		t.Error(ErrAUF)
	}

	// 1 станица
	cc, lc, ErrPage := bz.ParsePage("BAIC MOTOR", "Официальные", 1)
	if ErrPage != nil {
		t.Error(ErrPage)
	}
	fmt.Println(lc, "Всего данных на странице -", len(cc))

	// 2 станица
	cc2, lc2, ErrPage2 := bz.ParsePage("BAIC MOTOR", "Официальные", 2)
	if ErrPage2 != nil {
		t.Error(ErrPage2)
	}
	fmt.Println(lc2, "Всего данных на странице -", len(cc2))
}
func TestListsCountOfColors(t *testing.T) {
	fmt.Println(ListsCountOfColors(1))
	fmt.Println(ListsCountOfColors(2))
	fmt.Println(ListsCountOfColors(4))
	fmt.Println(ListsCountOfColors(10))
	fmt.Println(ListsCountOfColors(100))
	fmt.Println(ListsCountOfColors(15))
	fmt.Println(ListsCountOfColors(0))
}