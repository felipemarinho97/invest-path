package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/felipemarinho97/go-mock-gen
 */

// ILambdaClient generic client
type ILambdaClient interface {
	// Adds permissions to the resource-based policy of a version of an Lambda layer
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). Use
	// this action to grant layer usage permission to other accounts. You can grant
	// permission to a single account, all accounts in an organization, or all Amazon
	// Web Services accounts. To revoke permission, call RemoveLayerVersionPermission
	// with the statement ID that you specified when you added it.
	//
	AddLayerVersionPermission(arg1 context.Context, arg2 *lambda.AddLayerVersionPermissionInput, arg3 ...func(*lambda.Options)) (*lambda.AddLayerVersionPermissionOutput, error)
	// Grants an Amazon Web Services service or another account permission to use a
	// function. You can apply the policy at the function level, or specify a qualifier
	// to restrict access to a single version or alias. If you use a qualifier, the
	// invoker must use the full Amazon Resource Name (ARN) of that version or alias to
	// invoke the function. To grant permission to another account, specify the account
	// ID as the Principal. For Amazon Web Services services, the principal is a
	// domain-style identifier defined by the service, like s3.amazonaws.com or
	// sns.amazonaws.com. For Amazon Web Services services, you can also specify the
	// ARN of the associated resource as the SourceArn. If you grant permission to a
	// service principal without specifying the source, other accounts could
	// potentially configure resources in their account to invoke your Lambda function.
	// This action adds a statement to a resource-based permissions policy for the
	// function. For more information about function policies, see Lambda Function
	// Policies
	// (https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html).
	//
	AddPermission(arg1 context.Context, arg2 *lambda.AddPermissionInput, arg3 ...func(*lambda.Options)) (*lambda.AddPermissionOutput, error)
	// Creates an alias
	// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html) for a
	// Lambda function version. Use aliases to provide clients with a function
	// identifier that you can update to invoke a different version. You can also map
	// an alias to split invocation requests between two versions. Use the
	// RoutingConfig parameter to specify a second version and the percentage of
	// invocation requests that it receives.
	//
	CreateAlias(arg1 context.Context, arg2 *lambda.CreateAliasInput, arg3 ...func(*lambda.Options)) (*lambda.CreateAliasOutput, error)
	// Creates a code signing configuration. A code signing configuration
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-trustedcode.html)
	// defines a list of allowed signing profiles and defines the code-signing
	// validation policy (action to be taken if deployment validation checks fail).
	//
	CreateCodeSigningConfig(arg1 context.Context, arg2 *lambda.CreateCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.CreateCodeSigningConfigOutput, error)
	// Creates a mapping between an event source and an Lambda function. Lambda reads
	// items from the event source and triggers the function. For details about each
	// event source type, see the following topics. In particular, each of the topics
	// describes the required and optional parameters for the specific event source.
	//
	// *
	// Configuring a Dynamo DB stream as an event source
	// (https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-dynamodb-eventsourcemapping)
	//
	// *
	// Configuring a Kinesis stream as an event source
	// (https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-eventsourcemapping)
	//
	// *
	// Configuring an SQS queue as an event source
	// (https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-eventsource)
	//
	// *
	// Configuring an MQ broker as an event source
	// (https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#services-mq-eventsourcemapping)
	//
	// *
	// Configuring MSK as an event source
	// (https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html)
	//
	// * Configuring
	// Self-Managed Apache Kafka as an event source
	// (https://docs.aws.amazon.com/lambda/latest/dg/kafka-smaa.html)
	//
	// The following
	// error handling options are only available for stream sources (DynamoDB and
	// Kinesis):
	//
	// * BisectBatchOnFunctionError - If the function returns an error,
	// split the batch in two and retry.
	//
	// * DestinationConfig - Send discarded records
	// to an Amazon SQS queue or Amazon SNS topic.
	//
	// * MaximumRecordAgeInSeconds -
	// Discard records older than the specified age. The default value is infinite
	// (-1). When set to infinite (-1), failed records are retried until the record
	// expires
	//
	// * MaximumRetryAttempts - Discard records after the specified number of
	// retries. The default value is infinite (-1). When set to infinite (-1), failed
	// records are retried until the record expires.
	//
	// * ParallelizationFactor - Process
	// multiple batches from each shard concurrently.
	//
	CreateEventSourceMapping(arg1 context.Context, arg2 *lambda.CreateEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.CreateEventSourceMappingOutput, error)
	// Creates a Lambda function. To create a function, you need a deployment package
	// (https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) and
	// an execution role
	// (https://docs.aws.amazon.com/lambda/latest/dg/intro-permission-model.html#lambda-intro-execution-role).
	// The deployment package is a .zip file archive or container image that contains
	// your function code. The execution role grants the function permission to use
	// Amazon Web Services services, such as Amazon CloudWatch Logs for log streaming
	// and X-Ray for request tracing. You set the package type to Image if the
	// deployment package is a container image
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html). For a
	// container image, the code property must include the URI of a container image in
	// the Amazon ECR registry. You do not need to specify the handler and runtime
	// properties. You set the package type to Zip if the deployment package is a .zip
	// file archive
	// (https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html#gettingstarted-package-zip).
	// For a .zip file archive, the code property specifies the location of the .zip
	// file. You must also specify the handler and runtime properties. When you create
	// a function, Lambda provisions an instance of the function and its supporting
	// resources. If your function connects to a VPC, this process can take a minute or
	// so. During this time, you can't invoke or modify the function. The State,
	// StateReason, and StateReasonCode fields in the response from
	// GetFunctionConfiguration indicate when the function is ready to invoke. For more
	// information, see Function States
	// (https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html). A function
	// has an unpublished version, and can have published versions and aliases. The
	// unpublished version changes when you update your function's code and
	// configuration. A published version is a snapshot of your function code and
	// configuration that can't be changed. An alias is a named resource that maps to a
	// version, and can be changed to map to a different version. Use the Publish
	// parameter to create version 1 of your function from its initial configuration.
	// The other parameters let you configure version-specific and function-level
	// settings. You can modify version-specific settings later with
	// UpdateFunctionConfiguration. Function-level settings apply to both the
	// unpublished and published versions of the function, and include tags
	// (TagResource) and per-function concurrency limits (PutFunctionConcurrency). You
	// can use code signing if your deployment package is a .zip file archive. To
	// enable code signing for this function, specify the ARN of a code-signing
	// configuration. When a user attempts to deploy a code package with
	// UpdateFunctionCode, Lambda checks that the code package has a valid signature
	// from a trusted publisher. The code-signing configuration includes set set of
	// signing profiles, which define the trusted publishers for this function. If
	// another account or an Amazon Web Services service invokes your function, use
	// AddPermission to grant permission by creating a resource-based IAM policy. You
	// can grant permissions at the function level, on a version, or on an alias. To
	// invoke your function directly, use Invoke. To invoke your function in response
	// to events in other Amazon Web Services services, create an event source mapping
	// (CreateEventSourceMapping), or configure a function trigger in the other
	// service. For more information, see Invoking Functions
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-invocation.html).
	//
	CreateFunction(arg1 context.Context, arg2 *lambda.CreateFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.CreateFunctionOutput, error)
	// Deletes a Lambda function alias
	// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
	//
	DeleteAlias(arg1 context.Context, arg2 *lambda.DeleteAliasInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteAliasOutput, error)
	// Deletes the code signing configuration. You can delete the code signing
	// configuration only if no function is using it.
	//
	DeleteCodeSigningConfig(arg1 context.Context, arg2 *lambda.DeleteCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteCodeSigningConfigOutput, error)
	// Deletes an event source mapping
	// (https://docs.aws.amazon.com/lambda/latest/dg/intro-invocation-modes.html). You
	// can get the identifier of a mapping from the output of ListEventSourceMappings.
	// When you delete an event source mapping, it enters a Deleting state and might
	// not be completely deleted for several seconds.
	//
	DeleteEventSourceMapping(arg1 context.Context, arg2 *lambda.DeleteEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteEventSourceMappingOutput, error)
	// Deletes a Lambda function. To delete a specific function version, use the
	// Qualifier parameter. Otherwise, all versions and aliases are deleted. To delete
	// Lambda event source mappings that invoke a function, use
	// DeleteEventSourceMapping. For Amazon Web Services services and resources that
	// invoke your function directly, delete the trigger in the service where you
	// originally configured it.
	//
	DeleteFunction(arg1 context.Context, arg2 *lambda.DeleteFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionOutput, error)
	// Removes the code signing configuration from the function.
	//
	DeleteFunctionCodeSigningConfig(arg1 context.Context, arg2 *lambda.DeleteFunctionCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionCodeSigningConfigOutput, error)
	// Removes a concurrent execution limit from a function.
	//
	DeleteFunctionConcurrency(arg1 context.Context, arg2 *lambda.DeleteFunctionConcurrencyInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionConcurrencyOutput, error)
	// Deletes the configuration for asynchronous invocation for a function, version,
	// or alias. To configure options for asynchronous invocation, use
	// PutFunctionEventInvokeConfig.
	//
	DeleteFunctionEventInvokeConfig(arg1 context.Context, arg2 *lambda.DeleteFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionEventInvokeConfigOutput, error)
	// Deletes a version of an Lambda layer
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	// Deleted versions can no longer be viewed or added to functions. To avoid
	// breaking functions, a copy of the version remains in Lambda until no functions
	// refer to it.
	//
	DeleteLayerVersion(arg1 context.Context, arg2 *lambda.DeleteLayerVersionInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteLayerVersionOutput, error)
	// Deletes the provisioned concurrency configuration for a function.
	//
	DeleteProvisionedConcurrencyConfig(arg1 context.Context, arg2 *lambda.DeleteProvisionedConcurrencyConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error)
	// Retrieves details about your account's limits
	// (https://docs.aws.amazon.com/lambda/latest/dg/limits.html) and usage in an
	// Amazon Web Services Region.
	//
	GetAccountSettings(arg1 context.Context, arg2 *lambda.GetAccountSettingsInput, arg3 ...func(*lambda.Options)) (*lambda.GetAccountSettingsOutput, error)
	// Returns details about a Lambda function alias
	// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
	//
	GetAlias(arg1 context.Context, arg2 *lambda.GetAliasInput, arg3 ...func(*lambda.Options)) (*lambda.GetAliasOutput, error)
	// Returns information about the specified code signing configuration.
	//
	GetCodeSigningConfig(arg1 context.Context, arg2 *lambda.GetCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetCodeSigningConfigOutput, error)
	// Returns details about an event source mapping. You can get the identifier of a
	// mapping from the output of ListEventSourceMappings.
	//
	GetEventSourceMapping(arg1 context.Context, arg2 *lambda.GetEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.GetEventSourceMappingOutput, error)
	// Returns information about the function or function version, with a link to
	// download the deployment package that's valid for 10 minutes. If you specify a
	// function version, only details that are specific to that version are returned.
	//
	GetFunction(arg1 context.Context, arg2 *lambda.GetFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error)
	// Returns the code signing configuration for the specified function.
	//
	GetFunctionCodeSigningConfig(arg1 context.Context, arg2 *lambda.GetFunctionCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionCodeSigningConfigOutput, error)
	// Returns details about the reserved concurrency configuration for a function. To
	// set a concurrency limit for a function, use PutFunctionConcurrency.
	//
	GetFunctionConcurrency(arg1 context.Context, arg2 *lambda.GetFunctionConcurrencyInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionConcurrencyOutput, error)
	// Returns the version-specific settings of a Lambda function or version. The
	// output includes only options that can vary between versions of a function. To
	// modify these settings, use UpdateFunctionConfiguration. To get all of a
	// function's details, including function-level settings, use GetFunction.
	//
	GetFunctionConfiguration(arg1 context.Context, arg2 *lambda.GetFunctionConfigurationInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionConfigurationOutput, error)
	// Retrieves the configuration for asynchronous invocation for a function, version,
	// or alias. To configure options for asynchronous invocation, use
	// PutFunctionEventInvokeConfig.
	//
	GetFunctionEventInvokeConfig(arg1 context.Context, arg2 *lambda.GetFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionEventInvokeConfigOutput, error)
	// Returns information about a version of an Lambda layer
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html), with a
	// link to download the layer archive that's valid for 10 minutes.
	//
	GetLayerVersion(arg1 context.Context, arg2 *lambda.GetLayerVersionInput, arg3 ...func(*lambda.Options)) (*lambda.GetLayerVersionOutput, error)
	// Returns information about a version of an Lambda layer
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html), with a
	// link to download the layer archive that's valid for 10 minutes.
	//
	GetLayerVersionByArn(arg1 context.Context, arg2 *lambda.GetLayerVersionByArnInput, arg3 ...func(*lambda.Options)) (*lambda.GetLayerVersionByArnOutput, error)
	// Returns the permission policy for a version of an Lambda layer
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). For
	// more information, see AddLayerVersionPermission.
	//
	GetLayerVersionPolicy(arg1 context.Context, arg2 *lambda.GetLayerVersionPolicyInput, arg3 ...func(*lambda.Options)) (*lambda.GetLayerVersionPolicyOutput, error)
	// Returns the resource-based IAM policy
	// (https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html)
	// for a function, version, or alias.
	//
	GetPolicy(arg1 context.Context, arg2 *lambda.GetPolicyInput, arg3 ...func(*lambda.Options)) (*lambda.GetPolicyOutput, error)
	// Retrieves the provisioned concurrency configuration for a function's alias or
	// version.
	//
	GetProvisionedConcurrencyConfig(arg1 context.Context, arg2 *lambda.GetProvisionedConcurrencyConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetProvisionedConcurrencyConfigOutput, error)
	// Invokes a Lambda function. You can invoke a function synchronously (and wait for
	// the response), or asynchronously. To invoke a function asynchronously, set
	// InvocationType to Event. For synchronous invocation
	// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-sync.html), details
	// about the function response, including errors, are included in the response body
	// and headers. For either invocation type, you can find more information in the
	// execution log
	// (https://docs.aws.amazon.com/lambda/latest/dg/monitoring-functions.html) and
	// trace (https://docs.aws.amazon.com/lambda/latest/dg/lambda-x-ray.html). When an
	// error occurs, your function may be invoked multiple times. Retry behavior varies
	// by error type, client, event source, and invocation type. For example, if you
	// invoke a function asynchronously and it returns an error, Lambda executes the
	// function up to two more times. For more information, see Retry Behavior
	// (https://docs.aws.amazon.com/lambda/latest/dg/retries-on-errors.html). For
	// asynchronous invocation
	// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html), Lambda
	// adds events to a queue before sending them to your function. If your function
	// does not have enough capacity to keep up with the queue, events may be lost.
	// Occasionally, your function may receive the same event multiple times, even if
	// no error occurs. To retain events that were not processed, configure your
	// function with a dead-letter queue
	// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq). The
	// status code in the API response doesn't reflect function errors. Error codes are
	// reserved for errors that prevent your function from executing, such as
	// permissions errors, limit errors
	// (https://docs.aws.amazon.com/lambda/latest/dg/limits.html), or issues with your
	// function's code and configuration. For example, Lambda returns
	// TooManyRequestsException if executing the function would cause you to exceed a
	// concurrency limit at either the account level
	// (ConcurrentInvocationLimitExceeded) or function level
	// (ReservedFunctionConcurrentInvocationLimitExceeded). For functions with a long
	// timeout, your client might be disconnected during synchronous invocation while
	// it waits for a response. Configure your HTTP client, SDK, firewall, proxy, or
	// operating system to allow for long connections with timeout or keep-alive
	// settings. This operation requires permission for the lambda:InvokeFunction
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/list_awslambda.html) action.
	//
	Invoke(arg1 context.Context, arg2 *lambda.InvokeInput, arg3 ...func(*lambda.Options)) (*lambda.InvokeOutput, error)
	// For asynchronous function invocation, use Invoke. Invokes a function
	// asynchronously.
	//
	// Deprecated: This operation has been deprecated.
	//
	InvokeAsync(arg1 context.Context, arg2 *lambda.InvokeAsyncInput, arg3 ...func(*lambda.Options)) (*lambda.InvokeAsyncOutput, error)
	// Returns a list of aliases
	// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html) for a
	// Lambda function.
	//
	ListAliases(arg1 context.Context, arg2 *lambda.ListAliasesInput, arg3 ...func(*lambda.Options)) (*lambda.ListAliasesOutput, error)
	// Returns a list of code signing configurations
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuring-codesigning.html). A
	// request returns up to 10,000 configurations per call. You can use the MaxItems
	// parameter to return fewer configurations per call.
	//
	ListCodeSigningConfigs(arg1 context.Context, arg2 *lambda.ListCodeSigningConfigsInput, arg3 ...func(*lambda.Options)) (*lambda.ListCodeSigningConfigsOutput, error)
	// Lists event source mappings. Specify an EventSourceArn to only show event source
	// mappings for a single event source.
	//
	ListEventSourceMappings(arg1 context.Context, arg2 *lambda.ListEventSourceMappingsInput, arg3 ...func(*lambda.Options)) (*lambda.ListEventSourceMappingsOutput, error)
	// Retrieves a list of configurations for asynchronous invocation for a function.
	// To configure options for asynchronous invocation, use
	// PutFunctionEventInvokeConfig.
	//
	ListFunctionEventInvokeConfigs(arg1 context.Context, arg2 *lambda.ListFunctionEventInvokeConfigsInput, arg3 ...func(*lambda.Options)) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
	// Returns a list of Lambda functions, with the version-specific configuration of
	// each. Lambda returns up to 50 functions per call. Set FunctionVersion to ALL to
	// include all published versions of each function in addition to the unpublished
	// version. The ListFunctions action returns a subset of the FunctionConfiguration
	// fields. To get the additional fields (State, StateReasonCode, StateReason,
	// LastUpdateStatus, LastUpdateStatusReason, LastUpdateStatusReasonCode) for a
	// function or version, use GetFunction.
	//
	ListFunctions(arg1 context.Context, arg2 *lambda.ListFunctionsInput, arg3 ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error)
	// List the functions that use the specified code signing configuration. You can
	// use this method prior to deleting a code signing configuration, to verify that
	// no functions are using it.
	//
	ListFunctionsByCodeSigningConfig(arg1 context.Context, arg2 *lambda.ListFunctionsByCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.ListFunctionsByCodeSigningConfigOutput, error)
	// Lists the versions of an Lambda layer
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	// Versions that have been deleted aren't listed. Specify a runtime identifier
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) to list only
	// versions that indicate that they're compatible with that runtime.
	//
	ListLayerVersions(arg1 context.Context, arg2 *lambda.ListLayerVersionsInput, arg3 ...func(*lambda.Options)) (*lambda.ListLayerVersionsOutput, error)
	// Lists Lambda layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) and
	// shows information about the latest version of each. Specify a runtime identifier
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html) to list only
	// layers that indicate that they're compatible with that runtime.
	//
	ListLayers(arg1 context.Context, arg2 *lambda.ListLayersInput, arg3 ...func(*lambda.Options)) (*lambda.ListLayersOutput, error)
	// Retrieves a list of provisioned concurrency configurations for a function.
	//
	ListProvisionedConcurrencyConfigs(arg1 context.Context, arg2 *lambda.ListProvisionedConcurrencyConfigsInput, arg3 ...func(*lambda.Options)) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
	// Returns a function's tags
	// (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html). You can also view
	// tags with GetFunction.
	//
	ListTags(arg1 context.Context, arg2 *lambda.ListTagsInput, arg3 ...func(*lambda.Options)) (*lambda.ListTagsOutput, error)
	// Returns a list of versions
	// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html), with the
	// version-specific configuration of each. Lambda returns up to 50 versions per
	// call.
	//
	ListVersionsByFunction(arg1 context.Context, arg2 *lambda.ListVersionsByFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.ListVersionsByFunctionOutput, error)
	// Creates an Lambda layer
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) from a
	// ZIP archive. Each time you call PublishLayerVersion with the same layer name, a
	// new version is created. Add layers to your function with CreateFunction or
	// UpdateFunctionConfiguration.
	//
	PublishLayerVersion(arg1 context.Context, arg2 *lambda.PublishLayerVersionInput, arg3 ...func(*lambda.Options)) (*lambda.PublishLayerVersionOutput, error)
	// Creates a version
	// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html) from the
	// current code and configuration of a function. Use versions to create a snapshot
	// of your function code and configuration that doesn't change. Lambda doesn't
	// publish a version if the function's configuration and code haven't changed since
	// the last version. Use UpdateFunctionCode or UpdateFunctionConfiguration to
	// update the function before publishing a version. Clients can invoke versions
	// directly or with an alias. To create an alias, use CreateAlias.
	//
	PublishVersion(arg1 context.Context, arg2 *lambda.PublishVersionInput, arg3 ...func(*lambda.Options)) (*lambda.PublishVersionOutput, error)
	// Update the code signing configuration for the function. Changes to the code
	// signing configuration take effect the next time a user tries to deploy a code
	// package to the function.
	//
	PutFunctionCodeSigningConfig(arg1 context.Context, arg2 *lambda.PutFunctionCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.PutFunctionCodeSigningConfigOutput, error)
	// Sets the maximum number of simultaneous executions for a function, and reserves
	// capacity for that concurrency level. Concurrency settings apply to the function
	// as a whole, including all published versions and the unpublished version.
	// Reserving concurrency both ensures that your function has capacity to process
	// the specified number of events simultaneously, and prevents it from scaling
	// beyond that level. Use GetFunction to see the current setting for a function.
	// Use GetAccountSettings to see your Regional concurrency limit. You can reserve
	// concurrency for as many functions as you like, as long as you leave at least 100
	// simultaneous executions unreserved for functions that aren't configured with a
	// per-function limit. For more information, see Managing Concurrency
	// (https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html).
	//
	PutFunctionConcurrency(arg1 context.Context, arg2 *lambda.PutFunctionConcurrencyInput, arg3 ...func(*lambda.Options)) (*lambda.PutFunctionConcurrencyOutput, error)
	// Configures options for asynchronous invocation
	// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html) on a
	// function, version, or alias. If a configuration already exists for a function,
	// version, or alias, this operation overwrites it. If you exclude any settings,
	// they are removed. To set one option without affecting existing settings for
	// other options, use UpdateFunctionEventInvokeConfig. By default, Lambda retries
	// an asynchronous invocation twice if the function returns an error. It retains
	// events in a queue for up to six hours. When an event fails all processing
	// attempts or stays in the asynchronous invocation queue for too long, Lambda
	// discards it. To retain discarded events, configure a dead-letter queue with
	// UpdateFunctionConfiguration. To send an invocation record to a queue, topic,
	// function, or event bus, specify a destination
	// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations).
	// You can configure separate destinations for successful invocations (on-success)
	// and events that fail all processing attempts (on-failure). You can configure
	// destinations in addition to or instead of a dead-letter queue.
	//
	PutFunctionEventInvokeConfig(arg1 context.Context, arg2 *lambda.PutFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.PutFunctionEventInvokeConfigOutput, error)
	// Adds a provisioned concurrency configuration to a function's alias or version.
	//
	PutProvisionedConcurrencyConfig(arg1 context.Context, arg2 *lambda.PutProvisionedConcurrencyConfigInput, arg3 ...func(*lambda.Options)) (*lambda.PutProvisionedConcurrencyConfigOutput, error)
	// Removes a statement from the permissions policy for a version of an Lambda layer
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). For
	// more information, see AddLayerVersionPermission.
	//
	RemoveLayerVersionPermission(arg1 context.Context, arg2 *lambda.RemoveLayerVersionPermissionInput, arg3 ...func(*lambda.Options)) (*lambda.RemoveLayerVersionPermissionOutput, error)
	// Revokes function-use permission from an Amazon Web Services service or another
	// account. You can get the ID of the statement from the output of GetPolicy.
	//
	RemovePermission(arg1 context.Context, arg2 *lambda.RemovePermissionInput, arg3 ...func(*lambda.Options)) (*lambda.RemovePermissionOutput, error)
	// Adds tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to a
	// function.
	//
	TagResource(arg1 context.Context, arg2 *lambda.TagResourceInput, arg3 ...func(*lambda.Options)) (*lambda.TagResourceOutput, error)
	// Removes tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) from a
	// function.
	//
	UntagResource(arg1 context.Context, arg2 *lambda.UntagResourceInput, arg3 ...func(*lambda.Options)) (*lambda.UntagResourceOutput, error)
	// Updates the configuration of a Lambda function alias
	// (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
	//
	UpdateAlias(arg1 context.Context, arg2 *lambda.UpdateAliasInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateAliasOutput, error)
	// Update the code signing configuration. Changes to the code signing configuration
	// take effect the next time a user tries to deploy a code package to the function.
	//
	UpdateCodeSigningConfig(arg1 context.Context, arg2 *lambda.UpdateCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateCodeSigningConfigOutput, error)
	// Updates an event source mapping. You can change the function that Lambda
	// invokes, or pause invocation and resume later from the same location. The
	// following error handling options are only available for stream sources (DynamoDB
	// and Kinesis):
	//
	// * BisectBatchOnFunctionError - If the function returns an error,
	// split the batch in two and retry.
	//
	// * DestinationConfig - Send discarded records
	// to an Amazon SQS queue or Amazon SNS topic.
	//
	// * MaximumRecordAgeInSeconds -
	// Discard records older than the specified age. The default value is infinite
	// (-1). When set to infinite (-1), failed records are retried until the record
	// expires
	//
	// * MaximumRetryAttempts - Discard records after the specified number of
	// retries. The default value is infinite (-1). When set to infinite (-1), failed
	// records are retried until the record expires.
	//
	// * ParallelizationFactor - Process
	// multiple batches from each shard concurrently.
	//
	UpdateEventSourceMapping(arg1 context.Context, arg2 *lambda.UpdateEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateEventSourceMappingOutput, error)
	// Updates a Lambda function's code. If code signing is enabled for the function,
	// the code package must be signed by a trusted publisher. For more information,
	// see Configuring code signing
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-trustedcode.html).
	// The function's code is locked when you publish a version. You can't modify the
	// code of a published version, only the unpublished version. For a function
	// defined as a container image, Lambda resolves the image tag to an image digest.
	// In Amazon ECR, if you update the image tag to a new image, Lambda does not
	// automatically update the function.
	//
	UpdateFunctionCode(arg1 context.Context, arg2 *lambda.UpdateFunctionCodeInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateFunctionCodeOutput, error)
	// Modify the version-specific settings of a Lambda function. When you update a
	// function, Lambda provisions an instance of the function and its supporting
	// resources. If your function connects to a VPC, this process can take a minute.
	// During this time, you can't modify the function, but you can still invoke it.
	// The LastUpdateStatus, LastUpdateStatusReason, and LastUpdateStatusReasonCode
	// fields in the response from GetFunctionConfiguration indicate when the update is
	// complete and the function is processing events with the new configuration. For
	// more information, see Function States
	// (https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html). These
	// settings can vary between versions of a function and are locked when you publish
	// a version. You can't modify the configuration of a published version, only the
	// unpublished version. To configure function concurrency, use
	// PutFunctionConcurrency. To grant invoke permissions to an account or Amazon Web
	// Services service, use AddPermission.
	//
	UpdateFunctionConfiguration(arg1 context.Context, arg2 *lambda.UpdateFunctionConfigurationInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateFunctionConfigurationOutput, error)
	// Updates the configuration for asynchronous invocation for a function, version,
	// or alias. To configure options for asynchronous invocation, use
	// PutFunctionEventInvokeConfig.
	//
	UpdateFunctionEventInvokeConfig(arg1 context.Context, arg2 *lambda.UpdateFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateFunctionEventInvokeConfigOutput, error)
}

