// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Access configuration parameters.
type AccessConfiguration struct {

	// The type of authentication used to access content from
	// HttpConfiguration::BaseUrl on your source location. Accepted value: S3_SIGV4.
	// S3_SIGV4 - AWS Signature Version 4 authentication for Amazon S3 hosted
	// virtual-style access. If your source location base URL is an Amazon S3 bucket,
	// MediaTailor can use AWS Signature Version 4 (SigV4) authentication to access the
	// bucket where your source content is stored. Your MediaTailor source location
	// baseURL must follow the S3 virtual hosted-style request URL format. For example,
	// https://bucket-name.s3.Region.amazonaws.com/key-name. Before you can use
	// S3_SIGV4, you must meet these requirements: • You must allow MediaTailor to
	// access your S3 bucket by granting mediatailor.amazonaws.com principal access in
	// IAM. For information about configuring access in IAM, see Access management in
	// the IAM User Guide. • The mediatailor.amazonaws.com service principal must have
	// permissions to read all top level manifests referenced by the VodSource
	// packaging configurations. • The caller of the API must have s3:GetObject IAM
	// permissions to read all top level manifests referenced by your MediaTailor
	// VodSource packaging configurations.
	AccessType AccessType

	// AWS Secrets Manager access token configuration parameters.
	SecretsManagerAccessTokenConfiguration *SecretsManagerAccessTokenConfiguration

	noSmithyDocumentSerde
}

// Ad break configuration parameters.
type AdBreak struct {

	// The SCTE-35 ad insertion type. Accepted value: SPLICE_INSERT.
	MessageType MessageType

	// How long (in milliseconds) after the beginning of the program that an ad starts.
	// This value must fall within 100ms of a segment boundary, otherwise the ad break
	// will be skipped.
	OffsetMillis int64

	// Ad break slate configuration.
	Slate *SlateSource

	// This defines the SCTE-35 splice_insert() message inserted around the ad. For
	// information about using splice_insert(), see the SCTE-35 specficiaiton, section
	// 9.7.3.1.
	SpliceInsertMessage *SpliceInsertMessage

	noSmithyDocumentSerde
}

// For HLS, when set to true, MediaTailor passes through EXT-X-CUE-IN,
// EXT-X-CUE-OUT, and EXT-X-SPLICEPOINT-SCTE35 ad markers from the origin manifest
// to the MediaTailor personalized manifest. No logic is applied to these ad
// markers. For example, if EXT-X-CUE-OUT has a value of 60, but no ads are filled
// for that ad break, MediaTailor will not set the value to 0.
type AdMarkerPassthrough struct {

	// Enables ad marker passthrough for your configuration.
	Enabled bool

	noSmithyDocumentSerde
}

