package product

type ReviewRes struct {
	Username   string `json:"username"`
	Gender     string `json:"gender"`
	SkinType   string `json:"skin_type"`
	SkinColor  string `json:"skin_color"`
	DescReview string `json:"desc_review"`
	Likes      int    `json:"likes"`
}

type ProductRes struct {
	ID          uint        `json:"id"`
	NameProduct string      `json:"name_product"`
	Price       float64     `json:"price"`
	Reviews     []ReviewRes `json:"reviews"`
}
