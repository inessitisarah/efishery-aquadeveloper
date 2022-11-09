package entity

type User struct {
	ID      int64  `json:"id" gorm:"column:id;type:bignit;primaryKey:autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	//Password string `json:"-" column:"password"`
	//Gender   string `json:"gender"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
