# logger

Simple extendable logging library.

## Quick start 

```go
log := logger.New(
    logger.SetLogLevel(logger.LevelDebug),
    logger.SetFormatter(logger.JSONFormatter{Pretty: true}),
    logger.SetWriter(os.Stdout),
    )

log.Info("msg", "info message")
log.Warn("msg", "warn message")
log.Debug("msg", "debug message")
log.Error("msg", "error message")
```
Result:
```json
    {
      "level": "INFO",
      "msg": "info message",
      "t": "2020-Apr-22 04:36:56.010"
    }
    {
      "level": "WARN",
      "msg": "warn message",
      "t": "2020-Apr-22 04:36:56.010"
    }
    {
      "level": "DEBUG",
      "msg": "debug message",
      "t": "2020-Apr-22 04:36:56.010"
    }
    {
      "level": "ERROR",
      "msg": "error message",
      "t": "2020-Apr-22 04:36:56.010"
    }
```

