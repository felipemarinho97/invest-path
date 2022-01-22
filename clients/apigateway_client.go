package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/felipemarinho97/go-mock-gen
 */

// IAPIGatewayClient generic client
type IAPIGatewayClient interface {
	// Create an ApiKey resource. AWS CLI
	// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/create-api-key.html)
	//
	CreateApiKey(arg1 context.Context, arg2 *apigateway.CreateApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateApiKeyOutput, error)
	// Adds a new Authorizer resource to an existing RestApi resource. AWS CLI
	// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/create-authorizer.html)
	//
	CreateAuthorizer(arg1 context.Context, arg2 *apigateway.CreateAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateAuthorizerOutput, error)
	// Creates a new BasePathMapping resource.
	//
	CreateBasePathMapping(arg1 context.Context, arg2 *apigateway.CreateBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateBasePathMappingOutput, error)
	// Creates a Deployment resource, which makes a specified RestApi callable over the
	// internet.
	//
	CreateDeployment(arg1 context.Context, arg2 *apigateway.CreateDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDeploymentOutput, error)
	CreateDocumentationPart(arg1 context.Context, arg2 *apigateway.CreateDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDocumentationPartOutput, error)
	CreateDocumentationVersion(arg1 context.Context, arg2 *apigateway.CreateDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDocumentationVersionOutput, error)
	// Creates a new domain name.
	//
	CreateDomainName(arg1 context.Context, arg2 *apigateway.CreateDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDomainNameOutput, error)
	// Adds a new Model resource to an existing RestApi resource.
	//
	CreateModel(arg1 context.Context, arg2 *apigateway.CreateModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateModelOutput, error)
	// Creates a ReqeustValidator of a given RestApi.
	//
	CreateRequestValidator(arg1 context.Context, arg2 *apigateway.CreateRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateRequestValidatorOutput, error)
	// Creates a Resource resource.
	//
	CreateResource(arg1 context.Context, arg2 *apigateway.CreateResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateResourceOutput, error)
	// Creates a new RestApi resource.
	//
	CreateRestApi(arg1 context.Context, arg2 *apigateway.CreateRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateRestApiOutput, error)
	// Creates a new Stage resource that references a pre-existing Deployment for the
	// API.
	//
	CreateStage(arg1 context.Context, arg2 *apigateway.CreateStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateStageOutput, error)
	// Creates a usage plan with the throttle and quota limits, as well as the
	// associated API stages, specified in the payload.
	//
	CreateUsagePlan(arg1 context.Context, arg2 *apigateway.CreateUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateUsagePlanOutput, error)
	// Creates a usage plan key for adding an existing API key to a usage plan.
	//
	CreateUsagePlanKey(arg1 context.Context, arg2 *apigateway.CreateUsagePlanKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateUsagePlanKeyOutput, error)
	// Creates a VPC link, under the caller's account in a selected region, in an
	// asynchronous operation that typically takes 2-4 minutes to complete and become
	// operational. The caller must have permissions to create and update VPC Endpoint
	// services.
	//
	CreateVpcLink(arg1 context.Context, arg2 *apigateway.CreateVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateVpcLinkOutput, error)
	// Deletes the ApiKey resource.
	//
	DeleteApiKey(arg1 context.Context, arg2 *apigateway.DeleteApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteApiKeyOutput, error)
	// Deletes an existing Authorizer resource. AWS CLI
	// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/delete-authorizer.html)
	//
	DeleteAuthorizer(arg1 context.Context, arg2 *apigateway.DeleteAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteAuthorizerOutput, error)
	// Deletes the BasePathMapping resource.
	//
	DeleteBasePathMapping(arg1 context.Context, arg2 *apigateway.DeleteBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteBasePathMappingOutput, error)
	// Deletes the ClientCertificate resource.
	//
	DeleteClientCertificate(arg1 context.Context, arg2 *apigateway.DeleteClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteClientCertificateOutput, error)
	// Deletes a Deployment resource. Deleting a deployment will only succeed if there
	// are no Stage resources associated with it.
	//
	DeleteDeployment(arg1 context.Context, arg2 *apigateway.DeleteDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDeploymentOutput, error)
	DeleteDocumentationPart(arg1 context.Context, arg2 *apigateway.DeleteDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDocumentationPartOutput, error)
	DeleteDocumentationVersion(arg1 context.Context, arg2 *apigateway.DeleteDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDocumentationVersionOutput, error)
	// Deletes the DomainName resource.
	//
	DeleteDomainName(arg1 context.Context, arg2 *apigateway.DeleteDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDomainNameOutput, error)
	// Clears any customization of a GatewayResponse of a specified response type on
	// the given RestApi and resets it with the default settings.
	//
	DeleteGatewayResponse(arg1 context.Context, arg2 *apigateway.DeleteGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteGatewayResponseOutput, error)
	// Represents a delete integration.
	//
	DeleteIntegration(arg1 context.Context, arg2 *apigateway.DeleteIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteIntegrationOutput, error)
	// Represents a delete integration response.
	//
	DeleteIntegrationResponse(arg1 context.Context, arg2 *apigateway.DeleteIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteIntegrationResponseOutput, error)
	// Deletes an existing Method resource.
	//
	DeleteMethod(arg1 context.Context, arg2 *apigateway.DeleteMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteMethodOutput, error)
	// Deletes an existing MethodResponse resource.
	//
	DeleteMethodResponse(arg1 context.Context, arg2 *apigateway.DeleteMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteMethodResponseOutput, error)
	// Deletes a model.
	//
	DeleteModel(arg1 context.Context, arg2 *apigateway.DeleteModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteModelOutput, error)
	// Deletes a RequestValidator of a given RestApi.
	//
	DeleteRequestValidator(arg1 context.Context, arg2 *apigateway.DeleteRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteRequestValidatorOutput, error)
	// Deletes a Resource resource.
	//
	DeleteResource(arg1 context.Context, arg2 *apigateway.DeleteResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteResourceOutput, error)
	// Deletes the specified API.
	//
	DeleteRestApi(arg1 context.Context, arg2 *apigateway.DeleteRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteRestApiOutput, error)
	// Deletes a Stage resource.
	//
	DeleteStage(arg1 context.Context, arg2 *apigateway.DeleteStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteStageOutput, error)
	// Deletes a usage plan of a given plan Id.
	//
	DeleteUsagePlan(arg1 context.Context, arg2 *apigateway.DeleteUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteUsagePlanOutput, error)
	// Deletes a usage plan key and remove the underlying API key from the associated
	// usage plan.
	//
	DeleteUsagePlanKey(arg1 context.Context, arg2 *apigateway.DeleteUsagePlanKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteUsagePlanKeyOutput, error)
	// Deletes an existing VpcLink of a specified identifier.
	//
	DeleteVpcLink(arg1 context.Context, arg2 *apigateway.DeleteVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteVpcLinkOutput, error)
	// Flushes all authorizer cache entries on a stage.
	//
	FlushStageAuthorizersCache(arg1 context.Context, arg2 *apigateway.FlushStageAuthorizersCacheInput, arg3 ...func(*apigateway.Options)) (*apigateway.FlushStageAuthorizersCacheOutput, error)
	// Flushes a stage's cache.
	//
	FlushStageCache(arg1 context.Context, arg2 *apigateway.FlushStageCacheInput, arg3 ...func(*apigateway.Options)) (*apigateway.FlushStageCacheOutput, error)
	// Generates a ClientCertificate resource.
	//
	GenerateClientCertificate(arg1 context.Context, arg2 *apigateway.GenerateClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.GenerateClientCertificateOutput, error)
	// Gets information about the current Account resource.
	//
	GetAccount(arg1 context.Context, arg2 *apigateway.GetAccountInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetAccountOutput, error)
	// Gets information about the current ApiKey resource.
	//
	GetApiKey(arg1 context.Context, arg2 *apigateway.GetApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetApiKeyOutput, error)
	// Gets information about the current ApiKeys resource.
	//
	GetApiKeys(arg1 context.Context, arg2 *apigateway.GetApiKeysInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetApiKeysOutput, error)
	// Describe an existing Authorizer resource. AWS CLI
	// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-authorizer.html)
	//
	GetAuthorizer(arg1 context.Context, arg2 *apigateway.GetAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetAuthorizerOutput, error)
	// Describe an existing Authorizers resource. AWS CLI
	// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-authorizers.html)
	//
	GetAuthorizers(arg1 context.Context, arg2 *apigateway.GetAuthorizersInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetAuthorizersOutput, error)
	// Describe a BasePathMapping resource.
	//
	GetBasePathMapping(arg1 context.Context, arg2 *apigateway.GetBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingOutput, error)
	// Represents a collection of BasePathMapping resources.
	//
	GetBasePathMappings(arg1 context.Context, arg2 *apigateway.GetBasePathMappingsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingsOutput, error)
	// Gets information about the current ClientCertificate resource.
	//
	GetClientCertificate(arg1 context.Context, arg2 *apigateway.GetClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetClientCertificateOutput, error)
	// Gets a collection of ClientCertificate resources.
	//
	GetClientCertificates(arg1 context.Context, arg2 *apigateway.GetClientCertificatesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetClientCertificatesOutput, error)
	// Gets information about a Deployment resource.
	//
	GetDeployment(arg1 context.Context, arg2 *apigateway.GetDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDeploymentOutput, error)
	// Gets information about a Deployments collection.
	//
	GetDeployments(arg1 context.Context, arg2 *apigateway.GetDeploymentsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDeploymentsOutput, error)
	GetDocumentationPart(arg1 context.Context, arg2 *apigateway.GetDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartOutput, error)
	GetDocumentationParts(arg1 context.Context, arg2 *apigateway.GetDocumentationPartsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartsOutput, error)
	GetDocumentationVersion(arg1 context.Context, arg2 *apigateway.GetDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionOutput, error)
	GetDocumentationVersions(arg1 context.Context, arg2 *apigateway.GetDocumentationVersionsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionsOutput, error)
	// Represents a domain name that is contained in a simpler, more intuitive URL that
	// can be called.
	//
	GetDomainName(arg1 context.Context, arg2 *apigateway.GetDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDomainNameOutput, error)
	// Represents a collection of DomainName resources.
	//
	GetDomainNames(arg1 context.Context, arg2 *apigateway.GetDomainNamesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDomainNamesOutput, error)
	// Exports a deployed version of a RestApi in a specified format.
	//
	GetExport(arg1 context.Context, arg2 *apigateway.GetExportInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetExportOutput, error)
	// Gets a GatewayResponse of a specified response type on the given RestApi.
	//
	GetGatewayResponse(arg1 context.Context, arg2 *apigateway.GetGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetGatewayResponseOutput, error)
	// Gets the GatewayResponses collection on the given RestApi. If an API developer
	// has not added any definitions for gateway responses, the result will be the API
	// Gateway-generated default GatewayResponses collection for the supported response
	// types.
	//
	GetGatewayResponses(arg1 context.Context, arg2 *apigateway.GetGatewayResponsesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetGatewayResponsesOutput, error)
	// Get the integration settings.
	//
	GetIntegration(arg1 context.Context, arg2 *apigateway.GetIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetIntegrationOutput, error)
	// Represents a get integration response.
	//
	GetIntegrationResponse(arg1 context.Context, arg2 *apigateway.GetIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetIntegrationResponseOutput, error)
	// Describe an existing Method resource.
	//
	GetMethod(arg1 context.Context, arg2 *apigateway.GetMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetMethodOutput, error)
	// Describes a MethodResponse resource.
	//
	GetMethodResponse(arg1 context.Context, arg2 *apigateway.GetMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetMethodResponseOutput, error)
	// Describes an existing model defined for a RestApi resource.
	//
	GetModel(arg1 context.Context, arg2 *apigateway.GetModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetModelOutput, error)
	// Generates a sample mapping template that can be used to transform a payload into
	// the structure of a model.
	//
	GetModelTemplate(arg1 context.Context, arg2 *apigateway.GetModelTemplateInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetModelTemplateOutput, error)
	// Describes existing Models defined for a RestApi resource.
	//
	GetModels(arg1 context.Context, arg2 *apigateway.GetModelsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetModelsOutput, error)
	// Gets a RequestValidator of a given RestApi.
	//
	GetRequestValidator(arg1 context.Context, arg2 *apigateway.GetRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorOutput, error)
	// Gets the RequestValidators collection of a given RestApi.
	//
	GetRequestValidators(arg1 context.Context, arg2 *apigateway.GetRequestValidatorsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorsOutput, error)
	// Lists information about a resource.
	//
	GetResource(arg1 context.Context, arg2 *apigateway.GetResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetResourceOutput, error)
	// Lists information about a collection of Resource resources.
	//
	GetResources(arg1 context.Context, arg2 *apigateway.GetResourcesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetResourcesOutput, error)
	// Lists the RestApi resource in the collection.
	//
	GetRestApi(arg1 context.Context, arg2 *apigateway.GetRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRestApiOutput, error)
	// Lists the RestApis resources for your collection.
	//
	GetRestApis(arg1 context.Context, arg2 *apigateway.GetRestApisInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRestApisOutput, error)
	// Generates a client SDK for a RestApi and Stage.
	//
	GetSdk(arg1 context.Context, arg2 *apigateway.GetSdkInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetSdkOutput, error)
	GetSdkType(arg1 context.Context, arg2 *apigateway.GetSdkTypeInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetSdkTypeOutput, error)
	GetSdkTypes(arg1 context.Context, arg2 *apigateway.GetSdkTypesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetSdkTypesOutput, error)
	// Gets information about a Stage resource.
	//
	GetStage(arg1 context.Context, arg2 *apigateway.GetStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetStageOutput, error)
	// Gets information about one or more Stage resources.
	//
	GetStages(arg1 context.Context, arg2 *apigateway.GetStagesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetStagesOutput, error)
	// Gets the Tags collection for a given resource.
	//
	GetTags(arg1 context.Context, arg2 *apigateway.GetTagsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetTagsOutput, error)
	// Gets the usage data of a usage plan in a specified time interval.
	//
	GetUsage(arg1 context.Context, arg2 *apigateway.GetUsageInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsageOutput, error)
	// Gets a usage plan of a given plan identifier.
	//
	GetUsagePlan(arg1 context.Context, arg2 *apigateway.GetUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanOutput, error)
	// Gets a usage plan key of a given key identifier.
	//
	GetUsagePlanKey(arg1 context.Context, arg2 *apigateway.GetUsagePlanKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeyOutput, error)
	// Gets all the usage plan keys representing the API keys added to a specified
	// usage plan.
	//
	GetUsagePlanKeys(arg1 context.Context, arg2 *apigateway.GetUsagePlanKeysInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeysOutput, error)
	// Gets all the usage plans of the caller's account.
	//
	GetUsagePlans(arg1 context.Context, arg2 *apigateway.GetUsagePlansInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlansOutput, error)
	// Gets a specified VPC link under the caller's account in a region.
	//
	GetVpcLink(arg1 context.Context, arg2 *apigateway.GetVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetVpcLinkOutput, error)
	// Gets the VpcLinks collection under the caller's account in a selected region.
	//
	GetVpcLinks(arg1 context.Context, arg2 *apigateway.GetVpcLinksInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetVpcLinksOutput, error)
	// Import API keys from an external source, such as a CSV-formatted file.
	//
	ImportApiKeys(arg1 context.Context, arg2 *apigateway.ImportApiKeysInput, arg3 ...func(*apigateway.Options)) (*apigateway.ImportApiKeysOutput, error)
	ImportDocumentationParts(arg1 context.Context, arg2 *apigateway.ImportDocumentationPartsInput, arg3 ...func(*apigateway.Options)) (*apigateway.ImportDocumentationPartsOutput, error)
	// A feature of the API Gateway control service for creating a new API from an
	// external API definition file.
	//
	ImportRestApi(arg1 context.Context, arg2 *apigateway.ImportRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.ImportRestApiOutput, error)
	// Creates a customization of a GatewayResponse of a specified response type and
	// status code on the given RestApi.
	//
	PutGatewayResponse(arg1 context.Context, arg2 *apigateway.PutGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutGatewayResponseOutput, error)
	// Sets up a method's integration.
	//
	PutIntegration(arg1 context.Context, arg2 *apigateway.PutIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutIntegrationOutput, error)
	// Represents a put integration.
	//
	PutIntegrationResponse(arg1 context.Context, arg2 *apigateway.PutIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutIntegrationResponseOutput, error)
	// Add a method to an existing Resource resource.
	//
	PutMethod(arg1 context.Context, arg2 *apigateway.PutMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutMethodOutput, error)
	// Adds a MethodResponse to an existing Method resource.
	//
	PutMethodResponse(arg1 context.Context, arg2 *apigateway.PutMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutMethodResponseOutput, error)
	// A feature of the API Gateway control service for updating an existing API with
	// an input of external API definitions. The update can take the form of merging
	// the supplied definition into the existing API or overwriting the existing API.
	//
	PutRestApi(arg1 context.Context, arg2 *apigateway.PutRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutRestApiOutput, error)
	// Adds or updates a tag on a given resource.
	//
	TagResource(arg1 context.Context, arg2 *apigateway.TagResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.TagResourceOutput, error)
	// Simulate the execution of an Authorizer in your RestApi with headers,
	// parameters, and an incoming request body. Use Lambda Function as Authorizer
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html)Use
	// Cognito User Pool as Authorizer
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html)
	//
	TestInvokeAuthorizer(arg1 context.Context, arg2 *apigateway.TestInvokeAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.TestInvokeAuthorizerOutput, error)
	// Simulate the execution of a Method in your RestApi with headers, parameters, and
	// an incoming request body.
	//
	TestInvokeMethod(arg1 context.Context, arg2 *apigateway.TestInvokeMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.TestInvokeMethodOutput, error)
	// Removes a tag from a given resource.
	//
	UntagResource(arg1 context.Context, arg2 *apigateway.UntagResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.UntagResourceOutput, error)
	// Changes information about the current Account resource.
	//
	UpdateAccount(arg1 context.Context, arg2 *apigateway.UpdateAccountInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateAccountOutput, error)
	// Changes information about an ApiKey resource.
	//
	UpdateApiKey(arg1 context.Context, arg2 *apigateway.UpdateApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateApiKeyOutput, error)
	// Updates an existing Authorizer resource. AWS CLI
	// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/update-authorizer.html)
	//
	UpdateAuthorizer(arg1 context.Context, arg2 *apigateway.UpdateAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateAuthorizerOutput, error)
	// Changes information about the BasePathMapping resource.
	//
	UpdateBasePathMapping(arg1 context.Context, arg2 *apigateway.UpdateBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateBasePathMappingOutput, error)
	// Changes information about an ClientCertificate resource.
	//
	UpdateClientCertificate(arg1 context.Context, arg2 *apigateway.UpdateClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateClientCertificateOutput, error)
	// Changes information about a Deployment resource.
	//
	UpdateDeployment(arg1 context.Context, arg2 *apigateway.UpdateDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDeploymentOutput, error)
	UpdateDocumentationPart(arg1 context.Context, arg2 *apigateway.UpdateDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDocumentationPartOutput, error)
	UpdateDocumentationVersion(arg1 context.Context, arg2 *apigateway.UpdateDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDocumentationVersionOutput, error)
	// Changes information about the DomainName resource.
	//
	UpdateDomainName(arg1 context.Context, arg2 *apigateway.UpdateDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDomainNameOutput, error)
	// Updates a GatewayResponse of a specified response type on the given RestApi.
	//
	UpdateGatewayResponse(arg1 context.Context, arg2 *apigateway.UpdateGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateGatewayResponseOutput, error)
	// Represents an update integration.
	//
	UpdateIntegration(arg1 context.Context, arg2 *apigateway.UpdateIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateIntegrationOutput, error)
	// Represents an update integration response.
	//
	UpdateIntegrationResponse(arg1 context.Context, arg2 *apigateway.UpdateIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateIntegrationResponseOutput, error)
	// Updates an existing Method resource.
	//
	UpdateMethod(arg1 context.Context, arg2 *apigateway.UpdateMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateMethodOutput, error)
	// Updates an existing MethodResponse resource.
	//
	UpdateMethodResponse(arg1 context.Context, arg2 *apigateway.UpdateMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateMethodResponseOutput, error)
	// Changes information about a model.
	//
	UpdateModel(arg1 context.Context, arg2 *apigateway.UpdateModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateModelOutput, error)
	// Updates a RequestValidator of a given RestApi.
	//
	UpdateRequestValidator(arg1 context.Context, arg2 *apigateway.UpdateRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateRequestValidatorOutput, error)
	// Changes information about a Resource resource.
	//
	UpdateResource(arg1 context.Context, arg2 *apigateway.UpdateResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateResourceOutput, error)
	// Changes information about the specified API.
	//
	UpdateRestApi(arg1 context.Context, arg2 *apigateway.UpdateRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateRestApiOutput, error)
	// Changes information about a Stage resource.
	//
	UpdateStage(arg1 context.Context, arg2 *apigateway.UpdateStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateStageOutput, error)
	// Grants a temporary extension to the remaining quota of a usage plan associated
	// with a specified API key.
	//
	UpdateUsage(arg1 context.Context, arg2 *apigateway.UpdateUsageInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateUsageOutput, error)
	// Updates a usage plan of a given plan Id.
	//
	UpdateUsagePlan(arg1 context.Context, arg2 *apigateway.UpdateUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateUsagePlanOutput, error)
	// Updates an existing VpcLink of a specified identifier.
	//
	UpdateVpcLink(arg1 context.Context, arg2 *apigateway.UpdateVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateVpcLinkOutput, error)
}

