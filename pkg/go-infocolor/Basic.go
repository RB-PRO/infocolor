package goinfocolor

import "time"

const URL string = "https://infocolor.ru"

// Стртуктура цвета, которая содержит сведения о:
//   - Марке (только одна)
//   - Формулах (может быть несколько)
type Color struct {
	Info ColorInfo // Информация по цвету
	Note string    // Примечание
	Link string    // Ссылка на цвет
	Type string    // Тип: "Официальные", "Уч. центра", "Колористов"

}

// Информация о цвете, она будет дублироваться как в формуле, так и в цвете
// Однако заполнять должна и там и там по отдельности
type ColorInfo struct {
	Brand     string // Марка авто
	ColorCode string // Код краски
	ColorName string // Название цвета
}

// Формула цвета
// Одновременно содержит возможную информацию по конкретной формуле
// и всё, что с ней связано
type ColorForm struct {
	Info ColorInfo // Информация по цвету

	// Официальные
	Color        string    // Цвет
	Number       string    // Номер панели
	Seria        string    // Серия
	Coverage     string    // Покрытие
	Region       string    // Регион
	Shade        string    // Оттенок
	Create       time.Time // Дата раз-ки формулы
	STD          string    // СТД
	Model        string    // Модель
	Year         int       // Год выпуска
	Manufacturer string    // Производитель

	// Уч. центра
	Add time.Time // Дата добавления формулы
	// Производитель

	// Колористов
	Autor string // Автор формулы
	// Дата добавления формулы

	Coast float64 // Сумма
}

// Структура комментария
type Comment struct {
	Autor   string
	Message string
}

// Формула цвета такая сложность поля связана с тем, что может быть несколько формул
type Formula struct {
}
