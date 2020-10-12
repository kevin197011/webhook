package providers

type Msg interface {
	send(string)
}

func SendMsg(m Msg, message string) {
	m.send(message)
}
