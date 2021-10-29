package consul_agent

type ConsulAgent interface {
	Register() error
	CreateSession() error
	AcquireSession(string) (bool, error)
	RenewSession() error
	DestroySession() error
	GetAddressByKey(string) (string, error)
	CloseAgent() error
}
