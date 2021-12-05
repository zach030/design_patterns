package proxy

// User 用户
type User struct {
	Name string
}

// IUser 用户接口
type IUser interface {
	Login(phone int, code int) (*User, error)
	Register(phone int, code int) (*User, error)
}

// IUserFacade 门面模式
type IUserFacade interface {
	LoginOrRegister(phone int, code int) (*User,error)
}

type UserService struct {}

func (u *UserService) Login(phone int, code int)(*User,error) {
	return nil,nil
}

func (u *UserService) Register(phone int, code int)(*User,error) {
	return nil,nil
}

func (u *UserService) LoginOrRegister(phone int, code int)(*User,error) {
	user,err := u.Login(phone,code)
	if err != nil {
		return nil,err
	}
	if user!=nil{
		return user,nil
	}
	return u.Register(phone,code)
}