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

package v1alpha1

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_AddonSetDefaults_DefaultParamsWithValues(t *testing.T) {

	ta := &TektonAddon{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "name",
			Namespace: "namespace",
		},
		Spec: TektonAddonSpec{
			CommonSpec: CommonSpec{
				TargetNamespace: "namespace",
			},
		},
	}

	ta.SetDefaults(context.TODO())
	assert.Equal(t, 3, len(ta.Spec.Params))

	params := ParseParams(ta.Spec.Params)
	value, ok := params[PipelineTemplatesParam]
	assert.Equal(t, true, ok)
	assert.Equal(t, "true", value)
}

func Test_AddonSetDefaults_ResolverTaskIsFalse(t *testing.T) {

	ta := &TektonAddon{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "name",
			Namespace: "namespace",
		},
		Spec: TektonAddonSpec{
			CommonSpec: CommonSpec{
				TargetNamespace: "namespace",
			},
			Addon: Addon{
				Params: []Param{
					{
						Name:  "resolverTasks",
						Value: "false",
					},
				},
			},
		},
	}

	ta.SetDefaults(context.TODO())
	assert.Equal(t, 3, len(ta.Spec.Params))

	params := ParseParams(ta.Spec.Params)
	value, ok := params[ResolverTasks]
	assert.Equal(t, true, ok)
	assert.Equal(t, "false", value)
}

func Test_AddonSetDefaults_ResolverStepActions(t *testing.T) {

	ta := &TektonAddon{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "name",
			Namespace: "namespace",
		},
		Spec: TektonAddonSpec{
			CommonSpec: CommonSpec{
				TargetNamespace: "namespace",
			},
			Addon: Addon{
				Params: []Param{
					{
						Name:  "resolverStepActions",
						Value: "false",
					},
				},
			},
		},
	}

	ta.SetDefaults(context.TODO())
	assert.Equal(t, 3, len(ta.Spec.Params))

	params := ParseParams(ta.Spec.Params)
	value, ok := params[ResolverStepActions]
	assert.Equal(t, true, ok)
	assert.Equal(t, "false", value)
}