// Alert configuration parameters.
type Alert struct {

	// The code for the alert. For example, NOT_PROCESSED.
	//
	// This member is required.
	AlertCode *string

	// If an alert is generated for a resource, an explanation of the reason for the
	// alert.
	//
	// This member is required.
	AlertMessage *string

	// The timestamp when the alert was last modified.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The Amazon Resource Names (ARNs) related to this alert.
	//
	// This member is required.
	RelatedResourceArns []string

	// The Amazon Resource Name (ARN) of the resource.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

// The configuration for avail suppression, also known as ad suppression. For more
// information about ad suppression, see Ad Suppression
// (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
type AvailSuppression struct {

	// Sets the ad suppression mode. By default, ad suppression is off and all ad
	// breaks are filled with ads or slate. When Mode is set to BEHIND_LIVE_EDGE, ad
	// suppression is active and MediaTailor won't fill ad breaks on or behind the ad
	// suppression Value time in the manifest lookback window.
	Mode Mode

	// A live edge offset time in HH:MM:SS. MediaTailor won't fill ad breaks on or
	// behind this time in the manifest lookback window. If Value is set to 00:00:00,
	// it is in sync with the live edge, and MediaTailor won't fill any ad breaks on or
	// behind the live edge. If you set a Value time, MediaTailor won't fill any ad
	// breaks on or behind this time in the manifest lookback window. For example, if
	// you set 00:45:00, then MediaTailor will fill ad breaks that occur within 45
	// minutes behind the live edge, but won't fill ad breaks on or behind 45 minutes
	// behind the live edge.
	Value *string

	noSmithyDocumentSerde
}

// The configuration for bumpers. Bumpers are short audio or video clips that play
// at the start or before the end of an ad break. To learn more about bumpers, see
// Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
type Bumper struct {

	// The URL for the end bumper asset.
	EndUrl *string

	// The URL for the start bumper asset.
	StartUrl *string

	noSmithyDocumentSerde
}

// The configuration for using a content delivery network (CDN), like Amazon
// CloudFront, for content and ad segment management.
type CdnConfiguration struct {

	// A non-default content delivery network (CDN) to serve ad segments. By default,
	// AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings as
	// its CDN for ad segments. To set up an alternate CDN, create a rule in your CDN
	// for the origin ads.mediatailor.&lt;region>.amazonaws.com. Then specify the
	// rule's name in this AdSegmentUrlPrefix. When AWS Elemental MediaTailor serves a
	// manifest, it reports your CDN as the source for ad segments.
	AdSegmentUrlPrefix *string

	// A content delivery network (CDN) to cache content segments, so that content
	// requests don’t always have to go to the origin server. First, create a rule in
	// your CDN for the content segment origin server. Then specify the rule's name in
	// this ContentSegmentUrlPrefix. When AWS Elemental MediaTailor serves a manifest,
	// it reports your CDN as the source for content segments.
	ContentSegmentUrlPrefix *string

	noSmithyDocumentSerde
}

// The configuration parameters for a channel.
type Channel struct {

	// The ARN of the channel.
	//
	// This member is required.
	Arn *string

	// The name of the channel.
	//
	// This member is required.
	ChannelName *string

	// Returns the state whether the channel is running or not.
	//
	// This member is required.
	ChannelState *string

	// The channel's output properties.
	//
	// This member is required.
	Outputs []ResponseOutputItem

	// The type of playback mode for this channel. LINEAR - Programs play back-to-back
	// only once. LOOP - Programs play back-to-back in an endless loop. When the last
	// program in the schedule plays, playback loops back to the first program in the
	// schedule.
	//
	// This member is required.
	PlaybackMode *string

	// The timestamp of when the channel was created.
	CreationTime *time.Time

	// Contains information about the slate used to fill gaps between programs in the
	// schedule. You must configure FillerSlate if your channel uses an LINEAR
	// PlaybackMode.
	FillerSlate *SlateSource

	// The timestamp of when the channel was last modified.
	LastModifiedTime *time.Time

	// The tags to assign to the channel.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The configuration for DASH content.
type DashConfiguration struct {

	// The URL generated by MediaTailor to initiate a playback session. The session
	// uses server-side reporting. This setting is ignored in PUT operations.
	ManifestEndpointPrefix *string

	// The setting that controls whether MediaTailor includes the Location tag in DASH
	// manifests. MediaTailor populates the Location tag with the URL for manifest
	// update requests, to be used by players that don't support sticky redirects.
	// Disable this if you have CDN routing rules set up for accessing MediaTailor
	// manifests, and you are either using client-side reporting or your players
	// support sticky HTTP redirects. Valid values are DISABLED and EMT_DEFAULT. The
	// EMT_DEFAULT setting enables the inclusion of the tag and is the default value.
	MpdLocation *string

	// The setting that controls whether MediaTailor handles manifests from the origin
	// server as multi-period manifests or single-period manifests. If your origin
	// server produces single-period manifests, set this to SINGLE_PERIOD. The default
	// setting is MULTI_PERIOD. For multi-period manifests, omit this setting or set it
	// to MULTI_PERIOD.
	OriginManifestType OriginManifestType

	noSmithyDocumentSerde
}

// The configuration for DASH PUT operations.
type DashConfigurationForPut struct {

	// The setting that controls whether MediaTailor includes the Location tag in DASH
	// manifests. MediaTailor populates the Location tag with the URL for manifest
	// update requests, to be used by players that don't support sticky redirects.
	// Disable this if you have CDN routing rules set up for accessing MediaTailor
	// manifests, and you are either using client-side reporting or your players
	// support sticky HTTP redirects. Valid values are DISABLED and EMT_DEFAULT. The
	// EMT_DEFAULT setting enables the inclusion of the tag and is the default value.
	MpdLocation *string

	// The setting that controls whether MediaTailor handles manifests from the origin
	// server as multi-period manifests or single-period manifests. If your origin
	// server produces single-period manifests, set this to SINGLE_PERIOD. The default
	// setting is MULTI_PERIOD. For multi-period manifests, omit this setting or set it
	// to MULTI_PERIOD.
	OriginManifestType OriginManifestType

	noSmithyDocumentSerde
}

// Dash manifest configuration parameters.
type DashPlaylistSettings struct {

	// The total duration (in seconds) of each manifest. Minimum value: 30 seconds.
	// Maximum value: 3600 seconds.
	ManifestWindowSeconds int32

	// Minimum amount of content (measured in seconds) that a player must keep
	// available in the buffer. Minimum value: 2 seconds. Maximum value: 60 seconds.
	MinBufferTimeSeconds int32

	// Minimum amount of time (in seconds) that the player should wait before
	// requesting updates to the manifest. Minimum value: 2 seconds. Maximum value: 60
	// seconds.
	MinUpdatePeriodSeconds int32

	// Amount of time (in seconds) that the player should be from the live point at the
	// end of the manifest. Minimum value: 2 seconds. Maximum value: 60 seconds.
	SuggestedPresentationDelaySeconds int32

	noSmithyDocumentSerde
}

// The optional configuration for a server that serves segments. Use this if you
// want the segment delivery server to be different from the source location
// server. For example, you can configure your source location server to be an
// origination server, such as MediaPackage, and the segment delivery server to be
// a content delivery network (CDN), such as CloudFront. If you don't specify a
// segment delivery server, then the source location server is used.
type DefaultSegmentDeliveryConfiguration struct {

	// The hostname of the server that will be used to serve segments. This string must
	// include the protocol, such as https://.
	BaseUrl *string

	noSmithyDocumentSerde
}

// The configuration for HLS content.
type HlsConfiguration struct {

	// The URL that is used to initiate a playback session for devices that support
	// Apple HLS. The session uses server-side reporting.
	ManifestEndpointPrefix *string

	noSmithyDocumentSerde
}

// HLS playlist configuration parameters.
type HlsPlaylistSettings struct {

	// The total duration (in seconds) of each manifest. Minimum value: 30 seconds.
	// Maximum value: 3600 seconds.
	ManifestWindowSeconds int32

	noSmithyDocumentSerde
}

// The HTTP configuration for the source location.
type HttpConfiguration struct {

	// The base URL for the source location host server. This string must include the
	// protocol, such as https://.
	//
	// This member is required.
	BaseUrl *string

	noSmithyDocumentSerde
}

// The HTTP package configuration properties for the requested VOD source.
type HttpPackageConfiguration struct {

	// The relative path to the URL for this VOD source. This is combined with
	// SourceLocation::HttpConfiguration::BaseUrl to form a valid URL.
	//
	// This member is required.
	Path *string

	// The name of the source group. This has to match one of the
	// Channel::Outputs::SourceGroup.
	//
	// This member is required.
	SourceGroup *string

	// The streaming protocol for this package configuration. Supported values are HLS
	// and DASH.
	//
	// This member is required.
	Type Type

	noSmithyDocumentSerde
}

// The configuration for pre-roll ad insertion.
type LivePreRollConfiguration struct {

	// The URL for the ad decision server (ADS) for pre-roll ads. This includes the
	// specification of static parameters and placeholders for dynamic parameters. AWS
	// Elemental MediaTailor substitutes player-specific and session-specific
	// parameters as needed when calling the ADS. Alternately, for testing, you can
	// provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string

	// The maximum allowed duration for the pre-roll ad avail. AWS Elemental
	// MediaTailor won't play pre-roll ads to exceed this duration, regardless of the
	// total duration of ads that the ADS returns.
	MaxDurationSeconds int32

	noSmithyDocumentSerde
}

// Returns Amazon CloudWatch log settings for a playback configuration.
type LogConfiguration struct {

	// The percentage of session logs that MediaTailor sends to your Cloudwatch Logs
	// account. For example, if your playback configuration has 1000 sessions and
	// percentEnabled is set to 60, MediaTailor sends logs for 600 of the sessions to
	// CloudWatch Logs. MediaTailor decides at random which of the playback
	// configuration sessions to send logs for. If you want to view logs for a specific
	// session, you can use the debug log mode
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/debug-log-mode.html). Valid
	// values: 0 - 100
	//
	// This member is required.
	PercentEnabled int32

	noSmithyDocumentSerde
}

// The configuration for manifest processing rules. Manifest processing rules
// enable customization of the personalized manifests created by MediaTailor.
type ManifestProcessingRules struct {

	// For HLS, when set to true, MediaTailor passes through EXT-X-CUE-IN,
	// EXT-X-CUE-OUT, and EXT-X-SPLICEPOINT-SCTE35 ad markers from the origin manifest
	// to the MediaTailor personalized manifest. No logic is applied to these ad
	// markers. For example, if EXT-X-CUE-OUT has a value of 60, but no ads are filled
	// for that ad break, MediaTailor will not set the value to 0.
	AdMarkerPassthrough *AdMarkerPassthrough

	noSmithyDocumentSerde
}

// Creates a playback configuration. For information about MediaTailor
// configurations, see Working with configurations in AWS Elemental MediaTailor
// (https://docs.aws.amazon.com/mediatailor/latest/ug/configurations.html).
type PlaybackConfiguration struct {

	// The URL for the ad decision server (ADS). This includes the specification of
	// static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing you can provide a static
	// VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string

	// The configuration for avail suppression, also known as ad suppression. For more
	// information about ad suppression, see Ad Suppression
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	AvailSuppression *AvailSuppression

	// The configuration for bumpers. Bumpers are short audio or video clips that play
	// at the start or before the end of an ad break. To learn more about bumpers, see
	// Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
	Bumper *Bumper

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration

	// The player parameters and aliases used as dynamic variables during session
	// initialization. For more information, see Domain Variables
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html).
	ConfigurationAliases map[string]map[string]string

	// The configuration for a DASH source.
	DashConfiguration *DashConfiguration

	// The configuration for HLS content.
	HlsConfiguration *HlsConfiguration

	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *LivePreRollConfiguration

	// The Amazon CloudWatch log settings for a playback configuration.
	LogConfiguration *LogConfiguration

	// The configuration for manifest processing rules. Manifest processing rules
	// enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules *ManifestProcessingRules

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
	// MediaTailor.
	PlaybackEndpointPrefix *string

	// The URL that the player uses to initialize a session that uses client-side
	// reporting.
	SessionInitializationEndpointPrefix *string

	// The URL for a video asset to transcode and use to fill in time that's not used
	// by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media
	// content. Configuring the slate is optional for non-VPAID playback
	// configurations. For VPAID, the slate is required because MediaTailor provides it
	// in the slots designated for dynamic ad content. The slate must be a high-quality
	// asset that contains both audio and video.
	SlateAdUrl *string

	// The tags to assign to the playback configuration.
	Tags map[string]string

	// The name that is used to associate this playback configuration with a custom
	// transcode profile. This overrides the dynamic transcoding defaults of
	// MediaTailor. Use this only if you have already set up custom profiles with the
	// help of AWS Support.
	TranscodeProfileName *string

	// The URL prefix for the parent manifest for the stream, minus the asset ID. The
	// maximum length is 512 characters.
	VideoContentSourceUrl *string

	noSmithyDocumentSerde
}

