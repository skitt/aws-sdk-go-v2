// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Describes an additional authentication provider.
type AdditionalAuthenticationProvider struct {

	// The authentication type: API key, Identity and Access Management, OIDC, Amazon
	// Cognito user pools, or Amazon Web Services Lambda.
	AuthenticationType AuthenticationType

	// Configuration for Amazon Web Services Lambda function authorization.
	LambdaAuthorizerConfig *LambdaAuthorizerConfig

	// The OpenID Connect configuration.
	OpenIDConnectConfig *OpenIDConnectConfig

	// The Amazon Cognito user pool configuration.
	UserPoolConfig *CognitoUserPoolConfig

	noSmithyDocumentSerde
}

// The ApiCache object.
type ApiCache struct {

	// Caching behavior.
	//
	// * FULL_REQUEST_CACHING: All requests are fully cached.
	//
	// *
	// PER_RESOLVER_CACHING: Individual resolvers that you specify are cached.
	ApiCachingBehavior ApiCachingBehavior

	// At rest encryption flag for cache. This setting cannot be updated after
	// creation.
	AtRestEncryptionEnabled bool

	// The cache instance status.
	//
	// * AVAILABLE: The instance is available for use.
	//
	// *
	// CREATING: The instance is currently creating.
	//
	// * DELETING: The instance is
	// currently deleting.
	//
	// * MODIFYING: The instance is currently modifying.
	//
	// *
	// FAILED: The instance has failed creation.
	Status ApiCacheStatus

	// Transit encryption flag when connecting to cache. This setting cannot be updated
	// after creation.
	TransitEncryptionEnabled bool

	// TTL in seconds for cache entries. Valid values are between 1 and 3600 seconds.
	Ttl int64

	// The cache instance type. Valid values are
	//
	// * SMALL
	//
	// * MEDIUM
	//
	// * LARGE
	//
	// *
	// XLARGE
	//
	// * LARGE_2X
	//
	// * LARGE_4X
	//
	// * LARGE_8X (not available in all regions)
	//
	// *
	// LARGE_12X
	//
	// Historically, instance types were identified by an EC2-style value.
	// As of July 2020, this is deprecated, and the generic identifiers above should be
	// used. The following legacy instance types are available, but their use is
	// discouraged:
	//
	// * T2_SMALL: A t2.small instance type.
	//
	// * T2_MEDIUM: A t2.medium
	// instance type.
	//
	// * R4_LARGE: A r4.large instance type.
	//
	// * R4_XLARGE: A r4.xlarge
	// instance type.
	//
	// * R4_2XLARGE: A r4.2xlarge instance type.
	//
	// * R4_4XLARGE: A
	// r4.4xlarge instance type.
	//
	// * R4_8XLARGE: A r4.8xlarge instance type.
	Type ApiCacheType

	noSmithyDocumentSerde
}

// Describes an API key. Customers invoke AppSync GraphQL API operations with API
// keys as an identity mechanism. There are two key versions: da1: This version was
// introduced at launch in November 2017. These keys always expire after 7 days.
// Key expiration is managed by Amazon DynamoDB TTL. The keys ceased to be valid
// after February 21, 2018 and should not be used after that date.
//
// * ListApiKeys
// returns the expiration time in milliseconds.
//
// * CreateApiKey returns the
// expiration time in milliseconds.
//
// * UpdateApiKey is not available for this key
// version.
//
// * DeleteApiKey deletes the item from the table.
//
// * Expiration is
// stored in Amazon DynamoDB as milliseconds. This results in a bug where keys are
// not automatically deleted because DynamoDB expects the TTL to be stored in
// seconds. As a one-time action, we will delete these keys from the table after
// February 21, 2018.
//
// da2: This version was introduced in February 2018 when
// AppSync added support to extend key expiration.
//
// * ListApiKeys returns the
// expiration time and deletion time in seconds.
//
// * CreateApiKey returns the
// expiration time and deletion time in seconds and accepts a user-provided
// expiration time in seconds.
//
// * UpdateApiKey returns the expiration time and and
// deletion time in seconds and accepts a user-provided expiration time in seconds.
// Expired API keys are kept for 60 days after the expiration time. Key expiration
// time can be updated while the key is not deleted.
//
// * DeleteApiKey deletes the
// item from the table.
//
// * Expiration is stored in Amazon DynamoDB as seconds.
// After the expiration time, using the key to authenticate will fail. But the key
// can be reinstated before deletion.
//
// * Deletion is stored in Amazon DynamoDB as
// seconds. The key will be deleted after deletion time.
type ApiKey struct {

	// The time after which the API key is deleted. The date is represented as seconds
	// since the epoch, rounded down to the nearest hour.
	Deletes int64

	// A description of the purpose of the API key.
	Description *string

	// The time after which the API key expires. The date is represented as seconds
	// since the epoch, rounded down to the nearest hour.
	Expires int64

	// The API key ID.
	Id *string

	noSmithyDocumentSerde
}

