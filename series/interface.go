// Licensed to LinDB under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. LinDB licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package series

import (
	"github.com/lindb/roaring"
)

//go:generate mockgen -source ./interface.go -destination=./interface_mock.go -package=series

// MetricMetaSuggester represents the suggest ability for metricNames and tagKeys.
// default max limit of suggestions is set in constants
type MetricMetaSuggester interface {
	// SuggestMetrics returns suggestions from a given prefix of metricName
	SuggestMetrics(namespace, metricPrefix string, limit int) ([]string, error)
	// SuggestTagKeys returns suggestions from given metricName and prefix of tagKey
	SuggestTagKeys(namespace, metricName, tagKeyPrefix string, limit int) ([]string, error)
}

// TagValueSuggester represents the suggest ability for tagValues.
// default max limit of suggestions is set in constants
type TagValueSuggester interface {
	// SuggestTagValues returns suggestions from given tag key id and prefix of tagValue
	SuggestTagValues(tagKeyID uint32, tagValuePrefix string, limit int) []string
}

// Filter represents the query ability for filtering seriesIDs by expr from an index of tags.
type Filter interface {
	// GetSeriesIDsByTagValueIDs gets series ids by tag value ids for spec metric's tag key
	GetSeriesIDsByTagValueIDs(tagKeyID uint32, tagValueIDs *roaring.Bitmap) (*roaring.Bitmap, error)
	// GetSeriesIDsForTag gets series ids for spec metric's tag key
	GetSeriesIDsForTag(tagKeyID uint32) (*roaring.Bitmap, error)
	// GetSeriesIDsForMetric gets series ids for spec metric name
	GetSeriesIDsForMetric(namespace, metricName string) (*roaring.Bitmap, error)
	// GetGroupingContext returns the context of group by
	GetGroupingContext(tagKeyIDs []uint32, seriesIDs *roaring.Bitmap) (GroupingContext, error)
}
