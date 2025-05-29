package domain

type User struct {
	Id        int64  `json:"id"`
	AccountId int64  `json:"account_id"`
	Gender    string `json:"gender"`
	Name      string `json:"name"`
	Address   string `json:"address"`
}

func (this *User) TableName() string {
	return "users"
}
