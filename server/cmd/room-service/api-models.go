package main

type RoomRequest struct {
	HotelRoomNumber string `json:"HotelRoomNumber"`
}

type RoomRespond struct {
	RoomID int64 `json:"RoomID"`
}
