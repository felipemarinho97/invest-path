package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3control"
)

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/felipemarinho97/go-mock-gen
 */

// IS3ControlClient generic client
type IS3ControlClient interface {
	// Creates an access point and associates it with the specified bucket. For more
	// information, see Managing Data Access with Amazon S3 Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html) in
	// the Amazon S3 User Guide. S3 on Outposts only supports VPC-style access points.
	// For more information, see  Accessing Amazon S3 on Outposts using virtual private
	// cloud (VPC) only access points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
	// Amazon S3 User Guide. All Amazon S3 on Outposts REST API requests for this
	// action require an additional parameter of x-amz-outpost-id to be passed with the
	// request and an S3 on Outposts endpoint hostname prefix instead of s3-control.
	// For an example of the request syntax for Amazon S3 on Outposts that uses the S3
	// on Outposts endpoint hostname prefix and the x-amz-outpost-id derived using the
	// access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html#API_control_CreateAccessPoint_Examples)
	// section. The following actions are related to CreateAccessPoint:
	//
	// *
	// GetAccessPoint
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html)
	//
	// *
	// DeleteAccessPoint
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html)
	//
	// *
	// ListAccessPoints
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html)
	//
	CreateAccessPoint(arg1 context.Context, arg2 *s3control.CreateAccessPointInput, arg3 ...func(*s3control.Options)) (*s3control.CreateAccessPointOutput, error)
	// Creates an Object Lambda Access Point. For more information, see Transforming
	// objects with Object Lambda Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/transforming-objects.html)
	// in the Amazon S3 User Guide. The following actions are related to
	// CreateAccessPointForObjectLambda:
	//
	// * DeleteAccessPointForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointForObjectLambda.html)
	//
	// *
	// GetAccessPointForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointForObjectLambda.html)
	//
	// *
	// ListAccessPointsForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPointsForObjectLambda.html)
	//
	CreateAccessPointForObjectLambda(arg1 context.Context, arg2 *s3control.CreateAccessPointForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.CreateAccessPointForObjectLambdaOutput, error)
	// This action creates an Amazon S3 on Outposts bucket. To create an S3 bucket, see
	// Create Bucket
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html) in the
	// Amazon S3 API Reference. Creates a new Outposts bucket. By creating the bucket,
	// you become the bucket owner. To create an Outposts bucket, you must have S3 on
	// Outposts. For more information, see Using Amazon S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in
	// Amazon S3 User Guide. Not every string is an acceptable bucket name. For
	// information on bucket naming restrictions, see Working with Amazon S3 Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/BucketRestrictions.html#bucketnamingrules).
	// S3 on Outposts buckets support:
	//
	// * Tags
	//
	// * LifecycleConfigurations for deleting
	// expired objects
	//
	// For a complete list of restrictions and Amazon S3 feature
	// limitations on S3 on Outposts, see  Amazon S3 on Outposts Restrictions and
	// Limitations
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OnOutpostsRestrictionsLimitations.html).
	// For an example of the request syntax for Amazon S3 on Outposts that uses the S3
	// on Outposts endpoint hostname prefix and x-amz-outpost-id in your API request,
	// see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html#API_control_CreateBucket_Examples)
	// section. The following actions are related to CreateBucket for Amazon S3 on
	// Outposts:
	//
	// * PutObject
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
	//
	// *
	// GetBucket
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html)
	//
	// *
	// DeleteBucket
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html)
	//
	// *
	// CreateAccessPoint
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html)
	//
	// *
	// PutAccessPointPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html)
	//
	CreateBucket(arg1 context.Context, arg2 *s3control.CreateBucketInput, arg3 ...func(*s3control.Options)) (*s3control.CreateBucketOutput, error)
	// You can use S3 Batch Operations to perform large-scale batch actions on Amazon
	// S3 objects. Batch Operations can run a single action on lists of Amazon S3
	// objects that you specify. For more information, see S3 Batch Operations
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html) in the
	// Amazon S3 User Guide. This action creates a S3 Batch Operations job. Related
	// actions include:
	//
	// * DescribeJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html)
	//
	// *
	// ListJobs
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html)
	//
	// *
	// UpdateJobPriority
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobPriority.html)
	//
	// *
	// UpdateJobStatus
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html)
	//
	// *
	// JobOperation
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_JobOperation.html)
	//
	CreateJob(arg1 context.Context, arg2 *s3control.CreateJobInput, arg3 ...func(*s3control.Options)) (*s3control.CreateJobOutput, error)
	// Deletes the specified access point. All Amazon S3 on Outposts REST API requests
	// for this action require an additional parameter of x-amz-outpost-id to be passed
	// with the request and an S3 on Outposts endpoint hostname prefix instead of
	// s3-control. For an example of the request syntax for Amazon S3 on Outposts that
	// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
	// derived using the access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html#API_control_DeleteAccessPoint_Examples)
	// section. The following actions are related to DeleteAccessPoint:
	//
	// *
	// CreateAccessPoint
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html)
	//
	// *
	// GetAccessPoint
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html)
	//
	// *
	// ListAccessPoints
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html)
	//
	DeleteAccessPoint(arg1 context.Context, arg2 *s3control.DeleteAccessPointInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointOutput, error)
	// Deletes the specified Object Lambda Access Point. The following actions are
	// related to DeleteAccessPointForObjectLambda:
	//
	// * CreateAccessPointForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPointForObjectLambda.html)
	//
	// *
	// GetAccessPointForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointForObjectLambda.html)
	//
	// *
	// ListAccessPointsForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPointsForObjectLambda.html)
	//
	DeleteAccessPointForObjectLambda(arg1 context.Context, arg2 *s3control.DeleteAccessPointForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointForObjectLambdaOutput, error)
	// Deletes the access point policy for the specified access point. All Amazon S3 on
	// Outposts REST API requests for this action require an additional parameter of
	// x-amz-outpost-id to be passed with the request and an S3 on Outposts endpoint
	// hostname prefix instead of s3-control. For an example of the request syntax for
	// Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and
	// the x-amz-outpost-id derived using the access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html#API_control_DeleteAccessPointPolicy_Examples)
	// section. The following actions are related to DeleteAccessPointPolicy:
	//
	// *
	// PutAccessPointPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html)
	//
	// *
	// GetAccessPointPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicy.html)
	//
	DeleteAccessPointPolicy(arg1 context.Context, arg2 *s3control.DeleteAccessPointPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointPolicyOutput, error)
	// Removes the resource policy for an Object Lambda Access Point. The following
	// actions are related to DeleteAccessPointPolicyForObjectLambda:
	//
	// *
	// GetAccessPointPolicyForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicyForObjectLambda.html)
	//
	// *
	// PutAccessPointPolicyForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicyForObjectLambda.html)
	//
	DeleteAccessPointPolicyForObjectLambda(arg1 context.Context, arg2 *s3control.DeleteAccessPointPolicyForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointPolicyForObjectLambdaOutput, error)
	// This action deletes an Amazon S3 on Outposts bucket. To delete an S3 bucket, see
	// DeleteBucket
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html) in the
	// Amazon S3 API Reference. Deletes the Amazon S3 on Outposts bucket. All objects
	// (including all object versions and delete markers) in the bucket must be deleted
	// before the bucket itself can be deleted. For more information, see Using Amazon
	// S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in
	// Amazon S3 User Guide. All Amazon S3 on Outposts REST API requests for this
	// action require an additional parameter of x-amz-outpost-id to be passed with the
	// request and an S3 on Outposts endpoint hostname prefix instead of s3-control.
	// For an example of the request syntax for Amazon S3 on Outposts that uses the S3
	// on Outposts endpoint hostname prefix and the x-amz-outpost-id derived using the
	// access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html#API_control_DeleteBucket_Examples)
	// section. Related Resources
	//
	// * CreateBucket
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html)
	//
	// *
	// GetBucket
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html)
	//
	// *
	// DeleteObject
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html)
	//
	DeleteBucket(arg1 context.Context, arg2 *s3control.DeleteBucketInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketOutput, error)
	// This action deletes an Amazon S3 on Outposts bucket's lifecycle configuration.
	// To delete an S3 bucket's lifecycle configuration, see DeleteBucketLifecycle
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html)
	// in the Amazon S3 API Reference. Deletes the lifecycle configuration from the
	// specified Outposts bucket. Amazon S3 on Outposts removes all the lifecycle
	// configuration rules in the lifecycle subresource associated with the bucket.
	// Your objects never expire, and Amazon S3 on Outposts no longer automatically
	// deletes any objects on the basis of rules contained in the deleted lifecycle
	// configuration. For more information, see Using Amazon S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the s3-outposts:DeleteLifecycleConfiguration action. By default, the bucket
	// owner has this permission and the Outposts bucket owner can grant this
	// permission to others. All Amazon S3 on Outposts REST API requests for this
	// action require an additional parameter of x-amz-outpost-id to be passed with the
	// request and an S3 on Outposts endpoint hostname prefix instead of s3-control.
	// For an example of the request syntax for Amazon S3 on Outposts that uses the S3
	// on Outposts endpoint hostname prefix and the x-amz-outpost-id derived using the
	// access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html#API_control_DeleteBucketLifecycleConfiguration_Examples)
	// section. For more information about object expiration, see Elements to Describe
	// Lifecycle Actions
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#intro-lifecycle-rules-actions).
	// Related actions include:
	//
	// * PutBucketLifecycleConfiguration
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html)
	//
	// *
	// GetBucketLifecycleConfiguration
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html)
	//
	DeleteBucketLifecycleConfiguration(arg1 context.Context, arg2 *s3control.DeleteBucketLifecycleConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketLifecycleConfigurationOutput, error)
	// This action deletes an Amazon S3 on Outposts bucket policy. To delete an S3
	// bucket policy, see DeleteBucketPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketPolicy.html) in
	// the Amazon S3 API Reference. This implementation of the DELETE action uses the
	// policy subresource to delete the policy of a specified Amazon S3 on Outposts
	// bucket. If you are using an identity other than the root user of the AWS account
	// that owns the bucket, the calling identity must have the
	// s3-outposts:DeleteBucketPolicy permissions on the specified Outposts bucket and
	// belong to the bucket owner's account to use this action. For more information,
	// see Using Amazon S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in
	// Amazon S3 User Guide. If you don't have DeleteBucketPolicy permissions, Amazon
	// S3 returns a 403 Access Denied error. If you have the correct permissions, but
	// you're not using an identity that belongs to the bucket owner's account, Amazon
	// S3 returns a 405 Method Not Allowed error. As a security precaution, the root
	// user of the AWS account that owns a bucket can always use this action, even if
	// the policy explicitly denies the root user the ability to perform this action.
	// For more information about bucket policies, see Using Bucket Policies and User
	// Policies
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html). All
	// Amazon S3 on Outposts REST API requests for this action require an additional
	// parameter of x-amz-outpost-id to be passed with the request and an S3 on
	// Outposts endpoint hostname prefix instead of s3-control. For an example of the
	// request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint
	// hostname prefix and the x-amz-outpost-id derived using the access point ARN, see
	// the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html#API_control_DeleteBucketPolicy_Examples)
	// section. The following actions are related to DeleteBucketPolicy:
	//
	// *
	// GetBucketPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html)
	//
	// *
	// PutBucketPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketPolicy.html)
	//
	DeleteBucketPolicy(arg1 context.Context, arg2 *s3control.DeleteBucketPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketPolicyOutput, error)
	// This action deletes an Amazon S3 on Outposts bucket's tags. To delete an S3
	// bucket tags, see DeleteBucketTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html)
	// in the Amazon S3 API Reference. Deletes the tags from the Outposts bucket. For
	// more information, see Using Amazon S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the PutBucketTagging action. By default, the bucket owner has this permission
	// and can grant this permission to others. All Amazon S3 on Outposts REST API
	// requests for this action require an additional parameter of x-amz-outpost-id to
	// be passed with the request and an S3 on Outposts endpoint hostname prefix
	// instead of s3-control. For an example of the request syntax for Amazon S3 on
	// Outposts that uses the S3 on Outposts endpoint hostname prefix and the
	// x-amz-outpost-id derived using the access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html#API_control_DeleteBucketTagging_Examples)
	// section. The following actions are related to DeleteBucketTagging:
	//
	// *
	// GetBucketTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html)
	//
	// *
	// PutBucketTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html)
	//
	DeleteBucketTagging(arg1 context.Context, arg2 *s3control.DeleteBucketTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketTaggingOutput, error)
	// Removes the entire tag set from the specified S3 Batch Operations job. To use
	// this operation, you must have permission to perform the s3:DeleteJobTagging
	// action. For more information, see Controlling access and labeling jobs using
	// tags
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags)
	// in the Amazon S3 User Guide. Related actions include:
	//
	// * CreateJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
	//
	// *
	// GetJobTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetJobTagging.html)
	//
	// *
	// PutJobTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutJobTagging.html)
	//
	DeleteJobTagging(arg1 context.Context, arg2 *s3control.DeleteJobTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteJobTaggingOutput, error)
	// Removes the PublicAccessBlock configuration for an AWS account. For more
	// information, see  Using Amazon S3 block public access
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).
	// Related actions include:
	//
	// * GetPublicAccessBlock
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetPublicAccessBlock.html)
	//
	// *
	// PutPublicAccessBlock
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutPublicAccessBlock.html)
	//
	DeletePublicAccessBlock(arg1 context.Context, arg2 *s3control.DeletePublicAccessBlockInput, arg3 ...func(*s3control.Options)) (*s3control.DeletePublicAccessBlockOutput, error)
	// Deletes the Amazon S3 Storage Lens configuration. For more information about S3
	// Storage Lens, see Assessing your storage activity and usage with Amazon S3
	// Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html) in the
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the s3:DeleteStorageLensConfiguration action. For more information, see Setting
	// permissions to use Amazon S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html)
	// in the Amazon S3 User Guide.
	//
	DeleteStorageLensConfiguration(arg1 context.Context, arg2 *s3control.DeleteStorageLensConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteStorageLensConfigurationOutput, error)
	// Deletes the Amazon S3 Storage Lens configuration tags. For more information
	// about S3 Storage Lens, see Assessing your storage activity and usage with Amazon
	// S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html) in the
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the s3:DeleteStorageLensConfigurationTagging action. For more information, see
	// Setting permissions to use Amazon S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html)
	// in the Amazon S3 User Guide.
	//
	DeleteStorageLensConfigurationTagging(arg1 context.Context, arg2 *s3control.DeleteStorageLensConfigurationTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteStorageLensConfigurationTaggingOutput, error)
	// Retrieves the configuration parameters and status for a Batch Operations job.
	// For more information, see S3 Batch Operations
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html) in the
	// Amazon S3 User Guide. Related actions include:
	//
	// * CreateJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
	//
	// *
	// ListJobs
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html)
	//
	// *
	// UpdateJobPriority
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobPriority.html)
	//
	// *
	// UpdateJobStatus
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html)
	//
	DescribeJob(arg1 context.Context, arg2 *s3control.DescribeJobInput, arg3 ...func(*s3control.Options)) (*s3control.DescribeJobOutput, error)
	// Returns configuration information about the specified access point. All Amazon
	// S3 on Outposts REST API requests for this action require an additional parameter
	// of x-amz-outpost-id to be passed with the request and an S3 on Outposts endpoint
	// hostname prefix instead of s3-control. For an example of the request syntax for
	// Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and
	// the x-amz-outpost-id derived using the access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html#API_control_GetAccessPoint_Examples)
	// section. The following actions are related to GetAccessPoint:
	//
	// *
	// CreateAccessPoint
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html)
	//
	// *
	// DeleteAccessPoint
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html)
	//
	// *
	// ListAccessPoints
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html)
	//
	GetAccessPoint(arg1 context.Context, arg2 *s3control.GetAccessPointInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointOutput, error)
	// Returns configuration for an Object Lambda Access Point. The following actions
	// are related to GetAccessPointConfigurationForObjectLambda:
	//
	// *
	// PutAccessPointConfigurationForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointConfigurationForObjectLambda.html)
	//
	GetAccessPointConfigurationForObjectLambda(arg1 context.Context, arg2 *s3control.GetAccessPointConfigurationForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error)
	// Returns configuration information about the specified Object Lambda Access Point
	// The following actions are related to GetAccessPointForObjectLambda:
	//
	// *
	// CreateAccessPointForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPointForObjectLambda.html)
	//
	// *
	// DeleteAccessPointForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointForObjectLambda.html)
	//
	// *
	// ListAccessPointsForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPointsForObjectLambda.html)
	//
	GetAccessPointForObjectLambda(arg1 context.Context, arg2 *s3control.GetAccessPointForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointForObjectLambdaOutput, error)
	// Returns the access point policy associated with the specified access point. The
	// following actions are related to GetAccessPointPolicy:
	//
	// * PutAccessPointPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html)
	//
	// *
	// DeleteAccessPointPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html)
	//
	GetAccessPointPolicy(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyOutput, error)
	// Returns the resource policy for an Object Lambda Access Point. The following
	// actions are related to GetAccessPointPolicyForObjectLambda:
	//
	// *
	// DeleteAccessPointPolicyForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicyForObjectLambda.html)
	//
	// *
	// PutAccessPointPolicyForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicyForObjectLambda.html)
	//
	GetAccessPointPolicyForObjectLambda(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error)
	// Indicates whether the specified access point currently has a policy that allows
	// public access. For more information about public access through access points,
	// see Managing Data Access with Amazon S3 access points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html) in
	// the Amazon S3 User Guide.
	//
	GetAccessPointPolicyStatus(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyStatusInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusOutput, error)
	// Returns the status of the resource policy associated with an Object Lambda
	// Access Point.
	//
	GetAccessPointPolicyStatusForObjectLambda(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyStatusForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error)
	// Gets an Amazon S3 on Outposts bucket. For more information, see  Using Amazon S3
	// on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
	// Amazon S3 User Guide. If you are using an identity other than the root user of
	// the AWS account that owns the Outposts bucket, the calling identity must have
	// the s3-outposts:GetBucket permissions on the specified Outposts bucket and
	// belong to the Outposts bucket owner's account in order to use this action. Only
	// users from Outposts bucket owner account with the right permissions can perform
	// actions on an Outposts bucket. If you don't have s3-outposts:GetBucket
	// permissions or you're not using an identity that belongs to the bucket owner's
	// account, Amazon S3 returns a 403 Access Denied error. The following actions are
	// related to GetBucket for Amazon S3 on Outposts: All Amazon S3 on Outposts REST
	// API requests for this action require an additional parameter of x-amz-outpost-id
	// to be passed with the request and an S3 on Outposts endpoint hostname prefix
	// instead of s3-control. For an example of the request syntax for Amazon S3 on
	// Outposts that uses the S3 on Outposts endpoint hostname prefix and the
	// x-amz-outpost-id derived using the access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html#API_control_GetBucket_Examples)
	// section.
	//
	// * PutObject
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
	//
	// *
	// CreateBucket
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html)
	//
	// *
	// DeleteBucket
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html)
	//
	GetBucket(arg1 context.Context, arg2 *s3control.GetBucketInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketOutput, error)
	// This action gets an Amazon S3 on Outposts bucket's lifecycle configuration. To
	// get an S3 bucket's lifecycle configuration, see GetBucketLifecycleConfiguration
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html)
	// in the Amazon S3 API Reference. Returns the lifecycle configuration information
	// set on the Outposts bucket. For more information, see Using Amazon S3 on
	// Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) and
	// for information about lifecycle configuration, see  Object Lifecycle Management
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the s3-outposts:GetLifecycleConfiguration action. The Outposts bucket owner has
	// this permission, by default. The bucket owner can grant this permission to
	// others. For more information about permissions, see Permissions Related to
	// Bucket Subresource Operations
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
	// and Managing Access Permissions to Your Amazon S3 Resources
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html).
	// All Amazon S3 on Outposts REST API requests for this action require an
	// additional parameter of x-amz-outpost-id to be passed with the request and an S3
	// on Outposts endpoint hostname prefix instead of s3-control. For an example of
	// the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts
	// endpoint hostname prefix and the x-amz-outpost-id derived using the access point
	// ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html#API_control_GetBucketLifecycleConfiguration_Examples)
	// section. GetBucketLifecycleConfiguration has the following special error:
	//
	// *
	// Error code: NoSuchLifecycleConfiguration
	//
	// * Description: The lifecycle
	// configuration does not exist.
	//
	// * HTTP Status Code: 404 Not Found
	//
	// * SOAP Fault
	// Code Prefix: Client
	//
	// The following actions are related to
	// GetBucketLifecycleConfiguration:
	//
	// * PutBucketLifecycleConfiguration
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html)
	//
	// *
	// DeleteBucketLifecycleConfiguration
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html)
	//
	GetBucketLifecycleConfiguration(arg1 context.Context, arg2 *s3control.GetBucketLifecycleConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketLifecycleConfigurationOutput, error)
	// This action gets a bucket policy for an Amazon S3 on Outposts bucket. To get a
	// policy for an S3 bucket, see GetBucketPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicy.html) in
	// the Amazon S3 API Reference. Returns the policy of a specified Outposts bucket.
	// For more information, see Using Amazon S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
	// Amazon S3 User Guide. If you are using an identity other than the root user of
	// the AWS account that owns the bucket, the calling identity must have the
	// GetBucketPolicy permissions on the specified bucket and belong to the bucket
	// owner's account in order to use this action. Only users from Outposts bucket
	// owner account with the right permissions can perform actions on an Outposts
	// bucket. If you don't have s3-outposts:GetBucketPolicy permissions or you're not
	// using an identity that belongs to the bucket owner's account, Amazon S3 returns
	// a 403 Access Denied error. As a security precaution, the root user of the AWS
	// account that owns a bucket can always use this action, even if the policy
	// explicitly denies the root user the ability to perform this action. For more
	// information about bucket policies, see Using Bucket Policies and User Policies
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html). All
	// Amazon S3 on Outposts REST API requests for this action require an additional
	// parameter of x-amz-outpost-id to be passed with the request and an S3 on
	// Outposts endpoint hostname prefix instead of s3-control. For an example of the
	// request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint
	// hostname prefix and the x-amz-outpost-id derived using the access point ARN, see
	// the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html#API_control_GetBucketPolicy_Examples)
	// section. The following actions are related to GetBucketPolicy:
	//
	// * GetObject
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
	//
	// *
	// PutBucketPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketPolicy.html)
	//
	// *
	// DeleteBucketPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html)
	//
	GetBucketPolicy(arg1 context.Context, arg2 *s3control.GetBucketPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketPolicyOutput, error)
	// This action gets an Amazon S3 on Outposts bucket's tags. To get an S3 bucket
	// tags, see GetBucketTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html) in
	// the Amazon S3 API Reference. Returns the tag set associated with the Outposts
	// bucket. For more information, see Using Amazon S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the GetBucketTagging action. By default, the bucket owner has this permission
	// and can grant this permission to others. GetBucketTagging has the following
	// special error:
	//
	// * Error code: NoSuchTagSetError
	//
	// * Description: There is no tag
	// set associated with the bucket.
	//
	// All Amazon S3 on Outposts REST API requests for
	// this action require an additional parameter of x-amz-outpost-id to be passed
	// with the request and an S3 on Outposts endpoint hostname prefix instead of
	// s3-control. For an example of the request syntax for Amazon S3 on Outposts that
	// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
	// derived using the access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html#API_control_GetBucketTagging_Examples)
	// section. The following actions are related to GetBucketTagging:
	//
	// *
	// PutBucketTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html)
	//
	// *
	// DeleteBucketTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html)
	//
	GetBucketTagging(arg1 context.Context, arg2 *s3control.GetBucketTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketTaggingOutput, error)
	// Returns the tags on an S3 Batch Operations job. To use this operation, you must
	// have permission to perform the s3:GetJobTagging action. For more information,
	// see Controlling access and labeling jobs using tags
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags)
	// in the Amazon S3 User Guide. Related actions include:
	//
	// * CreateJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
	//
	// *
	// PutJobTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutJobTagging.html)
	//
	// *
	// DeleteJobTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteJobTagging.html)
	//
	GetJobTagging(arg1 context.Context, arg2 *s3control.GetJobTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.GetJobTaggingOutput, error)
	// Retrieves the PublicAccessBlock configuration for an AWS account. For more
	// information, see  Using Amazon S3 block public access
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).
	// Related actions include:
	//
	// * DeletePublicAccessBlock
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeletePublicAccessBlock.html)
	//
	// *
	// PutPublicAccessBlock
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutPublicAccessBlock.html)
	//
	GetPublicAccessBlock(arg1 context.Context, arg2 *s3control.GetPublicAccessBlockInput, arg3 ...func(*s3control.Options)) (*s3control.GetPublicAccessBlockOutput, error)
	// Gets the Amazon S3 Storage Lens configuration. For more information, see
	// Assessing your storage activity and usage with Amazon S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html) in the
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the s3:GetStorageLensConfiguration action. For more information, see Setting
	// permissions to use Amazon S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html)
	// in the Amazon S3 User Guide.
	//
	GetStorageLensConfiguration(arg1 context.Context, arg2 *s3control.GetStorageLensConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationOutput, error)
	// Gets the tags of Amazon S3 Storage Lens configuration. For more information
	// about S3 Storage Lens, see Assessing your storage activity and usage with Amazon
	// S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html) in the
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the s3:GetStorageLensConfigurationTagging action. For more information, see
	// Setting permissions to use Amazon S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html)
	// in the Amazon S3 User Guide.
	//
	GetStorageLensConfigurationTagging(arg1 context.Context, arg2 *s3control.GetStorageLensConfigurationTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationTaggingOutput, error)
	// Returns a list of the access points currently associated with the specified
	// bucket. You can retrieve up to 1000 access points per call. If the specified
	// bucket has more than 1,000 access points (or the number specified in maxResults,
	// whichever is less), the response will include a continuation token that you can
	// use to list the additional access points. All Amazon S3 on Outposts REST API
	// requests for this action require an additional parameter of x-amz-outpost-id to
	// be passed with the request and an S3 on Outposts endpoint hostname prefix
	// instead of s3-control. For an example of the request syntax for Amazon S3 on
	// Outposts that uses the S3 on Outposts endpoint hostname prefix and the
	// x-amz-outpost-id derived using the access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html#API_control_GetAccessPoint_Examples)
	// section. The following actions are related to ListAccessPoints:
	//
	// *
	// CreateAccessPoint
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html)
	//
	// *
	// DeleteAccessPoint
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html)
	//
	// *
	// GetAccessPoint
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html)
	//
	ListAccessPoints(arg1 context.Context, arg2 *s3control.ListAccessPointsInput, arg3 ...func(*s3control.Options)) (*s3control.ListAccessPointsOutput, error)
	// Returns a list of the access points associated with the Object Lambda Access
	// Point. You can retrieve up to 1000 access points per call. If there are more
	// than 1,000 access points (or the number specified in maxResults, whichever is
	// less), the response will include a continuation token that you can use to list
	// the additional access points. The following actions are related to
	// ListAccessPointsForObjectLambda:
	//
	// * CreateAccessPointForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPointForObjectLambda.html)
	//
	// *
	// DeleteAccessPointForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointForObjectLambda.html)
	//
	// *
	// GetAccessPointForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointForObjectLambda.html)
	//
	ListAccessPointsForObjectLambda(arg1 context.Context, arg2 *s3control.ListAccessPointsForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.ListAccessPointsForObjectLambdaOutput, error)
	// Lists current S3 Batch Operations jobs and jobs that have ended within the last
	// 30 days for the AWS account making the request. For more information, see S3
	// Batch Operations
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html) in the
	// Amazon S3 User Guide. Related actions include:
	//
	// * CreateJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
	//
	// *
	// DescribeJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html)
	//
	// *
	// UpdateJobPriority
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobPriority.html)
	//
	// *
	// UpdateJobStatus
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html)
	//
	ListJobs(arg1 context.Context, arg2 *s3control.ListJobsInput, arg3 ...func(*s3control.Options)) (*s3control.ListJobsOutput, error)
	// Returns a list of all Outposts buckets in an Outpost that are owned by the
	// authenticated sender of the request. For more information, see Using Amazon S3
	// on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
	// Amazon S3 User Guide. For an example of the request syntax for Amazon S3 on
	// Outposts that uses the S3 on Outposts endpoint hostname prefix and
	// x-amz-outpost-id in your request, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListRegionalBuckets.html#API_control_ListRegionalBuckets_Examples)
	// section.
	//
	ListRegionalBuckets(arg1 context.Context, arg2 *s3control.ListRegionalBucketsInput, arg3 ...func(*s3control.Options)) (*s3control.ListRegionalBucketsOutput, error)
	// Gets a list of Amazon S3 Storage Lens configurations. For more information about
	// S3 Storage Lens, see Assessing your storage activity and usage with Amazon S3
	// Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html) in the
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the s3:ListStorageLensConfigurations action. For more information, see Setting
	// permissions to use Amazon S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html)
	// in the Amazon S3 User Guide.
	//
	ListStorageLensConfigurations(arg1 context.Context, arg2 *s3control.ListStorageLensConfigurationsInput, arg3 ...func(*s3control.Options)) (*s3control.ListStorageLensConfigurationsOutput, error)
	// Replaces configuration for an Object Lambda Access Point. The following actions
	// are related to PutAccessPointConfigurationForObjectLambda:
	//
	// *
	// GetAccessPointConfigurationForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointConfigurationForObjectLambda.html)
	//
	PutAccessPointConfigurationForObjectLambda(arg1 context.Context, arg2 *s3control.PutAccessPointConfigurationForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.PutAccessPointConfigurationForObjectLambdaOutput, error)
	// Associates an access policy with the specified access point. Each access point
	// can have only one policy, so a request made to this API replaces any existing
	// policy associated with the specified access point. All Amazon S3 on Outposts
	// REST API requests for this action require an additional parameter of
	// x-amz-outpost-id to be passed with the request and an S3 on Outposts endpoint
	// hostname prefix instead of s3-control. For an example of the request syntax for
	// Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and
	// the x-amz-outpost-id derived using the access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html#API_control_PutAccessPointPolicy_Examples)
	// section. The following actions are related to PutAccessPointPolicy:
	//
	// *
	// GetAccessPointPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicy.html)
	//
	// *
	// DeleteAccessPointPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html)
	//
	PutAccessPointPolicy(arg1 context.Context, arg2 *s3control.PutAccessPointPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.PutAccessPointPolicyOutput, error)
	// Creates or replaces resource policy for an Object Lambda Access Point. For an
	// example policy, see Creating Object Lambda Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/olap-create.html#olap-create-cli)
	// in the Amazon S3 User Guide. The following actions are related to
	// PutAccessPointPolicyForObjectLambda:
	//
	// * DeleteAccessPointPolicyForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicyForObjectLambda.html)
	//
	// *
	// GetAccessPointPolicyForObjectLambda
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicyForObjectLambda.html)
	//
	PutAccessPointPolicyForObjectLambda(arg1 context.Context, arg2 *s3control.PutAccessPointPolicyForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.PutAccessPointPolicyForObjectLambdaOutput, error)
	// This action puts a lifecycle configuration to an Amazon S3 on Outposts bucket.
	// To put a lifecycle configuration to an S3 bucket, see
	// PutBucketLifecycleConfiguration
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html)
	// in the Amazon S3 API Reference. Creates a new lifecycle configuration for the S3
	// on Outposts bucket or replaces an existing lifecycle configuration. Outposts
	// buckets only support lifecycle configurations that delete/expire objects after a
	// certain period of time and abort incomplete multipart uploads. All Amazon S3 on
	// Outposts REST API requests for this action require an additional parameter of
	// x-amz-outpost-id to be passed with the request and an S3 on Outposts endpoint
	// hostname prefix instead of s3-control. For an example of the request syntax for
	// Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and
	// the x-amz-outpost-id derived using the access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html#API_control_PutBucketLifecycleConfiguration_Examples)
	// section. The following actions are related to
	// PutBucketLifecycleConfiguration:
	//
	// * GetBucketLifecycleConfiguration
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html)
	//
	// *
	// DeleteBucketLifecycleConfiguration
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html)
	//
	PutBucketLifecycleConfiguration(arg1 context.Context, arg2 *s3control.PutBucketLifecycleConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.PutBucketLifecycleConfigurationOutput, error)
	// This action puts a bucket policy to an Amazon S3 on Outposts bucket. To put a
	// policy on an S3 bucket, see PutBucketPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketPolicy.html) in
	// the Amazon S3 API Reference. Applies an Amazon S3 bucket policy to an Outposts
	// bucket. For more information, see Using Amazon S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
	// Amazon S3 User Guide. If you are using an identity other than the root user of
	// the AWS account that owns the Outposts bucket, the calling identity must have
	// the PutBucketPolicy permissions on the specified Outposts bucket and belong to
	// the bucket owner's account in order to use this action. If you don't have
	// PutBucketPolicy permissions, Amazon S3 returns a 403 Access Denied error. If you
	// have the correct permissions, but you're not using an identity that belongs to
	// the bucket owner's account, Amazon S3 returns a 405 Method Not Allowed error. As
	// a security precaution, the root user of the AWS account that owns a bucket can
	// always use this action, even if the policy explicitly denies the root user the
	// ability to perform this action. For more information about bucket policies, see
	// Using Bucket Policies and User Policies
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html). All
	// Amazon S3 on Outposts REST API requests for this action require an additional
	// parameter of x-amz-outpost-id to be passed with the request and an S3 on
	// Outposts endpoint hostname prefix instead of s3-control. For an example of the
	// request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint
	// hostname prefix and the x-amz-outpost-id derived using the access point ARN, see
	// the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketPolicy.html#API_control_PutBucketPolicy_Examples)
	// section. The following actions are related to PutBucketPolicy:
	//
	// *
	// GetBucketPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html)
	//
	// *
	// DeleteBucketPolicy
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html)
	//
	PutBucketPolicy(arg1 context.Context, arg2 *s3control.PutBucketPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.PutBucketPolicyOutput, error)
	// This action puts tags on an Amazon S3 on Outposts bucket. To put tags on an S3
	// bucket, see PutBucketTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html) in
	// the Amazon S3 API Reference. Sets the tags for an S3 on Outposts bucket. For
	// more information, see Using Amazon S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
	// Amazon S3 User Guide. Use tags to organize your AWS bill to reflect your own
	// cost structure. To do this, sign up to get your AWS account bill with tag key
	// values included. Then, to see the cost of combined resources, organize your
	// billing information according to resources with the same tag key values. For
	// example, you can tag several resources with a specific application name, and
	// then organize your billing information to see the total cost of that application
	// across several services. For more information, see Cost allocation and tagging
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html).
	// Within a bucket, if you add a tag that has the same key as an existing tag, the
	// new value overwrites the old value. For more information, see  Using cost
	// allocation in Amazon S3 bucket tags
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/CostAllocTagging.html).
	// To use this action, you must have permissions to perform the
	// s3-outposts:PutBucketTagging action. The Outposts bucket owner has this
	// permission by default and can grant this permission to others. For more
	// information about permissions, see  Permissions Related to Bucket Subresource
	// Operations
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
	// and Managing access permissions to your Amazon S3 resources
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html).
	// PutBucketTagging has the following special errors:
	//
	// * Error code:
	// InvalidTagError
	//
	// * Description: The tag provided was not a valid tag. This error
	// can occur if the tag did not pass input validation. For information about tag
	// restrictions, see  User-Defined Tag Restrictions
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html)
	// and  AWS-Generated Cost Allocation Tag Restrictions
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/aws-tag-restrictions.html).
	//
	// *
	// Error code: MalformedXMLError
	//
	// * Description: The XML provided does not match
	// the schema.
	//
	// * Error code: OperationAbortedError
	//
	// * Description: A conflicting
	// conditional action is currently in progress against this resource. Try again.
	//
	// *
	// Error code: InternalError
	//
	// * Description: The service was unable to apply the
	// provided tag to the bucket.
	//
	// All Amazon S3 on Outposts REST API requests for
	// this action require an additional parameter of x-amz-outpost-id to be passed
	// with the request and an S3 on Outposts endpoint hostname prefix instead of
	// s3-control. For an example of the request syntax for Amazon S3 on Outposts that
	// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
	// derived using the access point ARN, see the Examples
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html#API_control_PutBucketTagging_Examples)
	// section. The following actions are related to PutBucketTagging:
	//
	// *
	// GetBucketTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html)
	//
	// *
	// DeleteBucketTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html)
	//
	PutBucketTagging(arg1 context.Context, arg2 *s3control.PutBucketTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.PutBucketTaggingOutput, error)
	// Sets the supplied tag-set on an S3 Batch Operations job. A tag is a key-value
	// pair. You can associate S3 Batch Operations tags with any job by sending a PUT
	// request against the tagging subresource that is associated with the job. To
	// modify the existing tag set, you can either replace the existing tag set
	// entirely, or make changes within the existing tag set by retrieving the existing
	// tag set using GetJobTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetJobTagging.html),
	// modify that tag set, and use this action to replace the tag set with the one you
	// modified. For more information, see Controlling access and labeling jobs using
	// tags
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags)
	// in the Amazon S3 User Guide.
	//
	// * If you send this request with an empty tag set,
	// Amazon S3 deletes the existing tag set on the Batch Operations job. If you use
	// this method, you are charged for a Tier 1 Request (PUT). For more information,
	// see Amazon S3 pricing (http://aws.amazon.com/s3/pricing/).
	//
	// * For deleting
	// existing tags for your Batch Operations job, a DeleteJobTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteJobTagging.html)
	// request is preferred because it achieves the same result without incurring
	// charges.
	//
	// * A few things to consider about using tags:
	//
	// * Amazon S3 limits the
	// maximum number of tags to 50 tags per job.
	//
	// * You can associate up to 50 tags
	// with a job as long as they have unique tag keys.
	//
	// * A tag key can be up to 128
	// Unicode characters in length, and tag values can be up to 256 Unicode characters
	// in length.
	//
	// * The key and values are case sensitive.
	//
	// * For tagging-related
	// restrictions related to characters and encodings, see User-Defined Tag
	// Restrictions
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html)
	// in the AWS Billing and Cost Management User Guide.
	//
	// To use this action, you must
	// have permission to perform the s3:PutJobTagging action. Related actions
	// include:
	//
	// * CreatJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
	//
	// *
	// GetJobTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetJobTagging.html)
	//
	// *
	// DeleteJobTagging
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteJobTagging.html)
	//
	PutJobTagging(arg1 context.Context, arg2 *s3control.PutJobTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.PutJobTaggingOutput, error)
	// Creates or modifies the PublicAccessBlock configuration for an AWS account. For
	// more information, see  Using Amazon S3 block public access
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).
	// Related actions include:
	//
	// * GetPublicAccessBlock
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetPublicAccessBlock.html)
	//
	// *
	// DeletePublicAccessBlock
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeletePublicAccessBlock.html)
	//
	PutPublicAccessBlock(arg1 context.Context, arg2 *s3control.PutPublicAccessBlockInput, arg3 ...func(*s3control.Options)) (*s3control.PutPublicAccessBlockOutput, error)
	// Puts an Amazon S3 Storage Lens configuration. For more information about S3
	// Storage Lens, see Working with Amazon S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html) in the
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the s3:PutStorageLensConfiguration action. For more information, see Setting
	// permissions to use Amazon S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html)
	// in the Amazon S3 User Guide.
	//
	PutStorageLensConfiguration(arg1 context.Context, arg2 *s3control.PutStorageLensConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.PutStorageLensConfigurationOutput, error)
	// Put or replace tags on an existing Amazon S3 Storage Lens configuration. For
	// more information about S3 Storage Lens, see Assessing your storage activity and
	// usage with Amazon S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html) in the
	// Amazon S3 User Guide. To use this action, you must have permission to perform
	// the s3:PutStorageLensConfigurationTagging action. For more information, see
	// Setting permissions to use Amazon S3 Storage Lens
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html)
	// in the Amazon S3 User Guide.
	//
	PutStorageLensConfigurationTagging(arg1 context.Context, arg2 *s3control.PutStorageLensConfigurationTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.PutStorageLensConfigurationTaggingOutput, error)
	// Updates an existing S3 Batch Operations job's priority. For more information,
	// see S3 Batch Operations
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html) in the
	// Amazon S3 User Guide. Related actions include:
	//
	// * CreateJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
	//
	// *
	// ListJobs
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html)
	//
	// *
	// DescribeJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html)
	//
	// *
	// UpdateJobStatus
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html)
	//
	UpdateJobPriority(arg1 context.Context, arg2 *s3control.UpdateJobPriorityInput, arg3 ...func(*s3control.Options)) (*s3control.UpdateJobPriorityOutput, error)
	// Updates the status for the specified job. Use this action to confirm that you
	// want to run a job or to cancel an existing job. For more information, see S3
	// Batch Operations
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html) in the
	// Amazon S3 User Guide. Related actions include:
	//
	// * CreateJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
	//
	// *
	// ListJobs
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html)
	//
	// *
	// DescribeJob
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html)
	//
	// *
	// UpdateJobStatus
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html)
	//
	UpdateJobStatus(arg1 context.Context, arg2 *s3control.UpdateJobStatusInput, arg3 ...func(*s3control.Options)) (*s3control.UpdateJobStatusOutput, error)
}

