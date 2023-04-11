package response

type Response interface {
	StatusCode() int
	GetBody() ([]byte, error) //terutn array bytes an error
	Error() string
	GetData() interface{}
}