// APIGatewayClientMock generic client mock
type APIGatewayClientMock struct {
	CreateApiKeyMock               func(arg1 context.Context, arg2 *apigateway.CreateApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateApiKeyOutput, error)
	CreateAuthorizerMock           func(arg1 context.Context, arg2 *apigateway.CreateAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateAuthorizerOutput, error)
	CreateBasePathMappingMock      func(arg1 context.Context, arg2 *apigateway.CreateBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateBasePathMappingOutput, error)
	CreateDeploymentMock           func(arg1 context.Context, arg2 *apigateway.CreateDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDeploymentOutput, error)
	CreateDocumentationPartMock    func(arg1 context.Context, arg2 *apigateway.CreateDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDocumentationPartOutput, error)
	CreateDocumentationVersionMock func(arg1 context.Context, arg2 *apigateway.CreateDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDocumentationVersionOutput, error)
	CreateDomainNameMock           func(arg1 context.Context, arg2 *apigateway.CreateDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDomainNameOutput, error)
	CreateModelMock                func(arg1 context.Context, arg2 *apigateway.CreateModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateModelOutput, error)
	CreateRequestValidatorMock     func(arg1 context.Context, arg2 *apigateway.CreateRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateRequestValidatorOutput, error)
	CreateResourceMock             func(arg1 context.Context, arg2 *apigateway.CreateResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateResourceOutput, error)
	CreateRestApiMock              func(arg1 context.Context, arg2 *apigateway.CreateRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateRestApiOutput, error)
	CreateStageMock                func(arg1 context.Context, arg2 *apigateway.CreateStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateStageOutput, error)
	CreateUsagePlanMock            func(arg1 context.Context, arg2 *apigateway.CreateUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateUsagePlanOutput, error)
	CreateUsagePlanKeyMock         func(arg1 context.Context, arg2 *apigateway.CreateUsagePlanKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateUsagePlanKeyOutput, error)
	CreateVpcLinkMock              func(arg1 context.Context, arg2 *apigateway.CreateVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateVpcLinkOutput, error)
	DeleteApiKeyMock               func(arg1 context.Context, arg2 *apigateway.DeleteApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteApiKeyOutput, error)
	DeleteAuthorizerMock           func(arg1 context.Context, arg2 *apigateway.DeleteAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteAuthorizerOutput, error)
	DeleteBasePathMappingMock      func(arg1 context.Context, arg2 *apigateway.DeleteBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteBasePathMappingOutput, error)
	DeleteClientCertificateMock    func(arg1 context.Context, arg2 *apigateway.DeleteClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteClientCertificateOutput, error)
	DeleteDeploymentMock           func(arg1 context.Context, arg2 *apigateway.DeleteDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDeploymentOutput, error)
	DeleteDocumentationPartMock    func(arg1 context.Context, arg2 *apigateway.DeleteDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDocumentationPartOutput, error)
	DeleteDocumentationVersionMock func(arg1 context.Context, arg2 *apigateway.DeleteDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDocumentationVersionOutput, error)
	DeleteDomainNameMock           func(arg1 context.Context, arg2 *apigateway.DeleteDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDomainNameOutput, error)
	DeleteGatewayResponseMock      func(arg1 context.Context, arg2 *apigateway.DeleteGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteGatewayResponseOutput, error)
	DeleteIntegrationMock          func(arg1 context.Context, arg2 *apigateway.DeleteIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteIntegrationOutput, error)
	DeleteIntegrationResponseMock  func(arg1 context.Context, arg2 *apigateway.DeleteIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteIntegrationResponseOutput, error)
	DeleteMethodMock               func(arg1 context.Context, arg2 *apigateway.DeleteMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteMethodOutput, error)
	DeleteMethodResponseMock       func(arg1 context.Context, arg2 *apigateway.DeleteMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteMethodResponseOutput, error)
	DeleteModelMock                func(arg1 context.Context, arg2 *apigateway.DeleteModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteModelOutput, error)
	DeleteRequestValidatorMock     func(arg1 context.Context, arg2 *apigateway.DeleteRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteRequestValidatorOutput, error)
	DeleteResourceMock             func(arg1 context.Context, arg2 *apigateway.DeleteResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteResourceOutput, error)
	DeleteRestApiMock              func(arg1 context.Context, arg2 *apigateway.DeleteRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteRestApiOutput, error)
	DeleteStageMock                func(arg1 context.Context, arg2 *apigateway.DeleteStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteStageOutput, error)
	DeleteUsagePlanMock            func(arg1 context.Context, arg2 *apigateway.DeleteUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteUsagePlanOutput, error)
	DeleteUsagePlanKeyMock         func(arg1 context.Context, arg2 *apigateway.DeleteUsagePlanKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteUsagePlanKeyOutput, error)
	DeleteVpcLinkMock              func(arg1 context.Context, arg2 *apigateway.DeleteVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteVpcLinkOutput, error)
	FlushStageAuthorizersCacheMock func(arg1 context.Context, arg2 *apigateway.FlushStageAuthorizersCacheInput, arg3 ...func(*apigateway.Options)) (*apigateway.FlushStageAuthorizersCacheOutput, error)
	FlushStageCacheMock            func(arg1 context.Context, arg2 *apigateway.FlushStageCacheInput, arg3 ...func(*apigateway.Options)) (*apigateway.FlushStageCacheOutput, error)
	GenerateClientCertificateMock  func(arg1 context.Context, arg2 *apigateway.GenerateClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.GenerateClientCertificateOutput, error)
	GetAccountMock                 func(arg1 context.Context, arg2 *apigateway.GetAccountInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetAccountOutput, error)
	GetApiKeyMock                  func(arg1 context.Context, arg2 *apigateway.GetApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetApiKeyOutput, error)
	GetApiKeysMock                 func(arg1 context.Context, arg2 *apigateway.GetApiKeysInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetApiKeysOutput, error)
	GetAuthorizerMock              func(arg1 context.Context, arg2 *apigateway.GetAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetAuthorizerOutput, error)
	GetAuthorizersMock             func(arg1 context.Context, arg2 *apigateway.GetAuthorizersInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetAuthorizersOutput, error)
	GetBasePathMappingMock         func(arg1 context.Context, arg2 *apigateway.GetBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingOutput, error)
	GetBasePathMappingsMock        func(arg1 context.Context, arg2 *apigateway.GetBasePathMappingsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingsOutput, error)
	GetClientCertificateMock       func(arg1 context.Context, arg2 *apigateway.GetClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetClientCertificateOutput, error)
	GetClientCertificatesMock      func(arg1 context.Context, arg2 *apigateway.GetClientCertificatesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetClientCertificatesOutput, error)
	GetDeploymentMock              func(arg1 context.Context, arg2 *apigateway.GetDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDeploymentOutput, error)
	GetDeploymentsMock             func(arg1 context.Context, arg2 *apigateway.GetDeploymentsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDeploymentsOutput, error)
	GetDocumentationPartMock       func(arg1 context.Context, arg2 *apigateway.GetDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartOutput, error)
	GetDocumentationPartsMock      func(arg1 context.Context, arg2 *apigateway.GetDocumentationPartsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartsOutput, error)
	GetDocumentationVersionMock    func(arg1 context.Context, arg2 *apigateway.GetDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionOutput, error)
	GetDocumentationVersionsMock   func(arg1 context.Context, arg2 *apigateway.GetDocumentationVersionsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionsOutput, error)
	GetDomainNameMock              func(arg1 context.Context, arg2 *apigateway.GetDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDomainNameOutput, error)
	GetDomainNamesMock             func(arg1 context.Context, arg2 *apigateway.GetDomainNamesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDomainNamesOutput, error)
	GetExportMock                  func(arg1 context.Context, arg2 *apigateway.GetExportInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetExportOutput, error)
	GetGatewayResponseMock         func(arg1 context.Context, arg2 *apigateway.GetGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetGatewayResponseOutput, error)
	GetGatewayResponsesMock        func(arg1 context.Context, arg2 *apigateway.GetGatewayResponsesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetGatewayResponsesOutput, error)
	GetIntegrationMock             func(arg1 context.Context, arg2 *apigateway.GetIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetIntegrationOutput, error)
	GetIntegrationResponseMock     func(arg1 context.Context, arg2 *apigateway.GetIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetIntegrationResponseOutput, error)
	GetMethodMock                  func(arg1 context.Context, arg2 *apigateway.GetMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetMethodOutput, error)
	GetMethodResponseMock          func(arg1 context.Context, arg2 *apigateway.GetMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetMethodResponseOutput, error)
	GetModelMock                   func(arg1 context.Context, arg2 *apigateway.GetModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetModelOutput, error)
	GetModelTemplateMock           func(arg1 context.Context, arg2 *apigateway.GetModelTemplateInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetModelTemplateOutput, error)
	GetModelsMock                  func(arg1 context.Context, arg2 *apigateway.GetModelsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetModelsOutput, error)
	GetRequestValidatorMock        func(arg1 context.Context, arg2 *apigateway.GetRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorOutput, error)
	GetRequestValidatorsMock       func(arg1 context.Context, arg2 *apigateway.GetRequestValidatorsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorsOutput, error)
	GetResourceMock                func(arg1 context.Context, arg2 *apigateway.GetResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetResourceOutput, error)
	GetResourcesMock               func(arg1 context.Context, arg2 *apigateway.GetResourcesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetResourcesOutput, error)
	GetRestApiMock                 func(arg1 context.Context, arg2 *apigateway.GetRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRestApiOutput, error)
	GetRestApisMock                func(arg1 context.Context, arg2 *apigateway.GetRestApisInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRestApisOutput, error)
	GetSdkMock                     func(arg1 context.Context, arg2 *apigateway.GetSdkInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetSdkOutput, error)
	GetSdkTypeMock                 func(arg1 context.Context, arg2 *apigateway.GetSdkTypeInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetSdkTypeOutput, error)
	GetSdkTypesMock                func(arg1 context.Context, arg2 *apigateway.GetSdkTypesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetSdkTypesOutput, error)
	GetStageMock                   func(arg1 context.Context, arg2 *apigateway.GetStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetStageOutput, error)
	GetStagesMock                  func(arg1 context.Context, arg2 *apigateway.GetStagesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetStagesOutput, error)
	GetTagsMock                    func(arg1 context.Context, arg2 *apigateway.GetTagsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetTagsOutput, error)
	GetUsageMock                   func(arg1 context.Context, arg2 *apigateway.GetUsageInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsageOutput, error)
	GetUsagePlanMock               func(arg1 context.Context, arg2 *apigateway.GetUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanOutput, error)
	GetUsagePlanKeyMock            func(arg1 context.Context, arg2 *apigateway.GetUsagePlanKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeyOutput, error)
	GetUsagePlanKeysMock           func(arg1 context.Context, arg2 *apigateway.GetUsagePlanKeysInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeysOutput, error)
	GetUsagePlansMock              func(arg1 context.Context, arg2 *apigateway.GetUsagePlansInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlansOutput, error)
	GetVpcLinkMock                 func(arg1 context.Context, arg2 *apigateway.GetVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetVpcLinkOutput, error)
	GetVpcLinksMock                func(arg1 context.Context, arg2 *apigateway.GetVpcLinksInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetVpcLinksOutput, error)
	ImportApiKeysMock              func(arg1 context.Context, arg2 *apigateway.ImportApiKeysInput, arg3 ...func(*apigateway.Options)) (*apigateway.ImportApiKeysOutput, error)
	ImportDocumentationPartsMock   func(arg1 context.Context, arg2 *apigateway.ImportDocumentationPartsInput, arg3 ...func(*apigateway.Options)) (*apigateway.ImportDocumentationPartsOutput, error)
	ImportRestApiMock              func(arg1 context.Context, arg2 *apigateway.ImportRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.ImportRestApiOutput, error)
	PutGatewayResponseMock         func(arg1 context.Context, arg2 *apigateway.PutGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutGatewayResponseOutput, error)
	PutIntegrationMock             func(arg1 context.Context, arg2 *apigateway.PutIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutIntegrationOutput, error)
	PutIntegrationResponseMock     func(arg1 context.Context, arg2 *apigateway.PutIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutIntegrationResponseOutput, error)
	PutMethodMock                  func(arg1 context.Context, arg2 *apigateway.PutMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutMethodOutput, error)
	PutMethodResponseMock          func(arg1 context.Context, arg2 *apigateway.PutMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutMethodResponseOutput, error)
	PutRestApiMock                 func(arg1 context.Context, arg2 *apigateway.PutRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutRestApiOutput, error)
	TagResourceMock                func(arg1 context.Context, arg2 *apigateway.TagResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.TagResourceOutput, error)
	TestInvokeAuthorizerMock       func(arg1 context.Context, arg2 *apigateway.TestInvokeAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.TestInvokeAuthorizerOutput, error)
	TestInvokeMethodMock           func(arg1 context.Context, arg2 *apigateway.TestInvokeMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.TestInvokeMethodOutput, error)
	UntagResourceMock              func(arg1 context.Context, arg2 *apigateway.UntagResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.UntagResourceOutput, error)
	UpdateAccountMock              func(arg1 context.Context, arg2 *apigateway.UpdateAccountInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateAccountOutput, error)
	UpdateApiKeyMock               func(arg1 context.Context, arg2 *apigateway.UpdateApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateApiKeyOutput, error)
	UpdateAuthorizerMock           func(arg1 context.Context, arg2 *apigateway.UpdateAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateAuthorizerOutput, error)
	UpdateBasePathMappingMock      func(arg1 context.Context, arg2 *apigateway.UpdateBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateBasePathMappingOutput, error)
	UpdateClientCertificateMock    func(arg1 context.Context, arg2 *apigateway.UpdateClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateClientCertificateOutput, error)
	UpdateDeploymentMock           func(arg1 context.Context, arg2 *apigateway.UpdateDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDeploymentOutput, error)
	UpdateDocumentationPartMock    func(arg1 context.Context, arg2 *apigateway.UpdateDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDocumentationPartOutput, error)
	UpdateDocumentationVersionMock func(arg1 context.Context, arg2 *apigateway.UpdateDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDocumentationVersionOutput, error)
	UpdateDomainNameMock           func(arg1 context.Context, arg2 *apigateway.UpdateDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDomainNameOutput, error)
	UpdateGatewayResponseMock      func(arg1 context.Context, arg2 *apigateway.UpdateGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateGatewayResponseOutput, error)
	UpdateIntegrationMock          func(arg1 context.Context, arg2 *apigateway.UpdateIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateIntegrationOutput, error)
	UpdateIntegrationResponseMock  func(arg1 context.Context, arg2 *apigateway.UpdateIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateIntegrationResponseOutput, error)
	UpdateMethodMock               func(arg1 context.Context, arg2 *apigateway.UpdateMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateMethodOutput, error)
	UpdateMethodResponseMock       func(arg1 context.Context, arg2 *apigateway.UpdateMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateMethodResponseOutput, error)
	UpdateModelMock                func(arg1 context.Context, arg2 *apigateway.UpdateModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateModelOutput, error)
	UpdateRequestValidatorMock     func(arg1 context.Context, arg2 *apigateway.UpdateRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateRequestValidatorOutput, error)
	UpdateResourceMock             func(arg1 context.Context, arg2 *apigateway.UpdateResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateResourceOutput, error)
	UpdateRestApiMock              func(arg1 context.Context, arg2 *apigateway.UpdateRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateRestApiOutput, error)
	UpdateStageMock                func(arg1 context.Context, arg2 *apigateway.UpdateStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateStageOutput, error)
	UpdateUsageMock                func(arg1 context.Context, arg2 *apigateway.UpdateUsageInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateUsageOutput, error)
	UpdateUsagePlanMock            func(arg1 context.Context, arg2 *apigateway.UpdateUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateUsagePlanOutput, error)
	UpdateVpcLinkMock              func(arg1 context.Context, arg2 *apigateway.UpdateVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateVpcLinkOutput, error)
}

