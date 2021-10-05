package AvitoTech

type Users struct {
	Id      int `json:"id" db:"id"`
	Balance int `json:"balance" db:"balance"`
}

type History struct {
	Id         int `json:"id" db:"id"`
	SenderId   int `json:"sender_id" db:"sender_id"`
	ReceiverId int `json:"receiver_id" db:"receiver_id"`
	Amount     int `json:"amount" db:"amount"`
}