// LambdaClientMock generic client mock
type LambdaClientMock struct {
	AddLayerVersionPermissionMock          func(arg1 context.Context, arg2 *lambda.AddLayerVersionPermissionInput, arg3 ...func(*lambda.Options)) (*lambda.AddLayerVersionPermissionOutput, error)
	AddPermissionMock                      func(arg1 context.Context, arg2 *lambda.AddPermissionInput, arg3 ...func(*lambda.Options)) (*lambda.AddPermissionOutput, error)
	CreateAliasMock                        func(arg1 context.Context, arg2 *lambda.CreateAliasInput, arg3 ...func(*lambda.Options)) (*lambda.CreateAliasOutput, error)
	CreateCodeSigningConfigMock            func(arg1 context.Context, arg2 *lambda.CreateCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.CreateCodeSigningConfigOutput, error)
	CreateEventSourceMappingMock           func(arg1 context.Context, arg2 *lambda.CreateEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.CreateEventSourceMappingOutput, error)
	CreateFunctionMock                     func(arg1 context.Context, arg2 *lambda.CreateFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.CreateFunctionOutput, error)
	DeleteAliasMock                        func(arg1 context.Context, arg2 *lambda.DeleteAliasInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteAliasOutput, error)
	DeleteCodeSigningConfigMock            func(arg1 context.Context, arg2 *lambda.DeleteCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteCodeSigningConfigOutput, error)
	DeleteEventSourceMappingMock           func(arg1 context.Context, arg2 *lambda.DeleteEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteEventSourceMappingOutput, error)
	DeleteFunctionMock                     func(arg1 context.Context, arg2 *lambda.DeleteFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionOutput, error)
	DeleteFunctionCodeSigningConfigMock    func(arg1 context.Context, arg2 *lambda.DeleteFunctionCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionCodeSigningConfigOutput, error)
	DeleteFunctionConcurrencyMock          func(arg1 context.Context, arg2 *lambda.DeleteFunctionConcurrencyInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionConcurrencyOutput, error)
	DeleteFunctionEventInvokeConfigMock    func(arg1 context.Context, arg2 *lambda.DeleteFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionEventInvokeConfigOutput, error)
	DeleteLayerVersionMock                 func(arg1 context.Context, arg2 *lambda.DeleteLayerVersionInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteLayerVersionOutput, error)
	DeleteProvisionedConcurrencyConfigMock func(arg1 context.Context, arg2 *lambda.DeleteProvisionedConcurrencyConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error)
	GetAccountSettingsMock                 func(arg1 context.Context, arg2 *lambda.GetAccountSettingsInput, arg3 ...func(*lambda.Options)) (*lambda.GetAccountSettingsOutput, error)
	GetAliasMock                           func(arg1 context.Context, arg2 *lambda.GetAliasInput, arg3 ...func(*lambda.Options)) (*lambda.GetAliasOutput, error)
	GetCodeSigningConfigMock               func(arg1 context.Context, arg2 *lambda.GetCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetCodeSigningConfigOutput, error)
	GetEventSourceMappingMock              func(arg1 context.Context, arg2 *lambda.GetEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.GetEventSourceMappingOutput, error)
	GetFunctionMock                        func(arg1 context.Context, arg2 *lambda.GetFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error)
	GetFunctionCodeSigningConfigMock       func(arg1 context.Context, arg2 *lambda.GetFunctionCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionCodeSigningConfigOutput, error)
	GetFunctionConcurrencyMock             func(arg1 context.Context, arg2 *lambda.GetFunctionConcurrencyInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionConcurrencyOutput, error)
	GetFunctionConfigurationMock           func(arg1 context.Context, arg2 *lambda.GetFunctionConfigurationInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionConfigurationOutput, error)
	GetFunctionEventInvokeConfigMock       func(arg1 context.Context, arg2 *lambda.GetFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionEventInvokeConfigOutput, error)
	GetLayerVersionMock                    func(arg1 context.Context, arg2 *lambda.GetLayerVersionInput, arg3 ...func(*lambda.Options)) (*lambda.GetLayerVersionOutput, error)
	GetLayerVersionByArnMock               func(arg1 context.Context, arg2 *lambda.GetLayerVersionByArnInput, arg3 ...func(*lambda.Options)) (*lambda.GetLayerVersionByArnOutput, error)
	GetLayerVersionPolicyMock              func(arg1 context.Context, arg2 *lambda.GetLayerVersionPolicyInput, arg3 ...func(*lambda.Options)) (*lambda.GetLayerVersionPolicyOutput, error)
	GetPolicyMock                          func(arg1 context.Context, arg2 *lambda.GetPolicyInput, arg3 ...func(*lambda.Options)) (*lambda.GetPolicyOutput, error)
	GetProvisionedConcurrencyConfigMock    func(arg1 context.Context, arg2 *lambda.GetProvisionedConcurrencyConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetProvisionedConcurrencyConfigOutput, error)
	InvokeMock                             func(arg1 context.Context, arg2 *lambda.InvokeInput, arg3 ...func(*lambda.Options)) (*lambda.InvokeOutput, error)
	InvokeAsyncMock                        func(arg1 context.Context, arg2 *lambda.InvokeAsyncInput, arg3 ...func(*lambda.Options)) (*lambda.InvokeAsyncOutput, error)
	ListAliasesMock                        func(arg1 context.Context, arg2 *lambda.ListAliasesInput, arg3 ...func(*lambda.Options)) (*lambda.ListAliasesOutput, error)
	ListCodeSigningConfigsMock             func(arg1 context.Context, arg2 *lambda.ListCodeSigningConfigsInput, arg3 ...func(*lambda.Options)) (*lambda.ListCodeSigningConfigsOutput, error)
	ListEventSourceMappingsMock            func(arg1 context.Context, arg2 *lambda.ListEventSourceMappingsInput, arg3 ...func(*lambda.Options)) (*lambda.ListEventSourceMappingsOutput, error)
	ListFunctionEventInvokeConfigsMock     func(arg1 context.Context, arg2 *lambda.ListFunctionEventInvokeConfigsInput, arg3 ...func(*lambda.Options)) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
	ListFunctionsMock                      func(arg1 context.Context, arg2 *lambda.ListFunctionsInput, arg3 ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error)
	ListFunctionsByCodeSigningConfigMock   func(arg1 context.Context, arg2 *lambda.ListFunctionsByCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.ListFunctionsByCodeSigningConfigOutput, error)
	ListLayerVersionsMock                  func(arg1 context.Context, arg2 *lambda.ListLayerVersionsInput, arg3 ...func(*lambda.Options)) (*lambda.ListLayerVersionsOutput, error)
	ListLayersMock                         func(arg1 context.Context, arg2 *lambda.ListLayersInput, arg3 ...func(*lambda.Options)) (*lambda.ListLayersOutput, error)
	ListProvisionedConcurrencyConfigsMock  func(arg1 context.Context, arg2 *lambda.ListProvisionedConcurrencyConfigsInput, arg3 ...func(*lambda.Options)) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
	ListTagsMock                           func(arg1 context.Context, arg2 *lambda.ListTagsInput, arg3 ...func(*lambda.Options)) (*lambda.ListTagsOutput, error)
	ListVersionsByFunctionMock             func(arg1 context.Context, arg2 *lambda.ListVersionsByFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.ListVersionsByFunctionOutput, error)
	PublishLayerVersionMock                func(arg1 context.Context, arg2 *lambda.PublishLayerVersionInput, arg3 ...func(*lambda.Options)) (*lambda.PublishLayerVersionOutput, error)
	PublishVersionMock                     func(arg1 context.Context, arg2 *lambda.PublishVersionInput, arg3 ...func(*lambda.Options)) (*lambda.PublishVersionOutput, error)
	PutFunctionCodeSigningConfigMock       func(arg1 context.Context, arg2 *lambda.PutFunctionCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.PutFunctionCodeSigningConfigOutput, error)
	PutFunctionConcurrencyMock             func(arg1 context.Context, arg2 *lambda.PutFunctionConcurrencyInput, arg3 ...func(*lambda.Options)) (*lambda.PutFunctionConcurrencyOutput, error)
	PutFunctionEventInvokeConfigMock       func(arg1 context.Context, arg2 *lambda.PutFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.PutFunctionEventInvokeConfigOutput, error)
	PutProvisionedConcurrencyConfigMock    func(arg1 context.Context, arg2 *lambda.PutProvisionedConcurrencyConfigInput, arg3 ...func(*lambda.Options)) (*lambda.PutProvisionedConcurrencyConfigOutput, error)
	RemoveLayerVersionPermissionMock       func(arg1 context.Context, arg2 *lambda.RemoveLayerVersionPermissionInput, arg3 ...func(*lambda.Options)) (*lambda.RemoveLayerVersionPermissionOutput, error)
	RemovePermissionMock                   func(arg1 context.Context, arg2 *lambda.RemovePermissionInput, arg3 ...func(*lambda.Options)) (*lambda.RemovePermissionOutput, error)
	TagResourceMock                        func(arg1 context.Context, arg2 *lambda.TagResourceInput, arg3 ...func(*lambda.Options)) (*lambda.TagResourceOutput, error)
	UntagResourceMock                      func(arg1 context.Context, arg2 *lambda.UntagResourceInput, arg3 ...func(*lambda.Options)) (*lambda.UntagResourceOutput, error)
	UpdateAliasMock                        func(arg1 context.Context, arg2 *lambda.UpdateAliasInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateAliasOutput, error)
	UpdateCodeSigningConfigMock            func(arg1 context.Context, arg2 *lambda.UpdateCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateCodeSigningConfigOutput, error)
	UpdateEventSourceMappingMock           func(arg1 context.Context, arg2 *lambda.UpdateEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateEventSourceMappingOutput, error)
	UpdateFunctionCodeMock                 func(arg1 context.Context, arg2 *lambda.UpdateFunctionCodeInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateFunctionCodeOutput, error)
	UpdateFunctionConfigurationMock        func(arg1 context.Context, arg2 *lambda.UpdateFunctionConfigurationInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateFunctionConfigurationOutput, error)
	UpdateFunctionEventInvokeConfigMock    func(arg1 context.Context, arg2 *lambda.UpdateFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateFunctionEventInvokeConfigOutput, error)
}

