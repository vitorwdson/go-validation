package schema

import (
	"errors"
	"fmt"
)

type Rule interface {
	Validate(value any) error
}

type required struct{}

func (r *required) Validate(value any) error {
	if value == nil {
		return errors.New("is required")
	}

	str, ok := value.(string)
	if ok && str == "" {
		return errors.New("is required")
	}

	return nil
}

func Required() Rule {
	return &required{}
}

type Number interface {
	int | float64
}

type min[T Number] struct {
	ref T
}

func (m *min[T]) Validate(value any) error {
	if value == nil {
		return nil
	}

	val, ok := value.(T)
	if !ok {
		return errors.New("must be a number")
	}

	if val < m.ref {
		return fmt.Errorf("must be greater than %v", m.ref)
	}

	return nil
}

func Min[T Number](ref T) Rule {
	return &min[T]{ref: ref}
}
