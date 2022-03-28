// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rdsmysqliface provides an interface to enable mocking the RDS_MYSQL service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rdsmysql

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
)

// RDSMYSQLAPI provides an interface to enable mocking the
// rdsmysql.RDSMYSQL service client's API operation,
//
//    // volcstack sdk func uses an SDK service client to make a request to
//    // RDS_MYSQL.
//    func myFunc(svc RDSMYSQLAPI) bool {
//        // Make svc.CreateAccount request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := rdsmysql.New(sess)
//
//        myFunc(svc)
//    }
//
type RDSMYSQLAPI interface {
	CreateAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAccountCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAccount(*CreateAccountInput) (*CreateAccountOutput, error)
	CreateAccountWithContext(volcstack.Context, *CreateAccountInput, ...request.Option) (*CreateAccountOutput, error)
	CreateAccountRequest(*CreateAccountInput) (*request.Request, *CreateAccountOutput)

	CreateBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateBackupCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateBackup(*CreateBackupInput) (*CreateBackupOutput, error)
	CreateBackupWithContext(volcstack.Context, *CreateBackupInput, ...request.Option) (*CreateBackupOutput, error)
	CreateBackupRequest(*CreateBackupInput) (*request.Request, *CreateBackupOutput)

	CreateDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDBInstanceCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDBInstance(*CreateDBInstanceInput) (*CreateDBInstanceOutput, error)
	CreateDBInstanceWithContext(volcstack.Context, *CreateDBInstanceInput, ...request.Option) (*CreateDBInstanceOutput, error)
	CreateDBInstanceRequest(*CreateDBInstanceInput) (*request.Request, *CreateDBInstanceOutput)

	CreateDBInstanceIPListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDBInstanceIPListCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDBInstanceIPListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDBInstanceIPList(*CreateDBInstanceIPListInput) (*CreateDBInstanceIPListOutput, error)
	CreateDBInstanceIPListWithContext(volcstack.Context, *CreateDBInstanceIPListInput, ...request.Option) (*CreateDBInstanceIPListOutput, error)
	CreateDBInstanceIPListRequest(*CreateDBInstanceIPListInput) (*request.Request, *CreateDBInstanceIPListOutput)

	CreateDatabaseCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDatabaseCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDatabaseCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDatabase(*CreateDatabaseInput) (*CreateDatabaseOutput, error)
	CreateDatabaseWithContext(volcstack.Context, *CreateDatabaseInput, ...request.Option) (*CreateDatabaseOutput, error)
	CreateDatabaseRequest(*CreateDatabaseInput) (*request.Request, *CreateDatabaseOutput)

	DeleteAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAccountCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAccount(*DeleteAccountInput) (*DeleteAccountOutput, error)
	DeleteAccountWithContext(volcstack.Context, *DeleteAccountInput, ...request.Option) (*DeleteAccountOutput, error)
	DeleteAccountRequest(*DeleteAccountInput) (*request.Request, *DeleteAccountOutput)

	DeleteDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDBInstanceCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDBInstance(*DeleteDBInstanceInput) (*DeleteDBInstanceOutput, error)
	DeleteDBInstanceWithContext(volcstack.Context, *DeleteDBInstanceInput, ...request.Option) (*DeleteDBInstanceOutput, error)
	DeleteDBInstanceRequest(*DeleteDBInstanceInput) (*request.Request, *DeleteDBInstanceOutput)

	DeleteDBInstanceIPListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDBInstanceIPListCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDBInstanceIPListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDBInstanceIPList(*DeleteDBInstanceIPListInput) (*DeleteDBInstanceIPListOutput, error)
	DeleteDBInstanceIPListWithContext(volcstack.Context, *DeleteDBInstanceIPListInput, ...request.Option) (*DeleteDBInstanceIPListOutput, error)
	DeleteDBInstanceIPListRequest(*DeleteDBInstanceIPListInput) (*request.Request, *DeleteDBInstanceIPListOutput)

	DeleteDatabaseCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDatabaseCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDatabaseCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDatabase(*DeleteDatabaseInput) (*DeleteDatabaseOutput, error)
	DeleteDatabaseWithContext(volcstack.Context, *DeleteDatabaseInput, ...request.Option) (*DeleteDatabaseOutput, error)
	DeleteDatabaseRequest(*DeleteDatabaseInput) (*request.Request, *DeleteDatabaseOutput)

	DescribeDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstanceCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstance(*DescribeDBInstanceInput) (*DescribeDBInstanceOutput, error)
	DescribeDBInstanceWithContext(volcstack.Context, *DescribeDBInstanceInput, ...request.Option) (*DescribeDBInstanceOutput, error)
	DescribeDBInstanceRequest(*DescribeDBInstanceInput) (*request.Request, *DescribeDBInstanceOutput)

	DescribeDBInstanceConnectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstanceConnectionCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstanceConnectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstanceConnection(*DescribeDBInstanceConnectionInput) (*DescribeDBInstanceConnectionOutput, error)
	DescribeDBInstanceConnectionWithContext(volcstack.Context, *DescribeDBInstanceConnectionInput, ...request.Option) (*DescribeDBInstanceConnectionOutput, error)
	DescribeDBInstanceConnectionRequest(*DescribeDBInstanceConnectionInput) (*request.Request, *DescribeDBInstanceConnectionOutput)

	DescribeRecoverableTimeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRecoverableTimeCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRecoverableTimeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRecoverableTime(*DescribeRecoverableTimeInput) (*DescribeRecoverableTimeOutput, error)
	DescribeRecoverableTimeWithContext(volcstack.Context, *DescribeRecoverableTimeInput, ...request.Option) (*DescribeRecoverableTimeOutput, error)
	DescribeRecoverableTimeRequest(*DescribeRecoverableTimeInput) (*request.Request, *DescribeRecoverableTimeOutput)

	GrantAccountPrivilegeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GrantAccountPrivilegeCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GrantAccountPrivilegeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GrantAccountPrivilege(*GrantAccountPrivilegeInput) (*GrantAccountPrivilegeOutput, error)
	GrantAccountPrivilegeWithContext(volcstack.Context, *GrantAccountPrivilegeInput, ...request.Option) (*GrantAccountPrivilegeOutput, error)
	GrantAccountPrivilegeRequest(*GrantAccountPrivilegeInput) (*request.Request, *GrantAccountPrivilegeOutput)

	ListAccountsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListAccountsCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListAccountsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListAccounts(*ListAccountsInput) (*ListAccountsOutput, error)
	ListAccountsWithContext(volcstack.Context, *ListAccountsInput, ...request.Option) (*ListAccountsOutput, error)
	ListAccountsRequest(*ListAccountsInput) (*request.Request, *ListAccountsOutput)

	ListBackupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListBackupsCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListBackupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListBackups(*ListBackupsInput) (*ListBackupsOutput, error)
	ListBackupsWithContext(volcstack.Context, *ListBackupsInput, ...request.Option) (*ListBackupsOutput, error)
	ListBackupsRequest(*ListBackupsInput) (*request.Request, *ListBackupsOutput)

	ListDBInstanceIPListsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListDBInstanceIPListsCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListDBInstanceIPListsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListDBInstanceIPLists(*ListDBInstanceIPListsInput) (*ListDBInstanceIPListsOutput, error)
	ListDBInstanceIPListsWithContext(volcstack.Context, *ListDBInstanceIPListsInput, ...request.Option) (*ListDBInstanceIPListsOutput, error)
	ListDBInstanceIPListsRequest(*ListDBInstanceIPListsInput) (*request.Request, *ListDBInstanceIPListsOutput)

	ListDBInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListDBInstancesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListDBInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListDBInstances(*ListDBInstancesInput) (*ListDBInstancesOutput, error)
	ListDBInstancesWithContext(volcstack.Context, *ListDBInstancesInput, ...request.Option) (*ListDBInstancesOutput, error)
	ListDBInstancesRequest(*ListDBInstancesInput) (*request.Request, *ListDBInstancesOutput)

	ListDatabasesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListDatabasesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListDatabasesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListDatabases(*ListDatabasesInput) (*ListDatabasesOutput, error)
	ListDatabasesWithContext(volcstack.Context, *ListDatabasesInput, ...request.Option) (*ListDatabasesOutput, error)
	ListDatabasesRequest(*ListDatabasesInput) (*request.Request, *ListDatabasesOutput)

	ModifyDBInstanceIPListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceIPListCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceIPListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceIPList(*ModifyDBInstanceIPListInput) (*ModifyDBInstanceIPListOutput, error)
	ModifyDBInstanceIPListWithContext(volcstack.Context, *ModifyDBInstanceIPListInput, ...request.Option) (*ModifyDBInstanceIPListOutput, error)
	ModifyDBInstanceIPListRequest(*ModifyDBInstanceIPListInput) (*request.Request, *ModifyDBInstanceIPListOutput)

	RecoveryDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RecoveryDBInstanceCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RecoveryDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RecoveryDBInstance(*RecoveryDBInstanceInput) (*RecoveryDBInstanceOutput, error)
	RecoveryDBInstanceWithContext(volcstack.Context, *RecoveryDBInstanceInput, ...request.Option) (*RecoveryDBInstanceOutput, error)
	RecoveryDBInstanceRequest(*RecoveryDBInstanceInput) (*request.Request, *RecoveryDBInstanceOutput)

	ResetAccountPasswordCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ResetAccountPasswordCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ResetAccountPasswordCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ResetAccountPassword(*ResetAccountPasswordInput) (*ResetAccountPasswordOutput, error)
	ResetAccountPasswordWithContext(volcstack.Context, *ResetAccountPasswordInput, ...request.Option) (*ResetAccountPasswordOutput, error)
	ResetAccountPasswordRequest(*ResetAccountPasswordInput) (*request.Request, *ResetAccountPasswordOutput)

	RestartDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestartDBInstanceCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestartDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestartDBInstance(*RestartDBInstanceInput) (*RestartDBInstanceOutput, error)
	RestartDBInstanceWithContext(volcstack.Context, *RestartDBInstanceInput, ...request.Option) (*RestartDBInstanceOutput, error)
	RestartDBInstanceRequest(*RestartDBInstanceInput) (*request.Request, *RestartDBInstanceOutput)

	RevokeAccountPrivilegeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RevokeAccountPrivilegeCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RevokeAccountPrivilegeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RevokeAccountPrivilege(*RevokeAccountPrivilegeInput) (*RevokeAccountPrivilegeOutput, error)
	RevokeAccountPrivilegeWithContext(volcstack.Context, *RevokeAccountPrivilegeInput, ...request.Option) (*RevokeAccountPrivilegeOutput, error)
	RevokeAccountPrivilegeRequest(*RevokeAccountPrivilegeInput) (*request.Request, *RevokeAccountPrivilegeOutput)
}

var _ RDSMYSQLAPI = (*RDSMYSQL)(nil)
