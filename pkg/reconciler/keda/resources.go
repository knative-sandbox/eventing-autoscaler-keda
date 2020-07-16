/*
Copyright 2020 The Knative Authors

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

package keda

import (
	"fmt"
	"strconv"

	kedav1aplha1 "github.com/kedacore/keda/pkg/apis/keda/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	awssqsv1alpha1 "knative.dev/eventing-contrib/awssqs/pkg/apis/sources/v1alpha1"
	kafkav1beta1 "knative.dev/eventing-contrib/kafka/source/pkg/apis/sources/v1beta1"
)

const (
	defaultPollingInterval = 30
	defaultCooldownPeriod  = 300
	defaultMinReplicaCount = 0
	defaultMaxReplicaCount = 100
)

var (
	KedaSchemeGroupVersion = schema.GroupVersion{Group: "keda.sh", Version: "v1alpha1"}
)

func GenerateScaledObject(obj metav1.Object, gvk schema.GroupVersionKind, scaleTargetName string, triggers []kedav1aplha1.ScaleTriggers) (*kedav1aplha1.ScaledObject, error) {

	cooldownPeriod, err := GetInt32ValueFromMap(obj.GetAnnotations(), KedaAutoscalingCooldownPeriodAnnotation, defaultCooldownPeriod)
	if err != nil {
		return nil, err
	}
	pollingInterval, err := GetInt32ValueFromMap(obj.GetAnnotations(), KedaAutoscalingPollingIntervalAnnotation, defaultPollingInterval)
	if err != nil {
		return nil, err
	}
	minReplicaCount, err := GetInt32ValueFromMap(obj.GetAnnotations(), AutoscalingMinScaleAnnotation, defaultMinReplicaCount)
	if err != nil {
		return nil, err
	}
	maxReplicaCount, err := GetInt32ValueFromMap(obj.GetAnnotations(), AutoscalingMaxScaleAnnotation, defaultMaxReplicaCount)
	if err != nil {
		return nil, err
	}

	return &kedav1aplha1.ScaledObject{
		ObjectMeta: metav1.ObjectMeta{
			Name:      generateScaledObjectName(obj),
			Namespace: obj.GetNamespace(),
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(obj, gvk),
			},
		},
		Spec: kedav1aplha1.ScaledObjectSpec{
			PollingInterval: pollingInterval,
			CooldownPeriod:  cooldownPeriod,
			MinReplicaCount: minReplicaCount,
			MaxReplicaCount: maxReplicaCount,
			ScaleTargetRef: &kedav1aplha1.ScaleTarget{
				Name: scaleTargetName,
			},
			Triggers: triggers,
		},
	}, nil
}

func generateScaledObjectName(obj metav1.Object) string {
	return fmt.Sprintf("so-%s", string(obj.GetUID()))
}

func GetInt32ValueFromMap(dict map[string]string, key string, defaultValue int32) (*int32, error) {
	val, ok := dict[key]
	if !ok {
		return &defaultValue, nil
	}
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Expected value for annotation: "+key+" should be integer but got "+val, err)
	}
	i32 := int32(i)
	return &i32, nil
}

func SupportedCRD(gvk schema.GroupVersionKind) bool {
	switch gvk {
	case kafkav1beta1.SchemeGroupVersion.WithKind("KafkaSource"):
		return true
	case awssqsv1alpha1.SchemeGroupVersion.WithKind("AwsSqsSource"):
		return true
	}
	return false
}
