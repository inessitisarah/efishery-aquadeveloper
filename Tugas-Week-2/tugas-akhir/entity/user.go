package entity

type Product struct {
	ID       int64  `json:"id" gorm:"column:id;type:bignit;primaryKey:autoIncrement"`
	Name     string `json:"product_name"`
	Photo    bytea  `json:"product_photo"`
	Price    string `json:"product_price"`
	Category string `json:"product_price"`
}

type Cart struct {
	ID        int64 `json:"id" gorm:"column:id;type:bignit;primaryKey:autoIncrement"`
	ProductID int   `json:"product_id"`
}
type Payment struct {
	ID     int64 `json:"id" gorm:"column:id;type:bignit;primaryKey:autoIncrement"`
	Photo  bytea `json:"payment_photo"`
	CartID int   `"json:cart_id"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	//Password string `json:"-" column:"password"`
	//Gender   string `json:"gender"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
