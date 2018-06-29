# Kubernetes 201

This project will demonstrate the basic programmable infrastructure features of Kubernetes by using custom controller.

## Getting Started

Download and install the pre-requisites.

To see how we made use of the code generators please see: [code generation](generate.md)

## Prerequisites

### Windows

Go is installed and on the path.
Docker for Windows Edge Version with Kubernetes turned on (or Docker Toolbox + Minikube + Kubectl is installed).
Kubectl is available on the path. e.g. `path=%path%;C:\Program Files\Docker\Docker\Resources\bin\`
Make is installed and is added to the path. e.g. `path=%path%;C:\Program Files (x86)\GnuWin32\bin`
Tools like `sleep` are added to the path. e.g. `path=%path%;C:\Program Files\Git\usr\bin`

### Mac

Go is installed and on the path.
Docker for Mac edge version with Kubernetes turned on.
Make is installed and is added to the path. 
Tools like `sleep` are added to the path.

## Demo

[![Watch the demo](/demo/Kubernetes-201_First_Frame.png?raw=true)](/demo/Kubernetes-201.mp4?raw=true)

In this demo we are going to:
* deploy a custom resource definition
* deploy a custom resource definition validation
* deploy a custom controller
* deploy a custom resource and show that the controller has generated a deployment for the custom resource
* delete the deployment and show the the custom controller has generated another deployment for the custom resource

Make sure that Kubernetes does not have any pre-existing demo configuration.

```
make cleanup
```

Then to run demo:

```
make demo-1
```

## Presentation

[![Presentation](/presentation/programmable-infrastructure-with-k8s.jpg?raw=true)](http://nbviewer.jupyter.org/github/contino/kubernetes-201/blob/master/presentation/programmable-infrastructure-with-k8s.pdf)

## Contributing

Please read [CONTRIBUTING.md](https://github.com/contino/kubernetes-201) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **Hibri Marzook** - *Initial work* - [hibri](https://github.com/hibri)
* **Taliesin Sisson** - *Initial work* - [taliesins](https://github.com/taliesins)

See also the list of [contributors](https://github.com/contino/kubernetes-201/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