// AddLayerVersionPermission mocked funcition
func (m LambdaClientMock) AddLayerVersionPermission(arg1 context.Context, arg2 *lambda.AddLayerVersionPermissionInput, arg3 ...func(*lambda.Options)) (*lambda.AddLayerVersionPermissionOutput, error) {
	return m.AddLayerVersionPermissionMock(arg1, arg2, arg3...)
}

// AddPermission mocked funcition
func (m LambdaClientMock) AddPermission(arg1 context.Context, arg2 *lambda.AddPermissionInput, arg3 ...func(*lambda.Options)) (*lambda.AddPermissionOutput, error) {
	return m.AddPermissionMock(arg1, arg2, arg3...)
}

// CreateAlias mocked funcition
func (m LambdaClientMock) CreateAlias(arg1 context.Context, arg2 *lambda.CreateAliasInput, arg3 ...func(*lambda.Options)) (*lambda.CreateAliasOutput, error) {
	return m.CreateAliasMock(arg1, arg2, arg3...)
}

// CreateCodeSigningConfig mocked funcition
func (m LambdaClientMock) CreateCodeSigningConfig(arg1 context.Context, arg2 *lambda.CreateCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.CreateCodeSigningConfigOutput, error) {
	return m.CreateCodeSigningConfigMock(arg1, arg2, arg3...)
}

