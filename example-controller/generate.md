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


## generate.go
```
package main
func main() {}
```

Initialize dep
```
dep init
```

## Gopkg.toml
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

## doc.go
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

## register.go

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

## types.go

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

## Generate scaffolding for client and api
```
vendor/k8s.io/code-generator/generate-groups.sh all github.com/<company>/<project name>/pkg/client github.com/<company>/<project name>/pkg/apis "<controller name>:<version>"
```

```
vendor/k8s.io/code-generator/generate-groups.sh all github.com/taliesins/t/pkg/client github.com/taliesins/t/pkg/apis "samplecontroller:v1alpha1"
```



