# net/http/pprof handler wrapper

A simple (no dependencies) http pprof all-in-one handler

## usage

```go
router := echo.New()
router.Any("/stage/api/debug/pprof/*", echo.WrapHandler(pprofhandler.Handler("/debug/pprof/")))
```
