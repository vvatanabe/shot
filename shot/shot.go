package shot

type Configure func(binder Binder)

func CreateInjectorIgnoreTag(configures ...Configure) (Injector, error) {
	return newInternalInjectorCreator(false).
		addConfigures(configures...).
		build()
}

func CreateInjector(configures ...Configure) (Injector, error) {
	return newInternalInjectorCreator(true).
		addConfigures(configures...).
		build()
}

func newInternalInjectorCreator(tagOnly bool) *internalInjectorCreator {
	return &internalInjectorCreator{
		binder:     newBinder(),
		configures: []Configure{},
		tagOnly:    tagOnly,
	}
}

type internalInjectorCreator struct {
	binder     Binder
	configures []Configure
	tagOnly    bool
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
		injectedBinding := binding.fill(injector, creator.tagOnly)
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
