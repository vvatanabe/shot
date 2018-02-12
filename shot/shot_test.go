package shot

import (
	"testing"
)

func NewProjectService(userRepository UserRepository, groupRepository GroupRepository) *ProjectService {
	return &ProjectService{
		userRepository, groupRepository,
	}
}

type ProjectService struct {
	UserRepository  UserRepository  `inject:""`
	GroupRepository GroupRepository `inject:""`
}

func (u *ProjectService) FindUser() []string {
	return u.UserRepository.FindAll()
}

func (u *ProjectService) FindGroup() []string {
	return u.GroupRepository.FindAll()
}

func NewStoreOnMemory() *StoreOnMemory {
	return &StoreOnMemory{
		[]string{ "user-1", "user-2", "user-3" },
		[]string{ "group-1", "group-2", "group-3" },
	}
}

type Store interface {
	GetUsers() []string
	GetGroups() []string
}

type StoreOnMemory struct {
	users []string
	groups []string
}

func (s *StoreOnMemory) GetUsers() []string {
	if s.users == nil {
		s.users = []string{}
	}
	return s.users
}

func (s *StoreOnMemory) GetGroups() []string {
	if s.groups == nil {
		s.groups = []string{}
	}
	return s.groups
}

type UserRepository interface {
	FindAll() []string
}

func NewUserRepositoryOnMemory(store Store) *UserRepositoryOnMemory {
	return &UserRepositoryOnMemory{store}
}

type UserRepositoryOnMemory struct {
	Store Store  `inject:""`
}

func (repository *UserRepositoryOnMemory) FindAll() []string {
	return repository.Store.GetUsers()
}

type GroupRepository interface {
	FindAll() []string
}

func NewGroupRepositoryOnMemory(store Store) *GroupRepositoryOnMemory {
	return &GroupRepositoryOnMemory{store}
}

type GroupRepositoryOnMemory struct {
	Store Store  `inject:""`
}

func (repository *GroupRepositoryOnMemory) FindAll() []string {
	return repository.Store.GetGroups()
}

func Test_it_should_be_inject_implementation_into_the_interface_via_struct(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(Store)).To(new(StoreOnMemory))
		binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory))
		binder.Bind(new(GroupRepository)).To(GroupRepositoryOnMemory{})
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
}

func Test_it_should_be_inject_implementation_into_the_interface_as_singleton_via_struct(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(Store)).To(new(StoreOnMemory)).In(SingletonInstance)
		binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).In(SingletonInstance)
		binder.Bind(new(GroupRepository)).To(GroupRepositoryOnMemory{}).In(SingletonInstance)
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
}

func Test_it_should_be_inject_implementation_into_the_interface_as_eager_singleton_via_struct(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(Store)).To(new(StoreOnMemory)).AsEagerSingleton()
		binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).AsEagerSingleton()
		binder.Bind(new(GroupRepository)).To(GroupRepositoryOnMemory{}).AsEagerSingleton()
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
}


// ------------------------------

func Test_it_should_be_inject_implementation_into_the_interface_via_constructor(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(Store)).ToConstructor(NewStoreOnMemory)
		binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory)
		binder.Bind(new(GroupRepository)).ToConstructor(NewGroupRepositoryOnMemory)
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
}

func Test_it_should_be_inject_implementation_into_the_interface_as_singleton_via_constructor(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(Store)).ToConstructor(NewStoreOnMemory).In(SingletonInstance)
		binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory).In(SingletonInstance)
		binder.Bind(new(GroupRepository)).ToConstructor(NewGroupRepositoryOnMemory).In(SingletonInstance)
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
}

func Test_it_should_be_inject_implementation_into_the_interface_as_eager_singleton_via_constructor(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(Store)).ToConstructor(NewStoreOnMemory).AsEagerSingleton()
		binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory).AsEagerSingleton()
		binder.Bind(new(GroupRepository)).ToConstructor(NewGroupRepositoryOnMemory).AsEagerSingleton()
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
}

// ---------------------

func Test_it_should_be_inject_struct_into_the_struct_via_struct(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(ProjectService)).To(new(ProjectService))
		binder.Bind(new(Store)).To(new(StoreOnMemory))
		binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory))
		binder.Bind(new(GroupRepository)).To(GroupRepositoryOnMemory{})
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
	projectService := injector.Get(new(ProjectService)).(*ProjectService)
	if projectService == nil {
		t.Fatal("not found a GroupRepository")
	}
	if projectService.FindUser() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
	if projectService.FindGroup() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
}

func Test_it_should_be_inject_struct_into_the_struct_as_singleton_via_struct(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(ProjectService)).To(new(ProjectService)).In(SingletonInstance)
		binder.Bind(new(Store)).To(new(StoreOnMemory)).In(SingletonInstance)
		binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).In(SingletonInstance)
		binder.Bind(new(GroupRepository)).To(GroupRepositoryOnMemory{}).In(SingletonInstance)
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
	projectService := injector.Get(new(ProjectService)).(*ProjectService)
	if projectService == nil {
		t.Fatal("not found a GroupRepository")
	}
	if projectService.FindUser() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
	if projectService.FindGroup() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
}

