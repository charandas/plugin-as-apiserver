/*
Copyright 2023.

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

package main

import (
	"k8s.io/klog"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"

	// +kubebuilder:scaffold:resource-imports
	charandasv0alpha1 "github.com/charandas/plugin-as-apiserver/pkg/apis/charandas/v0alpha1"
)

func main() {
	charandasv0alpha1.AddToScheme(Scheme)

	shHandler := &SubresourceHandler{}

	err := builder.APIServer.
		// +kubebuilder:scaffold:resource-register
		WithResource(&charandasv0alpha1.GoogleSheetConnection{}).
		WithLocalDebugExtension().
		WithServerFns(func(s *builder.GenericAPIServer) *builder.GenericAPIServer {
			s.Handler.NonGoRestfulMux.Handle("/apis/ext.charandas.example.com", shHandler)
			s.Handler.NonGoRestfulMux.HandlePrefix("/apis/ext.charandas.example.com/", shHandler)
			return s
		}).
		WithOptionsFns(func(options *builder.ServerOptions) *builder.ServerOptions {
			options.RecommendedOptions.CoreAPI = nil
			options.RecommendedOptions.Admission = nil
			options.RecommendedOptions.Authorization = nil
			return options
		}).
		Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
