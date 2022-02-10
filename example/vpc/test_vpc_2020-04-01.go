package main

import (
	"code.byted.org/iaasng/volcstack-go-sdk/service/ecs"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/session"
	"code.byted.org/iaasng/volcstack-go-sdk/volcstack/volcstackutil"
	"fmt"
)

func main() {
	ak := "ak"
	sk := "ak"
	var region = "cn-beijing"

	//if use env Credentials
	//please
	//export VOLCSTACK_ACCESS_KEY=AK
	//export VOLCSTACK_SECRET_KEY=SK
	// and WithCredentials(credentials.NewEnvCredentials())

	config := volcstack.NewConfig().
		WithRegion(region).
		WithAkSk(ak, sk).
		//WithCredentials(credentials.NewEnvCredentials()).
		WithDisableSSL(true).
		WithLogLevel(volcstack.LogDebugWithHTTPBody).
		WithEndpoint(volcstackutil.NewEndpoint().GetEndpoint())
	sess, _ := session.NewSession(config)
	svc := ecs.New(sess)

	query := &ecs.RunInstancesInput{}

	query.Volumes = []*ecs.VolumeForRunInstancesInput{
		{
			Size:               volcstack.Int32(32),
			DeleteWithInstance: volcstack.String("true"),
		},
	}
	query.CpuOptions = &ecs.CpuOptionsForRunInstancesInput{Numa: volcstack.Int32(111)}

	resp, err := svc.RunInstances(query)
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
