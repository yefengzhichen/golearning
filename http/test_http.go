package main

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"gitlab.xiaoduoai.com/golib/xd_sdk/httpclient"
	"gitlab.xiaoduoai.com/golib/xd_sdk/httpclient/hooks"
)

var webClient httpclient.Client

func InitWebClient() {
	webClient, _ = httpclient.NewClient(
		httpclient.WithTimeout(30*time.Second),
		httpclient.WithPreRequestHooks(hooks.LoggingRequest()),
		httpclient.WithAfterResponseHooks(hooks.LoggingResponse()),
	)
}

func WebClient() httpclient.Client {
	return webClient
}

func test() {
	InitWebClient()
	type Order struct {
		OrderId         string   `json:"order_id"`
		Status          string   `json:"status"`
		ItemIDs         []string `json:"item_ids,omitempty"`
		Payment         float64  `json:"payment,omitempty"`
		OriginStatus    string   `json:"origin_status"`
		StepTradeStatus string   `json:"step_trade_status"`
		CreateAt        int64    `json:"created_at"`
		UpdateAt        int64    `json:"updated_at"`
	}
	ctx := context.Background()
	uri := ""
	req := url.Values{}
	req.Set("platform", "tb")
	req.Set("buyer_id", "li1031069371")
	req.Set("seller_id", "杜可风按")
	uri = "http://10.102.188.187:8080" + "/trade/get_last_orders?" + req.Encode()
	fmt.Println(uri)
	adapterOrders := struct {
		Orders []Order `json:"orders"`
	}{}
	webClient := WebClient()
	//reqC := webClient.NewRequest(ctx)
	_, err := webClient.NewRequest(ctx).SetResult(&adapterOrders).Get(uri)
	if err != nil {
		//添加一个错误时日志
		fmt.Printf("fail to get adapterOrders from /trade/get_last_orders, err:%+v, adapterOrders:%+v", err, adapterOrders)
	}
	//fmt.Println(reqC, adapterOrders)

}

func main() {
	test()
}