// The authorization config in case the HTTP endpoint requires authorization.
type AuthorizationConfig struct {

	// The authorization type required by the HTTP endpoint.
	//
	// * AWS_IAM: The
	// authorization type is Sigv4.
	//
	// This member is required.
	AuthorizationType AuthorizationType

	// The Identity and Access Management settings.
	AwsIamConfig *AwsIamConfig

	noSmithyDocumentSerde
}

// The Identity and Access Management configuration.
type AwsIamConfig struct {

	// The signing region for Identity and Access Management authorization.
	SigningRegion *string

	// The signing service name for Identity and Access Management authorization.
	SigningServiceName *string

	noSmithyDocumentSerde
}

// The caching configuration for a resolver that has caching enabled.
type CachingConfig struct {

	// The caching keys for a resolver that has caching enabled. Valid values are
	// entries from the $context.arguments, $context.source, and $context.identity
	// maps.
	CachingKeys []string

	// The TTL in seconds for a resolver that has caching enabled. Valid values are
	// between 1 and 3600 seconds.
	Ttl int64

	noSmithyDocumentSerde
}

// Describes an Amazon Cognito user pool configuration.
type CognitoUserPoolConfig struct {

	// The Amazon Web Services Region in which the user pool was created.
	//
	// This member is required.
	AwsRegion *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// A regular expression for validating the incoming Amazon Cognito user pool app
	// client ID.
	AppIdClientRegex *string

	noSmithyDocumentSerde
}

// Describes a data source.
type DataSource struct {

	// The data source ARN.
	DataSourceArn *string

	// The description of the data source.
	Description *string

	// Amazon DynamoDB settings.
	DynamodbConfig *DynamodbDataSourceConfig

	// Amazon OpenSearch Service settings.
	ElasticsearchConfig *ElasticsearchDataSourceConfig

	// HTTP endpoint settings.
	HttpConfig *HttpDataSourceConfig

	// Amazon Web Services Lambda settings.
	LambdaConfig *LambdaDataSourceConfig

	// The name of the data source.
	Name *string

	// Amazon OpenSearch Service settings.
	OpenSearchServiceConfig *OpenSearchServiceDataSourceConfig

	// Relational database settings.
	RelationalDatabaseConfig *RelationalDatabaseDataSourceConfig

	// The Identity and Access Management service role ARN for the data source. The
	// system assumes this role when accessing the data source.
	ServiceRoleArn *string

	// The type of the data source.
	//
	// * AWS_LAMBDA: The data source is an Amazon Web
	// Services Lambda function.
	//
	// * AMAZON_DYNAMODB: The data source is an Amazon
	// DynamoDB table.
	//
	// * AMAZON_ELASTICSEARCH: The data source is an Amazon OpenSearch
	// Service domain.
	//
	// * AMAZON_OPENSEARCH_SERVICE: The data source is an Amazon
	// OpenSearch Service domain.
	//
	// * NONE: There is no data source. This type is used
	// when you wish to invoke a GraphQL operation without connecting to a data source,
	// such as performing data transformation with resolvers or triggering a
	// subscription to be invoked from a mutation.
	//
	// * HTTP: The data source is an HTTP
	// endpoint.
	//
	// * RELATIONAL_DATABASE: The data source is a relational database.
	Type DataSourceType

	noSmithyDocumentSerde
}

