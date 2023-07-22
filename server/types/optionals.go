package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Optional[T any] struct {
	val *T
}

func New[T any](v any) Optional[T] {
	if v == nil {
		return Optional[T]{
			val: nil,
		}
	}

	vObj, ok := v.(T)

	if !ok {
		panic(errors.New("type mismatch"))
	}

	vN := new(T)
	*vN = vObj

	return Optional[T]{
		val: vN,
	}
}

func OptionalOf[T any](v any) Optional[T] {
	return New[T](v)
}

func (o *Optional[T]) Value() (driver.Value, error) {
	if o.Empty() {
		return nil, nil
	}

	return o.Get(), nil
}

func (o *Optional[T]) Scan(data any) error {
	*o = New[T](data)
	return nil
}

func (o *Optional[T]) Empty() bool {
	return o.val == nil
}

func (o *Optional[T]) NotEmpty() bool {
	return !o.Empty()
}

func (o *Optional[T]) Get() T {
	return *o.val
}

func (o *Optional[T]) GetOr(fallback T) T {
	if o.val == nil {
		return fallback
	}

	return o.Get()
}

func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &o.val)
}

func (o *Optional[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.val)
}

func OptionalBool(v any) Optional[bool] {
	return New[bool](v)
}

func OptionalString(v any) Optional[string] {
	return New[string](v)
}

func OptionalInt64(v any) Optional[int64] {
	return New[int64](v)
}

func OptionalInt(v any) Optional[int] {
	return New[int](v)
}

func OptionalFloat64(v any) Optional[float64] {
	return New[float64](v)
}

func OptionalUint(v any) Optional[uint] {
	return New[uint](v)
}
