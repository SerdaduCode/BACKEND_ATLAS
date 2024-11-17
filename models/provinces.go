package models

type Province struct {
	ID         string     `json:"id" bson:"_id,omitempty"`
	Name       string     `json:"name" bson:"name"`
	Slug       string     `json:"slug" bson:"slug"`
	LogoUrl    string     `json:"logoUrl" bson:"logoUrl"`
	Coordinate Coordinate `json:"coordinate" bson:"coordinate"`
}

type Coordinate struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}
