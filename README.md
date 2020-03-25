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


Install some stuff on the VM instance.
```bash
sudo apt install git docker.io
```

Clone the repo.
```bash
git clone https://github.com/pevab/k8s-101.git
cd k8s-101
```

Les images suivantes sont dispo sur docker hub
```bash
pevablanchard/k8s-101:1.0
pevablanchard/k8s-101:2.0
```

## Utiles

Dans un terminal:
- ``Ctrl+C`` stop le process courant.
- ``Ctrl+D`` vous déconnecte du shell courant.
