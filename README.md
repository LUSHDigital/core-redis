# Core Redis
This package is used to wrap redis to satisfy the [health checking server](https://github.com/LUSHDigital/core/tree/master/workers/readysrv#ready-server) in the [LUSH core service library](https://github.com/LUSHDigital/core)

## Examples

### Use in conjunction with [readysrv](https://github.com/LUSHDigital/core/tree/master/workers/readysrv)

```go
client := redis.NewDefaultClient()
readysrv.New(readysrv.Checks{
    "redis": client,
})
```
