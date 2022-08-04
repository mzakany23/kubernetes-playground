# Kubernetes Playground

The Goal is to be able to spin up a local kubernetes dev environment to test one off projects.

## Tecnologies
- [local development with kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [yaml mgmt with kustomize](https://kubectl.docs.kubernetes.io/guides/introduction/kustomize/)
- [opt in lightweight service mesh](https://linkerd.io/2.11/getting-started/)
- [gitops with flux](https://fluxcd.io/docs/get-started/)
- [local docker registries with kind](https://kind.sigs.k8s.io/docs/user/local-registry/)

## Build the sample app

1. Build and tag it
```bash
docker build . -t hello-server
docker tag hello-server:latest localhost:5001/hello-server:latest
```

2. Push to the local repo (created in below step)
```bash
docker push localhost:5001/hello-server:latest
```

## Setup a kubernetes test environment

1. Create a local cluster with [kind](https://kind.sigs.k8s.io/docs/user/quick-start/) and a local registry so you can test images locally
```bash
./bin/kind-create-cluster.sh
```

2. Expose ingress to it (e.g. nginx ingress)
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

3. Add your deployment
```bash 
kb apply -k ./deploy/test
```

## Light weight service mesh with linkderd

1. Add a service mesh (linkerd)
```bash
kb apply -k ./deploy/service-mesh
```

2. Opt a deployment into it
```bash
kb get deploy -n default hello -o yaml | linkerd inject - | kb apply -f -
```

3. Open dashboard
```bash
linkerd viz install | kubectl apply -f -
linkerd check
linkerd viz dashboard &
```