// Describes a Delta Sync configuration.
type DeltaSyncConfig struct {

	// The number of minutes an Item is stored in the datasource.
	BaseTableTTL int64

	// The Delta Sync table name.
	DeltaSyncTableName *string

	// The number of minutes a Delta Sync log entry is stored in the Delta Sync table.
	DeltaSyncTableTTL int64

	noSmithyDocumentSerde
}

// Describes an Amazon DynamoDB data source configuration.
type DynamodbDataSourceConfig struct {

	// The Amazon Web Services Region.
	//
	// This member is required.
	AwsRegion *string

	// The table name.
	//
	// This member is required.
	TableName *string

	// The DeltaSyncConfig for a versioned datasource.
	DeltaSyncConfig *DeltaSyncConfig

	// Set to TRUE to use Amazon Cognito credentials with this data source.
	UseCallerCredentials bool

	// Set to TRUE to use Conflict Detection and Resolution with this data source.
	Versioned bool

	noSmithyDocumentSerde
}

// Describes an OpenSearch data source configuration. As of September 2021, Amazon
// Elasticsearch service is Amazon OpenSearch Service. This configuration is
// deprecated. For new data sources, use OpenSearchServiceDataSourceConfig to
// specify an OpenSearch data source.
type ElasticsearchDataSourceConfig struct {

	// The Amazon Web Services Region.
	//
	// This member is required.
	AwsRegion *string

	// The endpoint.
	//
	// This member is required.
	Endpoint *string

	noSmithyDocumentSerde
}

// A function is a reusable entity. Multiple functions can be used to compose the
// resolver logic.
type FunctionConfiguration struct {

	// The name of the DataSource.
	DataSourceName *string

	// The Function description.
	Description *string

	// The ARN of the Function object.
	FunctionArn *string

	// A unique ID representing the Function object.
	FunctionId *string

	// The version of the request mapping template. Currently only the 2018-05-29
	// version of the template is supported.
	FunctionVersion *string

	// The name of the Function object.
	Name *string

	// The Function request mapping template. Functions support only the 2018-05-29
	// version of the request mapping template.
	RequestMappingTemplate *string

	// The Function response mapping template.
	ResponseMappingTemplate *string

	// Describes a Sync configuration for a resolver. Contains information on which
	// Conflict Detection as well as Resolution strategy should be performed when the
	// resolver is invoked.
	SyncConfig *SyncConfig

	noSmithyDocumentSerde
}

// Describes a GraphQL API.
type GraphqlApi struct {

	// A list of additional authentication providers for the GraphqlApi API.
	AdditionalAuthenticationProviders []AdditionalAuthenticationProvider

	// The API ID.
	ApiId *string

	// The ARN.
	Arn *string

	// The authentication type.
	AuthenticationType AuthenticationType

	// Configuration for Amazon Web Services Lambda function authorization.
	LambdaAuthorizerConfig *LambdaAuthorizerConfig

	// The Amazon CloudWatch Logs configuration.
	LogConfig *LogConfig

	// The API name.
	Name *string

	// The OpenID Connect configuration.
	OpenIDConnectConfig *OpenIDConnectConfig

	// The tags.
	Tags map[string]string

	// The URIs.
	Uris map[string]string

	// The Amazon Cognito user pool configuration.
	UserPoolConfig *UserPoolConfig

	// The ARN of the WAF ACL associated with this GraphqlApi, if one exists.
	WafWebAclArn *string

	// A flag representing whether X-Ray tracing is enabled for this GraphqlApi.
	XrayEnabled bool

	noSmithyDocumentSerde
}

// Describes an HTTP data source configuration.
type HttpDataSourceConfig struct {

	// The authorization config in case the HTTP endpoint requires authorization.
	AuthorizationConfig *AuthorizationConfig

	// The HTTP URL endpoint. You can either specify the domain name or IP, and port
	// combination, and the URL scheme must be HTTP or HTTPS. If the port is not
	// specified, AppSync uses the default port 80 for the HTTP endpoint and port 443
	// for HTTPS endpoints.
	Endpoint *string

	noSmithyDocumentSerde
}

