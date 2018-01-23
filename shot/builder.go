package shot

type BindingBuilder interface {
	To(implementation interface{}) BindingBuilder
	ToConstructor(constructor interface{}) BindingBuilder
	In(scope Scope)
	AsEagerSingleton()
}

func newLinkedBindingBuilder(binder Binder, key Key) BindingBuilder {
	size := binder.addBinding(newUntargettedBinding(key))
	position := size - 1
	return &linkedBindingBuilder{
		binder:   binder,
		position: position,
	}
}

type linkedBindingBuilder struct {
	binder   Binder
	position int
}

func (builder *linkedBindingBuilder) To(implementation interface{}) BindingBuilder {
	base := builder.getBinding()
	builder.setBinding(newLinkedBinding(base.getKey(), base.getScope(), implementation))
	return builder
}

func (builder *linkedBindingBuilder) ToConstructor(constructor interface{}) BindingBuilder {
	base := builder.getBinding()
	builder.setBinding(newConstructorBinding(base.getKey(), base.getScope(), constructor))
	return builder
}

func (builder *linkedBindingBuilder) In(scope Scope) {
	binding := builder.getBinding().withScope(scope)
	builder.binder.setBinding(builder.position, binding)
}

func (builder *linkedBindingBuilder) AsEagerSingleton() {
	binding := builder.getBinding().withScope(EagerSingleton)
	builder.binder.setBinding(builder.position, binding)
}

func (builder *linkedBindingBuilder) getBinding() binding {
	return builder.binder.getBinding(builder.position)
}

func (builder *linkedBindingBuilder) setBinding(binding binding) {
	builder.binder.setBinding(builder.position, binding)
}
