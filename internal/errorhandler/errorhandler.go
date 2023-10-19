package errorhandler

import "Bank-INA/domain"

func FailedFromRepo(err error) domain.ErrorData {
	return domain.ErrorData{
		ResponseCode: domain.FailedRepoCode,
		Message:      domain.FailedRepo,
		Data:         err,
	}
}

func FailedFromService(RC string, err error) domain.ErrorData {
	return domain.ErrorData{
		ResponseCode: RC,
		Message:      domain.FailedServ,
		Data:         err,
	}
}
