package models

type BlogGeeks struct {
	ID               int    `json:"id"`
	Title            string `json:"title" validate:"required"`
	Image            string `json:"image"`
	DescriptionShort string `json:"description_short"`
	DescriptionLong  string `json:"description_long"`
	Tag              string `json:"tag"`
	LikeCount        int    `json:"likeCount"`
}
