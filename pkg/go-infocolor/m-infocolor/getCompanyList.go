package minfocolor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const URL string = "http://m.infocolor.ru"

func GetCompanyList() (Brands []string, Err error) {
	//`{"color_system":"Spies Hecker"}`
	// Делаем запрос на получение категорий
	req, err := http.NewRequest(http.MethodPost, URL+"/formulas/getCompanyList.php",
		bytes.NewReader([]byte(`{"color_system":"Spies Hecker"}`)))
	if err != nil {
		return nil, err
	}

	// Добавляем необходимые атрибуты
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "_ga=GA1.2.1482780989.1699506639; _ga_1NELMSCYY0=GS1.2.1699506641.1.0.1699506641.60.0.0; openerType=browser; BITRIX_SM_LOGIN=Stepice; BITRIX_SM_SOUND_LOGIN_PLAYED=Y; Spies Hecker_discount=0; PHPSESSID=b93a4ad1f478c6034964a62e6212df7d; _ym_isad=1; _ym_visorc=w")
	req.Header.Add("Origin", "http://m.infocolor.ru")
	req.Header.Add("Referer", "http://m.infocolor.ru/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.837 Mobile Safari/537.36")

	// Выполнить запрос
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", res.StatusCode)
	}

	// Читаем ответ в массив байтов
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(bodyBytes))

	// Декодируем полученный json и получаем данные
	err = json.Unmarshal(bodyBytes, &Brands)
	if err != nil {
		return nil, err
	}

	return Brands, nil
}
