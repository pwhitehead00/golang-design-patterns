# Go Design Patterns
This is a collection of patterns and implementations I've found while learning Go

## Dependency Injection
https://golangforall.com/en/post/dependency-injection.html#dependency-injection-container-(injector)

## Constructor
Useful when there's a lot of options that must be configured.  See [internanl/constructor](internal/constructor/constructor.go)

### Implemtations
https://www.theelements.org/2016/05/30/the-builder-pattern-in-go/
https://golang.cafe/blog/golang-functional-options-pattern.html - Option 2
https://golangbot.com/structs-instead-of-classes/

## Options
Extensible and easy to read, used by the [AWS Go SDK](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/config#LoadDefaultConfig).  See [internal/options](internal/options/options.go)

### Whitepapers
https://dave.cheney.net/2016/11/13/do-not-fear-first-class-functions
https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

### Implementations
https://golang.cafe/blog/golang-functional-options-pattern.html - Option 3
https://www.sohamkamani.com/golang/options-pattern/
https://medium.com/star-gazers/go-options-pattern-da49185a2526
https://www.calhoun.io/using-functional-options-instead-of-method-chaining-in-go/

## Builder: Functional
Similar to a factory pattern.  See [internal/functional](internal/functional/functional.go)

### Implementations
https://devcharmander.medium.com/go-builder-pattern-the-functional-way-e40f347017ce

## Interfaces
Constructor pattern with interfaces

### Implementations
https://refactoring.guru/design-patterns/builder/go/example#example-0--iBuilder-go
https://golangbot.com/interfaces-part-2/
https://gist.github.com/vaskoz/10073335