// The output configuration for this channel.
type RequestOutputItem struct {

	// The name of the manifest for the channel. The name appears in the PlaybackUrl.
	//
	// This member is required.
	ManifestName *string

	// A string used to match which HttpPackageConfiguration is used for each
	// VodSource.
	//
	// This member is required.
	SourceGroup *string

	// DASH manifest configuration parameters.
	DashPlaylistSettings *DashPlaylistSettings

	// HLS playlist configuration parameters.
	HlsPlaylistSettings *HlsPlaylistSettings

	noSmithyDocumentSerde
}

// This response includes only the "property" : "type" property.
type ResponseOutputItem struct {

	// The name of the manifest for the channel that will appear in the channel
	// output's playback URL.
	//
	// This member is required.
	ManifestName *string

	// The URL used for playback by content players.
	//
	// This member is required.
	PlaybackUrl *string

	// A string used to associate a package configuration source group with a channel
	// output.
	//
	// This member is required.
	SourceGroup *string

	// DASH manifest configuration settings.
	DashPlaylistSettings *DashPlaylistSettings

	// HLS manifest configuration settings.
	HlsPlaylistSettings *HlsPlaylistSettings

	noSmithyDocumentSerde
}

// The schedule's ad break properties.
type ScheduleAdBreak struct {

	// The approximate duration of the ad break, in seconds.
	ApproximateDurationSeconds int64

	// The approximate time that the ad will start playing.
	ApproximateStartTime *time.Time

	// The name of the source location containing the VOD source used for the ad break.
	SourceLocationName *string

	// The name of the VOD source used for the ad break.
	VodSourceName *string

	noSmithyDocumentSerde
}

