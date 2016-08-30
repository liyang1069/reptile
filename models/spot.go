package models

type Spot struct {
	BaseModel   `bson:",inline" jsonapi:"inline,"`
	Title       string  `jsonapi:"attr,title"`
	Description string  `jsonapi:"attr,description"`
	Longitude   float64 `jsonapi:"attr,longitude`
	Latitude    float64 `jsonapi:"attr,latitude`
}

func NewSpot(title, description string, longitude, latitude float64) *Spot {
	return &Spot{
		Title:       title,
		Description: description,
		Longitude:   longitude,
		Latitude:    latitude,
	}
}
