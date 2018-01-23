package shot

type Injector interface {
	Get(from interface{}) interface{}
	GetByKey(key Key) interface{}
	set(key Key, binding filledBinding)
	getBindings() map[Key]filledBinding
}

func newInjector() Injector {
	return &injector{make(map[Key]filledBinding)}
}

type injector struct {
	bindings map[Key]filledBinding
}

func (i *injector) Get(from interface{}) interface{} {
	return i.GetByKey(NewKey(from))
}

func (i *injector) GetByKey(key Key) interface{} {
	binding, ok := i.bindings[key]
	if !ok {
		return nil
	}
	return binding.get()
}

func (i *injector) set(key Key, binding filledBinding) {
	i.bindings[key] = binding
}

func (i *injector) getBindings() map[Key]filledBinding {
	return i.bindings
}