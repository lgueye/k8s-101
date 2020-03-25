# k8s-101

## Structure

Le repo est organisé comme suit:
- TPs : les TP en markdown. Il vaut mieux les lire depuis le browser.
    - docker.md : quelques rudiments de docker
    - k8s-00.md : commandes de base ``kubectl``, collecte d'information, format d'affichage
    - k8s-01.md : ressources de base, pod, replica set, deploiement, configmap, secret, ingress
    - k8s-02.md : persistence, network policies, service account, rbac
- docker
    - Dockerfile : un exemple de dockerfile
    - helloworld : un bout de code en go qu'on déploie dans le kube
- k8s : ressources k8s en yaml pour les TPs.

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