// CreateEventSourceMapping mocked funcition
func (m LambdaClientMock) CreateEventSourceMapping(arg1 context.Context, arg2 *lambda.CreateEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.CreateEventSourceMappingOutput, error) {
	return m.CreateEventSourceMappingMock(arg1, arg2, arg3...)
}

// CreateFunction mocked funcition
func (m LambdaClientMock) CreateFunction(arg1 context.Context, arg2 *lambda.CreateFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.CreateFunctionOutput, error) {
	return m.CreateFunctionMock(arg1, arg2, arg3...)
}

// DeleteAlias mocked funcition
func (m LambdaClientMock) DeleteAlias(arg1 context.Context, arg2 *lambda.DeleteAliasInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteAliasOutput, error) {
	return m.DeleteAliasMock(arg1, arg2, arg3...)
}

// DeleteCodeSigningConfig mocked funcition
func (m LambdaClientMock) DeleteCodeSigningConfig(arg1 context.Context, arg2 *lambda.DeleteCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteCodeSigningConfigOutput, error) {
	return m.DeleteCodeSigningConfigMock(arg1, arg2, arg3...)
}

// DeleteEventSourceMapping mocked funcition
func (m LambdaClientMock) DeleteEventSourceMapping(arg1 context.Context, arg2 *lambda.DeleteEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteEventSourceMappingOutput, error) {
	return m.DeleteEventSourceMappingMock(arg1, arg2, arg3...)
}

