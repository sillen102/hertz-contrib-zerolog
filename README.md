# hertz-contrib-zerolog
[Zerolog](https://github.com/rs/zerolog) wrapper for [Hertz](https://github.com/cloudwego/hertz) web framework.

Heavily influenced by [Lecho](https://github.com/ziflex/lecho)

The wrapper implements the hlog.FullLogger interface using zerolog.
It contains methods to map between zerolog and hlog log levels and can be created with various options (see options.go).

## Installation
    go get github.com/sillen102/hertz-contrib-zerolog

## Usage
```go
func main () {
    h := server.Default()
	
    hlog.SetLogger(hertzZerolog.New(
        os.Stdout,
        hertzZerolog.WithLevel(hlog.LevelWarn)))

    h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
        hlog.Info("test log")
        c.JSON(consts.StatusOK, utils.H{"ping": "pong"})
    })
	
    h.Spin()
}
```