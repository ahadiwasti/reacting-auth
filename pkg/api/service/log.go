package service

import (
	// "github.com/spf13/viper"
	"time"

	"./pkg/api/dao"
	"./pkg/api/dto"
	"./pkg/api/model"
	// "fmt"
)

type LoginLog = model.LoginLog
type OperationLog = model.OperationLog

var loginLogDao = dao.LoginLogDao{}
var operationLogDao = dao.OperationLogDao{}

// LogService
type LogService struct {
}

// LoginLogDetail log login detail
func (LogService) LoginLogDetail(dto dto.GeneralGetDto) LoginLog {
	return loginLogDao.Detail(dto.Id)
}

// List - users list with pagination
func (LogService) LoginLogLists(dto dto.LoginLogListDto) ([]dao.LoginLogList, int64) {
	return loginLogDao.Lists(dto)
}

func (LogService) OperationLogDetail(dto dto.GeneralGetDto) OperationLog {
	return operationLogDao.Detail(dto.Id)
}

// List - users list with pagination
func (LogService) OperationLogLists(dto dto.OperationLogListDto) ([]dao.OperationLogList, int64) {
	return operationLogDao.Lists(dto)
}

//Insert Operation Log
func (LogService) InsertOperationLog(orLogDto dto.OperationLogDto) error {
	return operationLogDao.Create(orLogDto)
}

//Insert Operation Log
//func (LogService) DeleteLatestPwdUpdate(gDto dto.OperationLogDto) error {
//	//return operationLogDao.Create(orLogDto)
//}

// CheckIdleTooLong check duration between now and  last action time
// true - means too long time user not doing anything,we should kick user out of admin pages
func (LogService) CheckAccountIdleTooLong(uDto dto.GeneralGetDto) bool {
	// if viper.GetInt("security.level") == 0 {
	// 	return false
	// }
	// pick the latest access record of account
	// then judge if it pass over 1 hour
	oLog := operationLogDao.GetLatestLogOfAccount(uDto.Id)
	// fmt.Println( "time now=>", time.Now().Sub(oLog.CreateTime))
	// fmt.Println( "time now=>", time.Now().Sub(oLog.CreateTime).Minutes())
	// fmt.Println( "time now=>", time.Now().Sub(oLog.CreateTime).Seconds())
	// fmt.Println( "time now=>", time.Now().Sub(oLog.CreateTime).Milliseconds())
	// fmt.Println( "oLog OperationTime=>", oLog)
	// fmt.Println( "oLog=>", oLog.CreateTime)
	// fmt.Println("time now=>", time.Now())
	// fmt.Println("time duration=>", time.Now().Sub(oLog.CreateTime).Minutes())
	// tn:=time.Now().Format(time.RFC3339)
	// tnepco:= tn.Unix()
	// lct:=oLog.CreateTime.Format(time.RFC3339)
	// lctepoc:= lct.Unix()

	// fmt.Println("time now in epoc", tnepco)
	// fmt.Println("create now in epoc", tnepco)
	if oLog.Id < 1 || time.Now().Sub(oLog.CreateTime).Minutes() > 20 {
		return true
	} else {
		return false
	}
}

// oLog=> 2021-08-30 20:47:31 +0400 +04
// time now=> 2021-08-30 20:47:37.379171594 +0400 +04 m=+223.184401563
// time duration=> 6379

// time viper=> 1140000
// oLog=> 2021-08-30 16:47:27 +0400 +04
// time now=> 2021-08-30 20:47:33.438730534 +0400 +04 m=+29.763666537
// time duration=> 14406438