// CreateApiKey mocked funcition
func (m APIGatewayClientMock) CreateApiKey(arg1 context.Context, arg2 *apigateway.CreateApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateApiKeyOutput, error) {
	return m.CreateApiKeyMock(arg1, arg2, arg3...)
}

// CreateAuthorizer mocked funcition
func (m APIGatewayClientMock) CreateAuthorizer(arg1 context.Context, arg2 *apigateway.CreateAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateAuthorizerOutput, error) {
	return m.CreateAuthorizerMock(arg1, arg2, arg3...)
}

// CreateBasePathMapping mocked funcition
func (m APIGatewayClientMock) CreateBasePathMapping(arg1 context.Context, arg2 *apigateway.CreateBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateBasePathMappingOutput, error) {
	return m.CreateBasePathMappingMock(arg1, arg2, arg3...)
}

// CreateDeployment mocked funcition
func (m APIGatewayClientMock) CreateDeployment(arg1 context.Context, arg2 *apigateway.CreateDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDeploymentOutput, error) {
	return m.CreateDeploymentMock(arg1, arg2, arg3...)
}

// CreateDocumentationPart mocked funcition
func (m APIGatewayClientMock) CreateDocumentationPart(arg1 context.Context, arg2 *apigateway.CreateDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDocumentationPartOutput, error) {
	return m.CreateDocumentationPartMock(arg1, arg2, arg3...)
}

