// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an entity recognizer using submitted files. After your
// CreateEntityRecognizer request is submitted, you can check job status using the
// API.
func (c *Client) CreateEntityRecognizer(ctx context.Context, params *CreateEntityRecognizerInput, optFns ...func(*Options)) (*CreateEntityRecognizerOutput, error) {
	if params == nil {
		params = &CreateEntityRecognizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEntityRecognizer", params, optFns, c.addOperationCreateEntityRecognizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEntityRecognizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEntityRecognizerInput struct {

	// The Amazon Resource Name (ARN) of the AWS Identity and Management (IAM) role
	// that grants Amazon Comprehend read access to your input data.
	//
	// This member is required.
	DataAccessRoleArn *string

	// Specifies the format and location of the input data. The S3 bucket containing
	// the input data must be located in the same region as the entity recognizer being
	// created.
	//
	// This member is required.
	InputDataConfig *types.EntityRecognizerInputDataConfig

	// You can specify any of the following languages supported by Amazon Comprehend:
	// English ("en"), Spanish ("es"), French ("fr"), Italian ("it"), German ("de"), or
	// Portuguese ("pt"). All documents must be in the same language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The name given to the newly created recognizer. Recognizer names can be a
	// maximum of 256 characters. Alphanumeric characters, hyphens (-) and underscores
	// (_) are allowed. The name must be unique in the account/region.
	//
	// This member is required.
	RecognizerName *string

	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend generates one.
	ClientRequestToken *string

	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to
	// encrypt trained custom models. The ModelKmsKeyId can be either of the following
	// formats
	//
	// * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	// * Amazon Resource
	// Name (ARN) of a KMS Key:
	// "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	ModelKmsKeyId *string

	// Tags to be associated with the entity recognizer being created. A tag is a
	// key-value pair that adds as a metadata to a resource used by Amazon Comprehend.
	// For example, a tag with "Sales" as the key might be added to a resource to
	// indicate its use by the sales department.
	Tags []types.Tag

	// The version name given to the newly created recognizer. Version names can be a
	// maximum of 256 characters. Alphanumeric characters, hyphens (-) and underscores
	// (_) are allowed. The version name must be unique among all models with the same
	// recognizer name in the account/ AWS Region.
	VersionName *string

	// ID for the AWS Key Management Service (KMS) key that Amazon Comprehend uses to
	// encrypt data on the storage volume attached to the ML compute instance(s) that
	// process the analysis job. The VolumeKmsKeyId can be either of the following
	// formats:
	//
	// * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	// * Amazon
	// Resource Name (ARN) of a KMS Key:
	// "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	VolumeKmsKeyId *string

	// Configuration parameters for an optional private Virtual Private Cloud (VPC)
	// containing the resources you are using for your custom entity recognizer. For
	// more information, see Amazon VPC
	// (https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html).
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type CreateEntityRecognizerOutput struct {

	// The Amazon Resource Name (ARN) that identifies the entity recognizer.
	EntityRecognizerArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEntityRecognizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEntityRecognizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEntityRecognizer{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateEntityRecognizerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEntityRecognizerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEntityRecognizer(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateEntityRecognizer struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateEntityRecognizer) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateEntityRecognizer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateEntityRecognizerInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateEntityRecognizerInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateEntityRecognizerMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateEntityRecognizer{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateEntityRecognizer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "CreateEntityRecognizer",
	}
}
