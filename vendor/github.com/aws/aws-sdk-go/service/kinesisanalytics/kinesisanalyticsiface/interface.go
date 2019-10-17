// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kinesisanalyticsiface provides an interface to enable mocking the Amazon Kinesis Analytics service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kinesisanalyticsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
)

// KinesisAnalyticsAPI provides an interface to enable mocking the
// kinesisanalytics.KinesisAnalytics service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis Analytics.
//    func myFunc(svc kinesisanalyticsiface.KinesisAnalyticsAPI) bool {
//        // Make svc.AddApplicationCloudWatchLoggingOption request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kinesisanalytics.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKinesisAnalyticsClient struct {
//        kinesisanalyticsiface.KinesisAnalyticsAPI
//    }
//    func (m *mockKinesisAnalyticsClient) AddApplicationCloudWatchLoggingOption(input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKinesisAnalyticsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type KinesisAnalyticsAPI interface {
	AddApplicationCloudWatchLoggingOption(*kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error)
	AddApplicationCloudWatchLoggingOptionWithContext(aws.Context, *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput, ...request.Option) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error)
	AddApplicationCloudWatchLoggingOptionRequest(*kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*request.Request, *kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput)

	AddApplicationInput(*kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error)
	AddApplicationInputWithContext(aws.Context, *kinesisanalytics.AddApplicationInputInput, ...request.Option) (*kinesisanalytics.AddApplicationInputOutput, error)
	AddApplicationInputRequest(*kinesisanalytics.AddApplicationInputInput) (*request.Request, *kinesisanalytics.AddApplicationInputOutput)

	AddApplicationInputProcessingConfiguration(*kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error)
	AddApplicationInputProcessingConfigurationWithContext(aws.Context, *kinesisanalytics.AddApplicationInputProcessingConfigurationInput, ...request.Option) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error)
	AddApplicationInputProcessingConfigurationRequest(*kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*request.Request, *kinesisanalytics.AddApplicationInputProcessingConfigurationOutput)

	AddApplicationOutput(*kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error)
	AddApplicationOutputWithContext(aws.Context, *kinesisanalytics.AddApplicationOutputInput, ...request.Option) (*kinesisanalytics.AddApplicationOutputOutput, error)
	AddApplicationOutputRequest(*kinesisanalytics.AddApplicationOutputInput) (*request.Request, *kinesisanalytics.AddApplicationOutputOutput)

	AddApplicationReferenceDataSource(*kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error)
	AddApplicationReferenceDataSourceWithContext(aws.Context, *kinesisanalytics.AddApplicationReferenceDataSourceInput, ...request.Option) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error)
	AddApplicationReferenceDataSourceRequest(*kinesisanalytics.AddApplicationReferenceDataSourceInput) (*request.Request, *kinesisanalytics.AddApplicationReferenceDataSourceOutput)

	CreateApplication(*kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error)
	CreateApplicationWithContext(aws.Context, *kinesisanalytics.CreateApplicationInput, ...request.Option) (*kinesisanalytics.CreateApplicationOutput, error)
	CreateApplicationRequest(*kinesisanalytics.CreateApplicationInput) (*request.Request, *kinesisanalytics.CreateApplicationOutput)

	DeleteApplication(*kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error)
	DeleteApplicationWithContext(aws.Context, *kinesisanalytics.DeleteApplicationInput, ...request.Option) (*kinesisanalytics.DeleteApplicationOutput, error)
	DeleteApplicationRequest(*kinesisanalytics.DeleteApplicationInput) (*request.Request, *kinesisanalytics.DeleteApplicationOutput)

	DeleteApplicationCloudWatchLoggingOption(*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error)
	DeleteApplicationCloudWatchLoggingOptionWithContext(aws.Context, *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput, ...request.Option) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error)
	DeleteApplicationCloudWatchLoggingOptionRequest(*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*request.Request, *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput)

	DeleteApplicationInputProcessingConfiguration(*kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error)
	DeleteApplicationInputProcessingConfigurationWithContext(aws.Context, *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput, ...request.Option) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error)
	DeleteApplicationInputProcessingConfigurationRequest(*kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*request.Request, *kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput)

	DeleteApplicationOutput(*kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error)
	DeleteApplicationOutputWithContext(aws.Context, *kinesisanalytics.DeleteApplicationOutputInput, ...request.Option) (*kinesisanalytics.DeleteApplicationOutputOutput, error)
	DeleteApplicationOutputRequest(*kinesisanalytics.DeleteApplicationOutputInput) (*request.Request, *kinesisanalytics.DeleteApplicationOutputOutput)

	DeleteApplicationReferenceDataSource(*kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error)
	DeleteApplicationReferenceDataSourceWithContext(aws.Context, *kinesisanalytics.DeleteApplicationReferenceDataSourceInput, ...request.Option) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error)
	DeleteApplicationReferenceDataSourceRequest(*kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*request.Request, *kinesisanalytics.DeleteApplicationReferenceDataSourceOutput)

	DescribeApplication(*kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error)
	DescribeApplicationWithContext(aws.Context, *kinesisanalytics.DescribeApplicationInput, ...request.Option) (*kinesisanalytics.DescribeApplicationOutput, error)
	DescribeApplicationRequest(*kinesisanalytics.DescribeApplicationInput) (*request.Request, *kinesisanalytics.DescribeApplicationOutput)

	DiscoverInputSchema(*kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error)
	DiscoverInputSchemaWithContext(aws.Context, *kinesisanalytics.DiscoverInputSchemaInput, ...request.Option) (*kinesisanalytics.DiscoverInputSchemaOutput, error)
	DiscoverInputSchemaRequest(*kinesisanalytics.DiscoverInputSchemaInput) (*request.Request, *kinesisanalytics.DiscoverInputSchemaOutput)

	ListApplications(*kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error)
	ListApplicationsWithContext(aws.Context, *kinesisanalytics.ListApplicationsInput, ...request.Option) (*kinesisanalytics.ListApplicationsOutput, error)
	ListApplicationsRequest(*kinesisanalytics.ListApplicationsInput) (*request.Request, *kinesisanalytics.ListApplicationsOutput)

	ListTagsForResource(*kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *kinesisanalytics.ListTagsForResourceInput, ...request.Option) (*kinesisanalytics.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*kinesisanalytics.ListTagsForResourceInput) (*request.Request, *kinesisanalytics.ListTagsForResourceOutput)

	StartApplication(*kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error)
	StartApplicationWithContext(aws.Context, *kinesisanalytics.StartApplicationInput, ...request.Option) (*kinesisanalytics.StartApplicationOutput, error)
	StartApplicationRequest(*kinesisanalytics.StartApplicationInput) (*request.Request, *kinesisanalytics.StartApplicationOutput)

	StopApplication(*kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error)
	StopApplicationWithContext(aws.Context, *kinesisanalytics.StopApplicationInput, ...request.Option) (*kinesisanalytics.StopApplicationOutput, error)
	StopApplicationRequest(*kinesisanalytics.StopApplicationInput) (*request.Request, *kinesisanalytics.StopApplicationOutput)

	TagResource(*kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *kinesisanalytics.TagResourceInput, ...request.Option) (*kinesisanalytics.TagResourceOutput, error)
	TagResourceRequest(*kinesisanalytics.TagResourceInput) (*request.Request, *kinesisanalytics.TagResourceOutput)

	UntagResource(*kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *kinesisanalytics.UntagResourceInput, ...request.Option) (*kinesisanalytics.UntagResourceOutput, error)
	UntagResourceRequest(*kinesisanalytics.UntagResourceInput) (*request.Request, *kinesisanalytics.UntagResourceOutput)

	UpdateApplication(*kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error)
	UpdateApplicationWithContext(aws.Context, *kinesisanalytics.UpdateApplicationInput, ...request.Option) (*kinesisanalytics.UpdateApplicationOutput, error)
	UpdateApplicationRequest(*kinesisanalytics.UpdateApplicationInput) (*request.Request, *kinesisanalytics.UpdateApplicationOutput)
}

var _ KinesisAnalyticsAPI = (*kinesisanalytics.KinesisAnalytics)(nil)
