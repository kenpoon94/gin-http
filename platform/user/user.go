package user

type Getter interface {
	Get(username string) *User
	GetAll() []User
}

type Adder interface {
	Add(u User)
}

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedOn string `json:"createdOn"`
}

type Users struct {
	Users []User
}

func New() *Users {
	return &Users{
		Users: []User{},
	}
}

func (u *Users) GetAll() []User {
	return u.Users
}

func (u *Users) Add(user User) {
	u.Users = append(u.Users, user)
}

func (u *Users) Get(username string) *User {
	var userFound *User
	for _, user := range u.Users {
		if user.Username == username {
			userFound = &user
		}
	}
	return userFound
}