/*
Copyright 2024.

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

// Package sdk imports all sub packages to build all of them when calling `go install', `go build'
// or `go get' commands.

package sdk

import (
	_ "github.com/capitalonline/cds-cloudos-go-sdk/services/bandwidthpackage"
	_ "github.com/capitalonline/cds-cloudos-go-sdk/services/eip"
	_ "github.com/capitalonline/cds-cloudos-go-sdk/services/eks"
	_ "github.com/capitalonline/cds-cloudos-go-sdk/services/natgateway"
	_ "github.com/capitalonline/cds-cloudos-go-sdk/services/subnet"
	_ "github.com/capitalonline/cds-cloudos-go-sdk/services/vpc"
)
