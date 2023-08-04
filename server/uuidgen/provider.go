package uuid

type UuidProvider interface {
	GetUuid() string
}
