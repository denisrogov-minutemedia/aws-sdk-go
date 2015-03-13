package cognitoidentity

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// CreateIdentityPoolRequest generates a request for the CreateIdentityPool operation.
func (c *CognitoIdentity) CreateIdentityPoolRequest(input *CreateIdentityPoolInput) (req *aws.Request, output *IdentityPool) {
	if opCreateIdentityPool == nil {
		opCreateIdentityPool = &aws.Operation{
			Name:       "CreateIdentityPool",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreateIdentityPool, input, output)
	output = &IdentityPool{}
	req.Data = output
	return
}

func (c *CognitoIdentity) CreateIdentityPool(input *CreateIdentityPoolInput) (output *IdentityPool, err error) {
	req, out := c.CreateIdentityPoolRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateIdentityPool *aws.Operation

// DeleteIdentityPoolRequest generates a request for the DeleteIdentityPool operation.
func (c *CognitoIdentity) DeleteIdentityPoolRequest(input *DeleteIdentityPoolInput) (req *aws.Request, output *DeleteIdentityPoolOutput) {
	if opDeleteIdentityPool == nil {
		opDeleteIdentityPool = &aws.Operation{
			Name:       "DeleteIdentityPool",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteIdentityPool, input, output)
	output = &DeleteIdentityPoolOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) DeleteIdentityPool(input *DeleteIdentityPoolInput) (output *DeleteIdentityPoolOutput, err error) {
	req, out := c.DeleteIdentityPoolRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteIdentityPool *aws.Operation

// DescribeIdentityRequest generates a request for the DescribeIdentity operation.
func (c *CognitoIdentity) DescribeIdentityRequest(input *DescribeIdentityInput) (req *aws.Request, output *IdentityDescription) {
	if opDescribeIdentity == nil {
		opDescribeIdentity = &aws.Operation{
			Name:       "DescribeIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeIdentity, input, output)
	output = &IdentityDescription{}
	req.Data = output
	return
}

func (c *CognitoIdentity) DescribeIdentity(input *DescribeIdentityInput) (output *IdentityDescription, err error) {
	req, out := c.DescribeIdentityRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeIdentity *aws.Operation

// DescribeIdentityPoolRequest generates a request for the DescribeIdentityPool operation.
func (c *CognitoIdentity) DescribeIdentityPoolRequest(input *DescribeIdentityPoolInput) (req *aws.Request, output *IdentityPool) {
	if opDescribeIdentityPool == nil {
		opDescribeIdentityPool = &aws.Operation{
			Name:       "DescribeIdentityPool",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeIdentityPool, input, output)
	output = &IdentityPool{}
	req.Data = output
	return
}

func (c *CognitoIdentity) DescribeIdentityPool(input *DescribeIdentityPoolInput) (output *IdentityPool, err error) {
	req, out := c.DescribeIdentityPoolRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeIdentityPool *aws.Operation

// GetCredentialsForIdentityRequest generates a request for the GetCredentialsForIdentity operation.
func (c *CognitoIdentity) GetCredentialsForIdentityRequest(input *GetCredentialsForIdentityInput) (req *aws.Request, output *GetCredentialsForIdentityOutput) {
	if opGetCredentialsForIdentity == nil {
		opGetCredentialsForIdentity = &aws.Operation{
			Name:       "GetCredentialsForIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetCredentialsForIdentity, input, output)
	output = &GetCredentialsForIdentityOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) GetCredentialsForIdentity(input *GetCredentialsForIdentityInput) (output *GetCredentialsForIdentityOutput, err error) {
	req, out := c.GetCredentialsForIdentityRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetCredentialsForIdentity *aws.Operation

// GetIDRequest generates a request for the GetID operation.
func (c *CognitoIdentity) GetIDRequest(input *GetIDInput) (req *aws.Request, output *GetIDOutput) {
	if opGetID == nil {
		opGetID = &aws.Operation{
			Name:       "GetId",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetID, input, output)
	output = &GetIDOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) GetID(input *GetIDInput) (output *GetIDOutput, err error) {
	req, out := c.GetIDRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetID *aws.Operation

// GetIdentityPoolRolesRequest generates a request for the GetIdentityPoolRoles operation.
func (c *CognitoIdentity) GetIdentityPoolRolesRequest(input *GetIdentityPoolRolesInput) (req *aws.Request, output *GetIdentityPoolRolesOutput) {
	if opGetIdentityPoolRoles == nil {
		opGetIdentityPoolRoles = &aws.Operation{
			Name:       "GetIdentityPoolRoles",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetIdentityPoolRoles, input, output)
	output = &GetIdentityPoolRolesOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) GetIdentityPoolRoles(input *GetIdentityPoolRolesInput) (output *GetIdentityPoolRolesOutput, err error) {
	req, out := c.GetIdentityPoolRolesRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetIdentityPoolRoles *aws.Operation

// GetOpenIDTokenRequest generates a request for the GetOpenIDToken operation.
func (c *CognitoIdentity) GetOpenIDTokenRequest(input *GetOpenIDTokenInput) (req *aws.Request, output *GetOpenIDTokenOutput) {
	if opGetOpenIDToken == nil {
		opGetOpenIDToken = &aws.Operation{
			Name:       "GetOpenIdToken",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetOpenIDToken, input, output)
	output = &GetOpenIDTokenOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) GetOpenIDToken(input *GetOpenIDTokenInput) (output *GetOpenIDTokenOutput, err error) {
	req, out := c.GetOpenIDTokenRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetOpenIDToken *aws.Operation

// GetOpenIDTokenForDeveloperIdentityRequest generates a request for the GetOpenIDTokenForDeveloperIdentity operation.
func (c *CognitoIdentity) GetOpenIDTokenForDeveloperIdentityRequest(input *GetOpenIDTokenForDeveloperIdentityInput) (req *aws.Request, output *GetOpenIDTokenForDeveloperIdentityOutput) {
	if opGetOpenIDTokenForDeveloperIdentity == nil {
		opGetOpenIDTokenForDeveloperIdentity = &aws.Operation{
			Name:       "GetOpenIdTokenForDeveloperIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetOpenIDTokenForDeveloperIdentity, input, output)
	output = &GetOpenIDTokenForDeveloperIdentityOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) GetOpenIDTokenForDeveloperIdentity(input *GetOpenIDTokenForDeveloperIdentityInput) (output *GetOpenIDTokenForDeveloperIdentityOutput, err error) {
	req, out := c.GetOpenIDTokenForDeveloperIdentityRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetOpenIDTokenForDeveloperIdentity *aws.Operation

// ListIdentitiesRequest generates a request for the ListIdentities operation.
func (c *CognitoIdentity) ListIdentitiesRequest(input *ListIdentitiesInput) (req *aws.Request, output *ListIdentitiesOutput) {
	if opListIdentities == nil {
		opListIdentities = &aws.Operation{
			Name:       "ListIdentities",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListIdentities, input, output)
	output = &ListIdentitiesOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) ListIdentities(input *ListIdentitiesInput) (output *ListIdentitiesOutput, err error) {
	req, out := c.ListIdentitiesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListIdentities *aws.Operation

// ListIdentityPoolsRequest generates a request for the ListIdentityPools operation.
func (c *CognitoIdentity) ListIdentityPoolsRequest(input *ListIdentityPoolsInput) (req *aws.Request, output *ListIdentityPoolsOutput) {
	if opListIdentityPools == nil {
		opListIdentityPools = &aws.Operation{
			Name:       "ListIdentityPools",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListIdentityPools, input, output)
	output = &ListIdentityPoolsOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) ListIdentityPools(input *ListIdentityPoolsInput) (output *ListIdentityPoolsOutput, err error) {
	req, out := c.ListIdentityPoolsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListIdentityPools *aws.Operation

// LookupDeveloperIdentityRequest generates a request for the LookupDeveloperIdentity operation.
func (c *CognitoIdentity) LookupDeveloperIdentityRequest(input *LookupDeveloperIdentityInput) (req *aws.Request, output *LookupDeveloperIdentityOutput) {
	if opLookupDeveloperIdentity == nil {
		opLookupDeveloperIdentity = &aws.Operation{
			Name:       "LookupDeveloperIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opLookupDeveloperIdentity, input, output)
	output = &LookupDeveloperIdentityOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) LookupDeveloperIdentity(input *LookupDeveloperIdentityInput) (output *LookupDeveloperIdentityOutput, err error) {
	req, out := c.LookupDeveloperIdentityRequest(input)
	output = out
	err = req.Send()
	return
}

var opLookupDeveloperIdentity *aws.Operation

// MergeDeveloperIdentitiesRequest generates a request for the MergeDeveloperIdentities operation.
func (c *CognitoIdentity) MergeDeveloperIdentitiesRequest(input *MergeDeveloperIdentitiesInput) (req *aws.Request, output *MergeDeveloperIdentitiesOutput) {
	if opMergeDeveloperIdentities == nil {
		opMergeDeveloperIdentities = &aws.Operation{
			Name:       "MergeDeveloperIdentities",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opMergeDeveloperIdentities, input, output)
	output = &MergeDeveloperIdentitiesOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) MergeDeveloperIdentities(input *MergeDeveloperIdentitiesInput) (output *MergeDeveloperIdentitiesOutput, err error) {
	req, out := c.MergeDeveloperIdentitiesRequest(input)
	output = out
	err = req.Send()
	return
}

var opMergeDeveloperIdentities *aws.Operation

// SetIdentityPoolRolesRequest generates a request for the SetIdentityPoolRoles operation.
func (c *CognitoIdentity) SetIdentityPoolRolesRequest(input *SetIdentityPoolRolesInput) (req *aws.Request, output *SetIdentityPoolRolesOutput) {
	if opSetIdentityPoolRoles == nil {
		opSetIdentityPoolRoles = &aws.Operation{
			Name:       "SetIdentityPoolRoles",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetIdentityPoolRoles, input, output)
	output = &SetIdentityPoolRolesOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) SetIdentityPoolRoles(input *SetIdentityPoolRolesInput) (output *SetIdentityPoolRolesOutput, err error) {
	req, out := c.SetIdentityPoolRolesRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetIdentityPoolRoles *aws.Operation

// UnlinkDeveloperIdentityRequest generates a request for the UnlinkDeveloperIdentity operation.
func (c *CognitoIdentity) UnlinkDeveloperIdentityRequest(input *UnlinkDeveloperIdentityInput) (req *aws.Request, output *UnlinkDeveloperIdentityOutput) {
	if opUnlinkDeveloperIdentity == nil {
		opUnlinkDeveloperIdentity = &aws.Operation{
			Name:       "UnlinkDeveloperIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUnlinkDeveloperIdentity, input, output)
	output = &UnlinkDeveloperIdentityOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) UnlinkDeveloperIdentity(input *UnlinkDeveloperIdentityInput) (output *UnlinkDeveloperIdentityOutput, err error) {
	req, out := c.UnlinkDeveloperIdentityRequest(input)
	output = out
	err = req.Send()
	return
}

var opUnlinkDeveloperIdentity *aws.Operation

// UnlinkIdentityRequest generates a request for the UnlinkIdentity operation.
func (c *CognitoIdentity) UnlinkIdentityRequest(input *UnlinkIdentityInput) (req *aws.Request, output *UnlinkIdentityOutput) {
	if opUnlinkIdentity == nil {
		opUnlinkIdentity = &aws.Operation{
			Name:       "UnlinkIdentity",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUnlinkIdentity, input, output)
	output = &UnlinkIdentityOutput{}
	req.Data = output
	return
}

func (c *CognitoIdentity) UnlinkIdentity(input *UnlinkIdentityInput) (output *UnlinkIdentityOutput, err error) {
	req, out := c.UnlinkIdentityRequest(input)
	output = out
	err = req.Send()
	return
}

var opUnlinkIdentity *aws.Operation

// UpdateIdentityPoolRequest generates a request for the UpdateIdentityPool operation.
func (c *CognitoIdentity) UpdateIdentityPoolRequest(input *IdentityPool) (req *aws.Request, output *IdentityPool) {
	if opUpdateIdentityPool == nil {
		opUpdateIdentityPool = &aws.Operation{
			Name:       "UpdateIdentityPool",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateIdentityPool, input, output)
	output = &IdentityPool{}
	req.Data = output
	return
}

func (c *CognitoIdentity) UpdateIdentityPool(input *IdentityPool) (output *IdentityPool, err error) {
	req, out := c.UpdateIdentityPoolRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateIdentityPool *aws.Operation

type CreateIdentityPoolInput struct {
	AllowUnauthenticatedIdentities *bool               `type:"boolean" json:",omitempty"`
	DeveloperProviderName          *string             `type:"string" json:",omitempty"`
	IdentityPoolName               *string             `type:"string" json:",omitempty"`
	OpenIDConnectProviderARNs      []*string           `locationName:"OpenIdConnectProviderARNs" type:"list" json:"OpenIdConnectProviderARNs,omitempty"`
	SupportedLoginProviders        *map[string]*string `type:"map" json:",omitempty"`

	metadataCreateIdentityPoolInput `json:"-", xml:"-"`
}

type metadataCreateIdentityPoolInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityPoolName,AllowUnauthenticatedIdentities" json:",omitempty"`
}

type Credentials struct {
	AccessKeyID  *string    `locationName:"AccessKeyId" type:"string" json:"AccessKeyId,omitempty"`
	Expiration   *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	SecretKey    *string    `type:"string" json:",omitempty"`
	SessionToken *string    `type:"string" json:",omitempty"`

	metadataCredentials `json:"-", xml:"-"`
}

type metadataCredentials struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteIdentityPoolInput struct {
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`

	metadataDeleteIdentityPoolInput `json:"-", xml:"-"`
}

type metadataDeleteIdentityPoolInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityPoolId" json:",omitempty"`
}

type DeleteIdentityPoolOutput struct {
	metadataDeleteIdentityPoolOutput `json:"-", xml:"-"`
}

type metadataDeleteIdentityPoolOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DescribeIdentityInput struct {
	IdentityID *string `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`

	metadataDescribeIdentityInput `json:"-", xml:"-"`
}

type metadataDescribeIdentityInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityId" json:",omitempty"`
}

type DescribeIdentityPoolInput struct {
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`

	metadataDescribeIdentityPoolInput `json:"-", xml:"-"`
}

type metadataDescribeIdentityPoolInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityPoolId" json:",omitempty"`
}

type GetCredentialsForIdentityInput struct {
	IdentityID *string             `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	Logins     *map[string]*string `type:"map" json:",omitempty"`

	metadataGetCredentialsForIdentityInput `json:"-", xml:"-"`
}

type metadataGetCredentialsForIdentityInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityId" json:",omitempty"`
}

type GetCredentialsForIdentityOutput struct {
	Credentials *Credentials `type:"structure" json:",omitempty"`
	IdentityID  *string      `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`

	metadataGetCredentialsForIdentityOutput `json:"-", xml:"-"`
}

type metadataGetCredentialsForIdentityOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetIDInput struct {
	AccountID      *string             `locationName:"AccountId" type:"string" json:"AccountId,omitempty"`
	IdentityPoolID *string             `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	Logins         *map[string]*string `type:"map" json:",omitempty"`

	metadataGetIDInput `json:"-", xml:"-"`
}

type metadataGetIDInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityPoolId" json:",omitempty"`
}

type GetIDOutput struct {
	IdentityID *string `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`

	metadataGetIDOutput `json:"-", xml:"-"`
}

type metadataGetIDOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetIdentityPoolRolesInput struct {
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`

	metadataGetIdentityPoolRolesInput `json:"-", xml:"-"`
}

type metadataGetIdentityPoolRolesInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetIdentityPoolRolesOutput struct {
	IdentityPoolID *string             `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	Roles          *map[string]*string `type:"map" json:",omitempty"`

	metadataGetIdentityPoolRolesOutput `json:"-", xml:"-"`
}

type metadataGetIdentityPoolRolesOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetOpenIDTokenForDeveloperIdentityInput struct {
	IdentityID     *string             `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	IdentityPoolID *string             `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	Logins         *map[string]*string `type:"map" json:",omitempty"`
	TokenDuration  *int64              `type:"long" json:",omitempty"`

	metadataGetOpenIDTokenForDeveloperIdentityInput `json:"-", xml:"-"`
}

type metadataGetOpenIDTokenForDeveloperIdentityInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityPoolId,Logins" json:",omitempty"`
}

type GetOpenIDTokenForDeveloperIdentityOutput struct {
	IdentityID *string `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	Token      *string `type:"string" json:",omitempty"`

	metadataGetOpenIDTokenForDeveloperIdentityOutput `json:"-", xml:"-"`
}

type metadataGetOpenIDTokenForDeveloperIdentityOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetOpenIDTokenInput struct {
	IdentityID *string             `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	Logins     *map[string]*string `type:"map" json:",omitempty"`

	metadataGetOpenIDTokenInput `json:"-", xml:"-"`
}

type metadataGetOpenIDTokenInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityId" json:",omitempty"`
}

type GetOpenIDTokenOutput struct {
	IdentityID *string `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	Token      *string `type:"string" json:",omitempty"`

	metadataGetOpenIDTokenOutput `json:"-", xml:"-"`
}

type metadataGetOpenIDTokenOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type IdentityDescription struct {
	CreationDate     *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	IdentityID       *string    `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	LastModifiedDate *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	Logins           []*string  `type:"list" json:",omitempty"`

	metadataIdentityDescription `json:"-", xml:"-"`
}

type metadataIdentityDescription struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type IdentityPool struct {
	AllowUnauthenticatedIdentities *bool               `type:"boolean" json:",omitempty"`
	DeveloperProviderName          *string             `type:"string" json:",omitempty"`
	IdentityPoolID                 *string             `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	IdentityPoolName               *string             `type:"string" json:",omitempty"`
	OpenIDConnectProviderARNs      []*string           `locationName:"OpenIdConnectProviderARNs" type:"list" json:"OpenIdConnectProviderARNs,omitempty"`
	SupportedLoginProviders        *map[string]*string `type:"map" json:",omitempty"`

	metadataIdentityPool `json:"-", xml:"-"`
}

type metadataIdentityPool struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityPoolId,IdentityPoolName,AllowUnauthenticatedIdentities" json:",omitempty"`
}

type IdentityPoolShortDescription struct {
	IdentityPoolID   *string `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	IdentityPoolName *string `type:"string" json:",omitempty"`

	metadataIdentityPoolShortDescription `json:"-", xml:"-"`
}

type metadataIdentityPoolShortDescription struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListIdentitiesInput struct {
	IdentityPoolID *string `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	MaxResults     *int    `type:"integer" json:",omitempty"`
	NextToken      *string `type:"string" json:",omitempty"`

	metadataListIdentitiesInput `json:"-", xml:"-"`
}

type metadataListIdentitiesInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityPoolId,MaxResults" json:",omitempty"`
}

type ListIdentitiesOutput struct {
	Identities     []*IdentityDescription `type:"list" json:",omitempty"`
	IdentityPoolID *string                `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	NextToken      *string                `type:"string" json:",omitempty"`

	metadataListIdentitiesOutput `json:"-", xml:"-"`
}

type metadataListIdentitiesOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListIdentityPoolsInput struct {
	MaxResults *int    `type:"integer" json:",omitempty"`
	NextToken  *string `type:"string" json:",omitempty"`

	metadataListIdentityPoolsInput `json:"-", xml:"-"`
}

type metadataListIdentityPoolsInput struct {
	SDKShapeTraits bool `type:"structure" required:"MaxResults" json:",omitempty"`
}

type ListIdentityPoolsOutput struct {
	IdentityPools []*IdentityPoolShortDescription `type:"list" json:",omitempty"`
	NextToken     *string                         `type:"string" json:",omitempty"`

	metadataListIdentityPoolsOutput `json:"-", xml:"-"`
}

type metadataListIdentityPoolsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type LookupDeveloperIdentityInput struct {
	DeveloperUserIdentifier *string `type:"string" json:",omitempty"`
	IdentityID              *string `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	IdentityPoolID          *string `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	MaxResults              *int    `type:"integer" json:",omitempty"`
	NextToken               *string `type:"string" json:",omitempty"`

	metadataLookupDeveloperIdentityInput `json:"-", xml:"-"`
}

type metadataLookupDeveloperIdentityInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityPoolId" json:",omitempty"`
}

type LookupDeveloperIdentityOutput struct {
	DeveloperUserIdentifierList []*string `type:"list" json:",omitempty"`
	IdentityID                  *string   `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	NextToken                   *string   `type:"string" json:",omitempty"`

	metadataLookupDeveloperIdentityOutput `json:"-", xml:"-"`
}

type metadataLookupDeveloperIdentityOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type MergeDeveloperIdentitiesInput struct {
	DestinationUserIdentifier *string `type:"string" json:",omitempty"`
	DeveloperProviderName     *string `type:"string" json:",omitempty"`
	IdentityPoolID            *string `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	SourceUserIdentifier      *string `type:"string" json:",omitempty"`

	metadataMergeDeveloperIdentitiesInput `json:"-", xml:"-"`
}

type metadataMergeDeveloperIdentitiesInput struct {
	SDKShapeTraits bool `type:"structure" required:"SourceUserIdentifier,DestinationUserIdentifier,DeveloperProviderName,IdentityPoolId" json:",omitempty"`
}

type MergeDeveloperIdentitiesOutput struct {
	IdentityID *string `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`

	metadataMergeDeveloperIdentitiesOutput `json:"-", xml:"-"`
}

type metadataMergeDeveloperIdentitiesOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type SetIdentityPoolRolesInput struct {
	IdentityPoolID *string             `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`
	Roles          *map[string]*string `type:"map" json:",omitempty"`

	metadataSetIdentityPoolRolesInput `json:"-", xml:"-"`
}

type metadataSetIdentityPoolRolesInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityPoolId,Roles" json:",omitempty"`
}

type SetIdentityPoolRolesOutput struct {
	metadataSetIdentityPoolRolesOutput `json:"-", xml:"-"`
}

type metadataSetIdentityPoolRolesOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UnlinkDeveloperIdentityInput struct {
	DeveloperProviderName   *string `type:"string" json:",omitempty"`
	DeveloperUserIdentifier *string `type:"string" json:",omitempty"`
	IdentityID              *string `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	IdentityPoolID          *string `locationName:"IdentityPoolId" type:"string" json:"IdentityPoolId,omitempty"`

	metadataUnlinkDeveloperIdentityInput `json:"-", xml:"-"`
}

type metadataUnlinkDeveloperIdentityInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityId,IdentityPoolId,DeveloperProviderName,DeveloperUserIdentifier" json:",omitempty"`
}

type UnlinkDeveloperIdentityOutput struct {
	metadataUnlinkDeveloperIdentityOutput `json:"-", xml:"-"`
}

type metadataUnlinkDeveloperIdentityOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UnlinkIdentityInput struct {
	IdentityID     *string             `locationName:"IdentityId" type:"string" json:"IdentityId,omitempty"`
	Logins         *map[string]*string `type:"map" json:",omitempty"`
	LoginsToRemove []*string           `type:"list" json:",omitempty"`

	metadataUnlinkIdentityInput `json:"-", xml:"-"`
}

type metadataUnlinkIdentityInput struct {
	SDKShapeTraits bool `type:"structure" required:"IdentityId,Logins,LoginsToRemove" json:",omitempty"`
}

type UnlinkIdentityOutput struct {
	metadataUnlinkIdentityOutput `json:"-", xml:"-"`
}

type metadataUnlinkIdentityOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}