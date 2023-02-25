package errlog

type Error struct {
	Err  error
	Code int32
	Msg  string
}

func (e *Error) ToResponse() (*int32, *string) {
	if e == nil {
		return new(int32), new(string)
	}
	return &e.Code, &e.Msg
}
