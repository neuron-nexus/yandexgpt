package key

var (
	API_KEY = Type{"Api-Key"}
	Bearer  = Type{"Bearer"}
)

type Type struct {
	KeyType string
}

func (t *Type) String() string {
	return t.KeyType
}