// CreateDocumentationVersion mocked funcition
func (m APIGatewayClientMock) CreateDocumentationVersion(arg1 context.Context, arg2 *apigateway.CreateDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDocumentationVersionOutput, error) {
	return m.CreateDocumentationVersionMock(arg1, arg2, arg3...)
}

// CreateDomainName mocked funcition
func (m APIGatewayClientMock) CreateDomainName(arg1 context.Context, arg2 *apigateway.CreateDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateDomainNameOutput, error) {
	return m.CreateDomainNameMock(arg1, arg2, arg3...)
}

// CreateModel mocked funcition
func (m APIGatewayClientMock) CreateModel(arg1 context.Context, arg2 *apigateway.CreateModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateModelOutput, error) {
	return m.CreateModelMock(arg1, arg2, arg3...)
}

// CreateRequestValidator mocked funcition
func (m APIGatewayClientMock) CreateRequestValidator(arg1 context.Context, arg2 *apigateway.CreateRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateRequestValidatorOutput, error) {
	return m.CreateRequestValidatorMock(arg1, arg2, arg3...)
}

// CreateResource mocked funcition
func (m APIGatewayClientMock) CreateResource(arg1 context.Context, arg2 *apigateway.CreateResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateResourceOutput, error) {
	return m.CreateResourceMock(arg1, arg2, arg3...)
}

