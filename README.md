# net/http/pprof handler wrapper

A simple (no dependencies) http pprof all-in-one handler

## usage

```go
router := echo.New()
router.Any("/stage/api/debug/pprof/*", echo.WrapHandler(pprofhandler.Handler()))
router.Any("/stage/api/debug2/pprof/*", echo.WrapHandler(pprofhandler.Handler("debug2","pprof")))
router.Any("/stage/api/v2/debug/pprof/*", echo.WrapHandler(pprofhandler.Handler("")))
router.Any("/stage/api/debug-internal/secret/go-pprof/*", echo.WrapHandler(pprofhandler.Handler("debug-internal","secret", "go-pprof")))
```