// A LambdaAuthorizerConfig holds configuration on how to authorize AppSync API
// access when using the AWS_LAMBDA authorizer mode. Be aware that an AppSync API
// may have only one Lambda authorizer configured at a time.
type LambdaAuthorizerConfig struct {

	// The ARN of the Lambda function to be called for authorization. This may be a
	// standard Lambda ARN, a version ARN (.../v3) or alias ARN. Note: This Lambda
	// function must have the following resource-based policy assigned to it. When
	// configuring Lambda authorizers in the Console, this is done for you. To do so
	// with the Amazon Web Services CLI, run the following: aws lambda add-permission
	// --function-name "arn:aws:lambda:us-east-2:111122223333:function:my-function"
	// --statement-id "appsync" --principal appsync.amazonaws.com --action
	// lambda:InvokeFunction
	//
	// This member is required.
	AuthorizerUri *string

	// The number of seconds a response should be cached for. The default is 5 minutes
	// (300 seconds). The Lambda function can override this by returning a ttlOverride
	// key in its response. A value of 0 disables caching of responses.
	AuthorizerResultTtlInSeconds int32

	// A regular expression for validation of tokens before the Lambda function is
	// called.
	IdentityValidationExpression *string

	noSmithyDocumentSerde
}

// The LambdaConflictHandlerConfig object when configuring LAMBDA as the Conflict
// Handler.
type LambdaConflictHandlerConfig struct {

	// The Arn for the Lambda function to use as the Conflict Handler.
	LambdaConflictHandlerArn *string

	noSmithyDocumentSerde
}

// Describes an Amazon Web Services Lambda data source configuration.
type LambdaDataSourceConfig struct {

	// The ARN for the Lambda function.
	//
	// This member is required.
	LambdaFunctionArn *string

	noSmithyDocumentSerde
}

// The CloudWatch Logs configuration.
type LogConfig struct {

	// The service role that AppSync will assume to publish to Amazon CloudWatch logs
	// in your account.
	//
	// This member is required.
	CloudWatchLogsRoleArn *string

	// The field logging level. Values can be NONE, ERROR, or ALL.
	//
	// * NONE: No
	// field-level logs are captured.
	//
	// * ERROR: Logs the following information only for
	// the fields that are in error:
	//
	// * The error section in the server response.
	//
	// *
	// Field-level errors.
	//
	// * The generated request/response functions that got
	// resolved for error fields.
	//
	// * ALL: The following information is logged for all
	// fields in the query:
	//
	// * Field-level tracing information.
	//
	// * The generated
	// request/response functions that got resolved for each field.
	//
	// This member is required.
	FieldLogLevel FieldLogLevel

	// Set to TRUE to exclude sections that contain information such as headers,
	// context, and evaluated mapping templates, regardless of logging level.
	ExcludeVerboseContent bool

	noSmithyDocumentSerde
}

// Describes an OpenID Connect configuration.
type OpenIDConnectConfig struct {

	// The issuer for the OpenID Connect configuration. The issuer returned by
	// discovery must exactly match the value of iss in the ID token.
	//
	// This member is required.
	Issuer *string

	// The number of milliseconds a token is valid after being authenticated.
	AuthTTL int64

	// The client identifier of the Relying party at the OpenID identity provider. This
	// identifier is typically obtained when the Relying party is registered with the
	// OpenID identity provider. You can specify a regular expression so the AppSync
	// can validate against multiple client identifiers at a time.
	ClientId *string

	// The number of milliseconds a token is valid after being issued to a user.
	IatTTL int64

	noSmithyDocumentSerde
}

// Describes an OpenSearch data source configuration.
type OpenSearchServiceDataSourceConfig struct {

	// The Amazon Web Services Region.
	//
	// This member is required.
	AwsRegion *string

	// The endpoint.
	//
	// This member is required.
	Endpoint *string

	noSmithyDocumentSerde
}

// The pipeline configuration for a resolver of kind PIPELINE.
type PipelineConfig struct {

	// A list of Function objects.
	Functions []string

	noSmithyDocumentSerde
}

