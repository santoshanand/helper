# Helper Utility for Go Projects


This package contains logger and config modules

to use this package first need to initialize the helper module using 

> make sure keep `config.yml` file at root of your project.

```golang
  helper.Initialize()
```

Use some helper functions as follows

```golang
  helper.Log("Welcome to logger")
  helper.Logf("Welcome to logger with number %d", 1)
```