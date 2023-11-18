package tg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	minfocolor "github.com/RB-PRO/infocolor/pkg/go-infocolor/m-infocolor"
)

// Получить список файлов
func (tg *Telegram) FileList() ([]string, error) {
	files, err := os.ReadDir(tg.Folder)
	if err != nil {
		return nil, fmt.Errorf("os.ReadDir(%s): %v", tg.Folder, err)
	}
	FileNames := make([]string, 0, 55)
	for _, f := range files {
		FileName := strings.ReplaceAll(f.Name(), ".json", "")
		FileNames = append(FileNames, FileName)
	}
	return FileNames, nil
}

// Загрузить данные из файла
func (tg *Telegram) LoadFile(filename string) (data []minfocolor.Formulass, ErrorFile error) {
	// Открыть файл
	jsonFile, ErrorFile := os.Open(tg.Folder + filename + ".json")
	if ErrorFile != nil {
		return data, ErrorFile
	}
	defer jsonFile.Close()

	// Прочитать файл и получить массив byte
	jsonData, ErrorFile := io.ReadAll(jsonFile)
	if ErrorFile != nil {
		return data, ErrorFile
	}

	// Распарсить массив byte в структуру
	if ErrorFIle := json.Unmarshal(jsonData, &data); ErrorFIle != nil {
		return data, ErrorFile
	}
	return data, ErrorFile
}
