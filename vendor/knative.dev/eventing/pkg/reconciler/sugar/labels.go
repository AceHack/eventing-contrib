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

package sugar

const (
	// Label to enable Knative Eventing in a Namespace.
	// DeprecatedInjectionLabelKey, fully qualified label keys recommended.
	// Please use InjectionLabelKey.
	DeprecatedInjectionLabelKey = "knative-eventing-injection"
	InjectionLabelKey           = "eventing.knative.dev/injection"
	InjectionEnabledLabelValue  = "enabled"
	InjectionDisabledLabelValue = "disabled"
)

func InjectionLabelKeys() []string {
	// Note: InjectionLabelKey needs to be first, order matters for conflicts
	// with DeprecatedInjectionLabelKey. InjectionLabelKey should win.
	return []string{
		InjectionLabelKey,
		DeprecatedInjectionLabelKey,
	}
}

func InjectionEnabledLabels() map[string]string {
	return map[string]string{
		InjectionLabelKey: InjectionEnabledLabelValue,
	}
}

func InjectionDisabledLabels() map[string]string {
	return map[string]string{
		InjectionLabelKey: InjectionDisabledLabelValue,
	}
}
