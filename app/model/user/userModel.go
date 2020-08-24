package user

// 设置User的表名为`profiles`
func (User) TableName() string {
	return "MS_user_role"
}

type User struct {
	db.Model
	user_id int16
	role_id int16
}
