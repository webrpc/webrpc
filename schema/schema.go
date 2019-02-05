package schema

// schema of webrpc json file, and validations

type WebRPCSchema struct {
	Schema string `json:"schema"`
	App    string `json:"app"`

	Messages []Message `json:"messages"`
	Services []Service `json:"services"`
}

func (s *WebRPCSchema) Validate() error {
	return nil
}
