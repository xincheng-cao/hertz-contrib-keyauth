package main

import "data1/shiba/keyauth/client/to_bearer_option"

const (
	TRUE_BEARER  string = "innovationlabchaogeyangconggejinjinjiewuyuanqingchenxiaolu"
	FALSE_BEARER string = "unoriginalitylabchaogeyangconggejinjinjiewuyuanqingchenxiaolu"
)

func main() {
	to_bearer_option.DoGetRequest(TRUE_BEARER)
	to_bearer_option.DoGetRequest(FALSE_BEARER)
}
