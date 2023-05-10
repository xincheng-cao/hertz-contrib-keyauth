package to_bearer_option

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func DoGetRequest(bearer string) {
	c, err := client.NewClient()
	if err != nil {
		return
	}

	req := &protocol.Request{}
	res := &protocol.Response{}
	req.SetMethod(consts.MethodGet)
	(&(req.Header)).SetContentTypeBytes([]byte("application/json"))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearer))
	req.SetRequestURI("http://10.200.0.43:8888/ping")
	err = c.Do(context.Background(), req, res)
	if err != nil {
		return
	}
	fmt.Printf("%v\n\n", string(res.Body()))
}
