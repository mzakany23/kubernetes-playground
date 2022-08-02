## Setup a Cluster

1. Create a cluster 
```bash
cat <<EOF | kind create cluster --name $name --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
EOF
```

2. Add your deployment
```bash
kb apply -k ./deploy/test
```

3. Expose ingress to it (e.g. nginx ingress)
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

4. Add a service mesh (linkerd)
```bash
kb apply -k ./deploy/service-mesh
```

5. Opt a deployment into it
```bash
kb get deploy -n default hello -o yaml | linkerd inject - | kb apply -f -
```