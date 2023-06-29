package uid

type UIDProvider interface {
	GetId() string
}
