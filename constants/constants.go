/*
Copyright 2021 clusterpedia Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package constants

const (
	// ClusterPediaAPIPath is the extra path of apiserver
	ClusterPediaAPIPath = "/apis/clusterpedia.io/v1alpha1/resources"
	// ClusterAPIPath for a specific cluster
	ClusterAPIPath = "/clusters/"

	// all of the clusterpedia selectLabel
	SearchLabelOwner      = "search.clusterpedia.io/owner"
	SearchLabelNames      = "search.clusterpedia.io/names"
	SearchLabelClusters   = "search.clusterpedia.io/clusters"
	SearchLabelNamespaces = "search.clusterpedia.io/namespaces"
	SearchLabelOrderBy    = "search.clusterpedia.io/orderby"

	SearchLabelOwners         = "internalstorage.clusterpedia.io/owner-uid"
	SearchLabelOwnerSeniority = "internalstorage.clusterpedia.io/owner-seniority"

	SearchLabelWithContinue       = "search.clusterpedia.io/with-continue"
	SearchLabelWithRemainingCount = "search.clusterpedia.io/with-remaining-count"

	SearchLabelLimit  = "search.clusterpedia.io/limit"
	SearchLabelOffset = "search.clusterpedia.io/offset"

	ShadowAnnotationClusterName          = "shadow.clusterpedia.io/cluster-name"
	ShadowAnnotationGroupVersionResource = "shadow.clusterpedia.io/gvr"

	OrderByDesc = "_desc"
)
