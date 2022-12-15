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

package cloudvpc

import (
	"context"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/cmd/util"
	cloudvpcMgr "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/cloud_vpc"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

func newListCloudRegionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "listCloudRegions",
		Aliases: []string{"listRegions"},
		Short:   "list cloud regions from bcs-cluster-manager",
		Run:     listCloudRegions,
	}

	cmd.Flags().StringVarP(&cloudID, "cloudID", "c", "", `cloud ID (required)`)
	cmd.MarkFlagRequired("cloudID")

	return cmd
}

func listCloudRegions(cmd *cobra.Command, args []string) {
	resp, err := cloudvpcMgr.New(context.Background()).ListCloudRegions(types.ListCloudRegionsReq{
		CloudID: cloudID,
	})
	if err != nil {
		klog.Fatalf("list cloud regions failed: %v", err)
	}

	util.Output2Json(resp.Data)
}
