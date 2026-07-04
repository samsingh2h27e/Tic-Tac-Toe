package notification

type Observer interface {
	Update (msg string)
}