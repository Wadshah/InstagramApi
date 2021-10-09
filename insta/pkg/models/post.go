package models

type Post struct {
	ID        int    `json:"_id"`
	Caption   string `json:"caption"`
	Image_URL string `json:"url"`
	Posted    time   `json:"Date"`
}
