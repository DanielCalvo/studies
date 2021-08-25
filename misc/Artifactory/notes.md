
- https://www.jfrog.com/confluence/display/JFROG/Installing+Artifactory

```shell
export JFROG_HOME="/tmp/"
mkdir -p $JFROG_HOME/artifactory/var/etc/
cd $JFROG_HOME/artifactory/var/etc/
touch ./system.yaml
#chown -R $UID:$GID $JFROG_HOME/artifactory/var
chmod -R 777 $JFROG_HOME/artifactory/var
```

```shell
docker container prune --force #To remove old artifactory containers
docker run --name artifactory -v $JFROG_HOME/artifactory/var/:/var/opt/jfrog/artifactory -d -p 8081:8081 -p 8082:8082 releases-docker.jfrog.io/jfrog/artifactory-oss:latest
```

- http://localhost:8081/
- admin / password


Artifactory is not particularly exciting, it is just a registry, and a commercial one at that. Can't trial hosting docker containers on the OSS version. Boo!