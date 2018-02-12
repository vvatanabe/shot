package shot

import "sync"

type Binder interface {
	Bind(target interface{}) BindingBuilder
	size() int
	addBinding(binding binding) int
	setBinding(position int, binding binding)
	getBinding(position int) binding
	getBindingAll() []binding
}

func newBinder() Binder {
	return &binder{&sync.Mutex{}, []binding{}}
}

type binder struct {
	mux      *sync.Mutex
	bindings []binding
}

func (binder *binder) Bind(target interface{}) BindingBuilder {
	return newLinkedBindingBuilder(binder, NewKey(target))
}

func (binder *binder) size() int {
	return len(binder.bindings)
}

func (binder *binder) addBinding(binding binding) (size int) {
	binder.mux.Lock()
	defer binder.mux.Unlock()
	binder.bindings = append(binder.bindings, binding)
	size = binder.size()
	return
}

func (binder *binder) setBinding(position int, binding binding) {
	binder.bindings[position] = binding
}

func (binder *binder) getBinding(position int) binding {
	return binder.bindings[position]
}

func (binder *binder) getBindingAll() []binding {
	return binder.bindings
}
