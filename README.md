# Notice

This code has been merged into the [official hertz-contrib repository](https://github.com/hertz-contrib/logger) and will no longer be maintained here.

# hertz-contrib-zerolog
This is a logger library that uses zerolog to implement the [Hertz logger interface](https://www.cloudwego.io/docs/hertz/tutorials/framework-exten/log/)

## Usage

#### Download and install it:

```
go get github.com/sillen102/hertz-contrib-zerolog
```

#### Import it in your code:

```
import hertzZerolog "github.com/sillen102/hertz-contrib-zerolog"
```

#### Simple example:
```go
import (
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/hlog"
    "github.com/cloudwego/hertz/pkg/common/utils"
    "github.com/cloudwego/hertz/pkg/protocol/consts"

    hertzZerolog "github.com/sillen102/hertz-contrib-zerolog"
)

func main () {
    h := server.Default()
	
    hlog.SetLogger(hertzZerolog.New()) // set the Hertz logger

    h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
        hlog.Info("test log")
        c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
    })
	
    h.Spin()
}
```

### Options:

#### WithOutput:
- Allows to specify the output of the logger. By default, it is set to os.Stdout.

#### WithLevel:
- Allows to specify the level of the logger. By default, it is set to Warn.

#### WithField:
- Allows to specify a field that will always be in the logger.

#### WithFields:
- Same as WithField but allows to specify multiple fields.

#### WithTimestamp:
- Allows to specify if the timestamp should be logged. By default, it is set to false.

#### WithFormattedTimestamp:
- Same as WithTimeStamp but takes a time format string as parameter that allows to specify the format of the timestamp in the logs.

#### WithCaller:
- Allows to specify if the caller should be logged. By default, it is set to false.

#### WithHook:
- Allows to specify a hook that will be called when a log is written.

#### WithHookFunc:
- Allows to specify a hook function that will be called when a log is written.

#### Example:
```go
import (
    "os"
	
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/hlog"
    "github.com/cloudwego/hertz/pkg/common/utils"
    "github.com/cloudwego/hertz/pkg/protocol/consts"

    hertzZerolog "github.com/sillen102/hertz-contrib-zerolog"
)

func main () {
    h := server.Default()
	
    hlog.SetLogger(hertzZerolog.New(
        hertzZerolog.WithOutput(zerolog.ConsoleWriter{Out: os.Stdout}),
        hertzZerolog.WithLevel(hlog.LevelWarn),
	hertzZerolog.WithTimestamp(),
	hertzZerolog.WithCaller()))

    h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
        hlog.Info("test log")
        c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
    })
	
    h.Spin()
}
```

## Advanced usage

#### Implementing a request logging middleware:
```go
import (
    "context"
    "time"
    
    "github.com/cloudwego/hertz/pkg/app"

    hertzZerolog "github.com/sillen102/hertz-contrib-zerolog"
)

// RequestIDHeaderValue value for the request id header
const RequestIDHeaderValue = "X-Request-Id"

// LoggerMiddleware middleware for logging incoming requests
func LoggerMiddleware() app.HandlerFunc {
    return func(c context.Context, ctx *app.RequestContext) {
        start := time.Now()
        
        reqId := ctx.Request.Header.Get(RequestIDHeaderValue)
        if reqId == "" {
            reqId = c.Value(RequestIDHeaderValue).(string)
        }
        
        logger := hertzZerolog.GetLogger()
        
        if reqId != "" {
            logger = logger.WithField("request_id", reqId)
        }
        
        c = logger.WithContext(c)
        
        defer func() {
            stop := time.Now()
            
            logger.Unwrap().Info().
                Str("remote_ip", ctx.ClientIP()).
                Str("method", string(ctx.Method())).
                Str("path", string(ctx.Path())).
                Str("user_agent", string(ctx.UserAgent())).
                Int("status", ctx.Response.StatusCode()).
                Dur("latency", stop.Sub(start)).
                Str("latency_human", stop.Sub(start).String()).
                Msg("request processed")
        }()
        
        ctx.Next(c)
    }
}
```
