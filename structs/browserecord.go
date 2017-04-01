package structs

//BrowseRecord : type of return browserecord
type BrowseRecord struct {
	browseName string
	browseTime string
}

//GetBrowseName : return BrowseRecord 's browseName
func (b *BrowseRecord) GetBrowseName() *string {
	return &b.browseName
}

//SetBrowseName : return BrowseRecord 's browseName
func (b *BrowseRecord) SetBrowseName(browsename string) {
	b.browseName = browsename
}

//GetBrowseTime : return BrowseRecord 's browseTime
func (b *BrowseRecord) GetBrowseTime() *string {
	return &b.browseTime
}

//SetBrowseTime : return BrowseRecord 's browseTime
func (b *BrowseRecord) SetBrowseTime(browsetime string) {
	b.browseTime = browsetime
}
