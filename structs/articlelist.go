package structs

//ArtcileList : type of ArtcileList
type ArtcileList struct {
	artcileTitle   string
	pictureURL     string
	articleTraffic int
	dataNum        int
	operatorStatus bool
}

//GetArticleTitle : return ArtcileList 's artcileTitle
func (a *ArtcileList) GetArticleTitle() *string {
	return &a.artcileTitle
}

//SetArticleTitle : return ArtcileList 's artcileTitle
func (a *ArtcileList) SetArticleTitle(artcileTitle string) {
	a.artcileTitle = artcileTitle
}

//GetPictureURL : return ArtcileList 's pictureURL
func (a *ArtcileList) GetPictureURL() *string {
	return &a.pictureURL
}

//SetPictureURL : return ArtcileList 's pictureURL
func (a *ArtcileList) SetPictureURL(pictureURL string) {
	a.pictureURL = pictureURL
}

//GetDataNum : return ArtcileList 's dataNum
func (a *ArtcileList) GetDataNum() *int {
	return &a.dataNum
}

//SetDataNum : return ArtcileList 's dataNum
func (a *ArtcileList) SetDataNum(dataNum int) {
	a.dataNum = dataNum
}

//GetArticleTraffic : return ArtcileList 's articletraffic
func (a *ArtcileList) GetArticleTraffic() *int {
	return &a.articleTraffic
}

//SetArticleTraffic : return ArtcileList 's articletraffic
func (a *ArtcileList) SetArticleTraffic(articleTraffic int) {
	a.articleTraffic = articleTraffic
}

//GetOperatorStatus : return ArtcileList 's operatorStatus
func (a *ArtcileList) GetOperatorStatus() *bool {
	return &a.operatorStatus
}

//SetOperatorStatus : return ArtcileList 's operatorStatus
func (a *ArtcileList) SetOperatorStatus(operatorStatus bool) {
	a.operatorStatus = operatorStatus
}
