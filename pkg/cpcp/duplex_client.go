package cpcp

type DuplexClient interface {
	Start() error
	Send(line string)
	Receive() <-chan string
	Errors() <-chan error
	Stop() error
}
