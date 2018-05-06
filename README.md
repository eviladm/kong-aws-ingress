# K8s Kong ingress

Goal: setup a k8s cluster in AWS and provision services with kong as ingress controller

Technology used:
Linux
AWS (ec2, bucket, elb)
Kops (or kubespray)
Kubernetes
Kong
DNS outside AWS

Flow-Setup:
kong.kain.tech/* -> kong-proxy ELB -> kong-ingress -> backend services

Steps:
Install k8s with kops or kubespray, both tools will assist with creating your .kube/config
(Kops is easier in cloud deployments.) lmgtfy kops aws or kubespray aws are good starting points.
Verify with kubectl cluster-status or kops verify cluster $CLUSTERNAME
Once your cluster is ready create some namespace and deploy kong
Download the official kong setup yaml for minikube and change the kong-proxy service to LoadBalancer to get a public ELB IP:
```curl -o kong.yaml https://raw.githubusercontent.com/Kong/kubernetes-ingress-controller/master/deploy/single/all-in-one-postgres.yaml```

```
name: kong-proxy
…
  type: LoadBalancer
```

Deploy kong:
```
kubectl create -f kong.yaml
```

Wait a few minutes, verify with kubectl get svc -nkong
You should see something like:
```
kong-proxy                LoadBalancer   100.68.24.58     a282753fb514f...   80:30676/TCP,443:31513/TCP   1h
```

Now we got our public ELB and we can create a CNAME kong.kain.tech to point to a282753fb514f…

## Deploy our sample applications, services and ingress rules

ingress flow:
kong.kain.tech/* -> default backend container
kong.kain.tech/container -> go backend container

```
kubectl create -f app.yaml
kubectl create -f ingress.yaml
```


Notes:
Nano not enough RAM
Micro also to small


