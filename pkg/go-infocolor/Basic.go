package goinfocolor

import "time"

const URL string = "https://infocolor.ru"

// Информация о цвете, она будет дублироваться как в формуле, так и в цвете
// Однако заполнять должна и там и там по отдельности
type ColorInfo struct {
	Brand string // Марка авто
	Code  string // Код краски
	Name  string // Название цвета
}

// Стртуктура цвета, которая содержит сведения о:
//   - Марке (только одна)
//   - Формулах (может быть несколько)
type Color struct {
	Info ColorInfo // Информация по цвету
	Note string    // Примечание
	Link string    // Ссылка на цвет
	Type string    // Тип: "Официальные", "Уч. центра", "Колористов"
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

	Comments []Comment
}

// Комментарий к ColorForm
type Comment struct {
	Autor   string
	Message string
	Data    time.Time
}

// Рецепт цвета, может иметь несколько слоёв
type Recipe struct {
	LayerNumber int       // Номер слоя
	Formuls     []Formula // Формулы
	Note        string    // Примечание
	Coast       float64   // Сумма
}

// Формула цвета такая сложность поля связана с тем, что может быть несколько формул
type Formula struct {
	Code   string // Код компонента
	Name   string // Название компонента
	Weight string // Кол-во в (г), чтобы получился выбранный объем
}
