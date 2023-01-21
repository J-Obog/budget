package cache

type Cache interface {
	Get(key string) *string
	Set(key string, value string, ttl int64)
}
