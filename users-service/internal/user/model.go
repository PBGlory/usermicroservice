package user

type User struct {
	ID       uint32 `gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex"`
	Password string
}
