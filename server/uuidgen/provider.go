package uid

type UUIDProvider interface {
	GetUUID() string
}
