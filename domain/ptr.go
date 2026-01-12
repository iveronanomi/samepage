package domain

func Ptr[C any](value C) *C {
	return &value
}
