package models

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Track struct {
	ID           int        `json:"id"`
	Name         NullString `json:"name"`
	Album        NullString `json:"album"`
	MediaType    NullString `json:"media_type"`
	Genre        NullString `json:"genre"`
	Composer     NullString `json:"composer"`
	Milliseconds int        `json:"milliseconds"`
	Bytes        int        `json:"bytes"`
	UnitPrice    float64    `json:"unit_price"`
}
