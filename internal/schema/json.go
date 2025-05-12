package schema

import (
	"encoding/json"
	"fmt"

	"github.com/invopop/jsonschema"
)

// JSONSchemaString generates a JSON schema string from the provided value.
func JSONSchemaString(v any) (string, error) {
	s := jsonschema.Reflect(v)
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON schema: %w", err)
	}

	return string(b), nil
}
