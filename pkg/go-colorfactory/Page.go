package gocolorfactory

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const URL string = "https://colorfactory.pro"

// перевести к-во цветов  в к-во страниц в запросе
// Например для 2 в 1, 15 в 2, 100 в 10
func ListsCountOfColors(a int) int {
	b := a / 12
	if a-b*12 != 0 {
		return b + 1
	}
	return b
}
func Pages(brand string) (ColorCodes []string, err error) {
	MaxList := 2
	for iList := 1; iList <= MaxList; iList++ {
		url := fmt.Sprintf("https://colorfactory.pro/json/avtokraski/%s/?page=%d", brand, iList)
		ps, ErrPage := Page(url)
		if ErrPage != nil {
			panic(ErrPage)
		}
		MaxList = ListsCountOfColors(ps.Total) // максимальное к-во страниц
		for i := range ps.Items {
			ColorCodes = append(ColorCodes, ps.Items[i].ColorCodeAlias) // Сохраняем резы по бренду
		}
	}
	return ColorCodes, err
}
func Page(url string) (pp PageStruct, ErrorLine error) {

	// fmt.Println("URLS", fmt.Sprintf(TouchURL, id))

	// Делаем запрос на получение категорий
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return PageStruct{}, err
	}

	// Добавляем необходимые атрибуты
	req.Header.Add("authority", "colorfactory.pro")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "ru,en;q=0.9,lt;q=0.8,it;q=0.7")
	req.Header.Add("cookie", "country=UA; color_views_history=eyJpdiI6IjFFWWxVUi9KeHlGS2ZibmJvNEUwV0E9PSIsInZhbHVlIjoiTjNyQjhta3BHN09vVk1GWmJ5ZWE4Z1kxL2NjTU1wS2FVZW9mZFlRZVRZWUNxck8xRm5nOXBJbnFpL3A1aGtUZkZxTWJGMGZnLzlGQkNMQ0E0VHJJcWc9PSIsIm1hYyI6IjFhMGM3MGE2MjI1YTY5OTNiNzE5ODNjNzY2YjZlZDU1YmM0MjdmOGEwZjNiYzA5NzI4OWZiNzNmNmY4YzA1Y2QifQ%3D%3D; colorfactory_session=eyJpdiI6InZSNFpheEtpcTJWdlZXN0dkWVo2RVE9PSIsInZhbHVlIjoicG5oQUl6dlIxdUhJZTRBMUpSRmh4T3R2Y0hDalVwengxSGh1RmlaZ2tMR3E1SjVRUXExRGFRZ0JxdllBdll3VVp1NERtZEo5cGFyS0sxQndQdExoUmI3Uk9lMDVOVTdVZDFHcm1NU3JrOHJ6WjZXbFcwcEhxTkRFejgxeENyTkwiLCJtYWMiOiIzY2JjNDM2YmJhYmZiYjQyODg4NDk2NDA4MGZlY2IwNGQzZDY1YTBiNDg0MzAxZjU2MDM5MzVlMjc1MmI3ZGM3In0%3D")
	req.Header.Add("referer", "https://colorfactory.pro/avtokraski/acura/?page=2")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"116\", \"Not)A;Brand\";v=\"24\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.837 YaBrowser/23.9.4.837 Yowser/2.5 Safari/537.36")

	// Выполнить запрос
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return PageStruct{}, err
	}

	defer res.Body.Close() // Закрыть ответ в конце выполнения функции

	// В случае положительного результата
	if res.StatusCode != http.StatusOK {
		return PageStruct{}, fmt.Errorf("wrong status code: %d", res.StatusCode)
	}

	// Читаем ответ в массив байтов
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return PageStruct{}, err
	}

	// fmt.Println(string(bodyBytes))

	// Декодируем полученный json и получаем данные
	err = json.Unmarshal(bodyBytes, &pp)
	if err != nil {
		return PageStruct{}, err
	}

	return pp, nil
}
