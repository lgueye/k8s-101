# k8s-101

## Requirements

MacOS. Install multipass, and start a VM instance.
```bash
brew cask install multipass
multipass launch --name k3s --mem 1G --disk 5G
```

Get a shell on the instance, get and start k3s node.
```
multipass shell k3s
curl -sfL https://get.k3s.io | sh -
```

## Docker

```bash
cd docker
docker image build -t helloworld:1.0 .
```

```bash
docker container run --publish 8080:8080 --name hw helloworld:1.0
```

Tester dans le browser à l'adresse ``http://localhost:8080``.

## Kubernetes

### TP 1

### TP 2

### TP 3 - Premier pod sur Kubernetes

Les images suivantes sont dispo sur docker hub
```bash
pevablanchard/k8s-101:1.0
pevablanchard/k8s-101:2.0
```
