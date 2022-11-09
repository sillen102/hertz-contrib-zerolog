# hertz-contrib-zerolog
[Zerolog](https://github.com/rs/zerolog) wrapper for [Hertz](https://github.com/cloudwego/hertz) web framework.  Heavily influenced by [Lecho](https://github.com/ziflex/lecho)

The wrapper implements the hlog.FullLogger interface using Zerolog.
It can be created with various options (see [options](###Options) below).

## Installation
    go get github.com/sillen102/hertz-contrib-zerolog

## Usage
### Simple example:
```go
import (
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/hlog"

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
```go
import (
    "os"
	
    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/hlog"

    hertzZerolog "github.com/sillen102/hertz-contrib-zerolog"
)

func main () {
    h := server.Default()
	
    hlog.SetLogger(hertzZerolog.New(
        hertzZerolog.WithOutput(os.Stdout), // allows to specify output
        hertzZerolog.WithLevel(hlog.LevelWarn), // allows to specify log level
	hertzZerolog.WithTimestamp(), // option to add timestamp
	hertzZerolog.WithCaller())) // option to add caller

    h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
        hlog.Info("test log")
        c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
    })
	
    h.Spin()
}
```
