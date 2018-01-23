# shot

This library is a reflection based tiny DI container. It was inspired by the interface of the Google Juice.

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
		binder.Bind(new(Store)).ToConstructor(NewStoreOnMemory)
		binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).In(shot.SingletonInstance)
		binder.Bind(new(GroupRepository)).To(new(GroupRepositoryOnMemory)).In(shot.SingletonInstance)
		binder.Bind(new(ProjectService)).AsEagerSingleton()
	})

	if err != nil {
		panic(err)
	}

	userRepository := injector.Get(new(UserRepository)).(UserRepository)
	groupRepository := injector.Get(new(GroupRepository)).(GroupRepository)
	projectService := injector.Get(new(ProjectService)).(*ProjectService)
}
```

### To inject implementation into the interface via struct.
``` go
binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory))

// as singleton
binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).In(SingletonInstance)

// as eager singleton
binder.Bind(new(UserRepository)).To(new(UserRepositoryOnMemory)).AsEagerSingleton()
```

### To inject implementation into the interface via constructor.
``` go
binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory)

// as singleton
binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory).In(SingletonInstance)

// as eager singleton
binder.Bind(new(UserRepository)).ToConstructor(NewUserRepositoryOnMemory).AsEagerSingleton()
```

### To inject struct into the struct directly.
``` go
binder.Bind(new(ProjectService)).In(NoScope)

// as singleton
binder.Bind(new(ProjectService)).In(SingletonInstance)

// as eager singleton
binder.Bind(new(ProjectService)).AsEagerSingleton()
```

## Acknowledgments

[google/guice](https://github.com/google/guice) really inspired me. I appreciate it.

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

## License

Apache License 2.0