func Test_it_should_be_inject_struct_into_the_struct_as_eager_struct_via_struct(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(ProjectService)).To(new(ProjectService)).AsEagerSingleton()
		binder.Bind(new(Store)).To(new(StoreOnMemory)).AsEagerSingleton()
		binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).AsEagerSingleton()
		binder.Bind(new(GroupRepository)).To(GroupRepositoryOnMemory{}).AsEagerSingleton()
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
	projectService := injector.Get(new(ProjectService)).(*ProjectService)
	if projectService == nil {
		t.Fatal("not found a GroupRepository")
	}
	if projectService.FindUser() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
	if projectService.FindGroup() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
}

// ------------------------------

func Test_it_should_be_inject_struct_into_the_struct_via_constructor(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(ProjectService)).ToConstructor(NewProjectService)
		binder.Bind(new(Store)).ToConstructor(NewStoreOnMemory)
		binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory)
		binder.Bind(new(GroupRepository)).ToConstructor(NewGroupRepositoryOnMemory)
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
	projectService := injector.Get(new(ProjectService)).(*ProjectService)
	if projectService == nil {
		t.Fatal("not found a GroupRepository")
	}
	if projectService.FindUser() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
	if projectService.FindGroup() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
}

func Test_it_should_be_inject_struct_into_the_struct_as_singleton_via_constructor(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(ProjectService)).ToConstructor(NewProjectService).In(SingletonInstance)
		binder.Bind(new(Store)).ToConstructor(NewStoreOnMemory).In(SingletonInstance)
		binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory).In(SingletonInstance)
		binder.Bind(new(GroupRepository)).ToConstructor(NewGroupRepositoryOnMemory).In(SingletonInstance)
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
	projectService := injector.Get(new(ProjectService)).(*ProjectService)
	if projectService == nil {
		t.Fatal("not found a GroupRepository")
	}
	if projectService.FindUser() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
	if projectService.FindGroup() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
}

func Test_it_should_be_inject_struct_into_the_struct_as_eager_singleton_via_constructor(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(ProjectService)).ToConstructor(NewProjectService).AsEagerSingleton()
		binder.Bind(new(Store)).ToConstructor(NewStoreOnMemory).AsEagerSingleton()
		binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory).AsEagerSingleton()
		binder.Bind(new(GroupRepository)).ToConstructor(NewGroupRepositoryOnMemory).AsEagerSingleton()
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
	projectService := injector.Get(new(ProjectService)).(*ProjectService)
	if projectService == nil {
		t.Fatal("not found a GroupRepository")
	}
	if projectService.FindUser() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
	if projectService.FindGroup() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
}

func Test_it_should_be_inject_struct_into_the_struct_directly(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		binder.Bind(new(ProjectService)).In(NoScope)
		binder.Bind(new(Store)).ToConstructor(NewStoreOnMemory)
		binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory)
		binder.Bind(new(GroupRepository)).ToConstructor(NewGroupRepositoryOnMemory)
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
	projectService := injector.Get(new(ProjectService)).(*ProjectService)
	if projectService == nil {
		t.Fatal("not found a GroupRepository")
	}
	if projectService.FindUser() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
	if projectService.FindGroup() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
}

func Test_it_should_be_inject_instance_directly(t *testing.T) {
	injector, err := CreateInjector(func(binder Binder) {
		store := NewStoreOnMemory()
		userRepository := NewUserRepositoryOnMemory(store)
		groupRepository := NewGroupRepositoryOnMemory(store)
		projectService := NewProjectService(userRepository, groupRepository)
		binder.Bind(new(UserRepository)).ToInstance(userRepository)
		binder.Bind(new(GroupRepository)).ToInstance(groupRepository).In(SingletonInstance)
		binder.Bind(new(ProjectService)).ToInstance(projectService).AsEagerSingleton()
	})
	if err != nil {
		t.Fatalf("fatal: %v", err)
	}
	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	if userRepository == nil {
		t.Fatal("not found a UserRepository")
	}
	if userRepository.FindAll() == nil {
		t.Fatal("could not inject field of UserRepository")
	}
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	if groupRepository == nil {
		t.Fatal("not found a GroupRepository")
	}
	if groupRepository.FindAll() == nil {
		t.Fatal("could not inject field of GroupRepository")
	}
	projectService := injector.Get(new(ProjectService)).(*ProjectService)
	if projectService == nil {
		t.Fatal("not found a GroupRepository")
	}
	if projectService.FindUser() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
	if projectService.FindGroup() == nil {
		t.Fatal("could not inject field of ProjectService")
	}
}