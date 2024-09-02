// Copyright 2020 The Tekton Authors
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

//go:generate rm -rf results_go_proto
//go:generate mkdir results_go_proto
//go:generate protoc --go_out=results_go_proto --go-grpc_out=results_go_proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --grpc-gateway_out=results_go_proto --grpc-gateway_opt paths=source_relative -I$GOPATH/src/github.com/googleapis/googleapis -I../pipeline/v1beta1 -I. api.proto resources.proto

package v1alpha2
