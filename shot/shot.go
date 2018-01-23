package shot

type Configure func(binder Binder)

func CreateInjector(configures ...Configure) (Injector, error) {
	return newInternalInjectorCreator().
		addConfigures(configures...).
		build()
}

func newInternalInjectorCreator() *internalInjectorCreator {
	return &internalInjectorCreator{
		binder:     newBinder(),
		configures: []Configure{},
	}
}

type internalInjectorCreator struct {
	binder     Binder
	configures []Configure
}

func (creator *internalInjectorCreator) addConfigures(configures ...Configure) *internalInjectorCreator {
	creator.configures = append(creator.configures, configures...)
	return creator
}

func (creator *internalInjectorCreator) build() (Injector, error) {

	for _, configure := range creator.configures {
		configure(creator.binder)
	}

	injector := newInjector()

	for _, binding := range creator.binder.getBindingAll() {
		injectedBinding := binding.fill(injector)
		injector.set(binding.getKey(), injectedBinding)
	}

	for _, binding := range injector.getBindings() {
		if err := binding.ok(); err != nil {
			return nil, err
		}
	}

	loadEagerSingletons(injector)

	return injector, nil
}

func loadEagerSingletons(injector Injector) {
	for _, binding := range injector.getBindings() {
		if _, ok := binding.(*eagerSingletonBinding); ok {
			binding.get()
		}
	}
}