// S3ControlClientMock generic client mock
type S3ControlClientMock struct {
	CreateAccessPointMock                          func(arg1 context.Context, arg2 *s3control.CreateAccessPointInput, arg3 ...func(*s3control.Options)) (*s3control.CreateAccessPointOutput, error)
	CreateAccessPointForObjectLambdaMock           func(arg1 context.Context, arg2 *s3control.CreateAccessPointForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.CreateAccessPointForObjectLambdaOutput, error)
	CreateBucketMock                               func(arg1 context.Context, arg2 *s3control.CreateBucketInput, arg3 ...func(*s3control.Options)) (*s3control.CreateBucketOutput, error)
	CreateJobMock                                  func(arg1 context.Context, arg2 *s3control.CreateJobInput, arg3 ...func(*s3control.Options)) (*s3control.CreateJobOutput, error)
	DeleteAccessPointMock                          func(arg1 context.Context, arg2 *s3control.DeleteAccessPointInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointOutput, error)
	DeleteAccessPointForObjectLambdaMock           func(arg1 context.Context, arg2 *s3control.DeleteAccessPointForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointForObjectLambdaOutput, error)
	DeleteAccessPointPolicyMock                    func(arg1 context.Context, arg2 *s3control.DeleteAccessPointPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointPolicyOutput, error)
	DeleteAccessPointPolicyForObjectLambdaMock     func(arg1 context.Context, arg2 *s3control.DeleteAccessPointPolicyForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointPolicyForObjectLambdaOutput, error)
	DeleteBucketMock                               func(arg1 context.Context, arg2 *s3control.DeleteBucketInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketOutput, error)
	DeleteBucketLifecycleConfigurationMock         func(arg1 context.Context, arg2 *s3control.DeleteBucketLifecycleConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketLifecycleConfigurationOutput, error)
	DeleteBucketPolicyMock                         func(arg1 context.Context, arg2 *s3control.DeleteBucketPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketPolicyOutput, error)
	DeleteBucketTaggingMock                        func(arg1 context.Context, arg2 *s3control.DeleteBucketTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketTaggingOutput, error)
	DeleteJobTaggingMock                           func(arg1 context.Context, arg2 *s3control.DeleteJobTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteJobTaggingOutput, error)
	DeletePublicAccessBlockMock                    func(arg1 context.Context, arg2 *s3control.DeletePublicAccessBlockInput, arg3 ...func(*s3control.Options)) (*s3control.DeletePublicAccessBlockOutput, error)
	DeleteStorageLensConfigurationMock             func(arg1 context.Context, arg2 *s3control.DeleteStorageLensConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteStorageLensConfigurationOutput, error)
	DeleteStorageLensConfigurationTaggingMock      func(arg1 context.Context, arg2 *s3control.DeleteStorageLensConfigurationTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteStorageLensConfigurationTaggingOutput, error)
	DescribeJobMock                                func(arg1 context.Context, arg2 *s3control.DescribeJobInput, arg3 ...func(*s3control.Options)) (*s3control.DescribeJobOutput, error)
	GetAccessPointMock                             func(arg1 context.Context, arg2 *s3control.GetAccessPointInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointOutput, error)
	GetAccessPointConfigurationForObjectLambdaMock func(arg1 context.Context, arg2 *s3control.GetAccessPointConfigurationForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error)
	GetAccessPointForObjectLambdaMock              func(arg1 context.Context, arg2 *s3control.GetAccessPointForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointForObjectLambdaOutput, error)
	GetAccessPointPolicyMock                       func(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyOutput, error)
	GetAccessPointPolicyForObjectLambdaMock        func(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error)
	GetAccessPointPolicyStatusMock                 func(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyStatusInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusOutput, error)
	GetAccessPointPolicyStatusForObjectLambdaMock  func(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyStatusForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error)
	GetBucketMock                                  func(arg1 context.Context, arg2 *s3control.GetBucketInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketOutput, error)
	GetBucketLifecycleConfigurationMock            func(arg1 context.Context, arg2 *s3control.GetBucketLifecycleConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketLifecycleConfigurationOutput, error)
	GetBucketPolicyMock                            func(arg1 context.Context, arg2 *s3control.GetBucketPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketPolicyOutput, error)
	GetBucketTaggingMock                           func(arg1 context.Context, arg2 *s3control.GetBucketTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketTaggingOutput, error)
	GetJobTaggingMock                              func(arg1 context.Context, arg2 *s3control.GetJobTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.GetJobTaggingOutput, error)
	GetPublicAccessBlockMock                       func(arg1 context.Context, arg2 *s3control.GetPublicAccessBlockInput, arg3 ...func(*s3control.Options)) (*s3control.GetPublicAccessBlockOutput, error)
	GetStorageLensConfigurationMock                func(arg1 context.Context, arg2 *s3control.GetStorageLensConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationOutput, error)
	GetStorageLensConfigurationTaggingMock         func(arg1 context.Context, arg2 *s3control.GetStorageLensConfigurationTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationTaggingOutput, error)
	ListAccessPointsMock                           func(arg1 context.Context, arg2 *s3control.ListAccessPointsInput, arg3 ...func(*s3control.Options)) (*s3control.ListAccessPointsOutput, error)
	ListAccessPointsForObjectLambdaMock            func(arg1 context.Context, arg2 *s3control.ListAccessPointsForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.ListAccessPointsForObjectLambdaOutput, error)
	ListJobsMock                                   func(arg1 context.Context, arg2 *s3control.ListJobsInput, arg3 ...func(*s3control.Options)) (*s3control.ListJobsOutput, error)
	ListRegionalBucketsMock                        func(arg1 context.Context, arg2 *s3control.ListRegionalBucketsInput, arg3 ...func(*s3control.Options)) (*s3control.ListRegionalBucketsOutput, error)
	ListStorageLensConfigurationsMock              func(arg1 context.Context, arg2 *s3control.ListStorageLensConfigurationsInput, arg3 ...func(*s3control.Options)) (*s3control.ListStorageLensConfigurationsOutput, error)
	PutAccessPointConfigurationForObjectLambdaMock func(arg1 context.Context, arg2 *s3control.PutAccessPointConfigurationForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.PutAccessPointConfigurationForObjectLambdaOutput, error)
	PutAccessPointPolicyMock                       func(arg1 context.Context, arg2 *s3control.PutAccessPointPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.PutAccessPointPolicyOutput, error)
	PutAccessPointPolicyForObjectLambdaMock        func(arg1 context.Context, arg2 *s3control.PutAccessPointPolicyForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.PutAccessPointPolicyForObjectLambdaOutput, error)
	PutBucketLifecycleConfigurationMock            func(arg1 context.Context, arg2 *s3control.PutBucketLifecycleConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.PutBucketLifecycleConfigurationOutput, error)
	PutBucketPolicyMock                            func(arg1 context.Context, arg2 *s3control.PutBucketPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.PutBucketPolicyOutput, error)
	PutBucketTaggingMock                           func(arg1 context.Context, arg2 *s3control.PutBucketTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.PutBucketTaggingOutput, error)
	PutJobTaggingMock                              func(arg1 context.Context, arg2 *s3control.PutJobTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.PutJobTaggingOutput, error)
	PutPublicAccessBlockMock                       func(arg1 context.Context, arg2 *s3control.PutPublicAccessBlockInput, arg3 ...func(*s3control.Options)) (*s3control.PutPublicAccessBlockOutput, error)
	PutStorageLensConfigurationMock                func(arg1 context.Context, arg2 *s3control.PutStorageLensConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.PutStorageLensConfigurationOutput, error)
	PutStorageLensConfigurationTaggingMock         func(arg1 context.Context, arg2 *s3control.PutStorageLensConfigurationTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.PutStorageLensConfigurationTaggingOutput, error)
	UpdateJobPriorityMock                          func(arg1 context.Context, arg2 *s3control.UpdateJobPriorityInput, arg3 ...func(*s3control.Options)) (*s3control.UpdateJobPriorityOutput, error)
	UpdateJobStatusMock                            func(arg1 context.Context, arg2 *s3control.UpdateJobStatusInput, arg3 ...func(*s3control.Options)) (*s3control.UpdateJobStatusOutput, error)
}

