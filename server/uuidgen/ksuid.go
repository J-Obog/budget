package uid

import "github.com/segmentio/ksuid"

type KSUIDProvider struct{}

func NewKSUIDProvider() *KSUIDProvider {
	return &KSUIDProvider{}
}

func (k *KSUIDProvider) GetUUID() string {
	return ksuid.New().String()
}
