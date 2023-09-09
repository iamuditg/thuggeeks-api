package models

type BlogGeeks struct {
	ID               int    `json:"id"`
	Title            string `json:"title" validate:"required"`
	Image            string `json:"image"`
	DescriptionShort string `json:"descriptionShort"`
	DescriptionLong  string `json:"descriptionLong"`
	Tag              string `json:"tag"`
	LikeCount        int    `json:"likeCount"`
}
