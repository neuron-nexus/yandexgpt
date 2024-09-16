package role

var (
	User      = Model{RoleName: "user"}
	Assistant = Model{RoleName: "assistant"}
)

type Model struct {
	RoleName string
}

func (m *Model) String() string {
	return m.RoleName
}
