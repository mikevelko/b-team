package main

type createReviewModel struct {
	Content string `json:"content"`
	Rating  int32  `json:"rating"`
}

type createReviewRespondModel struct {
	ReviewID int64 `json:"reviewID"`
}
