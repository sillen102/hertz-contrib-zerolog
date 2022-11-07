# hertz-contrib-zerolog
[Zerolog](https://github.com/rs/zerolog) wrapper for [Hertz](https://github.com/cloudwego/hertz) web framework.

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