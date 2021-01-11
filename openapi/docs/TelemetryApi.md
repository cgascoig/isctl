# \TelemetryApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryTelemetryDatasourceMetadata**](TelemetryApi.md#QueryTelemetryDatasourceMetadata) | **Post** /api/v1/telemetry/DataSourceMetadata | Perform a Druid DatasourceMetadata request.
[**QueryTelemetryGroupBy**](TelemetryApi.md#QueryTelemetryGroupBy) | **Post** /api/v1/telemetry/GroupBys | Perform a Druid GroupBy request.
[**QueryTelemetryScan**](TelemetryApi.md#QueryTelemetryScan) | **Post** /api/v1/telemetry/Scans | Perform a Druid Scan request.
[**QueryTelemetrySearch**](TelemetryApi.md#QueryTelemetrySearch) | **Post** /api/v1/telemetry/Searches | Perform a Druid Search request.
[**QueryTelemetrySegmentMetadata**](TelemetryApi.md#QueryTelemetrySegmentMetadata) | **Post** /api/v1/telemetry/SegmentMetadata | Perform a Druid SegmentMetadata request.
[**QueryTelemetryTimeBoundary**](TelemetryApi.md#QueryTelemetryTimeBoundary) | **Post** /api/v1/telemetry/TimeBoundaries | Perform a Druid TimeBoundary request.
[**QueryTelemetryTimeSeries**](TelemetryApi.md#QueryTelemetryTimeSeries) | **Post** /api/v1/telemetry/TimeSeries | Perform a Druid TimeSeries request.
[**QueryTelemetryTopN**](TelemetryApi.md#QueryTelemetryTopN) | **Post** /api/v1/telemetry/Topns | Perform a Druid TopN request.



## QueryTelemetryDatasourceMetadata

> []TelemetryDruidDataSourceMetadataResult QueryTelemetryDatasourceMetadata(ctx).TelemetryDruidDataSourceMetadataRequest(telemetryDruidDataSourceMetadataRequest).Execute()

Perform a Druid DatasourceMetadata request.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    telemetryDruidDataSourceMetadataRequest := openapiclient.telemetry.DruidDataSourceMetadataRequest{QueryType: "QueryType_example", DataSource: openapiclient.telemetry.DruidDataSource{Type: "Type_example", Name: "Name_example", DataSources: []string{"DataSources_example"), Query: openapiclient.telemetry.DruidGroupByRequest{QueryType: "QueryType_example", DataSource: openapiclient.telemetry.DruidDataSource{Type: "Type_example", Name: "Name_example", DataSources: []string{"DataSources_example"), Query: openapiclient.telemetry.DruidGroupByRequest{QueryType: "QueryType_example", DataSource: , Dimensions: []TelemetryDruidDimensionSpec{openapiclient.telemetry.DruidDimensionSpec{Type: "Type_example", Dimension: "Dimension_example", OutputName: "OutputName_example", OutputType: "OutputType_example", ExtractionFn: 123}), LimitSpec: openapiclient.telemetry.DruidDefaultLimitSpec{Type: "Type_example", Limit: 123, Columns: []TelemetryDruidOrderByColumnSpec{openapiclient.telemetry.DruidOrderByColumnSpec{Dimension: "Dimension_example", Direction: "Direction_example", DimensionOrder: "DimensionOrder_example"})}, Having: openapiclient.telemetry.DruidHavingFilter{Type: "Type_example", Filter: openapiclient.telemetry.DruidFilter{Type: "Type_example", ExtractionFn: 123, Dimension: "Dimension_example", Value: "Value_example", Dimensions: []TelemetryDruidDimensionSpec{openapiclient.telemetry.DruidDimensionSpec{Type: "Type_example", Dimension: "Dimension_example", OutputName: "OutputName_example", OutputType: "OutputType_example", ExtractionFn: 123}), Pattern: "Pattern_example", Fields: []TelemetryDruidFilter{openapiclient.telemetry.DruidFilter{Type: "Type_example", ExtractionFn: 123, Dimension: "Dimension_example", Value: "Value_example", Dimensions: []TelemetryDruidDimensionSpec{), Pattern: "Pattern_example", Fields: []TelemetryDruidFilter{), Field: }), Field: }, Aggregation: "Aggregation_example", Value: 123, Dimension: "Dimension_example"}, Granularity: openapiclient.telemetry.DruidGranularity{Type: "Type_example", Duration: int64(123), Origin: "TODO", Period: "Period_example", TimeZone: "TimeZone_example"}, Filter: , Aggregations: []TelemetryDruidAggregator{openapiclient.telemetry.DruidAggregator{Type: "Type_example", Name: "Name_example", FieldName: "FieldName_example", MaxStringBytes: 123, Size: 123, Filter: , Aggregator: openapiclient.telemetry.DruidAggregator{Type: "Type_example", Name: "Name_example", FieldName: "FieldName_example", MaxStringBytes: 123, Size: 123, Filter: , Aggregator: }}), PostAggregations: []TelemetryDruidPostAggregator{openapiclient.telemetry.DruidPostAggregator{Type: "Type_example", Name: "Name_example", Fn: "Fn_example", Fields: []string{"Fields_example"), Ordering: "Ordering_example", FieldName: "FieldName_example", Value: 123, Field: "Field_example", Func: "Func_example", Size: 123}), Intervals: []string{"Intervals_example"), SubtotalsSpec: 123, Context: openapiclient.telemetry.DruidQueryContext{GrandTotal: false, SkipEmptyBuckets: false, Timeout: 123, Priority: 123, QueryId: "QueryId_example", UseCache: false, PopulateCache: false, UseResultLevelCache: false, PopulateResultLevelCache: false, BySegment: false, Finalize: false, ChunkPeriod: "ChunkPeriod_example", MaxScatterGatherBytes: 123, MaxQueuedBytes: 123, SerializeDateTimeAsLong: false, SerializeDateTimeAsLongInner: false, EnableParallelMerge: false, ParallelMergeParallelism: 123, ParallelMergeInitialYieldRows: 123, ParallelMergeSmallBatchRows: 123}}, Lookup: "Lookup_example", Left: "Left_example", Right: "Right_example", RightPrefix: "RightPrefix_example", Condition: "Condition_example", JoinType: "JoinType_example", ColumnNames: []string{"ColumnNames_example"), Rows: [][]string{[]string{"Rows_example"))}, Dimensions: []TelemetryDruidDimensionSpec{), LimitSpec: openapiclient.telemetry.DruidDefaultLimitSpec{Type: "Type_example", Limit: 123, Columns: []TelemetryDruidOrderByColumnSpec{openapiclient.telemetry.DruidOrderByColumnSpec{Dimension: "Dimension_example", Direction: "Direction_example", DimensionOrder: "DimensionOrder_example"})}, Having: openapiclient.telemetry.DruidHavingFilter{Type: "Type_example", Filter: , Aggregation: "Aggregation_example", Value: 123, Dimension: "Dimension_example"}, Granularity: openapiclient.telemetry.DruidGranularity{Type: "Type_example", Duration: int64(123), Origin: "TODO", Period: "Period_example", TimeZone: "TimeZone_example"}, Filter: , Aggregations: []TelemetryDruidAggregator{), PostAggregations: []TelemetryDruidPostAggregator{openapiclient.telemetry.DruidPostAggregator{Type: "Type_example", Name: "Name_example", Fn: "Fn_example", Fields: []string{"Fields_example"), Ordering: "Ordering_example", FieldName: "FieldName_example", Value: 123, Field: "Field_example", Func: "Func_example", Size: 123}), Intervals: []string{"Intervals_example"), SubtotalsSpec: 123, Context: openapiclient.telemetry.DruidQueryContext{GrandTotal: false, SkipEmptyBuckets: false, Timeout: 123, Priority: 123, QueryId: "QueryId_example", UseCache: false, PopulateCache: false, UseResultLevelCache: false, PopulateResultLevelCache: false, BySegment: false, Finalize: false, ChunkPeriod: "ChunkPeriod_example", MaxScatterGatherBytes: 123, MaxQueuedBytes: 123, SerializeDateTimeAsLong: false, SerializeDateTimeAsLongInner: false, EnableParallelMerge: false, ParallelMergeParallelism: 123, ParallelMergeInitialYieldRows: 123, ParallelMergeSmallBatchRows: 123}}, Lookup: "Lookup_example", Left: "Left_example", Right: "Right_example", RightPrefix: "RightPrefix_example", Condition: "Condition_example", JoinType: "JoinType_example", ColumnNames: []string{"ColumnNames_example"), Rows: [][]string{[]string{"Rows_example"))}, Context: } // TelemetryDruidDataSourceMetadataRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryDatasourceMetadata(context.Background(), telemetryDruidDataSourceMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryDatasourceMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryDatasourceMetadata`: []TelemetryDruidDataSourceMetadataResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryDatasourceMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryDatasourceMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidDataSourceMetadataRequest** | [**TelemetryDruidDataSourceMetadataRequest**](TelemetryDruidDataSourceMetadataRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidDataSourceMetadataResult**](telemetry.DruidDataSourceMetadataResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetryGroupBy

> []TelemetryDruidGroupByResult QueryTelemetryGroupBy(ctx).TelemetryDruidGroupByRequest(telemetryDruidGroupByRequest).Execute()

Perform a Druid GroupBy request.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    telemetryDruidGroupByRequest :=  // TelemetryDruidGroupByRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryGroupBy(context.Background(), telemetryDruidGroupByRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryGroupBy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryGroupBy`: []TelemetryDruidGroupByResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryGroupBy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryGroupByRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidGroupByRequest** | [**TelemetryDruidGroupByRequest**](TelemetryDruidGroupByRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidGroupByResult**](telemetry.DruidGroupByResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetryScan

> []TelemetryDruidScanResult QueryTelemetryScan(ctx).TelemetryDruidScanRequest(telemetryDruidScanRequest).Execute()

Perform a Druid Scan request.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    telemetryDruidScanRequest := openapiclient.telemetry.DruidScanRequest{QueryType: "QueryType_example", DataSource: , Intervals: []string{"Intervals_example"), ResultFormat: "ResultFormat_example", Filter: , Columns: []string{"Columns_example"), BatchSize: 123, Limit: 123, Order: "Order_example", Legacy: false, Context: } // TelemetryDruidScanRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryScan(context.Background(), telemetryDruidScanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryScan`: []TelemetryDruidScanResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryScan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidScanRequest** | [**TelemetryDruidScanRequest**](TelemetryDruidScanRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidScanResult**](telemetry.DruidScanResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetrySearch

> []TelemetryDruidSearchResult QueryTelemetrySearch(ctx).TelemetryDruidSearchRequest(telemetryDruidSearchRequest).Execute()

Perform a Druid Search request.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    telemetryDruidSearchRequest := openapiclient.telemetry.DruidSearchRequest{QueryType: "QueryType_example", DataSource: , Intervals: []string{"Intervals_example"), Granularity: , Filter: , Aggregations: []TelemetryDruidAggregator{), SearchDimensions: []string{"SearchDimensions_example"), Query: openapiclient.telemetry.DruidAggregateSearchSpec{Type: "Type_example", Value: "Value_example", Values: []string{"Values_example"), CaseSensitive: false, Regex: "Regex_example"}, Limit: 123, Context: } // TelemetryDruidSearchRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetrySearch(context.Background(), telemetryDruidSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetrySearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetrySearch`: []TelemetryDruidSearchResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetrySearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetrySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidSearchRequest** | [**TelemetryDruidSearchRequest**](TelemetryDruidSearchRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidSearchResult**](telemetry.DruidSearchResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetrySegmentMetadata

> []TelemetryDruidSegmentMetadataResult QueryTelemetrySegmentMetadata(ctx).TelemetryDruidSegmentMetadataRequest(telemetryDruidSegmentMetadataRequest).Execute()

Perform a Druid SegmentMetadata request.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    telemetryDruidSegmentMetadataRequest := openapiclient.telemetry.DruidSegmentMetadataRequest{QueryType: "QueryType_example", DataSource: , Intervals: []string{"Intervals_example"), ToInclude: 123, Merge: false, Context: , AnalysisTypes: []string{"AnalysisTypes_example"), LenientAggregatorMerge: false} // TelemetryDruidSegmentMetadataRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetrySegmentMetadata(context.Background(), telemetryDruidSegmentMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetrySegmentMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetrySegmentMetadata`: []TelemetryDruidSegmentMetadataResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetrySegmentMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetrySegmentMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidSegmentMetadataRequest** | [**TelemetryDruidSegmentMetadataRequest**](TelemetryDruidSegmentMetadataRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidSegmentMetadataResult**](telemetry.DruidSegmentMetadataResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetryTimeBoundary

> []TelemetryDruidTimeBoundaryResult QueryTelemetryTimeBoundary(ctx).TelemetryDruidTimeBoundaryRequest(telemetryDruidTimeBoundaryRequest).Execute()

Perform a Druid TimeBoundary request.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    telemetryDruidTimeBoundaryRequest := openapiclient.telemetry.DruidTimeBoundaryRequest{QueryType: "QueryType_example", DataSource: , Bound: "Bound_example", Filter: , Context: } // TelemetryDruidTimeBoundaryRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryTimeBoundary(context.Background(), telemetryDruidTimeBoundaryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryTimeBoundary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryTimeBoundary`: []TelemetryDruidTimeBoundaryResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryTimeBoundary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryTimeBoundaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidTimeBoundaryRequest** | [**TelemetryDruidTimeBoundaryRequest**](TelemetryDruidTimeBoundaryRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidTimeBoundaryResult**](telemetry.DruidTimeBoundaryResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetryTimeSeries

> []TelemetryDruidIntervalResult QueryTelemetryTimeSeries(ctx).TelemetryDruidTimeSeriesRequest(telemetryDruidTimeSeriesRequest).Execute()

Perform a Druid TimeSeries request.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    telemetryDruidTimeSeriesRequest := openapiclient.telemetry.DruidTimeSeriesRequest{QueryType: "QueryType_example", DataSource: , Descending: false, Intervals: []string{"Intervals_example"), Granularity: , Filter: , Aggregations: []TelemetryDruidAggregator{), PostAggregations: []TelemetryDruidPostAggregator{), Limit: 123, Context: } // TelemetryDruidTimeSeriesRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryTimeSeries(context.Background(), telemetryDruidTimeSeriesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryTimeSeries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryTimeSeries`: []TelemetryDruidIntervalResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryTimeSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryTimeSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidTimeSeriesRequest** | [**TelemetryDruidTimeSeriesRequest**](TelemetryDruidTimeSeriesRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidIntervalResult**](telemetry.DruidIntervalResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryTelemetryTopN

> []TelemetryDruidTopNResult QueryTelemetryTopN(ctx).TelemetryDruidTopNRequest(telemetryDruidTopNRequest).Execute()

Perform a Druid TopN request.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    telemetryDruidTopNRequest := openapiclient.telemetry.DruidTopNRequest{QueryType: "QueryType_example", DataSource: , Intervals: []string{"Intervals_example"), Granularity: , Filter: , Aggregations: []TelemetryDruidAggregator{), PostAggregations: []TelemetryDruidPostAggregator{), Dimension: , Threshold: 123, Metric: openapiclient.telemetry.DruidTopNMetricSpec{Type: "Type_example", Metric: openapiclient.telemetry.DruidTopNMetricSpec{Type: "Type_example", Metric: , Ordering: "Ordering_example", PreviousStop: "PreviousStop_example"}, Ordering: "Ordering_example", PreviousStop: "PreviousStop_example"}, Context: } // TelemetryDruidTopNRequest | The Druid request schema for time series queries.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TelemetryApi.QueryTelemetryTopN(context.Background(), telemetryDruidTopNRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryApi.QueryTelemetryTopN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryTelemetryTopN`: []TelemetryDruidTopNResult
    fmt.Fprintf(os.Stdout, "Response from `TelemetryApi.QueryTelemetryTopN`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryTelemetryTopNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryDruidTopNRequest** | [**TelemetryDruidTopNRequest**](TelemetryDruidTopNRequest.md) | The Druid request schema for time series queries. | 

### Return type

[**[]TelemetryDruidTopNResult**](telemetry.DruidTopNResult.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

