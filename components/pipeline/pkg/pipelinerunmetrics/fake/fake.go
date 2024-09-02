/*
Copyright 2021 The Tekton Authors

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

package fake

import (
	"context"

	_ "github.com/tektoncd/pipeline/pkg/client/injection/informers/pipeline/v1/pipelinerun/fake" // Make sure the fake pipelinerun informer is setup
	"github.com/tektoncd/pipeline/pkg/pipelinerunmetrics"
	"k8s.io/client-go/rest"
	"knative.dev/pkg/injection"
)

func init() {
	injection.Fake.RegisterClient(func(ctx context.Context, _ *rest.Config) context.Context { return pipelinerunmetrics.WithClient(ctx) })
	injection.Fake.RegisterInformer(pipelinerunmetrics.WithInformer)
}
