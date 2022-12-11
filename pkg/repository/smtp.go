package repository

type SMTPConfig struct {
	SenderEmail    string
	SenderPassword string
	ReceiverEmail  string
	Host           string
	Port           string
}
