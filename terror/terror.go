package terror

import "errors"

var (
	ErrGetPeerFailed          = errors.New("Get peer context failed")
	ErrClientNoExists         = errors.New("Client does not exists")
	ErrDoNotFindLeader        = errors.New("Do not find leader")
	ErrConnectionDoesNotReady = errors.New("Connection does not ready")
	ErrServerDoesNotReady     = errors.New("Server does not ready")
	ErrNotLeader              = errors.New("It is not leader")
	ErrAddressServer          = errors.New("Error Address Leader")
	ErrBuildExecutorFail      = errors.New("Build executor fail")
	ErrHaveNotTable           = errors.New("Have not table to select row")
	ErrUnique                 = errors.New("Duplicate value")
	ErrRowIsNotExists         = errors.New("Do not find data")
)
