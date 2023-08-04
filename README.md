# plugin-as-apiserver

Working on a POC for a Grafana Plugin as an aggregated API Server for the data plane.

## Important Intel

The project was initially setup with apiserver-builder-alpha but since there isn't much facility to add NonGoRestfulMux
routes other than just grabbing the GenericAPIServer in the build function chain, I started copying a bunch of
Subresource handling from grafana/grafana-apiserver to get something like a query endpoint working. Still a WIP.

I have also gone ahead and upgraded K8s deps to 0.27 as the builder gave me a pretty old 0.23 with go 1.17 go.mod file.

## Prerequisites

You will need the following tools for codegen and apiserver-boot utility.

```shell
go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.12.1
go install sigs.k8s.io/apiserver-builder-alpha/cmd/apiserver-boot@v1.23

## To add the DeepCopy implementations of included resources
controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."

## Also, ensure you have etcd install
brew install etcd
```

## To develop

```shell
## Note that this command is a hit or miss when it comes to performing cleanup of etcd. 
## If it leaves it running, any subsequent calls
## print un-obvious errors saying etcd is not able to determine default host. In reality, the problem can be easily resolved
## by killing the orphaned etcd process
apiserver-boot run local
```