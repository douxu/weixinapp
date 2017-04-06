package structs

//ReturnStructs : type of return info
type ReturnStructs struct {
	ErrorCode  int
	ErrorMsg   string
	DataStruct interface{}
}

//Insert :func of fill variable
func (r *ReturnStructs) Insert(errcode int, errmsg string, datastruct interface{}) {
	r.ErrorCode = errcode
	r.ErrorMsg = errmsg
	r.DataStruct = datastruct
}
