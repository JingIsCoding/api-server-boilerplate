package exceptions

type Error struct {
	Type       string `json:"type"`
	Title      string `json:"title"`
	Message    string `json:"message"`
	StackTrace string `json:"stack_trace"`
	Wrapped    error  `json:"-"`
}

var (
	DatbaseConnectionFailed Error = NewError("database.connecton.failed")
	RedisConnectionFailed   Error = NewError("redis.connecton.failed")
	HttpConnectionFailed    Error = NewError("http.connecton.failed")

	UUIDParseFailed Error = NewError("uuid.parse.failed")

	AuthFailed   Error = NewError("auth.failed")
	InvalidToken Error = NewError("auth.token.invalid")
	TokenExpires Error = NewError("auth.token.expires")

	UserCreateFaled  Error = NewError("user.create.failed")
	UserAlreadyExist Error = NewError("user.already.exists")
	UserNotExists    Error = NewError("user.not.exists")

	Unknown Error = NewError("unknown")
)

func (err Error) Error() string {
	return err.Title
}

func (e Error) Is(other error) bool {
	if err, ok := other.(Error); !ok {
		return false
	} else {
		return e.Type == err.Type
	}
}

func (e Error) Wrap(another error) Error {
	e.Wrapped = another
	return e
}

func (e Error) SetMessage(msg string) Error {
	e.Message = msg
	return e
}

func (e Error) Unwrap() error {
	return e.Wrapped
}

func NewError(errType string) Error {
	return Error{
		Type:  errType,
		Title: errType,
	}
}
