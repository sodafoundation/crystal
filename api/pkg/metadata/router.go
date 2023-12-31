// Copyright 2021 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metadata

import (
	"github.com/emicklei/go-restful"
	"github.com/micro/go-micro/v2/client"
)

func RegisterRouter(ws *restful.WebService) {
	handler := NewAPIService(client.DefaultClient)
	ws.Route(ws.POST("/backends/sync").To(handler.SyncMetadata)).Doc("Sync metdata from cloud")
	ws.Route(ws.POST("/backends/{backendID}/sync").To(handler.SyncMetadata)).Doc(
		"Sync metdata from cloud for a particular backend")
	ws.Route(ws.GET("/backends/metadata").To(handler.ListMetadata)).Doc("Show metdata details")
}
