package structs

//ReturnStructs : type of return info
type ReturnStructs struct {
	errorCode  int
	errorMsg   string
	dataStruct interface{}
}

//GetErrorCode : return ReturnStructs 's errorCode
func (r *ReturnStructs) GetErrorCode() int {
	return r.errorCode
}

//SetErrorCode : return ReturnStructs 's errorCode
func (r *ReturnStructs) SetErrorCode(errorCode int) {
	r.errorCode = errorCode
}

//GetErrorMsg : return ReturnStructs 's errorMsg
func (r *ReturnStructs) GetErrorMsg() string {
	return r.errorMsg
}

//SetErrorMsg : return ReturnStructs 's errorMsg
func (r *ReturnStructs) SetErrorMsg(errorMsg string) {
	r.errorMsg = errorMsg
}

//GetDataStruct : return ReturnStructs 's dataStruct
func (r *ReturnStructs) GetDataStruct() interface{} {
	return r.dataStruct
}

//SetDataStruct : return ReturnStructs 's dataStruct
func (r *ReturnStructs) SetDataStruct(dataStruct interface{}) {
	r.dataStruct = dataStruct
}
