// Copyright 2020 Red Hat, Inc. and/or its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"github.com/kiegroup/kogito-operator/apis"
	"github.com/kiegroup/kogito-operator/apis/app/v1beta1"
	"github.com/kiegroup/kogito-operator/core/kogitosupportingservice"
	"github.com/kiegroup/kogito-operator/core/test"
	"github.com/kiegroup/kogito-operator/meta"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestReconcileKogitoSupportingService_Reconcile(t *testing.T) {
	replicas := int32(1)
	instance := &v1beta1.KogitoSupportingService{
		ObjectMeta: v1.ObjectMeta{Name: kogitosupportingservice.DefaultJobsServiceName, Namespace: t.Name()},
		Spec: v1beta1.KogitoSupportingServiceSpec{
			ServiceType:       api.JobsService,
			KogitoServiceSpec: v1beta1.KogitoServiceSpec{Replicas: &replicas},
		},
	}
	cli := test.NewFakeClientBuilder().AddK8sObjects(instance).Build()

	r := NewKogitoSupportingServiceReconciler(cli, meta.GetRegisteredSchema())
	test.AssertReconcileMustNotRequeue(t, r, instance)
}

func TestContains(t *testing.T) {
	allServices := []api.ServiceType{
		api.MgmtConsole,
		api.JobsService,
		api.TrustyAI,
	}
	testService := api.DataIndex

	assert.False(t, contains(allServices, testService))
}

// Check is the testService is available in the slice of allServices
func contains(allServices []api.ServiceType, testService api.ServiceType) bool {
	for _, a := range allServices {
		if a == testService {
			return true
		}
	}
	return false
}
