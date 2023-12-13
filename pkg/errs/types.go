package errs

type Type string

const (
	ValidationType   Type = "validation"
	NotFoundType     Type = "not_found"
	InternalType     Type = "internal"
	BusinessRuleType Type = "business_rule"
	ConflictType     Type = "conflict"
)
