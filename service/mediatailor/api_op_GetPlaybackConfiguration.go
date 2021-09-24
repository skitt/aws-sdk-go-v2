// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the playback configuration for the specified name.
func (c *Client) GetPlaybackConfiguration(ctx context.Context, params *GetPlaybackConfigurationInput, optFns ...func(*Options)) (*GetPlaybackConfigurationOutput, error) {
	if params == nil {
		params = &GetPlaybackConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPlaybackConfiguration", params, optFns, c.addOperationGetPlaybackConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPlaybackConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPlaybackConfigurationInput struct {

	// The identifier for the playback configuration.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetPlaybackConfigurationOutput struct {

	// The URL for the ad decision server (ADS). This includes the specification of
	// static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing, you can provide a static
	// VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string

	// The configuration for avail suppression, also known as ad suppression. For more
	// information about ad suppression, see Ad Suppression
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	AvailSuppression *types.AvailSuppression

	// The configuration for bumpers. Bumpers are short audio or video clips that play
	// at the start or before the end of an ad break. To learn more about bumpers, see
	// Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
	Bumper *types.Bumper

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *types.CdnConfiguration

	// The player parameters and aliases used as dynamic variables during session
	// initialization. For more information, see Domain Variables
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html).
	ConfigurationAliases map[string]map[string]string

	// The configuration for DASH content.
	DashConfiguration *types.DashConfiguration

	// The configuration for HLS content.
	HlsConfiguration *types.HlsConfiguration

	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *types.LivePreRollConfiguration

	// The Amazon CloudWatch log settings for a playback configuration.
	LogConfiguration *types.LogConfiguration

	// The configuration for manifest processing rules. Manifest processing rules
	// enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules *types.ManifestProcessingRules

	// The identifier for the playback configuration.
	Name *string

	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an
	// ad break. If the duration of underfilled ad time exceeds the personalization
	// threshold, then the personalization of the ad break is abandoned and the
	// underlying content is shown. This feature applies to ad replacement in live and
	// VOD streams, rather than ad insertion, because it relies on an underlying
	// content stream. For more information about ad break behavior, including ad
	// replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	PersonalizationThresholdSeconds int32

	// The Amazon Resource Name (ARN) for the playback configuration.
	PlaybackConfigurationArn *string

	// The URL that the player accesses to get a manifest from AWS Elemental
	// MediaTailor. This session will use server-side reporting.
	PlaybackEndpointPrefix *string

	// The URL that the player uses to initialize a session that uses client-side
	// reporting.
	SessionInitializationEndpointPrefix *string

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in
	// gaps in media content. Configuring the slate is optional for non-VPAID playback
	// configurations. For VPAID, the slate is required because MediaTailor provides it
	// in the slots designated for dynamic ad content. The slate must be a high-quality
	// asset that contains both audio and video.
	SlateAdUrl *string

	// The tags assigned to the playback configuration.
	Tags map[string]string

	// The name that is used to associate this playback configuration with a custom
	// transcode profile. This overrides the dynamic transcoding defaults of
	// MediaTailor. Use this only if you have already set up custom profiles with the
	// help of AWS Support.
	TranscodeProfileName *string

	// The URL prefix for the parent manifest for the stream, minus the asset ID. The
	// maximum length is 512 characters.
	VideoContentSourceUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPlaybackConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPlaybackConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPlaybackConfiguration{}, middleware.After)
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
	if err = addOpGetPlaybackConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPlaybackConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPlaybackConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "GetPlaybackConfiguration",
	}
}
