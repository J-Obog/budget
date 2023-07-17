package optionals

import "encoding/json"

//Implement scanner
type Optional[T any] struct {
	val *T
}

func (o *Optional[T]) IsNull() bool {
	return o.val == nil
}

func (o *Optional[T]) IsNotNull() bool {
	return !o.IsNull()
}

func (o *Optional[T]) Get() T {
	return *o.val
}

func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &o.val)
}

func (o *Optional[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.val)
}