// CreateRestApi mocked funcition
func (m APIGatewayClientMock) CreateRestApi(arg1 context.Context, arg2 *apigateway.CreateRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateRestApiOutput, error) {
	return m.CreateRestApiMock(arg1, arg2, arg3...)
}

// CreateStage mocked funcition
func (m APIGatewayClientMock) CreateStage(arg1 context.Context, arg2 *apigateway.CreateStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateStageOutput, error) {
	return m.CreateStageMock(arg1, arg2, arg3...)
}

// CreateUsagePlan mocked funcition
func (m APIGatewayClientMock) CreateUsagePlan(arg1 context.Context, arg2 *apigateway.CreateUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateUsagePlanOutput, error) {
	return m.CreateUsagePlanMock(arg1, arg2, arg3...)
}

// CreateUsagePlanKey mocked funcition
func (m APIGatewayClientMock) CreateUsagePlanKey(arg1 context.Context, arg2 *apigateway.CreateUsagePlanKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateUsagePlanKeyOutput, error) {
	return m.CreateUsagePlanKeyMock(arg1, arg2, arg3...)
}

// CreateVpcLink mocked funcition
func (m APIGatewayClientMock) CreateVpcLink(arg1 context.Context, arg2 *apigateway.CreateVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.CreateVpcLinkOutput, error) {
	return m.CreateVpcLinkMock(arg1, arg2, arg3...)
}

// DeleteApiKey mocked funcition
func (m APIGatewayClientMock) DeleteApiKey(arg1 context.Context, arg2 *apigateway.DeleteApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteApiKeyOutput, error) {
	return m.DeleteApiKeyMock(arg1, arg2, arg3...)
}

// DeleteAuthorizer mocked funcition
func (m APIGatewayClientMock) DeleteAuthorizer(arg1 context.Context, arg2 *apigateway.DeleteAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteAuthorizerOutput, error) {
	return m.DeleteAuthorizerMock(arg1, arg2, arg3...)
}

// DeleteBasePathMapping mocked funcition
func (m APIGatewayClientMock) DeleteBasePathMapping(arg1 context.Context, arg2 *apigateway.DeleteBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteBasePathMappingOutput, error) {
	return m.DeleteBasePathMappingMock(arg1, arg2, arg3...)
}

// DeleteClientCertificate mocked funcition
func (m APIGatewayClientMock) DeleteClientCertificate(arg1 context.Context, arg2 *apigateway.DeleteClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteClientCertificateOutput, error) {
	return m.DeleteClientCertificateMock(arg1, arg2, arg3...)
}

// DeleteDeployment mocked funcition
func (m APIGatewayClientMock) DeleteDeployment(arg1 context.Context, arg2 *apigateway.DeleteDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDeploymentOutput, error) {
	return m.DeleteDeploymentMock(arg1, arg2, arg3...)
}

// DeleteDocumentationPart mocked funcition
func (m APIGatewayClientMock) DeleteDocumentationPart(arg1 context.Context, arg2 *apigateway.DeleteDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDocumentationPartOutput, error) {
	return m.DeleteDocumentationPartMock(arg1, arg2, arg3...)
}

// DeleteDocumentationVersion mocked funcition
func (m APIGatewayClientMock) DeleteDocumentationVersion(arg1 context.Context, arg2 *apigateway.DeleteDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDocumentationVersionOutput, error) {
	return m.DeleteDocumentationVersionMock(arg1, arg2, arg3...)
}

// DeleteDomainName mocked funcition
func (m APIGatewayClientMock) DeleteDomainName(arg1 context.Context, arg2 *apigateway.DeleteDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteDomainNameOutput, error) {
	return m.DeleteDomainNameMock(arg1, arg2, arg3...)
}

// DeleteGatewayResponse mocked funcition
func (m APIGatewayClientMock) DeleteGatewayResponse(arg1 context.Context, arg2 *apigateway.DeleteGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteGatewayResponseOutput, error) {
	return m.DeleteGatewayResponseMock(arg1, arg2, arg3...)
}

