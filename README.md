# shot [![Build Status](https://travis-ci.org/vvatanabe/shot.svg?branch=master)](https://travis-ci.org/vvatanabe/shot)

This library is a reflection based tiny DI container. It was inspired by the interface of the Google Guice.

## Requires

* Go 1.7+

## Installation

This package can be installed with the go get command:

``` zsh
$ go get github.com/vvatanabe/shot
```

## Usage

### Basically
``` go
package main

import (
	"github.com/vvatanabe/shot/shot"
)

func main() {
	injector, err := shot.CreateInjector(func(binder shot.Binder) {
		store := NewStoreOnMemory()
		binder.Bind(new(Store)).ToInstance(store)
		binder.Bind(new(GroupRepository)).ToConstructor(NewGroupRepositoryOnMemory)
		binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).In(shot.SingletonInstance)
		binder.Bind(new(ProjectService)).AsEagerSingleton()
	})

	if err != nil {
		panic(err)
	}

	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	projectService := injector.Get(new(ProjectService)).(*ProjectService)
}

func NewProjectService(userRepository UserRepository, groupRepository GroupRepository) *ProjectService {
	return &ProjectService{
		userRepository, groupRepository,
	}
}

type ProjectService struct {
	UserRepository  UserRepository  "inject"
	GroupRepository GroupRepository "inject"
}

func (u *ProjectService) FindUser() []string {
	return u.UserRepository.FindAll()
}

func (u *ProjectService) FindGroup() []string {
	return u.GroupRepository.FindAll()
}

func NewStoreOnMemory() *StoreOnMemory {
	return &StoreOnMemory{
		[]string{"user-1", "user-2", "user-3"},
		[]string{"group-1", "group-2", "group-3"},
	}
}

type Store interface {
	GetUsers() []string
	GetGroups() []string
}

type StoreOnMemory struct {
	users  []string
	groups []string
}

func (s *StoreOnMemory) GetUsers() []string {
	return s.users
}

func (s *StoreOnMemory) GetGroups() []string {
	return s.groups
}

type UserRepository interface {
	FindAll() []string
}

func NewUserRepositoryOnMemory(store Store) *UserRepositoryOnMemory {
	return &UserRepositoryOnMemory{store}
}

type UserRepositoryOnMemory struct {
	Store Store "inject"
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
	Store Store "inject"
}

func (repository *GroupRepositoryOnMemory) FindAll() []string {
	return repository.Store.GetGroups()
}
```

### To inject implementation into the interface via struct.
``` go
binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory))
```

### To inject implementation into the interface via constructor.
``` go
binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory)
```

### To inject instance into the struct directly.
``` go
binder.Bind(new(ProjectService)).ToInstance(store)
```

### To inject implementation into the struct directly.
``` go
binder.Bind(new(ProjectService)).In(shot.NoScope)
```

### The singleton binding
``` go
binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).In(shot.SingletonInstance)
```

### The eager singleton binding
``` go
binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).AsEagerSingleton()
```

## Acknowledgments

[google/guice](https://github.com/google/guice) really inspired me. I appreciate it.

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

## License

Apache License 2.0
