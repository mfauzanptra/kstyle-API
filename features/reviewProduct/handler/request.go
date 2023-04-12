package handler

import reviewproduct "kstyleAPI/features/reviewProduct"

type ReviewReq struct {
	IdProduct  uint   `json:"id_product"`
	IdMember   uint   `json:"id_member"`
	DescReview string `json:"desc_review"`
}

type LikeReq struct {
	IdReview uint `json:"id_review"`
	IdMember uint `json:"id_member"`
}

func ToCore(req ReviewReq) *reviewproduct.Core {
	return &reviewproduct.Core{
		IdProduct:  req.IdProduct,
		IdMember:   req.IdMember,
		DescReview: req.DescReview,
	}
}