// CreateAccessPoint mocked function
func (m S3ControlClientMock) CreateAccessPoint(arg1 context.Context, arg2 *s3control.CreateAccessPointInput, arg3 ...func(*s3control.Options)) (*s3control.CreateAccessPointOutput, error) {
	return m.CreateAccessPointMock(arg1, arg2, arg3...)
}

// CreateAccessPointForObjectLambda mocked function
func (m S3ControlClientMock) CreateAccessPointForObjectLambda(arg1 context.Context, arg2 *s3control.CreateAccessPointForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.CreateAccessPointForObjectLambdaOutput, error) {
	return m.CreateAccessPointForObjectLambdaMock(arg1, arg2, arg3...)
}

// CreateBucket mocked function
func (m S3ControlClientMock) CreateBucket(arg1 context.Context, arg2 *s3control.CreateBucketInput, arg3 ...func(*s3control.Options)) (*s3control.CreateBucketOutput, error) {
	return m.CreateBucketMock(arg1, arg2, arg3...)
}

// CreateJob mocked function
func (m S3ControlClientMock) CreateJob(arg1 context.Context, arg2 *s3control.CreateJobInput, arg3 ...func(*s3control.Options)) (*s3control.CreateJobOutput, error) {
	return m.CreateJobMock(arg1, arg2, arg3...)
}

