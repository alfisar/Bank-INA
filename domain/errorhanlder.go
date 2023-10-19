package domain

const (
	FailedControllCode           string = "1101"
	FailedValidationControllCode string = "1102"
	FailedServCode               string = "1201"
	FailedHashingPassCode        string = "1202"
	FailedRepoCode               string = "1301"

	FailedControll           string = "terjadi kesalahan - BINA - 1101"
	FailedValidationControll string = "terjadi kesalahan - BINA - 1102"
	FailedServ               string = "terjadi kesalahan - BINA - 1201"
	FailedHashingPass        string = "terjadi kesalahan - BINA - 1202"
	FailedRepo               string = "terjadi kesalahan - BINA - 1301"
)

type ErrorData struct {
	ResponseCode string      `json:"response_code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}