// DeleteFunction mocked funcition
func (m LambdaClientMock) DeleteFunction(arg1 context.Context, arg2 *lambda.DeleteFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionOutput, error) {
	return m.DeleteFunctionMock(arg1, arg2, arg3...)
}

// DeleteFunctionCodeSigningConfig mocked funcition
func (m LambdaClientMock) DeleteFunctionCodeSigningConfig(arg1 context.Context, arg2 *lambda.DeleteFunctionCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionCodeSigningConfigOutput, error) {
	return m.DeleteFunctionCodeSigningConfigMock(arg1, arg2, arg3...)
}

// DeleteFunctionConcurrency mocked funcition
func (m LambdaClientMock) DeleteFunctionConcurrency(arg1 context.Context, arg2 *lambda.DeleteFunctionConcurrencyInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionConcurrencyOutput, error) {
	return m.DeleteFunctionConcurrencyMock(arg1, arg2, arg3...)
}

// DeleteFunctionEventInvokeConfig mocked funcition
func (m LambdaClientMock) DeleteFunctionEventInvokeConfig(arg1 context.Context, arg2 *lambda.DeleteFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteFunctionEventInvokeConfigOutput, error) {
	return m.DeleteFunctionEventInvokeConfigMock(arg1, arg2, arg3...)
}

