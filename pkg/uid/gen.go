package uid

type UserIDGen interface {
	Load() error
	Next() int32
}
