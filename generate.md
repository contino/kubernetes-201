# Pre-requisites

## Download and install dep

```
go get -u github.com/golang/dep/cmd/dep
```

## Make project directory
Template:
```
mkdir C:\data\go\src\github.com\<company>\<project name>
cd C:\data\go\src\github.com\<company>\<project name>
```

Example:
```
mkdir C:\data\go\src\github.com\taliesins\t
cd C:\data\go\src\github.com\taliesins\t
```


## /generate.go
```
package main
func main() {}
```

Initialize dep
```
dep init
```

## /Gopkg.toml
```
required = [
  "k8s.io/code-generator/cmd/client-gen"
]
 
[[constraint]]
  name = "k8s.io/client-go"
  version = "kubernetes-1.10.0"
 
[[constraint]]
  name = "k8s.io/api"
  version = "kubernetes-1.10.0"
 
[[constraint]]
  name = "k8s.io/apimachinery"
  version = "kubernetes-1.10.0"
 
[[constraint]]
  name = "k8s.io/code-generator"
  version = "kubernetes-1.10.0"
 
[prune]
  non-go = true
  go-tests = true
  unused-packages = true
 
  [[prune.project]]
    name = "k8s.io/code-generator"
    unused-packages = false
    non-go = false
    go-tests = false
 
  [[prune.project]]
    name = "k8s.io/gengo"
    unused-packages = false
    non-go = false
    go-tests = false
```

## Run
```
dep ensure
```

# Generate codegen templates

Template:
```
mkdir -p C:\data\go\src\github.com\<company>\<project name>\pkg\apis\<controller name>\<version>
cd C:\data\go\src\github.com\<company>\<project name>\pkg\apis\<controller name>\<version>
```

Example:
```
mkdir -p C:\data\go\src\github.com\taliesins\t\pkg\apis\samplecontroller\v1alpha1
cd C:\data\go\src\github.com\taliesins\t\pkg\apis\samplecontroller\v1alpha1
```

## /pkg/apis/\<controller name>/\<version>/doc.go
Template:
```
// +k8s:deepcopy-gen=package
 
 
// Package <version> is the <version> version of the API.
// +groupName=<controller name><company namespace>
package <version>
```

Example:
```
// +k8s:deepcopy-gen=package
 
 
// Package v1alpha1 is the v1alpha1 version of the API.
// +groupName=samplecontroller.t.taliesins.github.com
package v1alpha1
```

## /pkg/apis/\<controller name>/\<version>/types.go

Template:
```
package v1alpha1
 
 
import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
 
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
 
// <Resource Name> is a specification for a <Resource Name> resource
type <Resource Name> struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
 
    Spec   <Resource Name>Spec   `json:"spec"`
    Status <Resource Name>Status `json:"status"`
}
 
// <Resource Name>Spec is the spec for a <Resource Name> resource
type <Resource Name>Spec struct {
    <Resource spec properties>
}
 
// <Resource Name>Status is the status for a <Resource Name> resource
type <Resource Name>Status struct {
    <Resource status properties>
}
 
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
 
// <Resource Name>List is a list of <Resource Name> resources
type <Resource Name>List struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata"`
 
    Items []<Resource Name> `json:"items"`
}
```

Example:
```
package v1alpha1
 
 
import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
 
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
 
// Foo is a specification for a Foo resource
type Foo struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`
 
    Spec   FooSpec   `json:"spec"`
    Status FooStatus `json:"status"`
}
 
// FooSpec is the spec for a Foo resource
type FooSpec struct {
    DeploymentName string `json:"deploymentName"`
    Replicas       *int32 `json:"replicas"`
}
 
// FooStatus is the status for a Foo resource
type FooStatus struct {
    AvailableReplicas int32 `json:"availableReplicas"`
}
 
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
 
// FooList is a list of Foo resources
type FooList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata"`
 
    Items []Foo `json:"items"`
}
```

## /pkg/apis/\<controller name>/\<version>/register.go

Template:
```
package <version>
 
 
import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/runtime/schema"
 
    <controller name> "github.com/<company>/<project name>/pkg/apis/<controller name>"
)
 
// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: <controller name>.GroupName, Version: "<version>"}
 
// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
    return SchemeGroupVersion.WithKind(kind).GroupKind()
}
 
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
    return SchemeGroupVersion.WithResource(resource).GroupResource()
}
 
var (
    SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
    AddToScheme   = SchemeBuilder.AddToScheme
)
 
// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
    scheme.AddKnownTypes(SchemeGroupVersion,
        &<Resource Name>{},
        &<Resource Name>List{},
        <Any other resources>
    )
    metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
    return nil
}
```

Example:
```
package v1alpha1
 
 
import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/apimachinery/pkg/runtime/schema"
 
    samplecontroller "github.com/taliesins/t/pkg/apis/samplecontroller"
)
 
// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: samplecontroller.GroupName, Version: "v1alpha1"}
 
// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
    return SchemeGroupVersion.WithKind(kind).GroupKind()
}
 
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
    return SchemeGroupVersion.WithResource(resource).GroupResource()
}
 
var (
    SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
    AddToScheme   = SchemeBuilder.AddToScheme
)
 
// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
    scheme.AddKnownTypes(SchemeGroupVersion,
        &Foo{},
        &FooList{},
    )
    metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
    return nil
}
```

# Hack scripts

## Make hack directory

Template:
```
mkdir C:\data\go\src\github.com\<company>\<project name>\hack
cd C:\data\go\src\github.com\<company>\<project name>\hack
```

Example:
```
mkdir C:\data\go\src\github.com\taliesins\t\hack
cd C:\data\go\src\github.com\taliesins\t\hack
```

## /hack/boilerplate.go.txt

Template:
```
<e.g. license>
```

Example:
```
/*
Copyright The Kubernetes Authors.

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
```

## /hack/update-codegen.sh

Template:
```
#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

#Run from project root
SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..

$SCRIPT_ROOT/vendor/k8s.io/code-generator/generate-groups.sh all github.com/<company>/<project name>/pkg/client github.com/<company>/<project name>/pkg/apis "<controller name>:<version>" --go-header-file $SCRIPT_ROOT/hack/boilerplate.go.txt
```

Example:
```
#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

#Run from project root
SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..

$SCRIPT_ROOT/vendor/k8s.io/code-generator/generate-groups.sh all github.com/taliesins/t/pkg/client github.com/taliesins/t/pkg/apis "samplecontroller:v1alpha1" --go-header-file $SCRIPT_ROOT/hack/boilerplate.go.txt
```

## /hack/verify-codegen.sh

```
#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..

DIFFROOT="${SCRIPT_ROOT}/pkg"
TMP_DIFFROOT="${SCRIPT_ROOT}/_tmp/pkg"
_tmp="${SCRIPT_ROOT}/_tmp"

cleanup() {
  rm -rf "${_tmp}"
}
trap "cleanup" EXIT SIGINT

cleanup

mkdir -p "${TMP_DIFFROOT}"
cp -a "${DIFFROOT}"/* "${TMP_DIFFROOT}"

"${SCRIPT_ROOT}/hack/update-codegen.sh"
echo "diffing ${DIFFROOT} against freshly generated codegen"
ret=0
diff -Naupr "${DIFFROOT}" "${TMP_DIFFROOT}" || ret=$?
cp -a "${TMP_DIFFROOT}"/* "${DIFFROOT}"
if [[ $ret -eq 0 ]]
then
  echo "${DIFFROOT} up to date."
else
  echo "${DIFFROOT} is out of date. Please run hack/update-codegen.sh"
  exit 1
fi

```

# Generate scaffolding for client and api

```
hack/update-codegen.sh
```

After the first run of `update-codegen.sh` you should run `dep ensure` again. 



# Use the generated code
# Create a main.go in the root
# Create a controller.go and a controller_test.go
Check if the imports map to the auto-generated client code

# Run a dep ensure

Run go test

# Make the resources available
# Create the custom resource definition

```
kubectl apply -f artifacts/crd.yaml
```

# Create the custom resource definition validation as well

```
kubectl apply -f artifacts/crd-validation.yaml
```