// DeleteAccessPoint mocked function
func (m S3ControlClientMock) DeleteAccessPoint(arg1 context.Context, arg2 *s3control.DeleteAccessPointInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointOutput, error) {
	return m.DeleteAccessPointMock(arg1, arg2, arg3...)
}

// DeleteAccessPointForObjectLambda mocked function
func (m S3ControlClientMock) DeleteAccessPointForObjectLambda(arg1 context.Context, arg2 *s3control.DeleteAccessPointForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointForObjectLambdaOutput, error) {
	return m.DeleteAccessPointForObjectLambdaMock(arg1, arg2, arg3...)
}

// DeleteAccessPointPolicy mocked function
func (m S3ControlClientMock) DeleteAccessPointPolicy(arg1 context.Context, arg2 *s3control.DeleteAccessPointPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointPolicyOutput, error) {
	return m.DeleteAccessPointPolicyMock(arg1, arg2, arg3...)
}

// DeleteAccessPointPolicyForObjectLambda mocked function
func (m S3ControlClientMock) DeleteAccessPointPolicyForObjectLambda(arg1 context.Context, arg2 *s3control.DeleteAccessPointPolicyForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteAccessPointPolicyForObjectLambdaOutput, error) {
	return m.DeleteAccessPointPolicyForObjectLambdaMock(arg1, arg2, arg3...)
}