// The Amazon RDS HTTP endpoint configuration.
type RdsHttpEndpointConfig struct {

	// Amazon Web Services Region for RDS HTTP endpoint.
	AwsRegion *string

	// Amazon Web Services secret store ARN for database credentials.
	AwsSecretStoreArn *string

	// Logical database name.
	DatabaseName *string

	// Amazon RDS cluster ARN.
	DbClusterIdentifier *string

	// Logical schema name.
	Schema *string

	noSmithyDocumentSerde
}

// Describes a relational database data source configuration.
type RelationalDatabaseDataSourceConfig struct {

	// Amazon RDS HTTP endpoint settings.
	RdsHttpEndpointConfig *RdsHttpEndpointConfig

	// Source type for the relational database.
	//
	// * RDS_HTTP_ENDPOINT: The relational
	// database source type is an Amazon RDS HTTP endpoint.
	RelationalDatabaseSourceType RelationalDatabaseSourceType

	noSmithyDocumentSerde
}

// Describes a resolver.
type Resolver struct {

	// The caching configuration for the resolver.
	CachingConfig *CachingConfig

	// The resolver data source name.
	DataSourceName *string

	// The resolver field name.
	FieldName *string

	// The resolver type.
	//
	// * UNIT: A UNIT resolver type. A UNIT resolver is the default
	// resolver type. A UNIT resolver enables you to execute a GraphQL query against a
	// single data source.
	//
	// * PIPELINE: A PIPELINE resolver type. A PIPELINE resolver
	// enables you to execute a series of Function in a serial manner. You can use a
	// pipeline resolver to execute a GraphQL query against multiple data sources.
	Kind ResolverKind

	// The PipelineConfig.
	PipelineConfig *PipelineConfig

	// The request mapping template.
	RequestMappingTemplate *string

	// The resolver ARN.
	ResolverArn *string

	// The response mapping template.
	ResponseMappingTemplate *string

	// The SyncConfig for a resolver attached to a versioned datasource.
	SyncConfig *SyncConfig

	// The resolver type name.
	TypeName *string

	noSmithyDocumentSerde
}

// Describes a Sync configuration for a resolver. Contains information on which
// Conflict Detection as well as Resolution strategy should be performed when the
// resolver is invoked.
type SyncConfig struct {

	// The Conflict Detection strategy to use.
	//
	// * VERSION: Detect conflicts based on
	// object versions for this resolver.
	//
	// * NONE: Do not detect conflicts when
	// executing this resolver.
	ConflictDetection ConflictDetectionType

	// The Conflict Resolution strategy to perform in the event of a conflict.
	//
	// *
	// OPTIMISTIC_CONCURRENCY: Resolve conflicts by rejecting mutations when versions
	// do not match the latest version at the server.
	//
	// * AUTOMERGE: Resolve conflicts
	// with the Automerge conflict resolution strategy.
	//
	// * LAMBDA: Resolve conflicts
	// with a Lambda function supplied in the LambdaConflictHandlerConfig.
	ConflictHandler ConflictHandlerType

	// The LambdaConflictHandlerConfig when configuring LAMBDA as the Conflict Handler.
	LambdaConflictHandlerConfig *LambdaConflictHandlerConfig

	noSmithyDocumentSerde
}

// Describes a type.
type Type struct {

	// The type ARN.
	Arn *string

	// The type definition.
	Definition *string

	// The type description.
	Description *string

	// The type format: SDL or JSON.
	Format TypeDefinitionFormat

	// The type name.
	Name *string

	noSmithyDocumentSerde
}

// Describes an Amazon Cognito user pool configuration.
type UserPoolConfig struct {

	// The Amazon Web Services Region in which the user pool was created.
	//
	// This member is required.
	AwsRegion *string

	// The action that you want your GraphQL API to take when a request that uses
	// Amazon Cognito user pool authentication doesn't match the Amazon Cognito user
	// pool configuration.
	//
	// This member is required.
	DefaultAction DefaultAction

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// A regular expression for validating the incoming Amazon Cognito user pool app
	// client ID.
	AppIdClientRegex *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
