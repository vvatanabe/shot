package shot

import "testing"

func Test_Scope(t *testing.T) {
	if NoScope.String() != "NoScope" {
		t.Fatalf("Does not match. result: %s", NoScope.String())
		return
	}
	if SingletonInstance.String() != "SingletonInstance" {
		t.Fatalf("Does not match. result: %s", SingletonInstance.String())
	}
	if EagerSingleton.String() != "EagerSingleton" {
		t.Fatalf("Does not match. result: %s", EagerSingleton.String())
	}
}