// Schedule configuration parameters. A channel must be stopped before changes can
// be made to the schedule.
type ScheduleConfiguration struct {

	// Program transition configurations.
	//
	// This member is required.
	Transition *Transition

	noSmithyDocumentSerde
}

// The properties for a schedule.
type ScheduleEntry struct {

	// The ARN of the program.
	//
	// This member is required.
	Arn *string

	// The name of the channel that uses this schedule.
	//
	// This member is required.
	ChannelName *string

	// The name of the program.
	//
	// This member is required.
	ProgramName *string

	// The name of the source location.
	//
	// This member is required.
	SourceLocationName *string

	// The name of the VOD source.
	//
	// This member is required.
	VodSourceName *string

	// The approximate duration of this program, in seconds.
	ApproximateDurationSeconds int64

	// The approximate time that the program will start playing.
	ApproximateStartTime *time.Time

	// The schedule's ad break properties.
	ScheduleAdBreaks []ScheduleAdBreak

	// The type of schedule entry. Valid values: PROGRAM or FILLER_SLATE.
	ScheduleEntryType ScheduleEntryType

	noSmithyDocumentSerde
}

// AWS Secrets Manager access token configuration parameters. For information about
// Secrets Manager access token authentication, see Working with AWS Secrets
// Manager access token authentication
// (https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-access-configuration-access-token.html).
type SecretsManagerAccessTokenConfiguration struct {

	// The name of the HTTP header used to supply the access token in requests to the
	// source location.
	HeaderName *string

	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains
	// the access token.
	SecretArn *string

	// The AWS Secrets Manager SecretString
	// (https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html#SecretsManager-CreateSecret-request-SecretString.html)
	// key associated with the access token. MediaTailor uses the key to look up
	// SecretString key and value pair containing the access token.
	SecretStringKey *string

	noSmithyDocumentSerde
}

