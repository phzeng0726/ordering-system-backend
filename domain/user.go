package domain

// TODO 修改成真正的User
// autoIncrement 會在insert之後自行塞回model，所以可以直接用Id取lastInsertId
type User struct {
	Id    int    `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	Email string `gorm:"column:store_id;" json:"storeId"`
}
