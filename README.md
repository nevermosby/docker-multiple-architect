# 如何通过docker构建多架构的容器镜像

## 问题
How can we put multiple Docker images, each supporting a different architecture, behind the sametag?

## 解决思路
What if this manifest file contained a list of manifests, so that the Docker Engine could pick the one that it matches at runtime? That’s exactly how the manifest is built for a multi-arch image. This type of manifest is called a manifest list.

容器镜像manifest提供不同架构镜像的信息，供docker engine在拉取镜像时判断。


## docker manifest方案
```bash
# AMD64
docker build -t robolwq/multiarch-example:manifest-amd64 --build-arg ARCH=amd64/ .
docker push robolwq/multiarch-example:manifest-amd64

# ARM64
docker build -t robolwq/multiarch-example:manifest-arm64v8 --build-arg ARCH=arm64v8/ .
docker push robolwq/multiarch-example:manifest-arm64v8

# combine the multiple image manifest into one
docker manifest create \
    robolwq/multiarch-example:manifest-latest \
    --amend robolwq/multiarch-example:manifest-amd64 \
    --amend robolwq/multiarch-example:manifest-arm64v8

# push the combined image to docker hub
docker manifest push robolwq/multiarch-example:manifest-latest                         

```

## docker buildx方案
```bash
# error case: cannot build with the driver "docker" 
> docker buildx build --push \
    --platform linux/arm64,linux/amd64 --tag robolwq/multiarch-example:buildx-latest .
                                                                                                   
ERROR: multiple platforms feature is currently not supported for docker driver. Please switch to a different driver (eg. "docker buildx create --use")

# create a new builder instance with the driver "docker-container"
docker buildx create --use --name build-node-example --driver docker-container 

# build with the specified builder instance
docker buildx build --push -t robolwq/multiarch-example:buildx --platform linux/amd64,linux/arm64 --builder=build-node-example .

```

