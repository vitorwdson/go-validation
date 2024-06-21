package schema_test

import (
	"testing"

	"github.com/vitorwdson/go-validation/schema"
)

func TestValidation(t *testing.T) {
	s := schema.Schema{
		"Name": schema.Required(),
		"Age":  schema.Min(18),
	}

	data := map[string]interface{}{
		"Name": "John Doe",
		"Age":  21,
	}

	errors := s.Validate(data)
	for field, err := range errors {
		t.Errorf("%s: %s", field, err)
	}

	if len(errors) != 0 {
		t.FailNow()
	}
}