// DeleteBucket mocked function
func (m S3ControlClientMock) DeleteBucket(arg1 context.Context, arg2 *s3control.DeleteBucketInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketOutput, error) {
	return m.DeleteBucketMock(arg1, arg2, arg3...)
}

// DeleteBucketLifecycleConfiguration mocked function
func (m S3ControlClientMock) DeleteBucketLifecycleConfiguration(arg1 context.Context, arg2 *s3control.DeleteBucketLifecycleConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketLifecycleConfigurationOutput, error) {
	return m.DeleteBucketLifecycleConfigurationMock(arg1, arg2, arg3...)
}

// DeleteBucketPolicy mocked function
func (m S3ControlClientMock) DeleteBucketPolicy(arg1 context.Context, arg2 *s3control.DeleteBucketPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketPolicyOutput, error) {
	return m.DeleteBucketPolicyMock(arg1, arg2, arg3...)
}

// DeleteBucketTagging mocked function
func (m S3ControlClientMock) DeleteBucketTagging(arg1 context.Context, arg2 *s3control.DeleteBucketTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteBucketTaggingOutput, error) {
	return m.DeleteBucketTaggingMock(arg1, arg2, arg3...)
}

// DeleteJobTagging mocked function
func (m S3ControlClientMock) DeleteJobTagging(arg1 context.Context, arg2 *s3control.DeleteJobTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteJobTaggingOutput, error) {
	return m.DeleteJobTaggingMock(arg1, arg2, arg3...)
}