// DeleteIntegration mocked funcition
func (m APIGatewayClientMock) DeleteIntegration(arg1 context.Context, arg2 *apigateway.DeleteIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteIntegrationOutput, error) {
	return m.DeleteIntegrationMock(arg1, arg2, arg3...)
}

// DeleteIntegrationResponse mocked funcition
func (m APIGatewayClientMock) DeleteIntegrationResponse(arg1 context.Context, arg2 *apigateway.DeleteIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteIntegrationResponseOutput, error) {
	return m.DeleteIntegrationResponseMock(arg1, arg2, arg3...)
}

// DeleteMethod mocked funcition
func (m APIGatewayClientMock) DeleteMethod(arg1 context.Context, arg2 *apigateway.DeleteMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteMethodOutput, error) {
	return m.DeleteMethodMock(arg1, arg2, arg3...)
}

// DeleteMethodResponse mocked funcition
func (m APIGatewayClientMock) DeleteMethodResponse(arg1 context.Context, arg2 *apigateway.DeleteMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteMethodResponseOutput, error) {
	return m.DeleteMethodResponseMock(arg1, arg2, arg3...)
}

// DeleteModel mocked funcition
func (m APIGatewayClientMock) DeleteModel(arg1 context.Context, arg2 *apigateway.DeleteModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteModelOutput, error) {
	return m.DeleteModelMock(arg1, arg2, arg3...)
}

// DeleteRequestValidator mocked funcition
func (m APIGatewayClientMock) DeleteRequestValidator(arg1 context.Context, arg2 *apigateway.DeleteRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteRequestValidatorOutput, error) {
	return m.DeleteRequestValidatorMock(arg1, arg2, arg3...)
}

// DeleteResource mocked funcition
func (m APIGatewayClientMock) DeleteResource(arg1 context.Context, arg2 *apigateway.DeleteResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteResourceOutput, error) {
	return m.DeleteResourceMock(arg1, arg2, arg3...)
}

// DeleteRestApi mocked funcition
func (m APIGatewayClientMock) DeleteRestApi(arg1 context.Context, arg2 *apigateway.DeleteRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteRestApiOutput, error) {
	return m.DeleteRestApiMock(arg1, arg2, arg3...)
}

// DeleteStage mocked funcition
func (m APIGatewayClientMock) DeleteStage(arg1 context.Context, arg2 *apigateway.DeleteStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteStageOutput, error) {
	return m.DeleteStageMock(arg1, arg2, arg3...)
}

// DeleteUsagePlan mocked funcition
func (m APIGatewayClientMock) DeleteUsagePlan(arg1 context.Context, arg2 *apigateway.DeleteUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteUsagePlanOutput, error) {
	return m.DeleteUsagePlanMock(arg1, arg2, arg3...)
}

// DeleteUsagePlanKey mocked funcition
func (m APIGatewayClientMock) DeleteUsagePlanKey(arg1 context.Context, arg2 *apigateway.DeleteUsagePlanKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteUsagePlanKeyOutput, error) {
	return m.DeleteUsagePlanKeyMock(arg1, arg2, arg3...)
}

// DeleteVpcLink mocked funcition
func (m APIGatewayClientMock) DeleteVpcLink(arg1 context.Context, arg2 *apigateway.DeleteVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.DeleteVpcLinkOutput, error) {
	return m.DeleteVpcLinkMock(arg1, arg2, arg3...)
}

// FlushStageAuthorizersCache mocked funcition
func (m APIGatewayClientMock) FlushStageAuthorizersCache(arg1 context.Context, arg2 *apigateway.FlushStageAuthorizersCacheInput, arg3 ...func(*apigateway.Options)) (*apigateway.FlushStageAuthorizersCacheOutput, error) {
	return m.FlushStageAuthorizersCacheMock(arg1, arg2, arg3...)
}

// FlushStageCache mocked funcition
func (m APIGatewayClientMock) FlushStageCache(arg1 context.Context, arg2 *apigateway.FlushStageCacheInput, arg3 ...func(*apigateway.Options)) (*apigateway.FlushStageCacheOutput, error) {
	return m.FlushStageCacheMock(arg1, arg2, arg3...)
}

// GenerateClientCertificate mocked funcition
func (m APIGatewayClientMock) GenerateClientCertificate(arg1 context.Context, arg2 *apigateway.GenerateClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.GenerateClientCertificateOutput, error) {
	return m.GenerateClientCertificateMock(arg1, arg2, arg3...)
}

// GetAccount mocked funcition
func (m APIGatewayClientMock) GetAccount(arg1 context.Context, arg2 *apigateway.GetAccountInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetAccountOutput, error) {
	return m.GetAccountMock(arg1, arg2, arg3...)
}

// GetApiKey mocked funcition
func (m APIGatewayClientMock) GetApiKey(arg1 context.Context, arg2 *apigateway.GetApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetApiKeyOutput, error) {
	return m.GetApiKeyMock(arg1, arg2, arg3...)
}

// GetApiKeys mocked funcition
func (m APIGatewayClientMock) GetApiKeys(arg1 context.Context, arg2 *apigateway.GetApiKeysInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetApiKeysOutput, error) {
	return m.GetApiKeysMock(arg1, arg2, arg3...)
}

// GetAuthorizer mocked funcition
func (m APIGatewayClientMock) GetAuthorizer(arg1 context.Context, arg2 *apigateway.GetAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetAuthorizerOutput, error) {
	return m.GetAuthorizerMock(arg1, arg2, arg3...)
}

// GetAuthorizers mocked funcition
func (m APIGatewayClientMock) GetAuthorizers(arg1 context.Context, arg2 *apigateway.GetAuthorizersInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetAuthorizersOutput, error) {
	return m.GetAuthorizersMock(arg1, arg2, arg3...)
}

// GetBasePathMapping mocked funcition
func (m APIGatewayClientMock) GetBasePathMapping(arg1 context.Context, arg2 *apigateway.GetBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingOutput, error) {
	return m.GetBasePathMappingMock(arg1, arg2, arg3...)
}

// GetBasePathMappings mocked funcition
func (m APIGatewayClientMock) GetBasePathMappings(arg1 context.Context, arg2 *apigateway.GetBasePathMappingsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetBasePathMappingsOutput, error) {
	return m.GetBasePathMappingsMock(arg1, arg2, arg3...)
}

// GetClientCertificate mocked funcition
func (m APIGatewayClientMock) GetClientCertificate(arg1 context.Context, arg2 *apigateway.GetClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetClientCertificateOutput, error) {
	return m.GetClientCertificateMock(arg1, arg2, arg3...)
}

// GetClientCertificates mocked funcition
func (m APIGatewayClientMock) GetClientCertificates(arg1 context.Context, arg2 *apigateway.GetClientCertificatesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetClientCertificatesOutput, error) {
	return m.GetClientCertificatesMock(arg1, arg2, arg3...)
}

// GetDeployment mocked funcition
func (m APIGatewayClientMock) GetDeployment(arg1 context.Context, arg2 *apigateway.GetDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDeploymentOutput, error) {
	return m.GetDeploymentMock(arg1, arg2, arg3...)
}

// GetDeployments mocked funcition
func (m APIGatewayClientMock) GetDeployments(arg1 context.Context, arg2 *apigateway.GetDeploymentsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDeploymentsOutput, error) {
	return m.GetDeploymentsMock(arg1, arg2, arg3...)
}

// GetDocumentationPart mocked funcition
func (m APIGatewayClientMock) GetDocumentationPart(arg1 context.Context, arg2 *apigateway.GetDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartOutput, error) {
	return m.GetDocumentationPartMock(arg1, arg2, arg3...)
}

// GetDocumentationParts mocked funcition
func (m APIGatewayClientMock) GetDocumentationParts(arg1 context.Context, arg2 *apigateway.GetDocumentationPartsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationPartsOutput, error) {
	return m.GetDocumentationPartsMock(arg1, arg2, arg3...)
}

// GetDocumentationVersion mocked funcition
func (m APIGatewayClientMock) GetDocumentationVersion(arg1 context.Context, arg2 *apigateway.GetDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionOutput, error) {
	return m.GetDocumentationVersionMock(arg1, arg2, arg3...)
}

// GetDocumentationVersions mocked funcition
func (m APIGatewayClientMock) GetDocumentationVersions(arg1 context.Context, arg2 *apigateway.GetDocumentationVersionsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDocumentationVersionsOutput, error) {
	return m.GetDocumentationVersionsMock(arg1, arg2, arg3...)
}

// GetDomainName mocked funcition
func (m APIGatewayClientMock) GetDomainName(arg1 context.Context, arg2 *apigateway.GetDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDomainNameOutput, error) {
	return m.GetDomainNameMock(arg1, arg2, arg3...)
}

// GetDomainNames mocked funcition
func (m APIGatewayClientMock) GetDomainNames(arg1 context.Context, arg2 *apigateway.GetDomainNamesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetDomainNamesOutput, error) {
	return m.GetDomainNamesMock(arg1, arg2, arg3...)
}

