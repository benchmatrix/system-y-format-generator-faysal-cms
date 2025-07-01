package GlobalHolder

var (
	IsReconnecting      = false
	ReconnectionWait    = make(chan string)
	ReconnectSignal     = make(chan string)
	ReconnectionSuccess = make(chan string)
)
