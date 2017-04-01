package structs

//DataInfo : type of fill in information
type DataInfo struct {
	articleID      int
	infomationList string
}

//GetArticleID : return DataInfo 's articleID
func (d *DataInfo) GetArticleID() *int {
	return &d.articleID
}

//SetArticleID : return DataInfo 's articleID
func (d *DataInfo) SetArticleID(articleID int) {
	d.articleID = articleID
}

//GetInfomationList : return DataInfo 's infomationList
func (d *DataInfo) GetInfomationList() *string {
	return &d.infomationList
}

//SetInfomationList : return DataInfo 's infomationList
func (d *DataInfo) SetInfomationList(infomationList string) {
	d.infomationList = infomationList
}
