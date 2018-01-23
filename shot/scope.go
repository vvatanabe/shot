package shot

type Scope int

func (scope Scope) String() string {
	names := [...]string{"NoScope", "SingletonInstance", "EagerSingleton"}
	if scope < NoScope || scope > EagerSingleton {
		return "Unknown"
	}
	return names[scope]
}

const (
	NoScope           Scope = 0
	SingletonInstance Scope = 1
	EagerSingleton    Scope = 2
)