// GetExport mocked funcition
func (m APIGatewayClientMock) GetExport(arg1 context.Context, arg2 *apigateway.GetExportInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetExportOutput, error) {
	return m.GetExportMock(arg1, arg2, arg3...)
}

// GetGatewayResponse mocked funcition
func (m APIGatewayClientMock) GetGatewayResponse(arg1 context.Context, arg2 *apigateway.GetGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetGatewayResponseOutput, error) {
	return m.GetGatewayResponseMock(arg1, arg2, arg3...)
}

// GetGatewayResponses mocked funcition
func (m APIGatewayClientMock) GetGatewayResponses(arg1 context.Context, arg2 *apigateway.GetGatewayResponsesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetGatewayResponsesOutput, error) {
	return m.GetGatewayResponsesMock(arg1, arg2, arg3...)
}

// GetIntegration mocked funcition
func (m APIGatewayClientMock) GetIntegration(arg1 context.Context, arg2 *apigateway.GetIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetIntegrationOutput, error) {
	return m.GetIntegrationMock(arg1, arg2, arg3...)
}

// GetIntegrationResponse mocked funcition
func (m APIGatewayClientMock) GetIntegrationResponse(arg1 context.Context, arg2 *apigateway.GetIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetIntegrationResponseOutput, error) {
	return m.GetIntegrationResponseMock(arg1, arg2, arg3...)
}

// GetMethod mocked funcition
func (m APIGatewayClientMock) GetMethod(arg1 context.Context, arg2 *apigateway.GetMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetMethodOutput, error) {
	return m.GetMethodMock(arg1, arg2, arg3...)
}

// GetMethodResponse mocked funcition
func (m APIGatewayClientMock) GetMethodResponse(arg1 context.Context, arg2 *apigateway.GetMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetMethodResponseOutput, error) {
	return m.GetMethodResponseMock(arg1, arg2, arg3...)
}

// GetModel mocked funcition
func (m APIGatewayClientMock) GetModel(arg1 context.Context, arg2 *apigateway.GetModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetModelOutput, error) {
	return m.GetModelMock(arg1, arg2, arg3...)
}

// GetModelTemplate mocked funcition
func (m APIGatewayClientMock) GetModelTemplate(arg1 context.Context, arg2 *apigateway.GetModelTemplateInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetModelTemplateOutput, error) {
	return m.GetModelTemplateMock(arg1, arg2, arg3...)
}

// GetModels mocked funcition
func (m APIGatewayClientMock) GetModels(arg1 context.Context, arg2 *apigateway.GetModelsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetModelsOutput, error) {
	return m.GetModelsMock(arg1, arg2, arg3...)
}

// GetRequestValidator mocked funcition
func (m APIGatewayClientMock) GetRequestValidator(arg1 context.Context, arg2 *apigateway.GetRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorOutput, error) {
	return m.GetRequestValidatorMock(arg1, arg2, arg3...)
}

// GetRequestValidators mocked funcition
func (m APIGatewayClientMock) GetRequestValidators(arg1 context.Context, arg2 *apigateway.GetRequestValidatorsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRequestValidatorsOutput, error) {
	return m.GetRequestValidatorsMock(arg1, arg2, arg3...)
}

// GetResource mocked funcition
func (m APIGatewayClientMock) GetResource(arg1 context.Context, arg2 *apigateway.GetResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetResourceOutput, error) {
	return m.GetResourceMock(arg1, arg2, arg3...)
}

// GetResources mocked funcition
func (m APIGatewayClientMock) GetResources(arg1 context.Context, arg2 *apigateway.GetResourcesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetResourcesOutput, error) {
	return m.GetResourcesMock(arg1, arg2, arg3...)
}

// GetRestApi mocked funcition
func (m APIGatewayClientMock) GetRestApi(arg1 context.Context, arg2 *apigateway.GetRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRestApiOutput, error) {
	return m.GetRestApiMock(arg1, arg2, arg3...)
}

// GetRestApis mocked funcition
func (m APIGatewayClientMock) GetRestApis(arg1 context.Context, arg2 *apigateway.GetRestApisInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetRestApisOutput, error) {
	return m.GetRestApisMock(arg1, arg2, arg3...)
}

// GetSdk mocked funcition
func (m APIGatewayClientMock) GetSdk(arg1 context.Context, arg2 *apigateway.GetSdkInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetSdkOutput, error) {
	return m.GetSdkMock(arg1, arg2, arg3...)
}

// GetSdkType mocked funcition
func (m APIGatewayClientMock) GetSdkType(arg1 context.Context, arg2 *apigateway.GetSdkTypeInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetSdkTypeOutput, error) {
	return m.GetSdkTypeMock(arg1, arg2, arg3...)
}

// GetSdkTypes mocked funcition
func (m APIGatewayClientMock) GetSdkTypes(arg1 context.Context, arg2 *apigateway.GetSdkTypesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetSdkTypesOutput, error) {
	return m.GetSdkTypesMock(arg1, arg2, arg3...)
}

// GetStage mocked funcition
func (m APIGatewayClientMock) GetStage(arg1 context.Context, arg2 *apigateway.GetStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetStageOutput, error) {
	return m.GetStageMock(arg1, arg2, arg3...)
}

// GetStages mocked funcition
func (m APIGatewayClientMock) GetStages(arg1 context.Context, arg2 *apigateway.GetStagesInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetStagesOutput, error) {
	return m.GetStagesMock(arg1, arg2, arg3...)
}

// GetTags mocked funcition
func (m APIGatewayClientMock) GetTags(arg1 context.Context, arg2 *apigateway.GetTagsInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetTagsOutput, error) {
	return m.GetTagsMock(arg1, arg2, arg3...)
}

// GetUsage mocked funcition
func (m APIGatewayClientMock) GetUsage(arg1 context.Context, arg2 *apigateway.GetUsageInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsageOutput, error) {
	return m.GetUsageMock(arg1, arg2, arg3...)
}

// GetUsagePlan mocked funcition
func (m APIGatewayClientMock) GetUsagePlan(arg1 context.Context, arg2 *apigateway.GetUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanOutput, error) {
	return m.GetUsagePlanMock(arg1, arg2, arg3...)
}

// GetUsagePlanKey mocked funcition
func (m APIGatewayClientMock) GetUsagePlanKey(arg1 context.Context, arg2 *apigateway.GetUsagePlanKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeyOutput, error) {
	return m.GetUsagePlanKeyMock(arg1, arg2, arg3...)
}

// GetUsagePlanKeys mocked funcition
func (m APIGatewayClientMock) GetUsagePlanKeys(arg1 context.Context, arg2 *apigateway.GetUsagePlanKeysInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlanKeysOutput, error) {
	return m.GetUsagePlanKeysMock(arg1, arg2, arg3...)
}

// GetUsagePlans mocked funcition
func (m APIGatewayClientMock) GetUsagePlans(arg1 context.Context, arg2 *apigateway.GetUsagePlansInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetUsagePlansOutput, error) {
	return m.GetUsagePlansMock(arg1, arg2, arg3...)
}

// GetVpcLink mocked funcition
func (m APIGatewayClientMock) GetVpcLink(arg1 context.Context, arg2 *apigateway.GetVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetVpcLinkOutput, error) {
	return m.GetVpcLinkMock(arg1, arg2, arg3...)
}

// GetVpcLinks mocked funcition
func (m APIGatewayClientMock) GetVpcLinks(arg1 context.Context, arg2 *apigateway.GetVpcLinksInput, arg3 ...func(*apigateway.Options)) (*apigateway.GetVpcLinksOutput, error) {
	return m.GetVpcLinksMock(arg1, arg2, arg3...)
}

// ImportApiKeys mocked funcition
func (m APIGatewayClientMock) ImportApiKeys(arg1 context.Context, arg2 *apigateway.ImportApiKeysInput, arg3 ...func(*apigateway.Options)) (*apigateway.ImportApiKeysOutput, error) {
	return m.ImportApiKeysMock(arg1, arg2, arg3...)
}

// ImportDocumentationParts mocked funcition
func (m APIGatewayClientMock) ImportDocumentationParts(arg1 context.Context, arg2 *apigateway.ImportDocumentationPartsInput, arg3 ...func(*apigateway.Options)) (*apigateway.ImportDocumentationPartsOutput, error) {
	return m.ImportDocumentationPartsMock(arg1, arg2, arg3...)
}

// ImportRestApi mocked funcition
func (m APIGatewayClientMock) ImportRestApi(arg1 context.Context, arg2 *apigateway.ImportRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.ImportRestApiOutput, error) {
	return m.ImportRestApiMock(arg1, arg2, arg3...)
}

// PutGatewayResponse mocked funcition
func (m APIGatewayClientMock) PutGatewayResponse(arg1 context.Context, arg2 *apigateway.PutGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutGatewayResponseOutput, error) {
	return m.PutGatewayResponseMock(arg1, arg2, arg3...)
}

