package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/felipemarinho97/go-mock-gen
 */

// IDynamoClient generic client
type IDynamoClient interface {
	// This operation allows you to perform batch reads and writes on data stored in
	// DynamoDB, using PartiQL.
	//
	BatchExecuteStatement(arg1 context.Context, arg2 *dynamodb.BatchExecuteStatementInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.BatchExecuteStatementOutput, error)
	// The BatchGetItem operation returns the attributes of one or more items from one
	// or more tables. You identify requested items by primary key. A single operation
	// can retrieve up to 16 MB of data, which can contain as many as 100 items.
	// BatchGetItem returns a partial result if the response size limit is exceeded,
	// the table's provisioned throughput is exceeded, or an internal processing
	// failure occurs. If a partial result is returned, the operation returns a value
	// for UnprocessedKeys. You can use this value to retry the operation starting with
	// the next item to get. If you request more than 100 items, BatchGetItem returns a
	// ValidationException with the message "Too many items requested for the
	// BatchGetItem call." For example, if you ask to retrieve 100 items, but each
	// individual item is 300 KB in size, the system returns 52 items (so as not to
	// exceed the 16 MB limit). It also returns an appropriate UnprocessedKeys value so
	// you can get the next page of results. If desired, your application can include
	// its own logic to assemble the pages of results into one dataset. If none of the
	// items can be processed due to insufficient provisioned throughput on all of the
	// tables in the request, then BatchGetItem returns a
	// ProvisionedThroughputExceededException. If at least one of the items is
	// successfully processed, then BatchGetItem completes successfully, while
	// returning the keys of the unread items in UnprocessedKeys. If DynamoDB returns
	// any unprocessed items, you should retry the batch operation on those items.
	// However, we strongly recommend that you use an exponential backoff algorithm. If
	// you retry the batch operation immediately, the underlying read or write requests
	// can still fail due to throttling on the individual tables. If you delay the
	// batch operation using exponential backoff, the individual requests in the batch
	// are much more likely to succeed. For more information, see Batch Operations and
	// Error Handling
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html#BatchOperations)
	// in the Amazon DynamoDB Developer Guide. By default, BatchGetItem performs
	// eventually consistent reads on every table in the request. If you want strongly
	// consistent reads instead, you can set ConsistentRead to true for any or all
	// tables. In order to minimize response latency, BatchGetItem retrieves items in
	// parallel. When designing your application, keep in mind that DynamoDB does not
	// return items in any particular order. To help parse the response by item,
	// include the primary key values for the items in your request in the
	// ProjectionExpression parameter. If a requested item does not exist, it is not
	// returned in the result. Requests for nonexistent items consume the minimum read
	// capacity units according to the type of read. For more information, see Working
	// with Tables
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#CapacityUnitCalculations)
	// in the Amazon DynamoDB Developer Guide.
	//
	BatchGetItem(arg1 context.Context, arg2 *dynamodb.BatchGetItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error)
	// The BatchWriteItem operation puts or deletes multiple items in one or more
	// tables. A single call to BatchWriteItem can write up to 16 MB of data, which can
	// comprise as many as 25 put or delete requests. Individual items to be written
	// can be as large as 400 KB. BatchWriteItem cannot update items. To update items,
	// use the UpdateItem action. The individual PutItem and DeleteItem operations
	// specified in BatchWriteItem are atomic; however BatchWriteItem as a whole is
	// not. If any requested operations fail because the table's provisioned throughput
	// is exceeded or an internal processing failure occurs, the failed operations are
	// returned in the UnprocessedItems response parameter. You can investigate and
	// optionally resend the requests. Typically, you would call BatchWriteItem in a
	// loop. Each iteration would check for unprocessed items and submit a new
	// BatchWriteItem request with those unprocessed items until all items have been
	// processed. If none of the items can be processed due to insufficient provisioned
	// throughput on all of the tables in the request, then BatchWriteItem returns a
	// ProvisionedThroughputExceededException. If DynamoDB returns any unprocessed
	// items, you should retry the batch operation on those items. However, we strongly
	// recommend that you use an exponential backoff algorithm. If you retry the batch
	// operation immediately, the underlying read or write requests can still fail due
	// to throttling on the individual tables. If you delay the batch operation using
	// exponential backoff, the individual requests in the batch are much more likely
	// to succeed. For more information, see Batch Operations and Error Handling
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html#Programming.Errors.BatchOperations)
	// in the Amazon DynamoDB Developer Guide. With BatchWriteItem, you can efficiently
	// write or delete large amounts of data, such as from Amazon EMR, or copy data
	// from another database into DynamoDB. In order to improve performance with these
	// large-scale operations, BatchWriteItem does not behave in the same way as
	// individual PutItem and DeleteItem calls would. For example, you cannot specify
	// conditions on individual put and delete requests, and BatchWriteItem does not
	// return deleted items in the response. If you use a programming language that
	// supports concurrency, you can use threads to write items in parallel. Your
	// application must include the necessary logic to manage the threads. With
	// languages that don't support threading, you must update or delete the specified
	// items one at a time. In both situations, BatchWriteItem performs the specified
	// put and delete operations in parallel, giving you the power of the thread pool
	// approach without having to introduce complexity into your application. Parallel
	// processing reduces latency, but each specified put and delete request consumes
	// the same number of write capacity units whether it is processed in parallel or
	// not. Delete operations on nonexistent items consume one write capacity unit. If
	// one or more of the following is true, DynamoDB rejects the entire batch write
	// operation:
	//
	// * One or more tables specified in the BatchWriteItem request does
	// not exist.
	//
	// * Primary key attributes specified on an item in the request do not
	// match those in the corresponding table's primary key schema.
	//
	// * You try to
	// perform multiple operations on the same item in the same BatchWriteItem request.
	// For example, you cannot put and delete the same item in the same BatchWriteItem
	// request.
	//
	// * Your request contains at least two items with identical hash and
	// range keys (which essentially is two put operations).
	//
	// * There are more than 25
	// requests in the batch.
	//
	// * Any individual item in a batch exceeds 400 KB.
	//
	// * The
	// total request size exceeds 16 MB.
	//
	BatchWriteItem(arg1 context.Context, arg2 *dynamodb.BatchWriteItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error)
	// Creates a backup for an existing table. Each time you create an on-demand
	// backup, the entire table data is backed up. There is no limit to the number of
	// on-demand backups that can be taken. When you create an on-demand backup, a time
	// marker of the request is cataloged, and the backup is created asynchronously, by
	// applying all changes until the time of the request to the last full table
	// snapshot. Backup requests are processed instantaneously and become available for
	// restore within minutes. You can call CreateBackup at a maximum rate of 50 times
	// per second. All backups in DynamoDB work without consuming any provisioned
	// throughput on the table. If you submit a backup request on 2018-12-14 at
	// 14:25:00, the backup is guaranteed to contain all data committed to the table up
	// to 14:24:00, and data committed after 14:26:00 will not be. The backup might
	// contain data modifications made between 14:24:00 and 14:26:00. On-demand backup
	// does not support causal consistency. Along with data, the following are also
	// included on the backups:
	//
	// * Global secondary indexes (GSIs)
	//
	// * Local secondary
	// indexes (LSIs)
	//
	// * Streams
	//
	// * Provisioned read and write capacity
	//
	CreateBackup(arg1 context.Context, arg2 *dynamodb.CreateBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.CreateBackupOutput, error)
	// Creates a global table from an existing table. A global table creates a
	// replication relationship between two or more DynamoDB tables with the same table
	// name in the provided Regions. This operation only applies to Version 2017.11.29
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html)
	// of global tables. If you want to add a new replica table to a global table, each
	// of the following conditions must be true:
	//
	// * The table must have the same
	// primary key as all of the other replicas.
	//
	// * The table must have the same name
	// as all of the other replicas.
	//
	// * The table must have DynamoDB Streams enabled,
	// with the stream containing both the new and the old images of the item.
	//
	// * None
	// of the replica tables in the global table can contain any data.
	//
	// If global
	// secondary indexes are specified, then the following conditions must also be
	// met:
	//
	// * The global secondary indexes must have the same name.
	//
	// * The global
	// secondary indexes must have the same hash key and sort key (if present).
	//
	// If
	// local secondary indexes are specified, then the following conditions must also
	// be met:
	//
	// * The local secondary indexes must have the same name.
	//
	// * The local
	// secondary indexes must have the same hash key and sort key (if present).
	//
	// Write
	// capacity settings should be set consistently across your replica tables and
	// secondary indexes. DynamoDB strongly recommends enabling auto scaling to manage
	// the write capacity settings for all of your global tables replicas and indexes.
	// If you prefer to manage write capacity settings manually, you should provision
	// equal replicated write capacity units to your replica tables. You should also
	// provision equal replicated write capacity units to matching secondary indexes
	// across your global table.
	//
	CreateGlobalTable(arg1 context.Context, arg2 *dynamodb.CreateGlobalTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.CreateGlobalTableOutput, error)
	// The CreateTable operation adds a new table to your account. In an AWS account,
	// table names must be unique within each Region. That is, you can have two tables
	// with same name if you create the tables in different Regions. CreateTable is an
	// asynchronous operation. Upon receiving a CreateTable request, DynamoDB
	// immediately returns a response with a TableStatus of CREATING. After the table
	// is created, DynamoDB sets the TableStatus to ACTIVE. You can perform read and
	// write operations only on an ACTIVE table. You can optionally define secondary
	// indexes on the new table, as part of the CreateTable operation. If you want to
	// create multiple tables with secondary indexes on them, you must create the
	// tables sequentially. Only one table with secondary indexes can be in the
	// CREATING state at any given time. You can use the DescribeTable action to check
	// the table status.
	//
	CreateTable(arg1 context.Context, arg2 *dynamodb.CreateTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)
	// Deletes an existing backup of a table. You can call DeleteBackup at a maximum
	// rate of 10 times per second.
	//
	DeleteBackup(arg1 context.Context, arg2 *dynamodb.DeleteBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DeleteBackupOutput, error)
	// Deletes a single item in a table by primary key. You can perform a conditional
	// delete operation that deletes the item if it exists, or if it has an expected
	// attribute value. In addition to deleting an item, you can also return the item's
	// attribute values in the same operation, using the ReturnValues parameter. Unless
	// you specify conditions, the DeleteItem is an idempotent operation; running it
	// multiple times on the same item or attribute does not result in an error
	// response. Conditional deletes are useful for deleting items only if specific
	// conditions are met. If those conditions are met, DynamoDB performs the delete.
	// Otherwise, the item is not deleted.
	//
	DeleteItem(arg1 context.Context, arg2 *dynamodb.DeleteItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
	// The DeleteTable operation deletes a table and all of its items. After a
	// DeleteTable request, the specified table is in the DELETING state until DynamoDB
	// completes the deletion. If the table is in the ACTIVE state, you can delete it.
	// If a table is in CREATING or UPDATING states, then DynamoDB returns a
	// ResourceInUseException. If the specified table does not exist, DynamoDB returns
	// a ResourceNotFoundException. If table is already in the DELETING state, no error
	// is returned. DynamoDB might continue to accept data read and write operations,
	// such as GetItem and PutItem, on a table in the DELETING state until the table
	// deletion is complete. When you delete a table, any indexes on that table are
	// also deleted. If you have DynamoDB Streams enabled on the table, then the
	// corresponding stream on that table goes into the DISABLED state, and the stream
	// is automatically deleted after 24 hours. Use the DescribeTable action to check
	// the status of the table.
	//
	DeleteTable(arg1 context.Context, arg2 *dynamodb.DeleteTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error)
	// Describes an existing backup of a table. You can call DescribeBackup at a
	// maximum rate of 10 times per second.
	//
	DescribeBackup(arg1 context.Context, arg2 *dynamodb.DescribeBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeBackupOutput, error)
	// Checks the status of continuous backups and point in time recovery on the
	// specified table. Continuous backups are ENABLED on all tables at table creation.
	// If point in time recovery is enabled, PointInTimeRecoveryStatus will be set to
	// ENABLED. After continuous backups and point in time recovery are enabled, you
	// can restore to any point in time within EarliestRestorableDateTime and
	// LatestRestorableDateTime. LatestRestorableDateTime is typically 5 minutes before
	// the current time. You can restore your table to any point in time during the
	// last 35 days. You can call DescribeContinuousBackups at a maximum rate of 10
	// times per second.
	//
	DescribeContinuousBackups(arg1 context.Context, arg2 *dynamodb.DescribeContinuousBackupsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeContinuousBackupsOutput, error)
	// Returns information about contributor insights, for a given table or global
	// secondary index.
	//
	DescribeContributorInsights(arg1 context.Context, arg2 *dynamodb.DescribeContributorInsightsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeContributorInsightsOutput, error)
	// Returns the regional endpoint information.
	//
	DescribeEndpoints(arg1 context.Context, arg2 *dynamodb.DescribeEndpointsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeEndpointsOutput, error)
	// Describes an existing table export.
	//
	DescribeExport(arg1 context.Context, arg2 *dynamodb.DescribeExportInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeExportOutput, error)
	// Returns information about the specified global table. This operation only
	// applies to Version 2017.11.29
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html)
	// of global tables. If you are using global tables Version 2019.11.21
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
	// you can use DescribeTable
	// (https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DescribeTable.html)
	// instead.
	//
	DescribeGlobalTable(arg1 context.Context, arg2 *dynamodb.DescribeGlobalTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableOutput, error)
	// Describes Region-specific settings for a global table. This operation only
	// applies to Version 2017.11.29
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html)
	// of global tables.
	//
	DescribeGlobalTableSettings(arg1 context.Context, arg2 *dynamodb.DescribeGlobalTableSettingsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
	// Returns information about the status of Kinesis streaming.
	//
	DescribeKinesisStreamingDestination(arg1 context.Context, arg2 *dynamodb.DescribeKinesisStreamingDestinationInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error)
	// Returns the current provisioned-capacity quotas for your AWS account in a
	// Region, both for the Region as a whole and for any one DynamoDB table that you
	// create there. When you establish an AWS account, the account has initial quotas
	// on the maximum read capacity units and write capacity units that you can
	// provision across all of your DynamoDB tables in a given Region. Also, there are
	// per-table quotas that apply when you create a table there. For more information,
	// see Service, Account, and Table Quotas
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// page in the Amazon DynamoDB Developer Guide. Although you can increase these
	// quotas by filing a case at AWS Support Center
	// (https://console.aws.amazon.com/support/home#/), obtaining the increase is not
	// instantaneous. The DescribeLimits action lets you write code to compare the
	// capacity you are currently using to those quotas imposed by your account so that
	// you have enough time to apply for an increase before you hit a quota. For
	// example, you could use one of the AWS SDKs to do the following:
	//
	// * Call
	// DescribeLimits for a particular Region to obtain your current account quotas on
	// provisioned capacity there.
	//
	// * Create a variable to hold the aggregate read
	// capacity units provisioned for all your tables in that Region, and one to hold
	// the aggregate write capacity units. Zero them both.
	//
	// * Call ListTables to obtain
	// a list of all your DynamoDB tables.
	//
	// * For each table name listed by ListTables,
	// do the following:
	//
	// * Call DescribeTable with the table name.
	//
	// * Use the data
	// returned by DescribeTable to add the read capacity units and write capacity
	// units provisioned for the table itself to your variables.
	//
	// * If the table has
	// one or more global secondary indexes (GSIs), loop over these GSIs and add their
	// provisioned capacity values to your variables as well.
	//
	// * Report the account
	// quotas for that Region returned by DescribeLimits, along with the total current
	// provisioned capacity levels you have calculated.
	//
	// This will let you see whether
	// you are getting close to your account-level quotas. The per-table quotas apply
	// only when you are creating a new table. They restrict the sum of the provisioned
	// capacity of the new table itself and all its global secondary indexes. For
	// existing tables and their GSIs, DynamoDB doesn't let you increase provisioned
	// capacity extremely rapidly, but the only quota that applies is that the
	// aggregate provisioned capacity over all your tables and GSIs cannot exceed
	// either of the per-account quotas. DescribeLimits should only be called
	// periodically. You can expect throttling errors if you call it more than once in
	// a minute. The DescribeLimits Request element has no content.
	//
	DescribeLimits(arg1 context.Context, arg2 *dynamodb.DescribeLimitsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeLimitsOutput, error)
	// Returns information about the table, including the current status of the table,
	// when it was created, the primary key schema, and any indexes on the table. If
	// you issue a DescribeTable request immediately after a CreateTable request,
	// DynamoDB might return a ResourceNotFoundException. This is because DescribeTable
	// uses an eventually consistent query, and the metadata for your table might not
	// be available at that moment. Wait for a few seconds, and then try the
	// DescribeTable request again.
	//
	DescribeTable(arg1 context.Context, arg2 *dynamodb.DescribeTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)
	// Describes auto scaling settings across replicas of the global table at once.
	// This operation only applies to Version 2019.11.21
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
	// of global tables.
	//
	DescribeTableReplicaAutoScaling(arg1 context.Context, arg2 *dynamodb.DescribeTableReplicaAutoScalingInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	// Gives a description of the Time to Live (TTL) status on the specified table.
	//
	DescribeTimeToLive(arg1 context.Context, arg2 *dynamodb.DescribeTimeToLiveInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeTimeToLiveOutput, error)
	// Stops replication from the DynamoDB table to the Kinesis data stream. This is
	// done without deleting either of the resources.
	//
	DisableKinesisStreamingDestination(arg1 context.Context, arg2 *dynamodb.DisableKinesisStreamingDestinationInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DisableKinesisStreamingDestinationOutput, error)
	// Starts table data replication to the specified Kinesis data stream at a
	// timestamp chosen during the enable workflow. If this operation doesn't return
	// results immediately, use DescribeKinesisStreamingDestination to check if
	// streaming to the Kinesis data stream is ACTIVE.
	//
	EnableKinesisStreamingDestination(arg1 context.Context, arg2 *dynamodb.EnableKinesisStreamingDestinationInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.EnableKinesisStreamingDestinationOutput, error)
	// This operation allows you to perform reads and singleton writes on data stored
	// in DynamoDB, using PartiQL.
	//
	ExecuteStatement(arg1 context.Context, arg2 *dynamodb.ExecuteStatementInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error)
	// This operation allows you to perform transactional reads or writes on data
	// stored in DynamoDB, using PartiQL.
	//
	ExecuteTransaction(arg1 context.Context, arg2 *dynamodb.ExecuteTransactionInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ExecuteTransactionOutput, error)
	// Exports table data to an S3 bucket. The table must have point in time recovery
	// enabled, and you can export data from any time within the point in time recovery
	// window.
	//
	ExportTableToPointInTime(arg1 context.Context, arg2 *dynamodb.ExportTableToPointInTimeInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ExportTableToPointInTimeOutput, error)
	// The GetItem operation returns a set of attributes for the item with the given
	// primary key. If there is no matching item, GetItem does not return any data and
	// there will be no Item element in the response. GetItem provides an eventually
	// consistent read by default. If your application requires a strongly consistent
	// read, set ConsistentRead to true. Although a strongly consistent read might take
	// more time than an eventually consistent read, it always returns the last updated
	// value.
	//
	GetItem(arg1 context.Context, arg2 *dynamodb.GetItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	// List backups associated with an AWS account. To list backups for a given table,
	// specify TableName. ListBackups returns a paginated list of results with at most
	// 1 MB worth of items in a page. You can also specify a maximum number of entries
	// to be returned in a page. In the request, start time is inclusive, but end time
	// is exclusive. Note that these boundaries are for the time at which the original
	// backup was requested. You can call ListBackups a maximum of five times per
	// second.
	//
	ListBackups(arg1 context.Context, arg2 *dynamodb.ListBackupsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListBackupsOutput, error)
	// Returns a list of ContributorInsightsSummary for a table and all its global
	// secondary indexes.
	//
	ListContributorInsights(arg1 context.Context, arg2 *dynamodb.ListContributorInsightsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListContributorInsightsOutput, error)
	// Lists completed exports within the past 90 days.
	//
	ListExports(arg1 context.Context, arg2 *dynamodb.ListExportsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListExportsOutput, error)
	// Lists all global tables that have a replica in the specified Region. This
	// operation only applies to Version 2017.11.29
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html)
	// of global tables.
	//
	ListGlobalTables(arg1 context.Context, arg2 *dynamodb.ListGlobalTablesInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListGlobalTablesOutput, error)
	// Returns an array of table names associated with the current account and
	// endpoint. The output from ListTables is paginated, with each page returning a
	// maximum of 100 table names.
	//
	ListTables(arg1 context.Context, arg2 *dynamodb.ListTablesInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)
	// List all tags on an Amazon DynamoDB resource. You can call ListTagsOfResource up
	// to 10 times per second, per account. For an overview on tagging DynamoDB
	// resources, see Tagging for DynamoDB
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html)
	// in the Amazon DynamoDB Developer Guide.
	//
	ListTagsOfResource(arg1 context.Context, arg2 *dynamodb.ListTagsOfResourceInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error)
	// Creates a new item, or replaces an old item with a new item. If an item that has
	// the same primary key as the new item already exists in the specified table, the
	// new item completely replaces the existing item. You can perform a conditional
	// put operation (add a new item if one with the specified primary key doesn't
	// exist), or replace an existing item if it has certain attribute values. You can
	// return the item's attribute values in the same operation, using the ReturnValues
	// parameter. This topic provides general information about the PutItem API. For
	// information on how to call the PutItem API using the AWS SDK in specific
	// languages, see the following:
	//
	// * PutItem in the AWS Command Line Interface
	// (http://docs.aws.amazon.com/goto/aws-cli/dynamodb-2012-08-10/PutItem)
	//
	// * PutItem
	// in the AWS SDK for .NET
	// (http://docs.aws.amazon.com/goto/DotNetSDKV3/dynamodb-2012-08-10/PutItem)
	//
	// *
	// PutItem in the AWS SDK for C++
	// (http://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/PutItem)
	//
	// *
	// PutItem in the AWS SDK for Go
	// (http://docs.aws.amazon.com/goto/SdkForGoV1/dynamodb-2012-08-10/PutItem)
	//
	// *
	// PutItem in the AWS SDK for Java
	// (http://docs.aws.amazon.com/goto/SdkForJava/dynamodb-2012-08-10/PutItem)
	//
	// *
	// PutItem in the AWS SDK for JavaScript
	// (http://docs.aws.amazon.com/goto/AWSJavaScriptSDK/dynamodb-2012-08-10/PutItem)
	//
	// *
	// PutItem in the AWS SDK for PHP V3
	// (http://docs.aws.amazon.com/goto/SdkForPHPV3/dynamodb-2012-08-10/PutItem)
	//
	// *
	// PutItem in the AWS SDK for Python
	// (http://docs.aws.amazon.com/goto/boto3/dynamodb-2012-08-10/PutItem)
	//
	// * PutItem
	// in the AWS SDK for Ruby V2
	// (http://docs.aws.amazon.com/goto/SdkForRubyV2/dynamodb-2012-08-10/PutItem)
	//
	// When
	// you add an item, the primary key attributes are the only required attributes.
	// Attribute values cannot be null. Empty String and Binary attribute values are
	// allowed. Attribute values of type String and Binary must have a length greater
	// than zero if the attribute is used as a key attribute for a table or index. Set
	// type attributes cannot be empty. Invalid Requests with empty values will be
	// rejected with a ValidationException exception. To prevent a new item from
	// replacing an existing item, use a conditional expression that contains the
	// attribute_not_exists function with the name of the attribute being used as the
	// partition key for the table. Since every record must contain that attribute, the
	// attribute_not_exists function will only succeed if no matching item exists. For
	// more information about PutItem, see Working with Items
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithItems.html)
	// in the Amazon DynamoDB Developer Guide.
	//
	PutItem(arg1 context.Context, arg2 *dynamodb.PutItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	// The Query operation finds items based on primary key values. You can query any
	// table or secondary index that has a composite primary key (a partition key and a
	// sort key). Use the KeyConditionExpression parameter to provide a specific value
	// for the partition key. The Query operation will return all of the items from the
	// table or index with that partition key value. You can optionally narrow the
	// scope of the Query operation by specifying a sort key value and a comparison
	// operator in KeyConditionExpression. To further refine the Query results, you can
	// optionally provide a FilterExpression. A FilterExpression determines which items
	// within the results should be returned to you. All of the other results are
	// discarded. A Query operation always returns a result set. If no matching items
	// are found, the result set will be empty. Queries that do not return results
	// consume the minimum number of read capacity units for that type of read
	// operation. DynamoDB calculates the number of read capacity units consumed based
	// on item size, not on the amount of data that is returned to an application. The
	// number of capacity units consumed will be the same whether you request all of
	// the attributes (the default behavior) or just some of them (using a projection
	// expression). The number will also be the same whether or not you use a
	// FilterExpression. Query results are always sorted by the sort key value. If the
	// data type of the sort key is Number, the results are returned in numeric order;
	// otherwise, the results are returned in order of UTF-8 bytes. By default, the
	// sort order is ascending. To reverse the order, set the ScanIndexForward
	// parameter to false. A single Query operation will read up to the maximum number
	// of items set (if using the Limit parameter) or a maximum of 1 MB of data and
	// then apply any filtering to the results using FilterExpression. If
	// LastEvaluatedKey is present in the response, you will need to paginate the
	// result set. For more information, see Paginating the Results
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Query.html#Query.Pagination)
	// in the Amazon DynamoDB Developer Guide. FilterExpression is applied after a
	// Query finishes, but before the results are returned. A FilterExpression cannot
	// contain partition key or sort key attributes. You need to specify those
	// attributes in the KeyConditionExpression. A Query operation can return an empty
	// result set and a LastEvaluatedKey if all the items read for the page of results
	// are filtered out. You can query a table, a local secondary index, or a global
	// secondary index. For a query on a table or on a local secondary index, you can
	// set the ConsistentRead parameter to true and obtain a strongly consistent
	// result. Global secondary indexes support eventually consistent reads only, so do
	// not specify ConsistentRead when querying a global secondary index.
	//
	Query(arg1 context.Context, arg2 *dynamodb.QueryInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	// Creates a new table from an existing backup. Any number of users can execute up
	// to 4 concurrent restores (any type of restore) in a given account. You can call
	// RestoreTableFromBackup at a maximum rate of 10 times per second. You must
	// manually set up the following on the restored table:
	//
	// * Auto scaling policies
	//
	// *
	// IAM policies
	//
	// * Amazon CloudWatch metrics and alarms
	//
	// * Tags
	//
	// * Stream
	// settings
	//
	// * Time to Live (TTL) settings
	//
	RestoreTableFromBackup(arg1 context.Context, arg2 *dynamodb.RestoreTableFromBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.RestoreTableFromBackupOutput, error)
	// Restores the specified table to the specified point in time within
	// EarliestRestorableDateTime and LatestRestorableDateTime. You can restore your
	// table to any point in time during the last 35 days. Any number of users can
	// execute up to 4 concurrent restores (any type of restore) in a given account.
	// When you restore using point in time recovery, DynamoDB restores your table data
	// to the state based on the selected date and time (day:hour:minute:second) to a
	// new table. Along with data, the following are also included on the new restored
	// table using point in time recovery:
	//
	// * Global secondary indexes (GSIs)
	//
	// * Local
	// secondary indexes (LSIs)
	//
	// * Provisioned read and write capacity
	//
	// * Encryption
	// settings All these settings come from the current settings of the source table
	// at the time of restore.
	//
	// You must manually set up the following on the restored
	// table:
	//
	// * Auto scaling policies
	//
	// * IAM policies
	//
	// * Amazon CloudWatch metrics and
	// alarms
	//
	// * Tags
	//
	// * Stream settings
	//
	// * Time to Live (TTL) settings
	//
	// * Point in
	// time recovery settings
	//
	RestoreTableToPointInTime(arg1 context.Context, arg2 *dynamodb.RestoreTableToPointInTimeInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.RestoreTableToPointInTimeOutput, error)
	// The Scan operation returns one or more items and item attributes by accessing
	// every item in a table or a secondary index. To have DynamoDB return fewer items,
	// you can provide a FilterExpression operation. If the total number of scanned
	// items exceeds the maximum dataset size limit of 1 MB, the scan stops and results
	// are returned to the user as a LastEvaluatedKey value to continue the scan in a
	// subsequent operation. The results also include the number of items exceeding the
	// limit. A scan can result in no table data meeting the filter criteria. A single
	// Scan operation reads up to the maximum number of items set (if using the Limit
	// parameter) or a maximum of 1 MB of data and then apply any filtering to the
	// results using FilterExpression. If LastEvaluatedKey is present in the response,
	// you need to paginate the result set. For more information, see Paginating the
	// Results
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Scan.html#Scan.Pagination)
	// in the Amazon DynamoDB Developer Guide. Scan operations proceed sequentially;
	// however, for faster performance on a large table or secondary index,
	// applications can request a parallel Scan operation by providing the Segment and
	// TotalSegments parameters. For more information, see Parallel Scan
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Scan.html#Scan.ParallelScan)
	// in the Amazon DynamoDB Developer Guide. Scan uses eventually consistent reads
	// when accessing the data in a table; therefore, the result set might not include
	// the changes to data in the table immediately before the operation began. If you
	// need a consistent copy of the data, as of the time that the Scan begins, you can
	// set the ConsistentRead parameter to true.
	//
	Scan(arg1 context.Context, arg2 *dynamodb.ScanInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
	// Associate a set of tags with an Amazon DynamoDB resource. You can then activate
	// these user-defined tags so that they appear on the Billing and Cost Management
	// console for cost allocation tracking. You can call TagResource up to five times
	// per second, per account. For an overview on tagging DynamoDB resources, see
	// Tagging for DynamoDB
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html)
	// in the Amazon DynamoDB Developer Guide.
	//
	TagResource(arg1 context.Context, arg2 *dynamodb.TagResourceInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.TagResourceOutput, error)
	// TransactGetItems is a synchronous operation that atomically retrieves multiple
	// items from one or more tables (but not from indexes) in a single account and
	// Region. A TransactGetItems call can contain up to 25 TransactGetItem objects,
	// each of which contains a Get structure that specifies an item to retrieve from a
	// table in the account and Region. A call to TransactGetItems cannot retrieve
	// items from tables in more than one AWS account or Region. The aggregate size of
	// the items in the transaction cannot exceed 4 MB. DynamoDB rejects the entire
	// TransactGetItems request if any of the following is true:
	//
	// * A conflicting
	// operation is in the process of updating an item to be read.
	//
	// * There is
	// insufficient provisioned capacity for the transaction to be completed.
	//
	// * There
	// is a user error, such as an invalid data format.
	//
	// * The aggregate size of the
	// items in the transaction cannot exceed 4 MB.
	//
	TransactGetItems(arg1 context.Context, arg2 *dynamodb.TransactGetItemsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.TransactGetItemsOutput, error)
	// TransactWriteItems is a synchronous write operation that groups up to 25 action
	// requests. These actions can target items in different tables, but not in
	// different AWS accounts or Regions, and no two actions can target the same item.
	// For example, you cannot both ConditionCheck and Update the same item. The
	// aggregate size of the items in the transaction cannot exceed 4 MB. The actions
	// are completed atomically so that either all of them succeed, or all of them
	// fail. They are defined by the following objects:
	//
	// * Put  Initiates a PutItem
	// operation to write a new item. This structure specifies the primary key of the
	// item to be written, the name of the table to write it in, an optional condition
	// expression that must be satisfied for the write to succeed, a list of the item's
	// attributes, and a field indicating whether to retrieve the item's attributes if
	// the condition is not met.
	//
	// * Update  Initiates an UpdateItem operation to
	// update an existing item. This structure specifies the primary key of the item to
	// be updated, the name of the table where it resides, an optional condition
	// expression that must be satisfied for the update to succeed, an expression that
	// defines one or more attributes to be updated, and a field indicating whether to
	// retrieve the item's attributes if the condition is not met.
	//
	// * Delete 
	// Initiates a DeleteItem operation to delete an existing item. This structure
	// specifies the primary key of the item to be deleted, the name of the table where
	// it resides, an optional condition expression that must be satisfied for the
	// deletion to succeed, and a field indicating whether to retrieve the item's
	// attributes if the condition is not met.
	//
	// * ConditionCheck  Applies a condition
	// to an item that is not being modified by the transaction. This structure
	// specifies the primary key of the item to be checked, the name of the table where
	// it resides, a condition expression that must be satisfied for the transaction to
	// succeed, and a field indicating whether to retrieve the item's attributes if the
	// condition is not met.
	//
	// DynamoDB rejects the entire TransactWriteItems request if
	// any of the following is true:
	//
	// * A condition in one of the condition expressions
	// is not met.
	//
	// * An ongoing operation is in the process of updating the same
	// item.
	//
	// * There is insufficient provisioned capacity for the transaction to be
	// completed.
	//
	// * An item size becomes too large (bigger than 400 KB), a local
	// secondary index (LSI) becomes too large, or a similar validation error occurs
	// because of changes made by the transaction.
	//
	// * The aggregate size of the items
	// in the transaction exceeds 4 MB.
	//
	// * There is a user error, such as an invalid
	// data format.
	//
	TransactWriteItems(arg1 context.Context, arg2 *dynamodb.TransactWriteItemsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error)
	// Removes the association of tags from an Amazon DynamoDB resource. You can call
	// UntagResource up to five times per second, per account. For an overview on
	// tagging DynamoDB resources, see Tagging for DynamoDB
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html)
	// in the Amazon DynamoDB Developer Guide.
	//
	UntagResource(arg1 context.Context, arg2 *dynamodb.UntagResourceInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UntagResourceOutput, error)
	// UpdateContinuousBackups enables or disables point in time recovery for the
	// specified table. A successful UpdateContinuousBackups call returns the current
	// ContinuousBackupsDescription. Continuous backups are ENABLED on all tables at
	// table creation. If point in time recovery is enabled, PointInTimeRecoveryStatus
	// will be set to ENABLED. Once continuous backups and point in time recovery are
	// enabled, you can restore to any point in time within EarliestRestorableDateTime
	// and LatestRestorableDateTime. LatestRestorableDateTime is typically 5 minutes
	// before the current time. You can restore your table to any point in time during
	// the last 35 days.
	//
	UpdateContinuousBackups(arg1 context.Context, arg2 *dynamodb.UpdateContinuousBackupsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateContinuousBackupsOutput, error)
	// Updates the status for contributor insights for a specific table or index.
	//
	UpdateContributorInsights(arg1 context.Context, arg2 *dynamodb.UpdateContributorInsightsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateContributorInsightsOutput, error)
	// Adds or removes replicas in the specified global table. The global table must
	// already exist to be able to use this operation. Any replica to be added must be
	// empty, have the same name as the global table, have the same key schema, have
	// DynamoDB Streams enabled, and have the same provisioned and maximum write
	// capacity units. Although you can use UpdateGlobalTable to add replicas and
	// remove replicas in a single request, for simplicity we recommend that you issue
	// separate requests for adding or removing replicas. If global secondary indexes
	// are specified, then the following conditions must also be met:
	//
	// * The global
	// secondary indexes must have the same name.
	//
	// * The global secondary indexes must
	// have the same hash key and sort key (if present).
	//
	// * The global secondary
	// indexes must have the same provisioned and maximum write capacity units.
	//
	UpdateGlobalTable(arg1 context.Context, arg2 *dynamodb.UpdateGlobalTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateGlobalTableOutput, error)
	// Updates settings for a global table.
	//
	UpdateGlobalTableSettings(arg1 context.Context, arg2 *dynamodb.UpdateGlobalTableSettingsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
	// Edits an existing item's attributes, or adds a new item to the table if it does
	// not already exist. You can put, delete, or add attribute values. You can also
	// perform a conditional update on an existing item (insert a new attribute
	// name-value pair if it doesn't exist, or replace an existing name-value pair if
	// it has certain expected attribute values). You can also return the item's
	// attribute values in the same UpdateItem operation using the ReturnValues
	// parameter.
	//
	UpdateItem(arg1 context.Context, arg2 *dynamodb.UpdateItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	// Modifies the provisioned throughput settings, global secondary indexes, or
	// DynamoDB Streams settings for a given table. You can only perform one of the
	// following operations at once:
	//
	// * Modify the provisioned throughput settings of
	// the table.
	//
	// * Enable or disable DynamoDB Streams on the table.
	//
	// * Remove a
	// global secondary index from the table.
	//
	// * Create a new global secondary index on
	// the table. After the index begins backfilling, you can use UpdateTable to
	// perform other operations.
	//
	// UpdateTable is an asynchronous operation; while it is
	// executing, the table status changes from ACTIVE to UPDATING. While it is
	// UPDATING, you cannot issue another UpdateTable request. When the table returns
	// to the ACTIVE state, the UpdateTable operation is complete.
	//
	UpdateTable(arg1 context.Context, arg2 *dynamodb.UpdateTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateTableOutput, error)
	// Updates auto scaling settings on your global tables at once. This operation only
	// applies to Version 2019.11.21
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
	// of global tables.
	//
	UpdateTableReplicaAutoScaling(arg1 context.Context, arg2 *dynamodb.UpdateTableReplicaAutoScalingInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error)
	// The UpdateTimeToLive method enables or disables Time to Live (TTL) for the
	// specified table. A successful UpdateTimeToLive call returns the current
	// TimeToLiveSpecification. It can take up to one hour for the change to fully
	// process. Any additional UpdateTimeToLive calls for the same table during this
	// one hour duration result in a ValidationException. TTL compares the current time
	// in epoch time format to the time stored in the TTL attribute of an item. If the
	// epoch time value stored in the attribute is less than the current time, the item
	// is marked as expired and subsequently deleted. The epoch time format is the
	// number of seconds elapsed since 12:00:00 AM January 1, 1970 UTC. DynamoDB
	// deletes expired items on a best-effort basis to ensure availability of
	// throughput for other data operations. DynamoDB typically deletes expired items
	// within two days of expiration. The exact duration within which an item gets
	// deleted after expiration is specific to the nature of the workload. Items that
	// have expired and not been deleted will still show up in reads, queries, and
	// scans. As items are deleted, they are removed from any local secondary index and
	// global secondary index immediately in the same eventually consistent way as a
	// standard delete operation. For more information, see Time To Live
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/TTL.html) in
	// the Amazon DynamoDB Developer Guide.
	//
	UpdateTimeToLive(arg1 context.Context, arg2 *dynamodb.UpdateTimeToLiveInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateTimeToLiveOutput, error)
}

// DynamoClientMock generic client mock
type DynamoClientMock struct {
	BatchExecuteStatementMock               func(arg1 context.Context, arg2 *dynamodb.BatchExecuteStatementInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.BatchExecuteStatementOutput, error)
	BatchGetItemMock                        func(arg1 context.Context, arg2 *dynamodb.BatchGetItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error)
	BatchWriteItemMock                      func(arg1 context.Context, arg2 *dynamodb.BatchWriteItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error)
	CreateBackupMock                        func(arg1 context.Context, arg2 *dynamodb.CreateBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.CreateBackupOutput, error)
	CreateGlobalTableMock                   func(arg1 context.Context, arg2 *dynamodb.CreateGlobalTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.CreateGlobalTableOutput, error)
	CreateTableMock                         func(arg1 context.Context, arg2 *dynamodb.CreateTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)
	DeleteBackupMock                        func(arg1 context.Context, arg2 *dynamodb.DeleteBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DeleteBackupOutput, error)
	DeleteItemMock                          func(arg1 context.Context, arg2 *dynamodb.DeleteItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
	DeleteTableMock                         func(arg1 context.Context, arg2 *dynamodb.DeleteTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error)
	DescribeBackupMock                      func(arg1 context.Context, arg2 *dynamodb.DescribeBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeBackupOutput, error)
	DescribeContinuousBackupsMock           func(arg1 context.Context, arg2 *dynamodb.DescribeContinuousBackupsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeContinuousBackupsOutput, error)
	DescribeContributorInsightsMock         func(arg1 context.Context, arg2 *dynamodb.DescribeContributorInsightsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeContributorInsightsOutput, error)
	DescribeEndpointsMock                   func(arg1 context.Context, arg2 *dynamodb.DescribeEndpointsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeEndpointsOutput, error)
	DescribeExportMock                      func(arg1 context.Context, arg2 *dynamodb.DescribeExportInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeExportOutput, error)
	DescribeGlobalTableMock                 func(arg1 context.Context, arg2 *dynamodb.DescribeGlobalTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableOutput, error)
	DescribeGlobalTableSettingsMock         func(arg1 context.Context, arg2 *dynamodb.DescribeGlobalTableSettingsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableSettingsOutput, error)
	DescribeKinesisStreamingDestinationMock func(arg1 context.Context, arg2 *dynamodb.DescribeKinesisStreamingDestinationInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error)
	DescribeLimitsMock                      func(arg1 context.Context, arg2 *dynamodb.DescribeLimitsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeLimitsOutput, error)
	DescribeTableMock                       func(arg1 context.Context, arg2 *dynamodb.DescribeTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)
	DescribeTableReplicaAutoScalingMock     func(arg1 context.Context, arg2 *dynamodb.DescribeTableReplicaAutoScalingInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error)
	DescribeTimeToLiveMock                  func(arg1 context.Context, arg2 *dynamodb.DescribeTimeToLiveInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeTimeToLiveOutput, error)
	DisableKinesisStreamingDestinationMock  func(arg1 context.Context, arg2 *dynamodb.DisableKinesisStreamingDestinationInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DisableKinesisStreamingDestinationOutput, error)
	EnableKinesisStreamingDestinationMock   func(arg1 context.Context, arg2 *dynamodb.EnableKinesisStreamingDestinationInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.EnableKinesisStreamingDestinationOutput, error)
	ExecuteStatementMock                    func(arg1 context.Context, arg2 *dynamodb.ExecuteStatementInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error)
	ExecuteTransactionMock                  func(arg1 context.Context, arg2 *dynamodb.ExecuteTransactionInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ExecuteTransactionOutput, error)
	ExportTableToPointInTimeMock            func(arg1 context.Context, arg2 *dynamodb.ExportTableToPointInTimeInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ExportTableToPointInTimeOutput, error)
	GetItemMock                             func(arg1 context.Context, arg2 *dynamodb.GetItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	ListBackupsMock                         func(arg1 context.Context, arg2 *dynamodb.ListBackupsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListBackupsOutput, error)
	ListContributorInsightsMock             func(arg1 context.Context, arg2 *dynamodb.ListContributorInsightsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListContributorInsightsOutput, error)
	ListExportsMock                         func(arg1 context.Context, arg2 *dynamodb.ListExportsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListExportsOutput, error)
	ListGlobalTablesMock                    func(arg1 context.Context, arg2 *dynamodb.ListGlobalTablesInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListGlobalTablesOutput, error)
	ListTablesMock                          func(arg1 context.Context, arg2 *dynamodb.ListTablesInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error)
	ListTagsOfResourceMock                  func(arg1 context.Context, arg2 *dynamodb.ListTagsOfResourceInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error)
	PutItemMock                             func(arg1 context.Context, arg2 *dynamodb.PutItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	QueryMock                               func(arg1 context.Context, arg2 *dynamodb.QueryInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
	RestoreTableFromBackupMock              func(arg1 context.Context, arg2 *dynamodb.RestoreTableFromBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.RestoreTableFromBackupOutput, error)
	RestoreTableToPointInTimeMock           func(arg1 context.Context, arg2 *dynamodb.RestoreTableToPointInTimeInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.RestoreTableToPointInTimeOutput, error)
	ScanMock                                func(arg1 context.Context, arg2 *dynamodb.ScanInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
	TagResourceMock                         func(arg1 context.Context, arg2 *dynamodb.TagResourceInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.TagResourceOutput, error)
	TransactGetItemsMock                    func(arg1 context.Context, arg2 *dynamodb.TransactGetItemsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.TransactGetItemsOutput, error)
	TransactWriteItemsMock                  func(arg1 context.Context, arg2 *dynamodb.TransactWriteItemsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error)
	UntagResourceMock                       func(arg1 context.Context, arg2 *dynamodb.UntagResourceInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UntagResourceOutput, error)
	UpdateContinuousBackupsMock             func(arg1 context.Context, arg2 *dynamodb.UpdateContinuousBackupsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateContinuousBackupsOutput, error)
	UpdateContributorInsightsMock           func(arg1 context.Context, arg2 *dynamodb.UpdateContributorInsightsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateContributorInsightsOutput, error)
	UpdateGlobalTableMock                   func(arg1 context.Context, arg2 *dynamodb.UpdateGlobalTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateGlobalTableOutput, error)
	UpdateGlobalTableSettingsMock           func(arg1 context.Context, arg2 *dynamodb.UpdateGlobalTableSettingsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateGlobalTableSettingsOutput, error)
	UpdateItemMock                          func(arg1 context.Context, arg2 *dynamodb.UpdateItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	UpdateTableMock                         func(arg1 context.Context, arg2 *dynamodb.UpdateTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateTableOutput, error)
	UpdateTableReplicaAutoScalingMock       func(arg1 context.Context, arg2 *dynamodb.UpdateTableReplicaAutoScalingInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error)
	UpdateTimeToLiveMock                    func(arg1 context.Context, arg2 *dynamodb.UpdateTimeToLiveInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateTimeToLiveOutput, error)
}

// BatchExecuteStatement mocked function
func (m DynamoClientMock) BatchExecuteStatement(arg1 context.Context, arg2 *dynamodb.BatchExecuteStatementInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.BatchExecuteStatementOutput, error) {
	return m.BatchExecuteStatementMock(arg1, arg2, arg3...)
}

// BatchGetItem mocked function
func (m DynamoClientMock) BatchGetItem(arg1 context.Context, arg2 *dynamodb.BatchGetItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error) {
	return m.BatchGetItemMock(arg1, arg2, arg3...)
}

// BatchWriteItem mocked function
func (m DynamoClientMock) BatchWriteItem(arg1 context.Context, arg2 *dynamodb.BatchWriteItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error) {
	return m.BatchWriteItemMock(arg1, arg2, arg3...)
}

// CreateBackup mocked function
func (m DynamoClientMock) CreateBackup(arg1 context.Context, arg2 *dynamodb.CreateBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.CreateBackupOutput, error) {
	return m.CreateBackupMock(arg1, arg2, arg3...)
}

// CreateGlobalTable mocked function
func (m DynamoClientMock) CreateGlobalTable(arg1 context.Context, arg2 *dynamodb.CreateGlobalTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.CreateGlobalTableOutput, error) {
	return m.CreateGlobalTableMock(arg1, arg2, arg3...)
}

// CreateTable mocked function
func (m DynamoClientMock) CreateTable(arg1 context.Context, arg2 *dynamodb.CreateTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
	return m.CreateTableMock(arg1, arg2, arg3...)
}

// DeleteBackup mocked function
func (m DynamoClientMock) DeleteBackup(arg1 context.Context, arg2 *dynamodb.DeleteBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DeleteBackupOutput, error) {
	return m.DeleteBackupMock(arg1, arg2, arg3...)
}

// DeleteItem mocked function
func (m DynamoClientMock) DeleteItem(arg1 context.Context, arg2 *dynamodb.DeleteItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	return m.DeleteItemMock(arg1, arg2, arg3...)
}

// DeleteTable mocked function
func (m DynamoClientMock) DeleteTable(arg1 context.Context, arg2 *dynamodb.DeleteTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DeleteTableOutput, error) {
	return m.DeleteTableMock(arg1, arg2, arg3...)
}

// DescribeBackup mocked function
func (m DynamoClientMock) DescribeBackup(arg1 context.Context, arg2 *dynamodb.DescribeBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeBackupOutput, error) {
	return m.DescribeBackupMock(arg1, arg2, arg3...)
}

// DescribeContinuousBackups mocked function
func (m DynamoClientMock) DescribeContinuousBackups(arg1 context.Context, arg2 *dynamodb.DescribeContinuousBackupsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	return m.DescribeContinuousBackupsMock(arg1, arg2, arg3...)
}

// DescribeContributorInsights mocked function
func (m DynamoClientMock) DescribeContributorInsights(arg1 context.Context, arg2 *dynamodb.DescribeContributorInsightsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeContributorInsightsOutput, error) {
	return m.DescribeContributorInsightsMock(arg1, arg2, arg3...)
}

// DescribeEndpoints mocked function
func (m DynamoClientMock) DescribeEndpoints(arg1 context.Context, arg2 *dynamodb.DescribeEndpointsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeEndpointsOutput, error) {
	return m.DescribeEndpointsMock(arg1, arg2, arg3...)
}

// DescribeExport mocked function
func (m DynamoClientMock) DescribeExport(arg1 context.Context, arg2 *dynamodb.DescribeExportInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeExportOutput, error) {
	return m.DescribeExportMock(arg1, arg2, arg3...)
}

// DescribeGlobalTable mocked function
func (m DynamoClientMock) DescribeGlobalTable(arg1 context.Context, arg2 *dynamodb.DescribeGlobalTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableOutput, error) {
	return m.DescribeGlobalTableMock(arg1, arg2, arg3...)
}

// DescribeGlobalTableSettings mocked function
func (m DynamoClientMock) DescribeGlobalTableSettings(arg1 context.Context, arg2 *dynamodb.DescribeGlobalTableSettingsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	return m.DescribeGlobalTableSettingsMock(arg1, arg2, arg3...)
}

// DescribeKinesisStreamingDestination mocked function
func (m DynamoClientMock) DescribeKinesisStreamingDestination(arg1 context.Context, arg2 *dynamodb.DescribeKinesisStreamingDestinationInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	return m.DescribeKinesisStreamingDestinationMock(arg1, arg2, arg3...)
}

// DescribeLimits mocked function
func (m DynamoClientMock) DescribeLimits(arg1 context.Context, arg2 *dynamodb.DescribeLimitsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeLimitsOutput, error) {
	return m.DescribeLimitsMock(arg1, arg2, arg3...)
}

// DescribeTable mocked function
func (m DynamoClientMock) DescribeTable(arg1 context.Context, arg2 *dynamodb.DescribeTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	return m.DescribeTableMock(arg1, arg2, arg3...)
}

// DescribeTableReplicaAutoScaling mocked function
func (m DynamoClientMock) DescribeTableReplicaAutoScaling(arg1 context.Context, arg2 *dynamodb.DescribeTableReplicaAutoScalingInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	return m.DescribeTableReplicaAutoScalingMock(arg1, arg2, arg3...)
}

// DescribeTimeToLive mocked function
func (m DynamoClientMock) DescribeTimeToLive(arg1 context.Context, arg2 *dynamodb.DescribeTimeToLiveInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DescribeTimeToLiveOutput, error) {
	return m.DescribeTimeToLiveMock(arg1, arg2, arg3...)
}

// DisableKinesisStreamingDestination mocked function
func (m DynamoClientMock) DisableKinesisStreamingDestination(arg1 context.Context, arg2 *dynamodb.DisableKinesisStreamingDestinationInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.DisableKinesisStreamingDestinationOutput, error) {
	return m.DisableKinesisStreamingDestinationMock(arg1, arg2, arg3...)
}

// EnableKinesisStreamingDestination mocked function
func (m DynamoClientMock) EnableKinesisStreamingDestination(arg1 context.Context, arg2 *dynamodb.EnableKinesisStreamingDestinationInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.EnableKinesisStreamingDestinationOutput, error) {
	return m.EnableKinesisStreamingDestinationMock(arg1, arg2, arg3...)
}

// ExecuteStatement mocked function
func (m DynamoClientMock) ExecuteStatement(arg1 context.Context, arg2 *dynamodb.ExecuteStatementInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error) {
	return m.ExecuteStatementMock(arg1, arg2, arg3...)
}

// ExecuteTransaction mocked function
func (m DynamoClientMock) ExecuteTransaction(arg1 context.Context, arg2 *dynamodb.ExecuteTransactionInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ExecuteTransactionOutput, error) {
	return m.ExecuteTransactionMock(arg1, arg2, arg3...)
}

// ExportTableToPointInTime mocked function
func (m DynamoClientMock) ExportTableToPointInTime(arg1 context.Context, arg2 *dynamodb.ExportTableToPointInTimeInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ExportTableToPointInTimeOutput, error) {
	return m.ExportTableToPointInTimeMock(arg1, arg2, arg3...)
}

// GetItem mocked function
func (m DynamoClientMock) GetItem(arg1 context.Context, arg2 *dynamodb.GetItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return m.GetItemMock(arg1, arg2, arg3...)
}

// ListBackups mocked function
func (m DynamoClientMock) ListBackups(arg1 context.Context, arg2 *dynamodb.ListBackupsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListBackupsOutput, error) {
	return m.ListBackupsMock(arg1, arg2, arg3...)
}

// ListContributorInsights mocked function
func (m DynamoClientMock) ListContributorInsights(arg1 context.Context, arg2 *dynamodb.ListContributorInsightsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListContributorInsightsOutput, error) {
	return m.ListContributorInsightsMock(arg1, arg2, arg3...)
}

// ListExports mocked function
func (m DynamoClientMock) ListExports(arg1 context.Context, arg2 *dynamodb.ListExportsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListExportsOutput, error) {
	return m.ListExportsMock(arg1, arg2, arg3...)
}

// ListGlobalTables mocked function
func (m DynamoClientMock) ListGlobalTables(arg1 context.Context, arg2 *dynamodb.ListGlobalTablesInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListGlobalTablesOutput, error) {
	return m.ListGlobalTablesMock(arg1, arg2, arg3...)
}

// ListTables mocked function
func (m DynamoClientMock) ListTables(arg1 context.Context, arg2 *dynamodb.ListTablesInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error) {
	return m.ListTablesMock(arg1, arg2, arg3...)
}

// ListTagsOfResource mocked function
func (m DynamoClientMock) ListTagsOfResource(arg1 context.Context, arg2 *dynamodb.ListTagsOfResourceInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error) {
	return m.ListTagsOfResourceMock(arg1, arg2, arg3...)
}

// PutItem mocked function
func (m DynamoClientMock) PutItem(arg1 context.Context, arg2 *dynamodb.PutItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	return m.PutItemMock(arg1, arg2, arg3...)
}

// Query mocked function
func (m DynamoClientMock) Query(arg1 context.Context, arg2 *dynamodb.QueryInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	return m.QueryMock(arg1, arg2, arg3...)
}

// RestoreTableFromBackup mocked function
func (m DynamoClientMock) RestoreTableFromBackup(arg1 context.Context, arg2 *dynamodb.RestoreTableFromBackupInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.RestoreTableFromBackupOutput, error) {
	return m.RestoreTableFromBackupMock(arg1, arg2, arg3...)
}

// RestoreTableToPointInTime mocked function
func (m DynamoClientMock) RestoreTableToPointInTime(arg1 context.Context, arg2 *dynamodb.RestoreTableToPointInTimeInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.RestoreTableToPointInTimeOutput, error) {
	return m.RestoreTableToPointInTimeMock(arg1, arg2, arg3...)
}

// Scan mocked function
func (m DynamoClientMock) Scan(arg1 context.Context, arg2 *dynamodb.ScanInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	return m.ScanMock(arg1, arg2, arg3...)
}

// TagResource mocked function
func (m DynamoClientMock) TagResource(arg1 context.Context, arg2 *dynamodb.TagResourceInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.TagResourceOutput, error) {
	return m.TagResourceMock(arg1, arg2, arg3...)
}

// TransactGetItems mocked function
func (m DynamoClientMock) TransactGetItems(arg1 context.Context, arg2 *dynamodb.TransactGetItemsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.TransactGetItemsOutput, error) {
	return m.TransactGetItemsMock(arg1, arg2, arg3...)
}

// TransactWriteItems mocked function
func (m DynamoClientMock) TransactWriteItems(arg1 context.Context, arg2 *dynamodb.TransactWriteItemsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error) {
	return m.TransactWriteItemsMock(arg1, arg2, arg3...)
}

// UntagResource mocked function
func (m DynamoClientMock) UntagResource(arg1 context.Context, arg2 *dynamodb.UntagResourceInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UntagResourceOutput, error) {
	return m.UntagResourceMock(arg1, arg2, arg3...)
}

// UpdateContinuousBackups mocked function
func (m DynamoClientMock) UpdateContinuousBackups(arg1 context.Context, arg2 *dynamodb.UpdateContinuousBackupsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateContinuousBackupsOutput, error) {
	return m.UpdateContinuousBackupsMock(arg1, arg2, arg3...)
}

// UpdateContributorInsights mocked function
func (m DynamoClientMock) UpdateContributorInsights(arg1 context.Context, arg2 *dynamodb.UpdateContributorInsightsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateContributorInsightsOutput, error) {
	return m.UpdateContributorInsightsMock(arg1, arg2, arg3...)
}

// UpdateGlobalTable mocked function
func (m DynamoClientMock) UpdateGlobalTable(arg1 context.Context, arg2 *dynamodb.UpdateGlobalTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateGlobalTableOutput, error) {
	return m.UpdateGlobalTableMock(arg1, arg2, arg3...)
}

// UpdateGlobalTableSettings mocked function
func (m DynamoClientMock) UpdateGlobalTableSettings(arg1 context.Context, arg2 *dynamodb.UpdateGlobalTableSettingsInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateGlobalTableSettingsOutput, error) {
	return m.UpdateGlobalTableSettingsMock(arg1, arg2, arg3...)
}

// UpdateItem mocked function
func (m DynamoClientMock) UpdateItem(arg1 context.Context, arg2 *dynamodb.UpdateItemInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	return m.UpdateItemMock(arg1, arg2, arg3...)
}

// UpdateTable mocked function
func (m DynamoClientMock) UpdateTable(arg1 context.Context, arg2 *dynamodb.UpdateTableInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateTableOutput, error) {
	return m.UpdateTableMock(arg1, arg2, arg3...)
}

// UpdateTableReplicaAutoScaling mocked function
func (m DynamoClientMock) UpdateTableReplicaAutoScaling(arg1 context.Context, arg2 *dynamodb.UpdateTableReplicaAutoScalingInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateTableReplicaAutoScalingOutput, error) {
	return m.UpdateTableReplicaAutoScalingMock(arg1, arg2, arg3...)
}

// UpdateTimeToLive mocked function
func (m DynamoClientMock) UpdateTimeToLive(arg1 context.Context, arg2 *dynamodb.UpdateTimeToLiveInput, arg3 ...func(*dynamodb.Options)) (*dynamodb.UpdateTimeToLiveOutput, error) {
	return m.UpdateTimeToLiveMock(arg1, arg2, arg3...)
}
