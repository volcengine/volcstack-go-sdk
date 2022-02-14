// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package vpciface provides an interface to enable mocking the VPC service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package vpc

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
)

// VPCAPI provides an interface to enable mocking the
// vpc.VPC service client's API operation,
//
//    // volcstack sdk func uses an SDK service client to make a request to
//    // VPC.
//    func myFunc(svc VPCAPI) bool {
//        // Make svc.AssociateRouteTable request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := vpc.New(sess)
//
//        myFunc(svc)
//    }
//
type VPCAPI interface {
	AssociateRouteTableCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateRouteTableCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateRouteTableCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateRouteTable(*AssociateRouteTableInput) (*AssociateRouteTableOutput, error)
	AssociateRouteTableWithContext(volcstack.Context, *AssociateRouteTableInput, ...request.Option) (*AssociateRouteTableOutput, error)
	AssociateRouteTableRequest(*AssociateRouteTableInput) (*request.Request, *AssociateRouteTableOutput)

	AttachNetworkInterfaceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachNetworkInterfaceCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachNetworkInterfaceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachNetworkInterface(*AttachNetworkInterfaceInput) (*AttachNetworkInterfaceOutput, error)
	AttachNetworkInterfaceWithContext(volcstack.Context, *AttachNetworkInterfaceInput, ...request.Option) (*AttachNetworkInterfaceOutput, error)
	AttachNetworkInterfaceRequest(*AttachNetworkInterfaceInput) (*request.Request, *AttachNetworkInterfaceOutput)

	AuthorizeSecurityGroupEgressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AuthorizeSecurityGroupEgressCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AuthorizeSecurityGroupEgressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AuthorizeSecurityGroupEgress(*AuthorizeSecurityGroupEgressInput) (*AuthorizeSecurityGroupEgressOutput, error)
	AuthorizeSecurityGroupEgressWithContext(volcstack.Context, *AuthorizeSecurityGroupEgressInput, ...request.Option) (*AuthorizeSecurityGroupEgressOutput, error)
	AuthorizeSecurityGroupEgressRequest(*AuthorizeSecurityGroupEgressInput) (*request.Request, *AuthorizeSecurityGroupEgressOutput)

	AuthorizeSecurityGroupIngressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AuthorizeSecurityGroupIngressCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AuthorizeSecurityGroupIngressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AuthorizeSecurityGroupIngress(*AuthorizeSecurityGroupIngressInput) (*AuthorizeSecurityGroupIngressOutput, error)
	AuthorizeSecurityGroupIngressWithContext(volcstack.Context, *AuthorizeSecurityGroupIngressInput, ...request.Option) (*AuthorizeSecurityGroupIngressOutput, error)
	AuthorizeSecurityGroupIngressRequest(*AuthorizeSecurityGroupIngressInput) (*request.Request, *AuthorizeSecurityGroupIngressOutput)

	CreateNetworkInterfaceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateNetworkInterfaceCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateNetworkInterfaceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateNetworkInterface(*CreateNetworkInterfaceInput) (*CreateNetworkInterfaceOutput, error)
	CreateNetworkInterfaceWithContext(volcstack.Context, *CreateNetworkInterfaceInput, ...request.Option) (*CreateNetworkInterfaceOutput, error)
	CreateNetworkInterfaceRequest(*CreateNetworkInterfaceInput) (*request.Request, *CreateNetworkInterfaceOutput)

	CreateRouteEntryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateRouteEntryCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateRouteEntryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateRouteEntry(*CreateRouteEntryInput) (*CreateRouteEntryOutput, error)
	CreateRouteEntryWithContext(volcstack.Context, *CreateRouteEntryInput, ...request.Option) (*CreateRouteEntryOutput, error)
	CreateRouteEntryRequest(*CreateRouteEntryInput) (*request.Request, *CreateRouteEntryOutput)

	CreateRouteTableCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateRouteTableCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateRouteTableCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateRouteTable(*CreateRouteTableInput) (*CreateRouteTableOutput, error)
	CreateRouteTableWithContext(volcstack.Context, *CreateRouteTableInput, ...request.Option) (*CreateRouteTableOutput, error)
	CreateRouteTableRequest(*CreateRouteTableInput) (*request.Request, *CreateRouteTableOutput)

	CreateSecurityGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSecurityGroupCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSecurityGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSecurityGroup(*CreateSecurityGroupInput) (*CreateSecurityGroupOutput, error)
	CreateSecurityGroupWithContext(volcstack.Context, *CreateSecurityGroupInput, ...request.Option) (*CreateSecurityGroupOutput, error)
	CreateSecurityGroupRequest(*CreateSecurityGroupInput) (*request.Request, *CreateSecurityGroupOutput)

	CreateSubnetCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSubnetCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSubnetCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSubnet(*CreateSubnetInput) (*CreateSubnetOutput, error)
	CreateSubnetWithContext(volcstack.Context, *CreateSubnetInput, ...request.Option) (*CreateSubnetOutput, error)
	CreateSubnetRequest(*CreateSubnetInput) (*request.Request, *CreateSubnetOutput)

	CreateVpcCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVpcCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVpcCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVpc(*CreateVpcInput) (*CreateVpcOutput, error)
	CreateVpcWithContext(volcstack.Context, *CreateVpcInput, ...request.Option) (*CreateVpcOutput, error)
	CreateVpcRequest(*CreateVpcInput) (*request.Request, *CreateVpcOutput)

	DeleteNetworkInterfaceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteNetworkInterfaceCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteNetworkInterfaceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteNetworkInterface(*DeleteNetworkInterfaceInput) (*DeleteNetworkInterfaceOutput, error)
	DeleteNetworkInterfaceWithContext(volcstack.Context, *DeleteNetworkInterfaceInput, ...request.Option) (*DeleteNetworkInterfaceOutput, error)
	DeleteNetworkInterfaceRequest(*DeleteNetworkInterfaceInput) (*request.Request, *DeleteNetworkInterfaceOutput)

	DeleteRouteEntryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRouteEntryCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRouteEntryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRouteEntry(*DeleteRouteEntryInput) (*DeleteRouteEntryOutput, error)
	DeleteRouteEntryWithContext(volcstack.Context, *DeleteRouteEntryInput, ...request.Option) (*DeleteRouteEntryOutput, error)
	DeleteRouteEntryRequest(*DeleteRouteEntryInput) (*request.Request, *DeleteRouteEntryOutput)

	DeleteRouteTableCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRouteTableCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRouteTableCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRouteTable(*DeleteRouteTableInput) (*DeleteRouteTableOutput, error)
	DeleteRouteTableWithContext(volcstack.Context, *DeleteRouteTableInput, ...request.Option) (*DeleteRouteTableOutput, error)
	DeleteRouteTableRequest(*DeleteRouteTableInput) (*request.Request, *DeleteRouteTableOutput)

	DeleteSecurityGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSecurityGroupCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSecurityGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSecurityGroup(*DeleteSecurityGroupInput) (*DeleteSecurityGroupOutput, error)
	DeleteSecurityGroupWithContext(volcstack.Context, *DeleteSecurityGroupInput, ...request.Option) (*DeleteSecurityGroupOutput, error)
	DeleteSecurityGroupRequest(*DeleteSecurityGroupInput) (*request.Request, *DeleteSecurityGroupOutput)

	DeleteSubnetCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSubnetCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSubnetCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSubnet(*DeleteSubnetInput) (*DeleteSubnetOutput, error)
	DeleteSubnetWithContext(volcstack.Context, *DeleteSubnetInput, ...request.Option) (*DeleteSubnetOutput, error)
	DeleteSubnetRequest(*DeleteSubnetInput) (*request.Request, *DeleteSubnetOutput)

	DeleteVpcCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVpcCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVpcCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVpc(*DeleteVpcInput) (*DeleteVpcOutput, error)
	DeleteVpcWithContext(volcstack.Context, *DeleteVpcInput, ...request.Option) (*DeleteVpcOutput, error)
	DeleteVpcRequest(*DeleteVpcInput) (*request.Request, *DeleteVpcOutput)

	DescribeNetworkInterfaceAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNetworkInterfaceAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNetworkInterfaceAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNetworkInterfaceAttributes(*DescribeNetworkInterfaceAttributesInput) (*DescribeNetworkInterfaceAttributesOutput, error)
	DescribeNetworkInterfaceAttributesWithContext(volcstack.Context, *DescribeNetworkInterfaceAttributesInput, ...request.Option) (*DescribeNetworkInterfaceAttributesOutput, error)
	DescribeNetworkInterfaceAttributesRequest(*DescribeNetworkInterfaceAttributesInput) (*request.Request, *DescribeNetworkInterfaceAttributesOutput)

	DescribeNetworkInterfacesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNetworkInterfacesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNetworkInterfacesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNetworkInterfaces(*DescribeNetworkInterfacesInput) (*DescribeNetworkInterfacesOutput, error)
	DescribeNetworkInterfacesWithContext(volcstack.Context, *DescribeNetworkInterfacesInput, ...request.Option) (*DescribeNetworkInterfacesOutput, error)
	DescribeNetworkInterfacesRequest(*DescribeNetworkInterfacesInput) (*request.Request, *DescribeNetworkInterfacesOutput)

	DescribeRouteEntryListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRouteEntryListCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRouteEntryListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRouteEntryList(*DescribeRouteEntryListInput) (*DescribeRouteEntryListOutput, error)
	DescribeRouteEntryListWithContext(volcstack.Context, *DescribeRouteEntryListInput, ...request.Option) (*DescribeRouteEntryListOutput, error)
	DescribeRouteEntryListRequest(*DescribeRouteEntryListInput) (*request.Request, *DescribeRouteEntryListOutput)

	DescribeRouteTableListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRouteTableListCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRouteTableListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRouteTableList(*DescribeRouteTableListInput) (*DescribeRouteTableListOutput, error)
	DescribeRouteTableListWithContext(volcstack.Context, *DescribeRouteTableListInput, ...request.Option) (*DescribeRouteTableListOutput, error)
	DescribeRouteTableListRequest(*DescribeRouteTableListInput) (*request.Request, *DescribeRouteTableListOutput)

	DescribeSecurityGroupAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSecurityGroupAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSecurityGroupAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSecurityGroupAttributes(*DescribeSecurityGroupAttributesInput) (*DescribeSecurityGroupAttributesOutput, error)
	DescribeSecurityGroupAttributesWithContext(volcstack.Context, *DescribeSecurityGroupAttributesInput, ...request.Option) (*DescribeSecurityGroupAttributesOutput, error)
	DescribeSecurityGroupAttributesRequest(*DescribeSecurityGroupAttributesInput) (*request.Request, *DescribeSecurityGroupAttributesOutput)

	DescribeSecurityGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSecurityGroupsCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSecurityGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSecurityGroups(*DescribeSecurityGroupsInput) (*DescribeSecurityGroupsOutput, error)
	DescribeSecurityGroupsWithContext(volcstack.Context, *DescribeSecurityGroupsInput, ...request.Option) (*DescribeSecurityGroupsOutput, error)
	DescribeSecurityGroupsRequest(*DescribeSecurityGroupsInput) (*request.Request, *DescribeSecurityGroupsOutput)

	DescribeSubnetAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSubnetAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSubnetAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSubnetAttributes(*DescribeSubnetAttributesInput) (*DescribeSubnetAttributesOutput, error)
	DescribeSubnetAttributesWithContext(volcstack.Context, *DescribeSubnetAttributesInput, ...request.Option) (*DescribeSubnetAttributesOutput, error)
	DescribeSubnetAttributesRequest(*DescribeSubnetAttributesInput) (*request.Request, *DescribeSubnetAttributesOutput)

	DescribeSubnetsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSubnetsCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSubnetsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSubnets(*DescribeSubnetsInput) (*DescribeSubnetsOutput, error)
	DescribeSubnetsWithContext(volcstack.Context, *DescribeSubnetsInput, ...request.Option) (*DescribeSubnetsOutput, error)
	DescribeSubnetsRequest(*DescribeSubnetsInput) (*request.Request, *DescribeSubnetsOutput)

	DescribeVpcAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpcAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpcAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpcAttributes(*DescribeVpcAttributesInput) (*DescribeVpcAttributesOutput, error)
	DescribeVpcAttributesWithContext(volcstack.Context, *DescribeVpcAttributesInput, ...request.Option) (*DescribeVpcAttributesOutput, error)
	DescribeVpcAttributesRequest(*DescribeVpcAttributesInput) (*request.Request, *DescribeVpcAttributesOutput)

	DescribeVpcsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpcsCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpcsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpcs(*DescribeVpcsInput) (*DescribeVpcsOutput, error)
	DescribeVpcsWithContext(volcstack.Context, *DescribeVpcsInput, ...request.Option) (*DescribeVpcsOutput, error)
	DescribeVpcsRequest(*DescribeVpcsInput) (*request.Request, *DescribeVpcsOutput)

	DetachNetworkInterfaceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachNetworkInterfaceCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachNetworkInterfaceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachNetworkInterface(*DetachNetworkInterfaceInput) (*DetachNetworkInterfaceOutput, error)
	DetachNetworkInterfaceWithContext(volcstack.Context, *DetachNetworkInterfaceInput, ...request.Option) (*DetachNetworkInterfaceOutput, error)
	DetachNetworkInterfaceRequest(*DetachNetworkInterfaceInput) (*request.Request, *DetachNetworkInterfaceOutput)

	DisassociateRouteTableCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateRouteTableCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateRouteTableCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateRouteTable(*DisassociateRouteTableInput) (*DisassociateRouteTableOutput, error)
	DisassociateRouteTableWithContext(volcstack.Context, *DisassociateRouteTableInput, ...request.Option) (*DisassociateRouteTableOutput, error)
	DisassociateRouteTableRequest(*DisassociateRouteTableInput) (*request.Request, *DisassociateRouteTableOutput)

	ModifyNetworkInterfaceAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNetworkInterfaceAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNetworkInterfaceAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNetworkInterfaceAttributes(*ModifyNetworkInterfaceAttributesInput) (*ModifyNetworkInterfaceAttributesOutput, error)
	ModifyNetworkInterfaceAttributesWithContext(volcstack.Context, *ModifyNetworkInterfaceAttributesInput, ...request.Option) (*ModifyNetworkInterfaceAttributesOutput, error)
	ModifyNetworkInterfaceAttributesRequest(*ModifyNetworkInterfaceAttributesInput) (*request.Request, *ModifyNetworkInterfaceAttributesOutput)

	ModifyRouteEntryCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyRouteEntryCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyRouteEntryCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyRouteEntry(*ModifyRouteEntryInput) (*ModifyRouteEntryOutput, error)
	ModifyRouteEntryWithContext(volcstack.Context, *ModifyRouteEntryInput, ...request.Option) (*ModifyRouteEntryOutput, error)
	ModifyRouteEntryRequest(*ModifyRouteEntryInput) (*request.Request, *ModifyRouteEntryOutput)

	ModifyRouteTableAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyRouteTableAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyRouteTableAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyRouteTableAttributes(*ModifyRouteTableAttributesInput) (*ModifyRouteTableAttributesOutput, error)
	ModifyRouteTableAttributesWithContext(volcstack.Context, *ModifyRouteTableAttributesInput, ...request.Option) (*ModifyRouteTableAttributesOutput, error)
	ModifyRouteTableAttributesRequest(*ModifyRouteTableAttributesInput) (*request.Request, *ModifyRouteTableAttributesOutput)

	ModifySecurityGroupAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySecurityGroupAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySecurityGroupAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySecurityGroupAttributes(*ModifySecurityGroupAttributesInput) (*ModifySecurityGroupAttributesOutput, error)
	ModifySecurityGroupAttributesWithContext(volcstack.Context, *ModifySecurityGroupAttributesInput, ...request.Option) (*ModifySecurityGroupAttributesOutput, error)
	ModifySecurityGroupAttributesRequest(*ModifySecurityGroupAttributesInput) (*request.Request, *ModifySecurityGroupAttributesOutput)

	ModifySecurityGroupRuleDescriptionsEgressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySecurityGroupRuleDescriptionsEgressCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySecurityGroupRuleDescriptionsEgressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySecurityGroupRuleDescriptionsEgress(*ModifySecurityGroupRuleDescriptionsEgressInput) (*ModifySecurityGroupRuleDescriptionsEgressOutput, error)
	ModifySecurityGroupRuleDescriptionsEgressWithContext(volcstack.Context, *ModifySecurityGroupRuleDescriptionsEgressInput, ...request.Option) (*ModifySecurityGroupRuleDescriptionsEgressOutput, error)
	ModifySecurityGroupRuleDescriptionsEgressRequest(*ModifySecurityGroupRuleDescriptionsEgressInput) (*request.Request, *ModifySecurityGroupRuleDescriptionsEgressOutput)

	ModifySecurityGroupRuleDescriptionsIngressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySecurityGroupRuleDescriptionsIngressCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySecurityGroupRuleDescriptionsIngressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySecurityGroupRuleDescriptionsIngress(*ModifySecurityGroupRuleDescriptionsIngressInput) (*ModifySecurityGroupRuleDescriptionsIngressOutput, error)
	ModifySecurityGroupRuleDescriptionsIngressWithContext(volcstack.Context, *ModifySecurityGroupRuleDescriptionsIngressInput, ...request.Option) (*ModifySecurityGroupRuleDescriptionsIngressOutput, error)
	ModifySecurityGroupRuleDescriptionsIngressRequest(*ModifySecurityGroupRuleDescriptionsIngressInput) (*request.Request, *ModifySecurityGroupRuleDescriptionsIngressOutput)

	ModifySubnetAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySubnetAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySubnetAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySubnetAttributes(*ModifySubnetAttributesInput) (*ModifySubnetAttributesOutput, error)
	ModifySubnetAttributesWithContext(volcstack.Context, *ModifySubnetAttributesInput, ...request.Option) (*ModifySubnetAttributesOutput, error)
	ModifySubnetAttributesRequest(*ModifySubnetAttributesInput) (*request.Request, *ModifySubnetAttributesOutput)

	ModifyVpcAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVpcAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVpcAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVpcAttributes(*ModifyVpcAttributesInput) (*ModifyVpcAttributesOutput, error)
	ModifyVpcAttributesWithContext(volcstack.Context, *ModifyVpcAttributesInput, ...request.Option) (*ModifyVpcAttributesOutput, error)
	ModifyVpcAttributesRequest(*ModifyVpcAttributesInput) (*request.Request, *ModifyVpcAttributesOutput)

	RevokeSecurityGroupEgressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RevokeSecurityGroupEgressCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RevokeSecurityGroupEgressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RevokeSecurityGroupEgress(*RevokeSecurityGroupEgressInput) (*RevokeSecurityGroupEgressOutput, error)
	RevokeSecurityGroupEgressWithContext(volcstack.Context, *RevokeSecurityGroupEgressInput, ...request.Option) (*RevokeSecurityGroupEgressOutput, error)
	RevokeSecurityGroupEgressRequest(*RevokeSecurityGroupEgressInput) (*request.Request, *RevokeSecurityGroupEgressOutput)

	RevokeSecurityGroupIngressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RevokeSecurityGroupIngressCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RevokeSecurityGroupIngressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RevokeSecurityGroupIngress(*RevokeSecurityGroupIngressInput) (*RevokeSecurityGroupIngressOutput, error)
	RevokeSecurityGroupIngressWithContext(volcstack.Context, *RevokeSecurityGroupIngressInput, ...request.Option) (*RevokeSecurityGroupIngressOutput, error)
	RevokeSecurityGroupIngressRequest(*RevokeSecurityGroupIngressInput) (*request.Request, *RevokeSecurityGroupIngressOutput)
}

var _ VPCAPI = (*VPC)(nil)
