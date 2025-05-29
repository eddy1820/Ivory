package setup

type User struct {
	Id   uint   `gorm:"column:id;primary_key;AUTO_INCREMENT;"`
	Name string `gorm:"column:name;"`
	Age  int    `gorm:"column:age;"`
}
