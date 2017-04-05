package structs

//BasicConfig : type of basic config
type BasicConfig struct {
	//userID        int
	mainPageTitle string
	categoryName  []string
	mode          string
	themeColor    string
}

/*
//GetUserID : return BasicConfig 's userID
func (b *BasicConfig) GetUserID() *int {
	return &b.userID
}

//SetUserID : return BasicConfig 's userID
func (b *BasicConfig) SetUserID(userID int) {
	b.userID = userID
}
*/
//GetMainPageTitle : return BasicConfig 's mainPageTitle
func (b *BasicConfig) GetMainPageTitle() *string {
	return &b.mainPageTitle
}

//SetMainPageTitle : return BasicConfig 's mainPageTitle
func (b *BasicConfig) SetMainPageTitle(mainPageTitle string) {
	b.mainPageTitle = mainPageTitle
}

//GetCategoryName : return BasicConfig 's categoryName
func (b *BasicConfig) GetCategoryName() *[]string {
	return &b.categoryName
}

//SetCategoryName : return BasicConfig 's categoryName
func (b *BasicConfig) SetCategoryName(categoryName []string) {
	b.categoryName = categoryName
}

//GetMode : return BasicConfig 's mode
func (b *BasicConfig) GetMode() *string {
	return &b.mode
}

//SetMode : return BasicConfig 's mode
func (b *BasicConfig) SetMode(mode string) {
	b.mode = mode
}

//GetThemeColor : return BasicConfig 's themeColor
func (b *BasicConfig) GetThemeColor() *string {
	return &b.themeColor
}

//SetThemeColor : return BasicConfig 's themeColor
func (b *BasicConfig) SetThemeColor(themeColor string) {
	b.themeColor = themeColor
}