// DeletePublicAccessBlock mocked function
func (m S3ControlClientMock) DeletePublicAccessBlock(arg1 context.Context, arg2 *s3control.DeletePublicAccessBlockInput, arg3 ...func(*s3control.Options)) (*s3control.DeletePublicAccessBlockOutput, error) {
	return m.DeletePublicAccessBlockMock(arg1, arg2, arg3...)
}

// DeleteStorageLensConfiguration mocked function
func (m S3ControlClientMock) DeleteStorageLensConfiguration(arg1 context.Context, arg2 *s3control.DeleteStorageLensConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteStorageLensConfigurationOutput, error) {
	return m.DeleteStorageLensConfigurationMock(arg1, arg2, arg3...)
}

// DeleteStorageLensConfigurationTagging mocked function
func (m S3ControlClientMock) DeleteStorageLensConfigurationTagging(arg1 context.Context, arg2 *s3control.DeleteStorageLensConfigurationTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.DeleteStorageLensConfigurationTaggingOutput, error) {
	return m.DeleteStorageLensConfigurationTaggingMock(arg1, arg2, arg3...)
}

// DescribeJob mocked function
func (m S3ControlClientMock) DescribeJob(arg1 context.Context, arg2 *s3control.DescribeJobInput, arg3 ...func(*s3control.Options)) (*s3control.DescribeJobOutput, error) {
	return m.DescribeJobMock(arg1, arg2, arg3...)
}