// DeleteLayerVersion mocked funcition
func (m LambdaClientMock) DeleteLayerVersion(arg1 context.Context, arg2 *lambda.DeleteLayerVersionInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteLayerVersionOutput, error) {
	return m.DeleteLayerVersionMock(arg1, arg2, arg3...)
}

// DeleteProvisionedConcurrencyConfig mocked funcition
func (m LambdaClientMock) DeleteProvisionedConcurrencyConfig(arg1 context.Context, arg2 *lambda.DeleteProvisionedConcurrencyConfigInput, arg3 ...func(*lambda.Options)) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error) {
	return m.DeleteProvisionedConcurrencyConfigMock(arg1, arg2, arg3...)
}

// GetAccountSettings mocked funcition
func (m LambdaClientMock) GetAccountSettings(arg1 context.Context, arg2 *lambda.GetAccountSettingsInput, arg3 ...func(*lambda.Options)) (*lambda.GetAccountSettingsOutput, error) {
	return m.GetAccountSettingsMock(arg1, arg2, arg3...)
}

// GetAlias mocked funcition
func (m LambdaClientMock) GetAlias(arg1 context.Context, arg2 *lambda.GetAliasInput, arg3 ...func(*lambda.Options)) (*lambda.GetAliasOutput, error) {
	return m.GetAliasMock(arg1, arg2, arg3...)
}

