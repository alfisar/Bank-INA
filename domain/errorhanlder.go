package domain

const (
	FailedRepoCode string = "1301"
	FailedRepo     string = "terjadi kesalahan - BINA - 1301"

	FailedServCode string = "1201"
	FailedServ     string = "terjadi kesalahan - BINA - 1201"
)

type ErrorData struct {
	ResponseCode string      `json:"response_code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}
