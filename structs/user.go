package structs

//Userinfo : type of User
type Userinfo struct {
	userName     string
	userPassword string
}

//GetUserName : return Userinfo 's userName
func (u *Userinfo) GetUserName() string {
	return u.userName
}

//SetUserName : return Userinfo 's userName
func (u *Userinfo) SetUserName(userName string) {
	u.userName = userName
}

//GetUserPassword : return Userinfo 's userPassword
func (u *Userinfo) GetUserPassword() string {
	return u.userPassword
}

//SetUserPassword : return Userinfo 's userPassword
func (u *Userinfo) SetUserPassword(userPassword string) {
	u.userPassword = userPassword
}
