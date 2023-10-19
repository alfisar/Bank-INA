package errorhandler

import "Bank-INA/domain"

func FailedFromRepo(err error) domain.ErrorData {
	return domain.ErrorData{
		ResponseCode: domain.FailedRepoCode,
		Message:      domain.FailedRepo,
		Data:         err,
	}
}

func FailedFirstFromRepo(err error) domain.ErrorData {
	return domain.ErrorData{
		ResponseCode: domain.FailedRepoCode,
		Message:      domain.FailedRepo,
		Data:         err.Error(),
	}
}

func FailedFromService(RC string, err error) domain.ErrorData {
	return domain.ErrorData{
		ResponseCode: RC,
		Message:      domain.FailedServ,
		Data:         err,
	}
}

func Failed(RC string, message string, err error) domain.ErrorData {
	return domain.ErrorData{
		ResponseCode: RC,
		Message:      message,
		Data:         err,
	}
}
