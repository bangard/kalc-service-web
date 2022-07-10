package models

type PublishContent struct {
	Content []PublishJSON `json:"contents"`
}

type Publishes struct {
	Id        int       `json:"id"`
	Date      string    `json:"date"`
	Days      *[]string `json:"days"`
	DateRange string    `json:"date_range"`
	Hours     string    `json:"hours"`
	Md5       string    `json:"md5"`
	Position  int       `json:"position"`
	TimeToAir string    `json:"time_to_air"`
	Url       string    `json:"url"`
}

type PublishJSON struct {
	Days      *[]string `json:"days"`
	DateRange string    `json:"date_range"`
	Hours     string    `json:"hours"`
	Md5       string    `json:"md5"`
	Position  int       `json:"position"`
	TimeToAir string    `json:"time_to_air"`
	Url       string    `json:"url"`
}

type PublishDate struct {
	Days *[]string `json:"days"`
}

type Info struct {
	DateRange string
}
