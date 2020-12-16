package api

type HelloReq struct {
	Id int64
}

type HelloResp struct {
	Content string
}

var (
	HelloUrl = "/test.v1/hello"
)
