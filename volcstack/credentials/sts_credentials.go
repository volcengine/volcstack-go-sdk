package credentials

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/volcengine/volc-sdk-golang/service/sts"
)

type StsAssumeRoleProvider struct {
	AccessKey       string
	SecurityKey     string
	RoleName        string
	AccountId       string
	Host            string
	Region          string
	Timeout         time.Duration
	DurationSeconds int
}

type StsAssumeRoleTime struct {
	CurrentTime string
	ExpiredTime string
}

func StsAssumeRole(p *StsAssumeRoleProvider) (*Credentials, *StsAssumeRoleTime, error) {
	ins := sts.NewInstance()
	if p.Region != "" {
		ins.SetRegion(p.Region)
	}
	if p.Host != "" {
		ins.SetHost(p.Host)
	}
	if p.Timeout > 0 {
		ins.Client.SetTimeout(p.Timeout)
	}

	ins.Client.SetAccessKey(p.AccessKey)
	ins.Client.SetSecretKey(p.SecurityKey)
	input := &sts.AssumeRoleRequest{
		DurationSeconds: p.DurationSeconds,
		RoleTrn:         fmt.Sprintf("trn:iam::%s:role/%s", p.AccountId, p.RoleName),
		RoleSessionName: uuid.New().String(),
	}
	output, _, err := ins.AssumeRole(input)
	if err != nil {
		return nil, nil, err
	}
	return NewCredentials(&StaticProvider{Value: Value{
			AccessKeyID:     output.Result.Credentials.AccessKeyId,
			SecretAccessKey: output.Result.Credentials.SecretAccessKey,
			SessionToken:    output.Result.Credentials.SessionToken,
		}}), &StsAssumeRoleTime{
			CurrentTime: output.Result.Credentials.CurrentTime,
			ExpiredTime: output.Result.Credentials.ExpiredTime,
		}, nil
}