// PutIntegration mocked funcition
func (m APIGatewayClientMock) PutIntegration(arg1 context.Context, arg2 *apigateway.PutIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutIntegrationOutput, error) {
	return m.PutIntegrationMock(arg1, arg2, arg3...)
}

// PutIntegrationResponse mocked funcition
func (m APIGatewayClientMock) PutIntegrationResponse(arg1 context.Context, arg2 *apigateway.PutIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutIntegrationResponseOutput, error) {
	return m.PutIntegrationResponseMock(arg1, arg2, arg3...)
}

// PutMethod mocked funcition
func (m APIGatewayClientMock) PutMethod(arg1 context.Context, arg2 *apigateway.PutMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutMethodOutput, error) {
	return m.PutMethodMock(arg1, arg2, arg3...)
}

// PutMethodResponse mocked funcition
func (m APIGatewayClientMock) PutMethodResponse(arg1 context.Context, arg2 *apigateway.PutMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutMethodResponseOutput, error) {
	return m.PutMethodResponseMock(arg1, arg2, arg3...)
}

// PutRestApi mocked funcition
func (m APIGatewayClientMock) PutRestApi(arg1 context.Context, arg2 *apigateway.PutRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.PutRestApiOutput, error) {
	return m.PutRestApiMock(arg1, arg2, arg3...)
}

// TagResource mocked funcition
func (m APIGatewayClientMock) TagResource(arg1 context.Context, arg2 *apigateway.TagResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.TagResourceOutput, error) {
	return m.TagResourceMock(arg1, arg2, arg3...)
}

// TestInvokeAuthorizer mocked funcition
func (m APIGatewayClientMock) TestInvokeAuthorizer(arg1 context.Context, arg2 *apigateway.TestInvokeAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.TestInvokeAuthorizerOutput, error) {
	return m.TestInvokeAuthorizerMock(arg1, arg2, arg3...)
}

// TestInvokeMethod mocked funcition
func (m APIGatewayClientMock) TestInvokeMethod(arg1 context.Context, arg2 *apigateway.TestInvokeMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.TestInvokeMethodOutput, error) {
	return m.TestInvokeMethodMock(arg1, arg2, arg3...)
}

// UntagResource mocked funcition
func (m APIGatewayClientMock) UntagResource(arg1 context.Context, arg2 *apigateway.UntagResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.UntagResourceOutput, error) {
	return m.UntagResourceMock(arg1, arg2, arg3...)
}

// UpdateAccount mocked funcition
func (m APIGatewayClientMock) UpdateAccount(arg1 context.Context, arg2 *apigateway.UpdateAccountInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateAccountOutput, error) {
	return m.UpdateAccountMock(arg1, arg2, arg3...)
}

// UpdateApiKey mocked funcition
func (m APIGatewayClientMock) UpdateApiKey(arg1 context.Context, arg2 *apigateway.UpdateApiKeyInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateApiKeyOutput, error) {
	return m.UpdateApiKeyMock(arg1, arg2, arg3...)
}

// UpdateAuthorizer mocked funcition
func (m APIGatewayClientMock) UpdateAuthorizer(arg1 context.Context, arg2 *apigateway.UpdateAuthorizerInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateAuthorizerOutput, error) {
	return m.UpdateAuthorizerMock(arg1, arg2, arg3...)
}

// UpdateBasePathMapping mocked funcition
func (m APIGatewayClientMock) UpdateBasePathMapping(arg1 context.Context, arg2 *apigateway.UpdateBasePathMappingInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateBasePathMappingOutput, error) {
	return m.UpdateBasePathMappingMock(arg1, arg2, arg3...)
}

// UpdateClientCertificate mocked funcition
func (m APIGatewayClientMock) UpdateClientCertificate(arg1 context.Context, arg2 *apigateway.UpdateClientCertificateInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateClientCertificateOutput, error) {
	return m.UpdateClientCertificateMock(arg1, arg2, arg3...)
}

// UpdateDeployment mocked funcition
func (m APIGatewayClientMock) UpdateDeployment(arg1 context.Context, arg2 *apigateway.UpdateDeploymentInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDeploymentOutput, error) {
	return m.UpdateDeploymentMock(arg1, arg2, arg3...)
}

// UpdateDocumentationPart mocked funcition
func (m APIGatewayClientMock) UpdateDocumentationPart(arg1 context.Context, arg2 *apigateway.UpdateDocumentationPartInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDocumentationPartOutput, error) {
	return m.UpdateDocumentationPartMock(arg1, arg2, arg3...)
}

// UpdateDocumentationVersion mocked funcition
func (m APIGatewayClientMock) UpdateDocumentationVersion(arg1 context.Context, arg2 *apigateway.UpdateDocumentationVersionInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDocumentationVersionOutput, error) {
	return m.UpdateDocumentationVersionMock(arg1, arg2, arg3...)
}

// UpdateDomainName mocked funcition
func (m APIGatewayClientMock) UpdateDomainName(arg1 context.Context, arg2 *apigateway.UpdateDomainNameInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateDomainNameOutput, error) {
	return m.UpdateDomainNameMock(arg1, arg2, arg3...)
}

// UpdateGatewayResponse mocked funcition
func (m APIGatewayClientMock) UpdateGatewayResponse(arg1 context.Context, arg2 *apigateway.UpdateGatewayResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateGatewayResponseOutput, error) {
	return m.UpdateGatewayResponseMock(arg1, arg2, arg3...)
}

// UpdateIntegration mocked funcition
func (m APIGatewayClientMock) UpdateIntegration(arg1 context.Context, arg2 *apigateway.UpdateIntegrationInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateIntegrationOutput, error) {
	return m.UpdateIntegrationMock(arg1, arg2, arg3...)
}

// UpdateIntegrationResponse mocked funcition
func (m APIGatewayClientMock) UpdateIntegrationResponse(arg1 context.Context, arg2 *apigateway.UpdateIntegrationResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateIntegrationResponseOutput, error) {
	return m.UpdateIntegrationResponseMock(arg1, arg2, arg3...)
}

// UpdateMethod mocked funcition
func (m APIGatewayClientMock) UpdateMethod(arg1 context.Context, arg2 *apigateway.UpdateMethodInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateMethodOutput, error) {
	return m.UpdateMethodMock(arg1, arg2, arg3...)
}

// UpdateMethodResponse mocked funcition
func (m APIGatewayClientMock) UpdateMethodResponse(arg1 context.Context, arg2 *apigateway.UpdateMethodResponseInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateMethodResponseOutput, error) {
	return m.UpdateMethodResponseMock(arg1, arg2, arg3...)
}

// UpdateModel mocked funcition
func (m APIGatewayClientMock) UpdateModel(arg1 context.Context, arg2 *apigateway.UpdateModelInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateModelOutput, error) {
	return m.UpdateModelMock(arg1, arg2, arg3...)
}

// UpdateRequestValidator mocked funcition
func (m APIGatewayClientMock) UpdateRequestValidator(arg1 context.Context, arg2 *apigateway.UpdateRequestValidatorInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateRequestValidatorOutput, error) {
	return m.UpdateRequestValidatorMock(arg1, arg2, arg3...)
}

// UpdateResource mocked funcition
func (m APIGatewayClientMock) UpdateResource(arg1 context.Context, arg2 *apigateway.UpdateResourceInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateResourceOutput, error) {
	return m.UpdateResourceMock(arg1, arg2, arg3...)
}

// UpdateRestApi mocked funcition
func (m APIGatewayClientMock) UpdateRestApi(arg1 context.Context, arg2 *apigateway.UpdateRestApiInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateRestApiOutput, error) {
	return m.UpdateRestApiMock(arg1, arg2, arg3...)
}

// UpdateStage mocked funcition
func (m APIGatewayClientMock) UpdateStage(arg1 context.Context, arg2 *apigateway.UpdateStageInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateStageOutput, error) {
	return m.UpdateStageMock(arg1, arg2, arg3...)
}

// UpdateUsage mocked funcition
func (m APIGatewayClientMock) UpdateUsage(arg1 context.Context, arg2 *apigateway.UpdateUsageInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateUsageOutput, error) {
	return m.UpdateUsageMock(arg1, arg2, arg3...)
}

// UpdateUsagePlan mocked funcition
func (m APIGatewayClientMock) UpdateUsagePlan(arg1 context.Context, arg2 *apigateway.UpdateUsagePlanInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateUsagePlanOutput, error) {
	return m.UpdateUsagePlanMock(arg1, arg2, arg3...)
}

// UpdateVpcLink mocked funcition
func (m APIGatewayClientMock) UpdateVpcLink(arg1 context.Context, arg2 *apigateway.UpdateVpcLinkInput, arg3 ...func(*apigateway.Options)) (*apigateway.UpdateVpcLinkOutput, error) {
	return m.UpdateVpcLinkMock(arg1, arg2, arg3...)
}
