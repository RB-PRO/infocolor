package minfocolor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type FormulasGroupByPaintCodeReq struct {
	ColorSystem                  string `json:"color_system"`
	SectionID                    int    `json:"SECTION_ID"`
	Company                      string `json:"company"`
	PaintCode                    string `json:"paint_code"`
	Name                         string `json:"NAME"`
	WTCode                       string `json:"WTCode"`
	CardReference                string `json:"CardReference"`
	GetFormulasByTypeFormulaList bool   `json:"getFormulasByTypeFormulaList"`
}
type FormulasGroupByPaintCode struct {
	Formulas                  []Formulass `json:"formulas"`
	FormulasByTypeFormulaList []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Selected bool   `json:"selected"`
	} `json:"formulasByTypeFormulaList"`
}

//	type Formulass struct {
//		Type        int    `json:"type"` // Тип этой формлуе офф, уч, кал
//		ID          string `json:"id"`
//		Company     string `json:"company"`
//		PaintCode   string `json:"paint_code"`
//		Name        string `json:"NAME"`
//		CoatOfPaint []any  `json:"coatOfPaint"`
//		Fields      []struct {
//			Text  string `json:"text"`
//			Value string `json:"value"`
//		} `json:"fields"`
//		Components []struct {
//			Code              string  `json:"code"`
//			CodeDisplay       string  `json:"codeDisplay"`
//			Name              string  `json:"name"`
//			Weight            string  `json:"weight"`
//			WeightDisplay     float64 `json:"weightDisplay"`
//			WeightDisplayInit float64 `json:"weightDisplayInit"`
//		} `json:"components"`
//		ComponentsPriceInfo struct {
//			Currency string `json:"currency"`
//			Discount int    `json:"discount"`
//		} `json:"componentsPriceInfo,omitempty"`
//		ComponentsPrice []struct {
//			Name     string `json:"NAME"`
//			Currency string `json:"CURRENCY"`
//			Density  string `json:"DENSITY"`
//			Rate     string `json:"RATE"`
//			Price    string `json:"PRICE"`
//		} `json:"componentsPrice,omitempty"`
//		Commentaries []any  `json:"commentaries"`
//		Quantity     string `json:"quantity"`
//		Thinning     string `json:"thinning"`
//		Unit         string `json:"unit"`
//		UnitSource   string `json:"unitSource"`
//		Price        string `json:"price"`
//		PriceLabel   string `json:"priceLabel"`
//		Sum          string `json:"sum"`
//	}
type Formulass struct {
	Type                int                 `json:"type"`
	ID                  string              `json:"id"`
	Company             string              `json:"company"`
	PaintCode           string              `json:"paint_code"`
	Name                string              `json:"NAME"`
	CoatOfPaint         []any               `json:"coatOfPaint"`
	Fields              []Fields            `json:"fields"`
	Components          []Components        `json:"components"`
	ComponentsPriceInfo ComponentsPriceInfo `json:"componentsPriceInfo"`
	ComponentsPrice     []ComponentsPrice   `json:"componentsPrice,omitempty"`
	Commentaries        []Commentaries      `json:"commentaries"`
	Quantity            string              `json:"quantity"`
	Thinning            string              `json:"thinning"`
	Unit                string              `json:"unit"`
	UnitSource          string              `json:"unitSource"`
	Price               string              `json:"price"`
	PriceLabel          string              `json:"priceLabel"`
	Sum                 string              `json:"sum"`
}
type Fields struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}
type Components struct {
	Code              string  `json:"code"`
	CodeDisplay       string  `json:"codeDisplay"`
	Name              string  `json:"name"`
	Weight            string  `json:"weight"`
	WeightDisplay     float64 `json:"weightDisplay"`
	WeightDisplayInit float64 `json:"weightDisplayInit"`
}

type ComponentsPriceInfo struct {
	Currency string `json:"currency"`
	Discount int    `json:"discount"`
}
type ComponentsPrice struct {
	Name     string `json:"NAME"`
	Currency string `json:"CURRENCY"`
	Density  string `json:"DENSITY"`
	Rate     string `json:"RATE"`
	Price    string `json:"PRICE"`
}
type Commentaries struct {
	Date string `json:"date"`
	Text string `json:"text"`
	User string `json:"user"`
}

func GetFormulasGroupByPaintCode(SectionID int, company, paintcode string) (formula FormulasGroupByPaintCode, Err error) {

	request := FormulasGroupByPaintCodeReq{
		SectionID: SectionID, Company: company, PaintCode: paintcode,
		ColorSystem: "Spies Hecker", Name: "", WTCode: "", CardReference: "", GetFormulasByTypeFormulaList: true,
	}
	// request := FormulasGroupByPaintCodeReq{
	// 	SectionID: SectionID, Company: company, PaintCode: "",
	// 	ColorSystem: "Spies Hecker", Name: paintcode, WTCode: "", CardReference: "", GetFormulasByTypeFormulaList: true,
	// }
	BytePayLoad, ErrMarshal := json.Marshal(request)
	if ErrMarshal != nil {
		return FormulasGroupByPaintCode{}, fmt.Errorf("json.Marshal: %v", ErrMarshal)
	}

	// Делаем запрос на получение категорий
	url := URL + "/formulas/getFormulasGroupByPaintCode.php"
	req, Err := http.NewRequest(http.MethodPost, url,
		bytes.NewReader(BytePayLoad))
	if Err != nil {
		return FormulasGroupByPaintCode{}, fmt.Errorf("http.NewRequest: %v", Err)
	}

	// Добавляем необходимые атрибуты
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("Cookie", "_ga=GA1.2.1482780989.1699506639; _ga_1NELMSCYY0=GS1.2.1699506641.1.0.1699506641.60.0.0; PHPSESSID=4q1r6ijlk9lj3n15fm0t58gvn1; openerType=browser; BITRIX_SM_LOGIN=Stepice; BITRIX_SM_SOUND_LOGIN_PLAYED=Y; Spies Hecker_discount=0; _ym_visorc=w")
	req.Header.Add("Origin", "http://m.infocolor.ru")
	req.Header.Add("Referer", "http://m.infocolor.ru/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.837 Mobile Safari/537.36")

	// Выполнить запрос
	client := &http.Client{Timeout: time.Second * 5}
	res, Err := client.Do(req)
	if Err != nil {
		return FormulasGroupByPaintCode{}, fmt.Errorf("client.Do: %v", Err)
	}

	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return FormulasGroupByPaintCode{}, fmt.Errorf("wrong status code: %d", res.StatusCode)
	}

	// Читаем ответ в массив байтов
	bodyBytes, Err := io.ReadAll(res.Body)
	if Err != nil {
		return FormulasGroupByPaintCode{}, fmt.Errorf("io.ReadAll: %v", Err)
	}

	// fmt.Println(string(bodyBytes))

	// Декодируем полученный json и получаем данные
	Err = json.Unmarshal(bodyBytes, &formula)
	if Err != nil {
		os.WriteFile(fmt.Sprintf("txt/%s_%d.txt", strings.ReplaceAll(company, "/", ""), SectionID), bodyBytes, 0666)
		return FormulasGroupByPaintCode{}, fmt.Errorf("json.Unmarshal: %v", Err)
	}

	for i := range formula.Formulas {
		formula.Formulas[i].Type = SectionID
	}

	return formula, nil
}
