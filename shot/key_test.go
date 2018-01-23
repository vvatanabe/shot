package shot

import "testing"

type User interface {
	GetName() string
}

type user struct {
	name string
}

func (u user) GetName() string {
	return u.name
}

func Test_Key(t *testing.T) {

	bindings := make(map[Key]interface{})

	key1 := NewKey(new(User))
	key2 := NewKey(new(User))


	bindings[key1] = &user{"john"}

	if bindings[key1] != bindings[key2] {
		t.Fatalf("Wrongs %v, %v", bindings[key1], bindings[key2])
	}

}