package uuid

import "github.com/segmentio/ksuid"

type KsuidProvider struct{}

func NewKsuidProvider() *KsuidProvider {
	return &KsuidProvider{}
}

func (k *KsuidProvider) GetUuid() string {
	return ksuid.New().String()
}
