package schema

type Schema map[string]Rule

func (s Schema) Validate(data map[string]interface{}) map[string]error {
	errors := make(map[string]error)

	for key, rule := range s {
		if err := rule.Validate(data[key]); err != nil {
			errors[key] = err
		}
	}

	return errors
}
