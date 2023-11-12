package gocolorfactory

import "time"

type PageStruct struct {
	Items []struct {
		ID             int    `json:"id"`
		IsThreeCoats   int    `json:"is_three_coats"`
		Alias          string `json:"alias"`
		CarBrandID     int    `json:"car_brand_id"`
		Rgb            string `json:"rgb"`
		ColorCodeAlias string `json:"color_code_alias"`
		ColorCode      struct {
			ID   int    `json:"id"`
			Code string `json:"code"`
		} `json:"color_code"`
		CarBrandName  string   `json:"car_brand_name"`
		CarBrandAlias string   `json:"car_brand_alias"`
		ColorNames    []any    `json:"color_names"`
		ColorCodes    []string `json:"color_codes"`
		CarBrand      struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			NameRu    string `json:"name_ru"`
			Alias     string `json:"alias"`
			IsPopular int    `json:"is_popular"`
			ImgPath   string `json:"imgPath"`
		} `json:"car_brand"`
		Images []any `json:"images"`
		// RgbData struct {
		// 	R int `json:"r"`
		// 	G int `json:"g"`
		// 	B int `json:"b"`
		// } `json:"rgb_data"`
		YearFrom        int      `json:"year_from"`
		YearTo          int      `json:"year_to"`
		ColorNamesClear []string `json:"color_names_clear"`
		ColorGroups     []struct {
			ID                int    `json:"id"`
			Name              string `json:"name"`
			NameRu            string `json:"name_ru"`
			Code              string `json:"code"`
			CodeText          string `json:"code_text"`
			Name2Ru           string `json:"name2_ru"`
			LaravelThroughKey int    `json:"laravel_through_key"`
		} `json:"color_groups"`
		URL string `json:"url"`
	} `json:"items"`
	PerPage  int `json:"per_page"`
	Page     int `json:"page"`
	Total    int `json:"total"`
	CarBrand struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		NameRu    string `json:"name_ru"`
		Alias     string `json:"alias"`
		IsPopular int    `json:"is_popular"`
		ImgPath   string `json:"imgPath"`
	} `json:"car_brand"`
	Post struct {
		ID              int       `json:"id"`
		Alias           string    `json:"alias"`
		Title           string    `json:"title"`
		MetaTitle       string    `json:"meta_title"`
		MetaKeywords    string    `json:"meta_keywords"`
		MetaDescription string    `json:"meta_description"`
		Text            string    `json:"text"`
		TextShort       any       `json:"text_short"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		URL             string    `json:"url"`
	} `json:"post"`
	Breadcrumbs []struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"breadcrumbs"`
	Meta struct {
		Title       string `json:"title"`
		Keywords    string `json:"keywords"`
		Description string `json:"description"`
	} `json:"meta"`
}