// GetCodeSigningConfig mocked funcition
func (m LambdaClientMock) GetCodeSigningConfig(arg1 context.Context, arg2 *lambda.GetCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetCodeSigningConfigOutput, error) {
	return m.GetCodeSigningConfigMock(arg1, arg2, arg3...)
}

// GetEventSourceMapping mocked funcition
func (m LambdaClientMock) GetEventSourceMapping(arg1 context.Context, arg2 *lambda.GetEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.GetEventSourceMappingOutput, error) {
	return m.GetEventSourceMappingMock(arg1, arg2, arg3...)
}

// GetFunction mocked funcition
func (m LambdaClientMock) GetFunction(arg1 context.Context, arg2 *lambda.GetFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error) {
	return m.GetFunctionMock(arg1, arg2, arg3...)
}

// GetFunctionCodeSigningConfig mocked funcition
func (m LambdaClientMock) GetFunctionCodeSigningConfig(arg1 context.Context, arg2 *lambda.GetFunctionCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionCodeSigningConfigOutput, error) {
	return m.GetFunctionCodeSigningConfigMock(arg1, arg2, arg3...)
}

// GetFunctionConcurrency mocked funcition
func (m LambdaClientMock) GetFunctionConcurrency(arg1 context.Context, arg2 *lambda.GetFunctionConcurrencyInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionConcurrencyOutput, error) {
	return m.GetFunctionConcurrencyMock(arg1, arg2, arg3...)
}

// GetFunctionConfiguration mocked funcition
func (m LambdaClientMock) GetFunctionConfiguration(arg1 context.Context, arg2 *lambda.GetFunctionConfigurationInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionConfigurationOutput, error) {
	return m.GetFunctionConfigurationMock(arg1, arg2, arg3...)
}

// GetFunctionEventInvokeConfig mocked funcition
func (m LambdaClientMock) GetFunctionEventInvokeConfig(arg1 context.Context, arg2 *lambda.GetFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetFunctionEventInvokeConfigOutput, error) {
	return m.GetFunctionEventInvokeConfigMock(arg1, arg2, arg3...)
}

// GetLayerVersion mocked funcition
func (m LambdaClientMock) GetLayerVersion(arg1 context.Context, arg2 *lambda.GetLayerVersionInput, arg3 ...func(*lambda.Options)) (*lambda.GetLayerVersionOutput, error) {
	return m.GetLayerVersionMock(arg1, arg2, arg3...)
}

// GetLayerVersionByArn mocked funcition
func (m LambdaClientMock) GetLayerVersionByArn(arg1 context.Context, arg2 *lambda.GetLayerVersionByArnInput, arg3 ...func(*lambda.Options)) (*lambda.GetLayerVersionByArnOutput, error) {
	return m.GetLayerVersionByArnMock(arg1, arg2, arg3...)
}

// GetLayerVersionPolicy mocked funcition
func (m LambdaClientMock) GetLayerVersionPolicy(arg1 context.Context, arg2 *lambda.GetLayerVersionPolicyInput, arg3 ...func(*lambda.Options)) (*lambda.GetLayerVersionPolicyOutput, error) {
	return m.GetLayerVersionPolicyMock(arg1, arg2, arg3...)
}

// GetPolicy mocked funcition
func (m LambdaClientMock) GetPolicy(arg1 context.Context, arg2 *lambda.GetPolicyInput, arg3 ...func(*lambda.Options)) (*lambda.GetPolicyOutput, error) {
	return m.GetPolicyMock(arg1, arg2, arg3...)
}

// GetProvisionedConcurrencyConfig mocked funcition
func (m LambdaClientMock) GetProvisionedConcurrencyConfig(arg1 context.Context, arg2 *lambda.GetProvisionedConcurrencyConfigInput, arg3 ...func(*lambda.Options)) (*lambda.GetProvisionedConcurrencyConfigOutput, error) {
	return m.GetProvisionedConcurrencyConfigMock(arg1, arg2, arg3...)
}

// Invoke mocked funcition
func (m LambdaClientMock) Invoke(arg1 context.Context, arg2 *lambda.InvokeInput, arg3 ...func(*lambda.Options)) (*lambda.InvokeOutput, error) {
	return m.InvokeMock(arg1, arg2, arg3...)
}

// InvokeAsync mocked funcition
func (m LambdaClientMock) InvokeAsync(arg1 context.Context, arg2 *lambda.InvokeAsyncInput, arg3 ...func(*lambda.Options)) (*lambda.InvokeAsyncOutput, error) {
	return m.InvokeAsyncMock(arg1, arg2, arg3...)
}

// ListAliases mocked funcition
func (m LambdaClientMock) ListAliases(arg1 context.Context, arg2 *lambda.ListAliasesInput, arg3 ...func(*lambda.Options)) (*lambda.ListAliasesOutput, error) {
	return m.ListAliasesMock(arg1, arg2, arg3...)
}

// ListCodeSigningConfigs mocked funcition
func (m LambdaClientMock) ListCodeSigningConfigs(arg1 context.Context, arg2 *lambda.ListCodeSigningConfigsInput, arg3 ...func(*lambda.Options)) (*lambda.ListCodeSigningConfigsOutput, error) {
	return m.ListCodeSigningConfigsMock(arg1, arg2, arg3...)
}

// ListEventSourceMappings mocked funcition
func (m LambdaClientMock) ListEventSourceMappings(arg1 context.Context, arg2 *lambda.ListEventSourceMappingsInput, arg3 ...func(*lambda.Options)) (*lambda.ListEventSourceMappingsOutput, error) {
	return m.ListEventSourceMappingsMock(arg1, arg2, arg3...)
}

// ListFunctionEventInvokeConfigs mocked funcition
func (m LambdaClientMock) ListFunctionEventInvokeConfigs(arg1 context.Context, arg2 *lambda.ListFunctionEventInvokeConfigsInput, arg3 ...func(*lambda.Options)) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
	return m.ListFunctionEventInvokeConfigsMock(arg1, arg2, arg3...)
}

// ListFunctions mocked funcition
func (m LambdaClientMock) ListFunctions(arg1 context.Context, arg2 *lambda.ListFunctionsInput, arg3 ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error) {
	return m.ListFunctionsMock(arg1, arg2, arg3...)
}

// ListFunctionsByCodeSigningConfig mocked funcition
func (m LambdaClientMock) ListFunctionsByCodeSigningConfig(arg1 context.Context, arg2 *lambda.ListFunctionsByCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.ListFunctionsByCodeSigningConfigOutput, error) {
	return m.ListFunctionsByCodeSigningConfigMock(arg1, arg2, arg3...)
}

