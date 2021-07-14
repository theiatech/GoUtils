# GoUtils
golang utils

### Use FontColor
```go
fmt.Println(FontColor.Green("hello GoUtils"))
 ```

### Use HttpUtil
```go
rsp := HttpUtil.Requests(HttpUtil.HttpRequestParams{
		Method: http.MethodPost,
		Url: "http://127.0.0.1:80",
		Headers: map[string]string{"Content-Type":"text/plain"},
		Body: []byte("a=1&b=2"),
		BasicAuth: struct {
			Username string
			Password string
		}{Username: "root", Password: "root"},
	})
fmt.Println(string(rsp.Body))
```