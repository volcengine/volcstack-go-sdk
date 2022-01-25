package main

import (
	"code.byted.org/iaasng/volcstack-go-sdk/service/ecs"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/session"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/volcstackutil"
	"fmt"
)

func main() {
	ak := "AKLTODI3MjQzMTk1NmZkNDgyMmE3YWRhMWUyOWQwOWViOWI"
	sk := "T1RWa09EUXhPVE14WVRZeE5EaGtObUZsTkRreE1XTTNOekV3TVRNeFltWQ=="
	var region = "cn-beijing"
	svc := ecs.New(session.SdkNew(ak, sk), &volcstack.Config{Region: &region}, &volcstackutil.UrlInfo{
		UseSSL: false,
	})

	query := &ecs.DescribeInstancesInput{}

	resp, err := svc.DescribeInstances(query)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %v ", *resp)
	//param := make(map[string]interface{})
	//param["CidrBlock"] = "172.20.0.0/16"
	//resp, err := svc.CreateVpcCommon(&param)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Response is %v ", *resp)
}
