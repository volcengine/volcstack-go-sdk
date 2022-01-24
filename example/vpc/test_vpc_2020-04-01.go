package main

import (
	"code.byted.org/iaasng/volcstack-go-sdk/service/vpc"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/session"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/volcstackutil"
	"fmt"
)

func main() {
	ak := "AK"
	sk := "SK"
	var region = "cn-beijing"
	svc := vpc.New(session.SdkNew(ak, sk), &volcstack.Config{Region: &region}, &volcstackutil.UrlInfo{
		UseSSL: false,
	})

	query := &vpc.DescribeVpcsInput{
		VpcIds: volcstack.StringSlice([]string{"vpc-rrrumxgi5khsv0x58kambxl"}),
	}

	resp, err := svc.DescribeVpcs(query)
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
