package parserinfocolor

import (
	"encoding/json"
	"io"
	"os"
)

// Загрузить данные из файла
func LoadLogPas(filename string) (Login, Password string, ErrorFIle error) {
	// Открыть файл
	jsonFile, ErrorFIle := os.Open(filename)
	if ErrorFIle != nil {
		return "", "", ErrorFIle
	}
	defer jsonFile.Close()

	// Прочитать файл и получить массив byte
	jsonData, ErrorFIle := io.ReadAll(jsonFile)
	if ErrorFIle != nil {
		return "", "", ErrorFIle
	}

	data := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}
	// Распарсить массив byte в структуру
	if ErrorFIle := json.Unmarshal(jsonData, &data); ErrorFIle != nil {
		return "", "", ErrorFIle
	}
	return data.Login, data.Password, ErrorFIle
}
