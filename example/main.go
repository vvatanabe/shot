package main

import (
	"github.com/vvatanabe/shot/shot"
	"fmt"
)

func main() {
	container := NewContainer()

	fmt.Println("UserRepository#FindAll =>", container.UserRepository().FindAll())
	fmt.Println("GroupRepository#FindAll =>", container.GroupRepository().FindAll())
	fmt.Println("ProjectService#FindUser =>", container.ProjectService().FindUser())
	fmt.Println("ProjectService#FindGroup =>", container.ProjectService().FindGroup())
}

func NewContainer() Container {

	injector, err := shot.CreateInjector(func(binder shot.Binder) {
		binder.Bind(new(Store)).ToConstructor(NewStoreOnMemory)
		binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).In(shot.SingletonInstance)
		binder.Bind(new(GroupRepository)).To(new(GroupRepositoryOnMemory)).In(shot.SingletonInstance)
		binder.Bind(new(ProjectService)).AsEagerSingleton()
	})

	if err != nil {
		panic(err)
	}

	return &container{injector}
}

type Container interface {
	Store() Store
	GroupRepository() GroupRepository
	UserRepository() UserRepository
	ProjectService() *ProjectService
}

type container struct {
	shot.Injector
}

func (c *container) Store() Store {
	return c.Get(new(Store)).(Store)
}

func (c *container) GroupRepository() GroupRepository {
	return c.Get(new(GroupRepository)).(GroupRepository)
}

func (c *container) UserRepository() UserRepository {
	return c.Get(new(UserRepository)).(UserRepository)
}

func (c *container) ProjectService() *ProjectService {
	return c.Get(new(ProjectService)).(*ProjectService)
}




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