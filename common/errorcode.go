package common

const (
	//ResultAsNormal : exec normal
	ResultAsNormal = iota
	//TokenFailure : token failure
	TokenFailure
	//ExecSQLFailed : exec sql failed
	ExecSQLFailed
	//NoDataInDB : no data in db
	NoDataInDB
)
