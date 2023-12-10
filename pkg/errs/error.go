package errs

type Error struct {
	Type    string   `json:"type,omitempty"`
	Code    string   `json:"code,omitempty"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}
