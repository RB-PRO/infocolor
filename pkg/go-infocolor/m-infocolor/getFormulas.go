package minfocolor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FormulasReq struct {
	ColorSystem   string `json:"color_system"`
	Company       string `json:"company"`
	PaintCode     string `json:"paint_code"`
	Name          string `json:"NAME"`
	WTCode        string `json:"WTCode"`
	CardReference string `json:"CardReference"`
}
type Formulas struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Formulas []struct {
		Company   string `json:"company"`
		PaintCode string `json:"paint_code"`
		Name      string `json:"NAME"`
	} `json:"formulas"`
}

func GetFormulas(company, paintcode string) (formulas []Formulas, Err error) {

	// request := FormulasReq{
	// 	Company: company, PaintCode: paintcode,
	// 	ColorSystem: "Spies Hecker", Name: "", WTCode: "", CardReference: "",
	// }
	request := FormulasReq{
		Company: company, PaintCode: "",
		ColorSystem: "Spies Hecker", Name: paintcode, WTCode: "", CardReference: "",
	}
	BytePayLoad, ErrMarshal := json.Marshal(request)
	if ErrMarshal != nil {
		return nil, fmt.Errorf("json.Marshal: %v", ErrMarshal)
	}

	// Делаем запрос на получение категорий
	url := URL + "/formulas/getFormulas.php"
	req, Err := http.NewRequest(http.MethodPost, url,
		bytes.NewReader(BytePayLoad))
	if Err != nil {
		return nil, Err
	}

	// Добавляем необходимые атрибуты
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "_ga=GA1.2.1482780989.1699506639; _ga_1NELMSCYY0=GS1.2.1699506641.1.0.1699506641.60.0.0; PHPSESSID=4q1r6ijlk9lj3n15fm0t58gvn1; openerType=browser; BITRIX_SM_LOGIN=Stepice; BITRIX_SM_SOUND_LOGIN_PLAYED=Y; Spies Hecker_discount=0; _ym_visorc=w")
	req.Header.Add("Origin", "http://m.infocolor.ru")
	req.Header.Add("Referer", "http://m.infocolor.ru/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.837 Mobile Safari/537.36")

	// Выполнить запрос
	res, Err := http.DefaultClient.Do(req)
	if Err != nil {
		return nil, Err
	}

	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", res.StatusCode)
	}

	// Читаем ответ в массив байтов
	bodyBytes, Err := io.ReadAll(res.Body)
	if Err != nil {
		return nil, Err
	}

	// fmt.Println(string(bodyBytes))

	// Декодируем полученный json и получаем данные
	Err = json.Unmarshal(bodyBytes, &formulas)
	if Err != nil {
		return nil, Err
	}

	return formulas, nil
}