// GetAccessPoint mocked function
func (m S3ControlClientMock) GetAccessPoint(arg1 context.Context, arg2 *s3control.GetAccessPointInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointOutput, error) {
	return m.GetAccessPointMock(arg1, arg2, arg3...)
}

// GetAccessPointConfigurationForObjectLambda mocked function
func (m S3ControlClientMock) GetAccessPointConfigurationForObjectLambda(arg1 context.Context, arg2 *s3control.GetAccessPointConfigurationForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error) {
	return m.GetAccessPointConfigurationForObjectLambdaMock(arg1, arg2, arg3...)
}

// GetAccessPointForObjectLambda mocked function
func (m S3ControlClientMock) GetAccessPointForObjectLambda(arg1 context.Context, arg2 *s3control.GetAccessPointForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointForObjectLambdaOutput, error) {
	return m.GetAccessPointForObjectLambdaMock(arg1, arg2, arg3...)
}

// GetAccessPointPolicy mocked function
func (m S3ControlClientMock) GetAccessPointPolicy(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyOutput, error) {
	return m.GetAccessPointPolicyMock(arg1, arg2, arg3...)
}

// GetAccessPointPolicyForObjectLambda mocked function
func (m S3ControlClientMock) GetAccessPointPolicyForObjectLambda(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error) {
	return m.GetAccessPointPolicyForObjectLambdaMock(arg1, arg2, arg3...)
}

// GetAccessPointPolicyStatus mocked function
func (m S3ControlClientMock) GetAccessPointPolicyStatus(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyStatusInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusOutput, error) {
	return m.GetAccessPointPolicyStatusMock(arg1, arg2, arg3...)
}

// GetAccessPointPolicyStatusForObjectLambda mocked function
func (m S3ControlClientMock) GetAccessPointPolicyStatusForObjectLambda(arg1 context.Context, arg2 *s3control.GetAccessPointPolicyStatusForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error) {
	return m.GetAccessPointPolicyStatusForObjectLambdaMock(arg1, arg2, arg3...)
}

// GetBucket mocked function
func (m S3ControlClientMock) GetBucket(arg1 context.Context, arg2 *s3control.GetBucketInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketOutput, error) {
	return m.GetBucketMock(arg1, arg2, arg3...)
}

