package domain

type User struct {
	Id        int64  `json:"id"`
	AccountId int64  `json:"accountId,omitempty" json:"account_id,omitempty"`
	Gender    string `json:"gender,omitempty" json:"gender,omitempty"`
	Name      string `json:"name,omitempty" json:"name,omitempty"`
	Address   string `json:"address,omitempty" json:"address,omitempty"`
}
