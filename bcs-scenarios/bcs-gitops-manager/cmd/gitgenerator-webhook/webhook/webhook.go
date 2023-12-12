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

package webhook

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	clusterclient "github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager/cmd/gitgenerator-webhook/options"
	"github.com/Tencent/bk-bcs/bcs-scenarios/bcs-gitops-manager/pkg/store"
)

// AdmissionWebhookServer defines the webhook server to check application
// generated by git-generator
type AdmissionWebhookServer struct {
	HttpHandler *gin.Engine
	cfg         *options.Config
	argoStore   store.Store
}

// NewAdmissionWebhookServer create the webhook-server instance
func NewAdmissionWebhookServer(cfg *options.Config) *AdmissionWebhookServer {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	server := &AdmissionWebhookServer{
		cfg:         cfg,
		HttpHandler: r,
	}
	server.registerRouter()
	return server
}

// Init will init the argocd client
func (s *AdmissionWebhookServer) Init() error {
	s.argoStore = store.NewStore(&store.Options{
		Service:      s.cfg.ArgoService,
		User:         s.cfg.ArgoUser,
		Pass:         s.cfg.ArgoPass,
		Cache:        false,
		CacheHistory: false,
	})
	if err := s.argoStore.Init(); err != nil {
		return errors.Wrapf(err, "init argocd stroe failed")
	}
	return nil
}

// Run the webhook server with http export
func (s *AdmissionWebhookServer) Run() error {
	pair, err := tls.LoadX509KeyPair(s.cfg.TlsCert, s.cfg.TlsKey)
	if err != nil {
		return err
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.cfg.ListenAddr, s.cfg.ListenPort),
		Handler: s.HttpHandler,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{pair},
		},
	}
	blog.Infof("Server serving on: %s:%d", s.cfg.ListenAddr, s.cfg.ListenPort)
	if err := srv.ListenAndServeTLS("", ""); err != nil {
		return err
	}
	return nil
}

func (s *AdmissionWebhookServer) webhookAllow(ctx *gin.Context, allowed bool, reqUID types.UID, errMsg string) {
	ctx.JSON(http.StatusOK, &v1.AdmissionReview{
		Response: &v1.AdmissionResponse{
			UID:     reqUID,
			Allowed: allowed,
			Result: &metav1.Status{
				Message: errMsg,
			},
		},
	})
}

func (s *AdmissionWebhookServer) registerRouter() {
	s.HttpHandler.Any("/check", s.check)
}

func (s *AdmissionWebhookServer) check(ctx *gin.Context) {
	ar := new(v1.AdmissionReview)
	if err := ctx.BindJSON(ar); err != nil {
		blog.Errorf("marshal request body failed, err: ", err.Error())
		s.webhookAllow(ctx, true, "", "")
		return
	}
	// Check request whether is nil.
	req := ar.Request
	if req != nil {
		blog.Infof("Received request. UID: %s, Name: %s, Operation: %s, Kind: %v.", req.UID, req.Name,
			req.Operation, req.Kind)
	} else {
		blog.Errorf("request is nil")
		s.webhookAllow(ctx, true, "", "")
		return
	}

	switch req.Kind.Kind {
	case application.ApplicationKind:
		if err := s.checkApplication(ctx, req.Object.Raw); err != nil {
			blog.Errorf("UID: %s, check application failed: %s", req.UID, err.Error())
			s.webhookAllow(ctx, false, req.UID, err.Error())
			return
		}
	default:
		blog.Warnf("Unknown request kind: %s", req.Kind.Kind)
	}
	s.webhookAllow(ctx, true, req.UID, "")
}

var (
	defaultTimeout = 15 * time.Second
)

func (s *AdmissionWebhookServer) checkApplication(ctx context.Context, bs []byte) error {
	app := new(v1alpha1.Application)
	if err := json.Unmarshal(bs, app); err != nil {
		return errors.Wrapf(err, "unmarshal failed with '%s'", string(bs))
	}
	belongApplicationSet := false
	for i := range app.ObjectMeta.OwnerReferences {
		owner := app.ObjectMeta.OwnerReferences[i]
		if owner.Kind == "ApplicationSet" {
			belongApplicationSet = true
			break
		}
	}
	if !belongApplicationSet {
		blog.Infof("application '%s' not belong to application, ignore it", app.Name)
		return nil
	}
	proj := app.Spec.Project
	cxt, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	argoProj, err := s.argoStore.GetProject(cxt, proj)
	if err != nil {
		return errors.Wrapf(err, "get project '%s' failed", proj)
	}
	if argoProj == nil {
		return errors.Errorf("project '%s' not exist", proj)
	}

	var repoBelong bool
	var repoProj string
	for i := range app.Spec.Sources {
		appSource := app.Spec.Sources[i]
		repoUrl := appSource.RepoURL
		repoProj, repoBelong, err = s.checkRepositoryBelongProject(ctx, repoUrl, proj)
		if err != nil {
			return errors.Wrapf(err, "check repo '%s' belong to project '%s' failed", repoUrl, repoProj)
		}
		if !repoBelong {
			return errors.Errorf("repo '%s' project is '%s', not same as '%s'", repoUrl, repoProj, proj)
		}
	}
	if app.Spec.Source != nil {
		repoUrl := app.Spec.Source.RepoURL
		repoProj, repoBelong, err = s.checkRepositoryBelongProject(ctx, repoUrl, proj)
		if err != nil {
			return errors.Wrapf(err, "check repo '%s' belong to project '%s' failed", repoUrl, repoProj)
		}
		if !repoBelong {
			return errors.Errorf("repo '%s' project is '%s', not same as '%s'", repoUrl, repoProj, proj)
		}
	}

	cls := app.Spec.Destination.Server
	var argoCls *v1alpha1.Cluster
	argoCls, err = s.argoStore.GetCluster(ctx, &clusterclient.ClusterQuery{
		Server: cls,
	})
	if err != nil {
		return errors.Wrapf(err, "get cluster '%s' failed", cls)
	}
	if argoCls == nil {
		return errors.Errorf("cluster '%s' not exist", cls)
	}
	if argoCls.Project != proj {
		return errors.Errorf("cluster '%s' project is '%s', not same as '%s'", cls, argoCls.Project, proj)
	}
	blog.Infof("check application '%s' success", app.Name)
	return nil
}

func (s *AdmissionWebhookServer) checkRepositoryBelongProject(ctx context.Context, repoUrl,
	project string) (string, bool, error) {
	repo, err := s.argoStore.GetRepository(ctx, repoUrl)
	if err != nil {
		return "", false, errors.Wrapf(err, "get repo '%s' failed", repoUrl)
	}
	if repo == nil {
		return "", false, fmt.Errorf("repo '%s' not found", repoUrl)
	}
	if repo.Project != project {
		return "", false, nil
	}
	return repo.Project, true, nil
}
