package schema

// schema of webrpc json file, and validations

type WebRPCSchema struct {
	Schema string `json:"schema"`
	App    string `json:"app"`

	Messages []*Message `json:"messages"`
	Services []*Service `json:"services"`
}

func (s *WebRPCSchema) Validate() error {
	return nil
}

// Parse parses the schema through the AST, intended to be called
// after the json has been unmarshalled
func (s *WebRPCSchema) Parse(schema *WebRPCSchema) error {
	if schema == nil {
		schema = s // we can take nil here, as this is the schema
	}

	for _, msg := range s.Messages {
		for _, field := range msg.Fields {
			err := field.Type.Parse(s)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
