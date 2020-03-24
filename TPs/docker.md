# Docker

Build d'image.
```bash
cd docker
sudo docker image build -t helloworld:1.0 .
```

Run le container.
```bash
sudo docker container run -d --publish 8080:8080 --name hw helloworld:1.0
```

Lister les conteneurs actifs.
```bash
sudo docker container list
```

Log.
```bash
sudo docker logs hw
```

Test
```bash
curl localhost:8080
```

Stop, start
```bash
sudo docker container stop hw
sudo docker container start hw
```