// ListLayerVersions mocked funcition
func (m LambdaClientMock) ListLayerVersions(arg1 context.Context, arg2 *lambda.ListLayerVersionsInput, arg3 ...func(*lambda.Options)) (*lambda.ListLayerVersionsOutput, error) {
	return m.ListLayerVersionsMock(arg1, arg2, arg3...)
}

// ListLayers mocked funcition
func (m LambdaClientMock) ListLayers(arg1 context.Context, arg2 *lambda.ListLayersInput, arg3 ...func(*lambda.Options)) (*lambda.ListLayersOutput, error) {
	return m.ListLayersMock(arg1, arg2, arg3...)
}

// ListProvisionedConcurrencyConfigs mocked funcition
func (m LambdaClientMock) ListProvisionedConcurrencyConfigs(arg1 context.Context, arg2 *lambda.ListProvisionedConcurrencyConfigsInput, arg3 ...func(*lambda.Options)) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
	return m.ListProvisionedConcurrencyConfigsMock(arg1, arg2, arg3...)
}

// ListTags mocked funcition
func (m LambdaClientMock) ListTags(arg1 context.Context, arg2 *lambda.ListTagsInput, arg3 ...func(*lambda.Options)) (*lambda.ListTagsOutput, error) {
	return m.ListTagsMock(arg1, arg2, arg3...)
}

// ListVersionsByFunction mocked funcition
func (m LambdaClientMock) ListVersionsByFunction(arg1 context.Context, arg2 *lambda.ListVersionsByFunctionInput, arg3 ...func(*lambda.Options)) (*lambda.ListVersionsByFunctionOutput, error) {
	return m.ListVersionsByFunctionMock(arg1, arg2, arg3...)
}

// PublishLayerVersion mocked funcition
func (m LambdaClientMock) PublishLayerVersion(arg1 context.Context, arg2 *lambda.PublishLayerVersionInput, arg3 ...func(*lambda.Options)) (*lambda.PublishLayerVersionOutput, error) {
	return m.PublishLayerVersionMock(arg1, arg2, arg3...)
}

// PublishVersion mocked funcition
func (m LambdaClientMock) PublishVersion(arg1 context.Context, arg2 *lambda.PublishVersionInput, arg3 ...func(*lambda.Options)) (*lambda.PublishVersionOutput, error) {
	return m.PublishVersionMock(arg1, arg2, arg3...)
}

// PutFunctionCodeSigningConfig mocked funcition
func (m LambdaClientMock) PutFunctionCodeSigningConfig(arg1 context.Context, arg2 *lambda.PutFunctionCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.PutFunctionCodeSigningConfigOutput, error) {
	return m.PutFunctionCodeSigningConfigMock(arg1, arg2, arg3...)
}

// PutFunctionConcurrency mocked funcition
func (m LambdaClientMock) PutFunctionConcurrency(arg1 context.Context, arg2 *lambda.PutFunctionConcurrencyInput, arg3 ...func(*lambda.Options)) (*lambda.PutFunctionConcurrencyOutput, error) {
	return m.PutFunctionConcurrencyMock(arg1, arg2, arg3...)
}

// PutFunctionEventInvokeConfig mocked funcition
func (m LambdaClientMock) PutFunctionEventInvokeConfig(arg1 context.Context, arg2 *lambda.PutFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.PutFunctionEventInvokeConfigOutput, error) {
	return m.PutFunctionEventInvokeConfigMock(arg1, arg2, arg3...)
}

// PutProvisionedConcurrencyConfig mocked funcition
func (m LambdaClientMock) PutProvisionedConcurrencyConfig(arg1 context.Context, arg2 *lambda.PutProvisionedConcurrencyConfigInput, arg3 ...func(*lambda.Options)) (*lambda.PutProvisionedConcurrencyConfigOutput, error) {
	return m.PutProvisionedConcurrencyConfigMock(arg1, arg2, arg3...)
}

// RemoveLayerVersionPermission mocked funcition
func (m LambdaClientMock) RemoveLayerVersionPermission(arg1 context.Context, arg2 *lambda.RemoveLayerVersionPermissionInput, arg3 ...func(*lambda.Options)) (*lambda.RemoveLayerVersionPermissionOutput, error) {
	return m.RemoveLayerVersionPermissionMock(arg1, arg2, arg3...)
}

// RemovePermission mocked funcition
func (m LambdaClientMock) RemovePermission(arg1 context.Context, arg2 *lambda.RemovePermissionInput, arg3 ...func(*lambda.Options)) (*lambda.RemovePermissionOutput, error) {
	return m.RemovePermissionMock(arg1, arg2, arg3...)
}

// TagResource mocked funcition
func (m LambdaClientMock) TagResource(arg1 context.Context, arg2 *lambda.TagResourceInput, arg3 ...func(*lambda.Options)) (*lambda.TagResourceOutput, error) {
	return m.TagResourceMock(arg1, arg2, arg3...)
}

// UntagResource mocked funcition
func (m LambdaClientMock) UntagResource(arg1 context.Context, arg2 *lambda.UntagResourceInput, arg3 ...func(*lambda.Options)) (*lambda.UntagResourceOutput, error) {
	return m.UntagResourceMock(arg1, arg2, arg3...)
}

// UpdateAlias mocked funcition
func (m LambdaClientMock) UpdateAlias(arg1 context.Context, arg2 *lambda.UpdateAliasInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateAliasOutput, error) {
	return m.UpdateAliasMock(arg1, arg2, arg3...)
}

// UpdateCodeSigningConfig mocked funcition
func (m LambdaClientMock) UpdateCodeSigningConfig(arg1 context.Context, arg2 *lambda.UpdateCodeSigningConfigInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateCodeSigningConfigOutput, error) {
	return m.UpdateCodeSigningConfigMock(arg1, arg2, arg3...)
}

// UpdateEventSourceMapping mocked funcition
func (m LambdaClientMock) UpdateEventSourceMapping(arg1 context.Context, arg2 *lambda.UpdateEventSourceMappingInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateEventSourceMappingOutput, error) {
	return m.UpdateEventSourceMappingMock(arg1, arg2, arg3...)
}

// UpdateFunctionCode mocked funcition
func (m LambdaClientMock) UpdateFunctionCode(arg1 context.Context, arg2 *lambda.UpdateFunctionCodeInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateFunctionCodeOutput, error) {
	return m.UpdateFunctionCodeMock(arg1, arg2, arg3...)
}

// UpdateFunctionConfiguration mocked funcition
func (m LambdaClientMock) UpdateFunctionConfiguration(arg1 context.Context, arg2 *lambda.UpdateFunctionConfigurationInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateFunctionConfigurationOutput, error) {
	return m.UpdateFunctionConfigurationMock(arg1, arg2, arg3...)
}

// UpdateFunctionEventInvokeConfig mocked funcition
func (m LambdaClientMock) UpdateFunctionEventInvokeConfig(arg1 context.Context, arg2 *lambda.UpdateFunctionEventInvokeConfigInput, arg3 ...func(*lambda.Options)) (*lambda.UpdateFunctionEventInvokeConfigOutput, error) {
	return m.UpdateFunctionEventInvokeConfigMock(arg1, arg2, arg3...)
}
