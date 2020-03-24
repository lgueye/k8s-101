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
docker image build -t bulletinboard:1.0 .
```

```bash
docker container run --publish 8080:8080 --name bb bulletinboard:1.0
```

Tester dans le browser à l'adresse ``http://localhost:8080``.
