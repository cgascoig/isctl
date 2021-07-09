/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// TelemetryDruidQueryContext The query context is used for various query configuration parameters. Can be used to modify query behavior, including grand totals and zero-filling.
type TelemetryDruidQueryContext struct {
	// Druid can include an extra \"grand totals\" row as the last row of a timeseries result set. To enable this, set \"grandTotal\" to true. The grand totals row will appear as the last row in the result array, and will have no timestamp. It will be the last row even if the query is run in \"descending\" mode. Post-aggregations in the grand totals row will be computed based upon the grand total aggregations.
	GrandTotal *bool `json:"grandTotal,omitempty"`
	// Timeseries queries normally fill empty interior time buckets with zeroes. Time buckets that lie completely outside the data interval are not zero-filled. You can disable all zero-filling with this flag. In this mode, the data point for empty buckets are omitted from the results.
	SkipEmptyBuckets *bool `json:"skipEmptyBuckets,omitempty"`
	// Query timeout in milliseconds, beyond which unfinished queries will be cancelled. 0 timeout means no timeout.
	Timeout *int32 `json:"timeout,omitempty"`
	// Query Priority. Queries with higher priority get precedence for computational resources.
	Priority *int32 `json:"priority,omitempty"`
	// Unique identifier given to this query. If a query ID is set or known, this can be used to cancel the query.
	QueryId *string `json:"queryId,omitempty"`
	// Flag indicating whether to leverage the query cache for this query. When set to false, it disables reading from the query cache for this query. When set to true, Apache Druid uses druid.broker.cache.useCache or druid.historical.cache.useCache to determine whether or not to read from the query cache.
	UseCache *bool `json:"useCache,omitempty"`
	// Flag indicating whether to save the results of the query to the query cache. Primarily used for debugging. When set to false, it disables saving the results of this query to the query cache. When set to true, Druid uses druid.broker.cache.populateCache or druid.historical.cache.populateCache to determine whether or not to save the results of this query to the query cache.
	PopulateCache *bool `json:"populateCache,omitempty"`
	// Flag indicating whether to leverage the result level cache for this query. When set to false, it disables reading from the query cache for this query. When set to true, Druid uses druid.broker.cache.useResultLevelCache to determine whether or not to read from the result-level query cache.
	UseResultLevelCache *bool `json:"useResultLevelCache,omitempty"`
	// Flag indicating whether to save the results of the query to the result level cache. Primarily used for debugging. When set to false, it disables saving the results of this query to the query cache. When set to true, Druid uses druid.broker.cache.populateResultLevelCache to determine whether or not to save the results of this query to the result-level query cache.
	PopulateResultLevelCache *bool `json:"populateResultLevelCache,omitempty"`
	// Return \"by segment\" results. Primarily used for debugging, setting it to true returns results associated with the data segment they came from.
	BySegment *bool `json:"bySegment,omitempty"`
	// Flag indicating whether to \"finalize\" aggregation results. Primarily used for debugging. For instance, the hyperUnique aggregator will return the full HyperLogLog sketch instead of the estimated cardinality when this flag is set to false.
	Finalize *bool `json:"finalize,omitempty"`
	// At the Broker process level, long interval queries (of any type) may be broken into shorter interval queries to parallelize merging more than normal. Broken up queries will use a larger share of cluster resources, but, if you use groupBy \"v1, it may be able to complete faster as a result. Use ISO 8601 periods. For example, if this property is set to P1M (one month), then a query covering a year would be broken into 12 smaller queries. The broker uses its query processing executor service to initiate processing for query chunks, so make sure druid.processing.numThreads is configured appropriately on the broker. groupBy queries do not support chunkPeriod by default, although they do if using the legacy \"v1\" engine. This context is deprecated since it's only useful for groupBy \"v1\", and will be removed in the future releases.
	ChunkPeriod *string `json:"chunkPeriod,omitempty"`
	// Maximum number of bytes gathered from data processes such as Historicals and realtime processes to execute a query. This parameter can be used to further reduce maxScatterGatherBytes limit at query time.
	MaxScatterGatherBytes *int32 `json:"maxScatterGatherBytes,omitempty"`
	// Maximum number of bytes queued per query before exerting backpressure on the channel to the data server. Similar to maxScatterGatherBytes, except unlike that configuration, this one will trigger backpressure rather than query failure. Zero means disabled.
	MaxQueuedBytes *int32 `json:"maxQueuedBytes,omitempty"`
	// If true, DateTime is serialized as long in the result returned by Broker and the data transportation between Broker and compute process.
	SerializeDateTimeAsLong *bool `json:"serializeDateTimeAsLong,omitempty"`
	// If true, DateTime is serialized as long in the data transportation between Broker and compute process.
	SerializeDateTimeAsLongInner *bool `json:"serializeDateTimeAsLongInner,omitempty"`
	// Enable parallel result merging on the Broker. Note that druid.processing.merge.useParallelMergePool must be enabled for this setting to be set to true.
	EnableParallelMerge *bool `json:"enableParallelMerge,omitempty"`
	// Maximum number of parallel threads to use for parallel result merging on the Broker.
	ParallelMergeParallelism *int32 `json:"parallelMergeParallelism,omitempty"`
	// Number of rows to yield per ForkJoinPool merge task for parallel result merging on the Broker, before forking off a new task to continue merging sequences.
	ParallelMergeInitialYieldRows *int32 `json:"parallelMergeInitialYieldRows,omitempty"`
	// Size of result batches to operate on in ForkJoinPool merge tasks for parallel result merging on the Broker.
	ParallelMergeSmallBatchRows *int32 `json:"parallelMergeSmallBatchRows,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _TelemetryDruidQueryContext TelemetryDruidQueryContext

// NewTelemetryDruidQueryContext instantiates a new TelemetryDruidQueryContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidQueryContext() *TelemetryDruidQueryContext {
	this := TelemetryDruidQueryContext{}
	return &this
}

// NewTelemetryDruidQueryContextWithDefaults instantiates a new TelemetryDruidQueryContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidQueryContextWithDefaults() *TelemetryDruidQueryContext {
	this := TelemetryDruidQueryContext{}
	return &this
}

// GetGrandTotal returns the GrandTotal field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetGrandTotal() bool {
	if o == nil || o.GrandTotal == nil {
		var ret bool
		return ret
	}
	return *o.GrandTotal
}

// GetGrandTotalOk returns a tuple with the GrandTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetGrandTotalOk() (*bool, bool) {
	if o == nil || o.GrandTotal == nil {
		return nil, false
	}
	return o.GrandTotal, true
}

// HasGrandTotal returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasGrandTotal() bool {
	if o != nil && o.GrandTotal != nil {
		return true
	}

	return false
}

// SetGrandTotal gets a reference to the given bool and assigns it to the GrandTotal field.
func (o *TelemetryDruidQueryContext) SetGrandTotal(v bool) {
	o.GrandTotal = &v
}

// GetSkipEmptyBuckets returns the SkipEmptyBuckets field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetSkipEmptyBuckets() bool {
	if o == nil || o.SkipEmptyBuckets == nil {
		var ret bool
		return ret
	}
	return *o.SkipEmptyBuckets
}

// GetSkipEmptyBucketsOk returns a tuple with the SkipEmptyBuckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetSkipEmptyBucketsOk() (*bool, bool) {
	if o == nil || o.SkipEmptyBuckets == nil {
		return nil, false
	}
	return o.SkipEmptyBuckets, true
}

// HasSkipEmptyBuckets returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasSkipEmptyBuckets() bool {
	if o != nil && o.SkipEmptyBuckets != nil {
		return true
	}

	return false
}

// SetSkipEmptyBuckets gets a reference to the given bool and assigns it to the SkipEmptyBuckets field.
func (o *TelemetryDruidQueryContext) SetSkipEmptyBuckets(v bool) {
	o.SkipEmptyBuckets = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetTimeout() int32 {
	if o == nil || o.Timeout == nil {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetTimeoutOk() (*int32, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *TelemetryDruidQueryContext) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *TelemetryDruidQueryContext) SetPriority(v int32) {
	o.Priority = &v
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetQueryId() string {
	if o == nil || o.QueryId == nil {
		var ret string
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetQueryIdOk() (*string, bool) {
	if o == nil || o.QueryId == nil {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasQueryId() bool {
	if o != nil && o.QueryId != nil {
		return true
	}

	return false
}

// SetQueryId gets a reference to the given string and assigns it to the QueryId field.
func (o *TelemetryDruidQueryContext) SetQueryId(v string) {
	o.QueryId = &v
}

// GetUseCache returns the UseCache field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetUseCache() bool {
	if o == nil || o.UseCache == nil {
		var ret bool
		return ret
	}
	return *o.UseCache
}

// GetUseCacheOk returns a tuple with the UseCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetUseCacheOk() (*bool, bool) {
	if o == nil || o.UseCache == nil {
		return nil, false
	}
	return o.UseCache, true
}

// HasUseCache returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasUseCache() bool {
	if o != nil && o.UseCache != nil {
		return true
	}

	return false
}

// SetUseCache gets a reference to the given bool and assigns it to the UseCache field.
func (o *TelemetryDruidQueryContext) SetUseCache(v bool) {
	o.UseCache = &v
}

// GetPopulateCache returns the PopulateCache field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetPopulateCache() bool {
	if o == nil || o.PopulateCache == nil {
		var ret bool
		return ret
	}
	return *o.PopulateCache
}

// GetPopulateCacheOk returns a tuple with the PopulateCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetPopulateCacheOk() (*bool, bool) {
	if o == nil || o.PopulateCache == nil {
		return nil, false
	}
	return o.PopulateCache, true
}

// HasPopulateCache returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasPopulateCache() bool {
	if o != nil && o.PopulateCache != nil {
		return true
	}

	return false
}

// SetPopulateCache gets a reference to the given bool and assigns it to the PopulateCache field.
func (o *TelemetryDruidQueryContext) SetPopulateCache(v bool) {
	o.PopulateCache = &v
}

// GetUseResultLevelCache returns the UseResultLevelCache field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetUseResultLevelCache() bool {
	if o == nil || o.UseResultLevelCache == nil {
		var ret bool
		return ret
	}
	return *o.UseResultLevelCache
}

// GetUseResultLevelCacheOk returns a tuple with the UseResultLevelCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetUseResultLevelCacheOk() (*bool, bool) {
	if o == nil || o.UseResultLevelCache == nil {
		return nil, false
	}
	return o.UseResultLevelCache, true
}

// HasUseResultLevelCache returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasUseResultLevelCache() bool {
	if o != nil && o.UseResultLevelCache != nil {
		return true
	}

	return false
}

// SetUseResultLevelCache gets a reference to the given bool and assigns it to the UseResultLevelCache field.
func (o *TelemetryDruidQueryContext) SetUseResultLevelCache(v bool) {
	o.UseResultLevelCache = &v
}

// GetPopulateResultLevelCache returns the PopulateResultLevelCache field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetPopulateResultLevelCache() bool {
	if o == nil || o.PopulateResultLevelCache == nil {
		var ret bool
		return ret
	}
	return *o.PopulateResultLevelCache
}

// GetPopulateResultLevelCacheOk returns a tuple with the PopulateResultLevelCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetPopulateResultLevelCacheOk() (*bool, bool) {
	if o == nil || o.PopulateResultLevelCache == nil {
		return nil, false
	}
	return o.PopulateResultLevelCache, true
}

// HasPopulateResultLevelCache returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasPopulateResultLevelCache() bool {
	if o != nil && o.PopulateResultLevelCache != nil {
		return true
	}

	return false
}

// SetPopulateResultLevelCache gets a reference to the given bool and assigns it to the PopulateResultLevelCache field.
func (o *TelemetryDruidQueryContext) SetPopulateResultLevelCache(v bool) {
	o.PopulateResultLevelCache = &v
}

// GetBySegment returns the BySegment field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetBySegment() bool {
	if o == nil || o.BySegment == nil {
		var ret bool
		return ret
	}
	return *o.BySegment
}

// GetBySegmentOk returns a tuple with the BySegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetBySegmentOk() (*bool, bool) {
	if o == nil || o.BySegment == nil {
		return nil, false
	}
	return o.BySegment, true
}

// HasBySegment returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasBySegment() bool {
	if o != nil && o.BySegment != nil {
		return true
	}

	return false
}

// SetBySegment gets a reference to the given bool and assigns it to the BySegment field.
func (o *TelemetryDruidQueryContext) SetBySegment(v bool) {
	o.BySegment = &v
}

// GetFinalize returns the Finalize field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetFinalize() bool {
	if o == nil || o.Finalize == nil {
		var ret bool
		return ret
	}
	return *o.Finalize
}

// GetFinalizeOk returns a tuple with the Finalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetFinalizeOk() (*bool, bool) {
	if o == nil || o.Finalize == nil {
		return nil, false
	}
	return o.Finalize, true
}

// HasFinalize returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasFinalize() bool {
	if o != nil && o.Finalize != nil {
		return true
	}

	return false
}

// SetFinalize gets a reference to the given bool and assigns it to the Finalize field.
func (o *TelemetryDruidQueryContext) SetFinalize(v bool) {
	o.Finalize = &v
}

// GetChunkPeriod returns the ChunkPeriod field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetChunkPeriod() string {
	if o == nil || o.ChunkPeriod == nil {
		var ret string
		return ret
	}
	return *o.ChunkPeriod
}

// GetChunkPeriodOk returns a tuple with the ChunkPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetChunkPeriodOk() (*string, bool) {
	if o == nil || o.ChunkPeriod == nil {
		return nil, false
	}
	return o.ChunkPeriod, true
}

// HasChunkPeriod returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasChunkPeriod() bool {
	if o != nil && o.ChunkPeriod != nil {
		return true
	}

	return false
}

// SetChunkPeriod gets a reference to the given string and assigns it to the ChunkPeriod field.
func (o *TelemetryDruidQueryContext) SetChunkPeriod(v string) {
	o.ChunkPeriod = &v
}

// GetMaxScatterGatherBytes returns the MaxScatterGatherBytes field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetMaxScatterGatherBytes() int32 {
	if o == nil || o.MaxScatterGatherBytes == nil {
		var ret int32
		return ret
	}
	return *o.MaxScatterGatherBytes
}

// GetMaxScatterGatherBytesOk returns a tuple with the MaxScatterGatherBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetMaxScatterGatherBytesOk() (*int32, bool) {
	if o == nil || o.MaxScatterGatherBytes == nil {
		return nil, false
	}
	return o.MaxScatterGatherBytes, true
}

// HasMaxScatterGatherBytes returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasMaxScatterGatherBytes() bool {
	if o != nil && o.MaxScatterGatherBytes != nil {
		return true
	}

	return false
}

// SetMaxScatterGatherBytes gets a reference to the given int32 and assigns it to the MaxScatterGatherBytes field.
func (o *TelemetryDruidQueryContext) SetMaxScatterGatherBytes(v int32) {
	o.MaxScatterGatherBytes = &v
}

// GetMaxQueuedBytes returns the MaxQueuedBytes field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetMaxQueuedBytes() int32 {
	if o == nil || o.MaxQueuedBytes == nil {
		var ret int32
		return ret
	}
	return *o.MaxQueuedBytes
}

// GetMaxQueuedBytesOk returns a tuple with the MaxQueuedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetMaxQueuedBytesOk() (*int32, bool) {
	if o == nil || o.MaxQueuedBytes == nil {
		return nil, false
	}
	return o.MaxQueuedBytes, true
}

// HasMaxQueuedBytes returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasMaxQueuedBytes() bool {
	if o != nil && o.MaxQueuedBytes != nil {
		return true
	}

	return false
}

// SetMaxQueuedBytes gets a reference to the given int32 and assigns it to the MaxQueuedBytes field.
func (o *TelemetryDruidQueryContext) SetMaxQueuedBytes(v int32) {
	o.MaxQueuedBytes = &v
}

// GetSerializeDateTimeAsLong returns the SerializeDateTimeAsLong field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetSerializeDateTimeAsLong() bool {
	if o == nil || o.SerializeDateTimeAsLong == nil {
		var ret bool
		return ret
	}
	return *o.SerializeDateTimeAsLong
}

// GetSerializeDateTimeAsLongOk returns a tuple with the SerializeDateTimeAsLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetSerializeDateTimeAsLongOk() (*bool, bool) {
	if o == nil || o.SerializeDateTimeAsLong == nil {
		return nil, false
	}
	return o.SerializeDateTimeAsLong, true
}

// HasSerializeDateTimeAsLong returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasSerializeDateTimeAsLong() bool {
	if o != nil && o.SerializeDateTimeAsLong != nil {
		return true
	}

	return false
}

// SetSerializeDateTimeAsLong gets a reference to the given bool and assigns it to the SerializeDateTimeAsLong field.
func (o *TelemetryDruidQueryContext) SetSerializeDateTimeAsLong(v bool) {
	o.SerializeDateTimeAsLong = &v
}

// GetSerializeDateTimeAsLongInner returns the SerializeDateTimeAsLongInner field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetSerializeDateTimeAsLongInner() bool {
	if o == nil || o.SerializeDateTimeAsLongInner == nil {
		var ret bool
		return ret
	}
	return *o.SerializeDateTimeAsLongInner
}

// GetSerializeDateTimeAsLongInnerOk returns a tuple with the SerializeDateTimeAsLongInner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetSerializeDateTimeAsLongInnerOk() (*bool, bool) {
	if o == nil || o.SerializeDateTimeAsLongInner == nil {
		return nil, false
	}
	return o.SerializeDateTimeAsLongInner, true
}

// HasSerializeDateTimeAsLongInner returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasSerializeDateTimeAsLongInner() bool {
	if o != nil && o.SerializeDateTimeAsLongInner != nil {
		return true
	}

	return false
}

// SetSerializeDateTimeAsLongInner gets a reference to the given bool and assigns it to the SerializeDateTimeAsLongInner field.
func (o *TelemetryDruidQueryContext) SetSerializeDateTimeAsLongInner(v bool) {
	o.SerializeDateTimeAsLongInner = &v
}

// GetEnableParallelMerge returns the EnableParallelMerge field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetEnableParallelMerge() bool {
	if o == nil || o.EnableParallelMerge == nil {
		var ret bool
		return ret
	}
	return *o.EnableParallelMerge
}

// GetEnableParallelMergeOk returns a tuple with the EnableParallelMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetEnableParallelMergeOk() (*bool, bool) {
	if o == nil || o.EnableParallelMerge == nil {
		return nil, false
	}
	return o.EnableParallelMerge, true
}

// HasEnableParallelMerge returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasEnableParallelMerge() bool {
	if o != nil && o.EnableParallelMerge != nil {
		return true
	}

	return false
}

// SetEnableParallelMerge gets a reference to the given bool and assigns it to the EnableParallelMerge field.
func (o *TelemetryDruidQueryContext) SetEnableParallelMerge(v bool) {
	o.EnableParallelMerge = &v
}

// GetParallelMergeParallelism returns the ParallelMergeParallelism field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetParallelMergeParallelism() int32 {
	if o == nil || o.ParallelMergeParallelism == nil {
		var ret int32
		return ret
	}
	return *o.ParallelMergeParallelism
}

// GetParallelMergeParallelismOk returns a tuple with the ParallelMergeParallelism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetParallelMergeParallelismOk() (*int32, bool) {
	if o == nil || o.ParallelMergeParallelism == nil {
		return nil, false
	}
	return o.ParallelMergeParallelism, true
}

// HasParallelMergeParallelism returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasParallelMergeParallelism() bool {
	if o != nil && o.ParallelMergeParallelism != nil {
		return true
	}

	return false
}

// SetParallelMergeParallelism gets a reference to the given int32 and assigns it to the ParallelMergeParallelism field.
func (o *TelemetryDruidQueryContext) SetParallelMergeParallelism(v int32) {
	o.ParallelMergeParallelism = &v
}

// GetParallelMergeInitialYieldRows returns the ParallelMergeInitialYieldRows field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetParallelMergeInitialYieldRows() int32 {
	if o == nil || o.ParallelMergeInitialYieldRows == nil {
		var ret int32
		return ret
	}
	return *o.ParallelMergeInitialYieldRows
}

// GetParallelMergeInitialYieldRowsOk returns a tuple with the ParallelMergeInitialYieldRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetParallelMergeInitialYieldRowsOk() (*int32, bool) {
	if o == nil || o.ParallelMergeInitialYieldRows == nil {
		return nil, false
	}
	return o.ParallelMergeInitialYieldRows, true
}

// HasParallelMergeInitialYieldRows returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasParallelMergeInitialYieldRows() bool {
	if o != nil && o.ParallelMergeInitialYieldRows != nil {
		return true
	}

	return false
}

// SetParallelMergeInitialYieldRows gets a reference to the given int32 and assigns it to the ParallelMergeInitialYieldRows field.
func (o *TelemetryDruidQueryContext) SetParallelMergeInitialYieldRows(v int32) {
	o.ParallelMergeInitialYieldRows = &v
}

// GetParallelMergeSmallBatchRows returns the ParallelMergeSmallBatchRows field value if set, zero value otherwise.
func (o *TelemetryDruidQueryContext) GetParallelMergeSmallBatchRows() int32 {
	if o == nil || o.ParallelMergeSmallBatchRows == nil {
		var ret int32
		return ret
	}
	return *o.ParallelMergeSmallBatchRows
}

// GetParallelMergeSmallBatchRowsOk returns a tuple with the ParallelMergeSmallBatchRows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidQueryContext) GetParallelMergeSmallBatchRowsOk() (*int32, bool) {
	if o == nil || o.ParallelMergeSmallBatchRows == nil {
		return nil, false
	}
	return o.ParallelMergeSmallBatchRows, true
}

// HasParallelMergeSmallBatchRows returns a boolean if a field has been set.
func (o *TelemetryDruidQueryContext) HasParallelMergeSmallBatchRows() bool {
	if o != nil && o.ParallelMergeSmallBatchRows != nil {
		return true
	}

	return false
}

// SetParallelMergeSmallBatchRows gets a reference to the given int32 and assigns it to the ParallelMergeSmallBatchRows field.
func (o *TelemetryDruidQueryContext) SetParallelMergeSmallBatchRows(v int32) {
	o.ParallelMergeSmallBatchRows = &v
}

func (o TelemetryDruidQueryContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GrandTotal != nil {
		toSerialize["grandTotal"] = o.GrandTotal
	}
	if o.SkipEmptyBuckets != nil {
		toSerialize["skipEmptyBuckets"] = o.SkipEmptyBuckets
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.QueryId != nil {
		toSerialize["queryId"] = o.QueryId
	}
	if o.UseCache != nil {
		toSerialize["useCache"] = o.UseCache
	}
	if o.PopulateCache != nil {
		toSerialize["populateCache"] = o.PopulateCache
	}
	if o.UseResultLevelCache != nil {
		toSerialize["useResultLevelCache"] = o.UseResultLevelCache
	}
	if o.PopulateResultLevelCache != nil {
		toSerialize["populateResultLevelCache"] = o.PopulateResultLevelCache
	}
	if o.BySegment != nil {
		toSerialize["bySegment"] = o.BySegment
	}
	if o.Finalize != nil {
		toSerialize["finalize"] = o.Finalize
	}
	if o.ChunkPeriod != nil {
		toSerialize["chunkPeriod"] = o.ChunkPeriod
	}
	if o.MaxScatterGatherBytes != nil {
		toSerialize["maxScatterGatherBytes"] = o.MaxScatterGatherBytes
	}
	if o.MaxQueuedBytes != nil {
		toSerialize["maxQueuedBytes"] = o.MaxQueuedBytes
	}
	if o.SerializeDateTimeAsLong != nil {
		toSerialize["serializeDateTimeAsLong"] = o.SerializeDateTimeAsLong
	}
	if o.SerializeDateTimeAsLongInner != nil {
		toSerialize["serializeDateTimeAsLongInner"] = o.SerializeDateTimeAsLongInner
	}
	if o.EnableParallelMerge != nil {
		toSerialize["enableParallelMerge"] = o.EnableParallelMerge
	}
	if o.ParallelMergeParallelism != nil {
		toSerialize["parallelMergeParallelism"] = o.ParallelMergeParallelism
	}
	if o.ParallelMergeInitialYieldRows != nil {
		toSerialize["parallelMergeInitialYieldRows"] = o.ParallelMergeInitialYieldRows
	}
	if o.ParallelMergeSmallBatchRows != nil {
		toSerialize["parallelMergeSmallBatchRows"] = o.ParallelMergeSmallBatchRows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidQueryContext) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidQueryContext := _TelemetryDruidQueryContext{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidQueryContext); err == nil {
		*o = TelemetryDruidQueryContext(varTelemetryDruidQueryContext)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "grandTotal")
		delete(additionalProperties, "skipEmptyBuckets")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "queryId")
		delete(additionalProperties, "useCache")
		delete(additionalProperties, "populateCache")
		delete(additionalProperties, "useResultLevelCache")
		delete(additionalProperties, "populateResultLevelCache")
		delete(additionalProperties, "bySegment")
		delete(additionalProperties, "finalize")
		delete(additionalProperties, "chunkPeriod")
		delete(additionalProperties, "maxScatterGatherBytes")
		delete(additionalProperties, "maxQueuedBytes")
		delete(additionalProperties, "serializeDateTimeAsLong")
		delete(additionalProperties, "serializeDateTimeAsLongInner")
		delete(additionalProperties, "enableParallelMerge")
		delete(additionalProperties, "parallelMergeParallelism")
		delete(additionalProperties, "parallelMergeInitialYieldRows")
		delete(additionalProperties, "parallelMergeSmallBatchRows")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidQueryContext struct {
	value *TelemetryDruidQueryContext
	isSet bool
}

func (v NullableTelemetryDruidQueryContext) Get() *TelemetryDruidQueryContext {
	return v.value
}

func (v *NullableTelemetryDruidQueryContext) Set(val *TelemetryDruidQueryContext) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidQueryContext) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidQueryContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidQueryContext(val *TelemetryDruidQueryContext) *NullableTelemetryDruidQueryContext {
	return &NullableTelemetryDruidQueryContext{value: val, isSet: true}
}

func (v NullableTelemetryDruidQueryContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidQueryContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
