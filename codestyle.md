# Code Style

## Linter Rules

Use golangci-lint with provided config and auto-fix all issues with it.

## fx modules Rules

All uber/fx related functions should be in `module.go` file in each package.

Module declaration should be in packages `Module` function.

Example:

```go
func Module() fx.Option {
	return fx.Module("module_name",
		fx.Provide(
			NewService1,
			NewService2,
			...
		),
	)
}
```

Module invokations should be in packages `Invoke` function.

Example:

```go
func Invoke() fx.Option {
	return fx.Invoke(
		func (*Service1) {},
		invokeService2Func,
	)
}
```

## Errors Rules

### Definition

All app-global errors should be defined in `pkg/models/errors.go` file.

Errors for single package only should be defined in that package in `errors.go` file.

### Wrapping

All errors should be wrapped with stack in point of creation with `github.com/pkg/errors.WithStack`.

If error is already wrapped in callee, it should be passed as is.
