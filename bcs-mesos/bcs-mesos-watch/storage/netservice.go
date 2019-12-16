/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package storage

import (
	"fmt"
	"time"

	"bk-bcs/bcs-common/common/blog"
	netservicetypes "bk-bcs/bcs-services/bcs-netservice/pkg/netservice/types"
)

// NetServiceHandler handle netservice resources for storage.
type NetServiceHandler struct {
	oper      DataOperator
	dataType  string
	ClusterID string
}

// GetType returns data type.
func (h *NetServiceHandler) GetType() string {
	return h.dataType
}

// dataNode := fmt.Sprintf("/bcsstorage/v1/mesos/dynamic/cluster_resources/clusters/%s/%s", h.ClusterID, h.dataType)

// CheckDirty cleans dirty data in remote bcs-storage.
func (h *NetServiceHandler) CheckDirty() error {
	// do nothing.
	return nil
}

// Add handle add event for netservice resources.
func (h *NetServiceHandler) Add(data interface{}) error {
	// do nothing.
	return nil
}

// Delete handle delete event for netservice resources.
func (h *NetServiceHandler) Delete(data interface{}) error {
	// do nothing.
	return nil
}

// Delete handle delete event for netservice resources.
func (h *NetServiceHandler) Update(data interface{}) error {
	ipStatic := data.(*netservicetypes.NetStatic)

	started := time.Now()
	dataNode := fmt.Sprintf("/bcsstorage/v1/mesos/dynamic/cluster_resources/clusters/%s/%s", h.ClusterID, h.dataType)

	if err := h.oper.CreateDCNode(dataNode, ipStatic, "PUT"); err != nil {
		blog.V(3).Infof("IPPoolStatic update node %s, err %+v", dataNode, err)
		reportStorageMetrics(dataTypeIPPoolStatic, actionPut, statusFailure, started)
		return err
	}
	reportStorageMetrics(dataTypeIPPoolStatic, actionPut, statusSuccess, started)

	return nil
}
