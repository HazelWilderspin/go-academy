package actor

type Response struct {
	Data []byte
	Err  error
}

type Request struct {
	ResponseChannel chan Response
	Action          string
	Args            []byte
}
