package todo

type User struct {
	Id int `json:"-"`
	// tags "binding required" for binding data from request. Realized in Gin
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
