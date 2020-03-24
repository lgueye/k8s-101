# Docker

Accéder au répertoire docker.
```bash
cd k8s-101/docker
```
Toutes les commandes ci-dessous, sauf mention contraire, supposent que l'on est situé dans ce répertoire.

## Images

Jeter un coup d'oeil au Dockerfile.
```bash
less Dockerfile
```

Build d'image.
```bash
sudo docker image build -t helloworld:1.0 .
```

Lister les images disponibles localement (registry local).
```bash
sudo docker image ls
```

## Conteneurs

Run le container.
```bash
sudo docker container run -d --publish 8080:8080 --name hw helloworld:1.0
```

Test
```bash
curl localhost:8080
```

Log.
```bash
sudo docker logs hw
```

Stopper le container.
```bash
sudo docker container stop hw
```

L'appli peut prendre un fichier de conf, situé dans ``helloworld/config/name``.
Pour run le container en montant le répertoire de config
```bash
sudo docker container run -d --publish 8080:8080 -v $(pwd)/helloworld/config:/app/config --name hw2 helloworld:1.0
```

### Misc.

Lister les conteneurs actifs.
```bash
sudo docker container ls
```

Lister tous les conteneurs, y compris ceux qui ne tournent pas.
```bash
sudo docker container ls -a
```

Supprimer le conteneur ``hw``.
```bash
sudo docker container rm hw
```

## Volumes

Lister les volumes
```bash
sudo docker volume ls
```

Créer un volume nommé
```bash
sudo docker volume create hw-volume
```

Run le container en montant le volume nommé, e.g., pour persister les logs
```bash
sudo docker container run -d --publish 8080:8080 -v hw-volume:/logs --name hw3 helloworld:1.0
curl localhost:8080
curl localhost:8080/sickz
```

Le volume peut être partagé, e.g.
```bash
sudo docker run -it --rm -v hw-volume:/hw-volume busybox
cat /hw-volume/out
```

## Réseau

Lister les configurations réseaux possibles pour docker.
```bash
sudo docker network ls
```

Le réseau par défaut pour tout conteneur est ``bridge``.
```bash
sudo docker network inspect bridge
```

On y voit notamment le sous-réseau des conteneurs. 
Si le conteneur hw est actif, on peut également voir son ip (entre autres).

Sur la VM, on peut voir l'interface virtuelle correspondante.
```bash
ip a | grep docker
```

