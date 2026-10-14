package mail

// Message is a queued email, with the name of its target server.
type Message struct {
	Server    string
	Recipient string
	Subject   string
	Body      string
}
