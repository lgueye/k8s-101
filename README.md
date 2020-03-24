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