// Slate VOD source configuration.
type SlateSource struct {

	// The name of the source location where the slate VOD source is stored.
	SourceLocationName *string

	// The slate VOD source name. The VOD source must already exist in a source
	// location before it can be used for slate.
	VodSourceName *string

	noSmithyDocumentSerde
}

// This response includes only the "type" : "object" property.
type SourceLocation struct {

	// The ARN of the SourceLocation.
	//
	// This member is required.
	Arn *string

	// The HTTP configuration for the source location.
	//
	// This member is required.
	HttpConfiguration *HttpConfiguration

	// The name of the source location.
	//
	// This member is required.
	SourceLocationName *string

	// The access configuration for the source location.
	AccessConfiguration *AccessConfiguration

	// The timestamp that indicates when the source location was created.
	CreationTime *time.Time

	// The default segment delivery configuration.
	DefaultSegmentDeliveryConfiguration *DefaultSegmentDeliveryConfiguration

	// The timestamp that indicates when the source location was last modified.
	LastModifiedTime *time.Time

	// The tags assigned to the source location.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Splice insert message configuration.
type SpliceInsertMessage struct {

	// This is written to splice_insert.avail_num, as defined in section 9.7.3.1 of the
	// SCTE-35 specification. The default value is 0. Values must be between 0 and 256,
	// inclusive.
	AvailNum int32

	// This is written to splice_insert.avails_expected, as defined in section 9.7.3.1
	// of the SCTE-35 specification. The default value is 0. Values must be between 0
	// and 256, inclusive.
	AvailsExpected int32

	// This is written to splice_insert.splice_event_id, as defined in section 9.7.3.1
	// of the SCTE-35 specification. The default value is 1.
	SpliceEventId int32

	// This is written to splice_insert.unique_program_id, as defined in section
	// 9.7.3.1 of the SCTE-35 specification. The default value is 0. Values must be
	// between 0 and 256, inclusive.
	UniqueProgramId int32

	noSmithyDocumentSerde
}

// Program transition configuration.
type Transition struct {

	// The position where this program will be inserted relative to the
	// RelativePosition.
	//
	// This member is required.
	RelativePosition RelativePosition

	// Defines when the program plays in the schedule. You can set the value to
	// ABSOLUTE or RELATIVE. ABSOLUTE - The program plays at a specific wall clock
	// time. This setting can only be used for channels using the LINEAR PlaybackMode.
	// Note the following considerations when using ABSOLUTE transitions: If the
	// preceding program in the schedule has a duration that extends past the wall
	// clock time, MediaTailor truncates the preceding program on a common segment
	// boundary. If there are gaps in playback, MediaTailor plays the FillerSlate you
	// configured for your linear channel. RELATIVE - The program is inserted into the
	// schedule either before or after a program that you specify via RelativePosition.
	//
	// This member is required.
	Type *string

	// The name of the program that this program will be inserted next to, as defined
	// by RelativePosition.
	RelativeProgram *string

	// The date and time that the program is scheduled to start, in epoch milliseconds.
	ScheduledStartTimeMillis int64

	noSmithyDocumentSerde
}

// VOD source configuration parameters.
type VodSource struct {

	// The ARN for the VOD source.
	//
	// This member is required.
	Arn *string

	// The HTTP package configurations for the VOD source.
	//
	// This member is required.
	HttpPackageConfigurations []HttpPackageConfiguration

	// The name of the source location that the VOD source is associated with.
	//
	// This member is required.
	SourceLocationName *string

	// The name of the VOD source.
	//
	// This member is required.
	VodSourceName *string

	// The timestamp that indicates when the VOD source was created.
	CreationTime *time.Time

	// The timestamp that indicates when the VOD source was last modified.
	LastModifiedTime *time.Time

	// The tags assigned to the VOD source.
	Tags map[string]string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