// GetBucketLifecycleConfiguration mocked function
func (m S3ControlClientMock) GetBucketLifecycleConfiguration(arg1 context.Context, arg2 *s3control.GetBucketLifecycleConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketLifecycleConfigurationOutput, error) {
	return m.GetBucketLifecycleConfigurationMock(arg1, arg2, arg3...)
}

// GetBucketPolicy mocked function
func (m S3ControlClientMock) GetBucketPolicy(arg1 context.Context, arg2 *s3control.GetBucketPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketPolicyOutput, error) {
	return m.GetBucketPolicyMock(arg1, arg2, arg3...)
}

// GetBucketTagging mocked function
func (m S3ControlClientMock) GetBucketTagging(arg1 context.Context, arg2 *s3control.GetBucketTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.GetBucketTaggingOutput, error) {
	return m.GetBucketTaggingMock(arg1, arg2, arg3...)
}

// GetJobTagging mocked function
func (m S3ControlClientMock) GetJobTagging(arg1 context.Context, arg2 *s3control.GetJobTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.GetJobTaggingOutput, error) {
	return m.GetJobTaggingMock(arg1, arg2, arg3...)
}

// GetPublicAccessBlock mocked function
func (m S3ControlClientMock) GetPublicAccessBlock(arg1 context.Context, arg2 *s3control.GetPublicAccessBlockInput, arg3 ...func(*s3control.Options)) (*s3control.GetPublicAccessBlockOutput, error) {
	return m.GetPublicAccessBlockMock(arg1, arg2, arg3...)
}

// GetStorageLensConfiguration mocked function
func (m S3ControlClientMock) GetStorageLensConfiguration(arg1 context.Context, arg2 *s3control.GetStorageLensConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationOutput, error) {
	return m.GetStorageLensConfigurationMock(arg1, arg2, arg3...)
}

// GetStorageLensConfigurationTagging mocked function
func (m S3ControlClientMock) GetStorageLensConfigurationTagging(arg1 context.Context, arg2 *s3control.GetStorageLensConfigurationTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.GetStorageLensConfigurationTaggingOutput, error) {
	return m.GetStorageLensConfigurationTaggingMock(arg1, arg2, arg3...)
}

// ListAccessPoints mocked function
func (m S3ControlClientMock) ListAccessPoints(arg1 context.Context, arg2 *s3control.ListAccessPointsInput, arg3 ...func(*s3control.Options)) (*s3control.ListAccessPointsOutput, error) {
	return m.ListAccessPointsMock(arg1, arg2, arg3...)
}

// ListAccessPointsForObjectLambda mocked function
func (m S3ControlClientMock) ListAccessPointsForObjectLambda(arg1 context.Context, arg2 *s3control.ListAccessPointsForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.ListAccessPointsForObjectLambdaOutput, error) {
	return m.ListAccessPointsForObjectLambdaMock(arg1, arg2, arg3...)
}

// ListJobs mocked function
func (m S3ControlClientMock) ListJobs(arg1 context.Context, arg2 *s3control.ListJobsInput, arg3 ...func(*s3control.Options)) (*s3control.ListJobsOutput, error) {
	return m.ListJobsMock(arg1, arg2, arg3...)
}

// ListRegionalBuckets mocked function
func (m S3ControlClientMock) ListRegionalBuckets(arg1 context.Context, arg2 *s3control.ListRegionalBucketsInput, arg3 ...func(*s3control.Options)) (*s3control.ListRegionalBucketsOutput, error) {
	return m.ListRegionalBucketsMock(arg1, arg2, arg3...)
}

// ListStorageLensConfigurations mocked function
func (m S3ControlClientMock) ListStorageLensConfigurations(arg1 context.Context, arg2 *s3control.ListStorageLensConfigurationsInput, arg3 ...func(*s3control.Options)) (*s3control.ListStorageLensConfigurationsOutput, error) {
	return m.ListStorageLensConfigurationsMock(arg1, arg2, arg3...)
}

// PutAccessPointConfigurationForObjectLambda mocked function
func (m S3ControlClientMock) PutAccessPointConfigurationForObjectLambda(arg1 context.Context, arg2 *s3control.PutAccessPointConfigurationForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.PutAccessPointConfigurationForObjectLambdaOutput, error) {
	return m.PutAccessPointConfigurationForObjectLambdaMock(arg1, arg2, arg3...)
}

// PutAccessPointPolicy mocked function
func (m S3ControlClientMock) PutAccessPointPolicy(arg1 context.Context, arg2 *s3control.PutAccessPointPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.PutAccessPointPolicyOutput, error) {
	return m.PutAccessPointPolicyMock(arg1, arg2, arg3...)
}

// PutAccessPointPolicyForObjectLambda mocked function
func (m S3ControlClientMock) PutAccessPointPolicyForObjectLambda(arg1 context.Context, arg2 *s3control.PutAccessPointPolicyForObjectLambdaInput, arg3 ...func(*s3control.Options)) (*s3control.PutAccessPointPolicyForObjectLambdaOutput, error) {
	return m.PutAccessPointPolicyForObjectLambdaMock(arg1, arg2, arg3...)
}

// PutBucketLifecycleConfiguration mocked function
func (m S3ControlClientMock) PutBucketLifecycleConfiguration(arg1 context.Context, arg2 *s3control.PutBucketLifecycleConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.PutBucketLifecycleConfigurationOutput, error) {
	return m.PutBucketLifecycleConfigurationMock(arg1, arg2, arg3...)
}

// PutBucketPolicy mocked function
func (m S3ControlClientMock) PutBucketPolicy(arg1 context.Context, arg2 *s3control.PutBucketPolicyInput, arg3 ...func(*s3control.Options)) (*s3control.PutBucketPolicyOutput, error) {
	return m.PutBucketPolicyMock(arg1, arg2, arg3...)
}

// PutBucketTagging mocked function
func (m S3ControlClientMock) PutBucketTagging(arg1 context.Context, arg2 *s3control.PutBucketTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.PutBucketTaggingOutput, error) {
	return m.PutBucketTaggingMock(arg1, arg2, arg3...)
}

// PutJobTagging mocked function
func (m S3ControlClientMock) PutJobTagging(arg1 context.Context, arg2 *s3control.PutJobTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.PutJobTaggingOutput, error) {
	return m.PutJobTaggingMock(arg1, arg2, arg3...)
}

// PutPublicAccessBlock mocked function
func (m S3ControlClientMock) PutPublicAccessBlock(arg1 context.Context, arg2 *s3control.PutPublicAccessBlockInput, arg3 ...func(*s3control.Options)) (*s3control.PutPublicAccessBlockOutput, error) {
	return m.PutPublicAccessBlockMock(arg1, arg2, arg3...)
}

// PutStorageLensConfiguration mocked function
func (m S3ControlClientMock) PutStorageLensConfiguration(arg1 context.Context, arg2 *s3control.PutStorageLensConfigurationInput, arg3 ...func(*s3control.Options)) (*s3control.PutStorageLensConfigurationOutput, error) {
	return m.PutStorageLensConfigurationMock(arg1, arg2, arg3...)
}

// PutStorageLensConfigurationTagging mocked function
func (m S3ControlClientMock) PutStorageLensConfigurationTagging(arg1 context.Context, arg2 *s3control.PutStorageLensConfigurationTaggingInput, arg3 ...func(*s3control.Options)) (*s3control.PutStorageLensConfigurationTaggingOutput, error) {
	return m.PutStorageLensConfigurationTaggingMock(arg1, arg2, arg3...)
}

// UpdateJobPriority mocked function
func (m S3ControlClientMock) UpdateJobPriority(arg1 context.Context, arg2 *s3control.UpdateJobPriorityInput, arg3 ...func(*s3control.Options)) (*s3control.UpdateJobPriorityOutput, error) {
	return m.UpdateJobPriorityMock(arg1, arg2, arg3...)
}

// UpdateJobStatus mocked function
func (m S3ControlClientMock) UpdateJobStatus(arg1 context.Context, arg2 *s3control.UpdateJobStatusInput, arg3 ...func(*s3control.Options)) (*s3control.UpdateJobStatusOutput, error) {
	return m.UpdateJobStatusMock(arg1, arg2, arg3...)
}
