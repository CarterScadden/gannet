package services

import "fmt"

const (
	NotFound = iota + 1
	Ok
	Conflict
	BadRequest
)

// getErrorStatus
// NOTE: ONLY USE THIS FOR TESTING
// this is used for testing to get the name of the status for easier debugging
func getErrorStatus(status int) string {
	switch status {
	case NotFound:
		return "NotFound"
	case Ok:
		return "Ok"
	case Conflict:
		return "Conflict"
	case BadRequest:
		return "BadRequest"
	default:
		return fmt.Sprintf("Not a error status: `%d`", status)
	}
}
