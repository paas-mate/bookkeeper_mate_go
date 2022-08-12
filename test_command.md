### start
```bash
docker run --rm -e REMOTE_MODE=false ttbb/bookkeeper:mate
```
### start with port expose
```bash
docker run -p 2181:2181 -p 3181:3181 -p 4181:4181 --rm -d -e REMOTE_MODE=false ttbb/bookkeeper:mate
```
### localhost
```bash
docker run --rm -e REMOTE_MODE=false -e BOOKKEEPER_ADVERTISED_ADDRESS="localhost" ttbb/bookkeeper:mate
```
### localhost, start with port expose
```bash
docker run -p 2181:2181 -p 3181:3181 -p 4181:4181 --rm -e REMOTE_MODE=false -e BOOKKEEPER_ADVERTISED_ADDRESS="localhost" -e BOOKIE_ROOT_LOGGER="DEBUG,INFO,CONSOLE" ttbb/bookkeeper:mate
```
### localhost, port expose, tls
```bash
docker run -p 2181:2181 -p 3181:3181 -p 4181:4181 --rm -e BOOKKEEPER_TLS_ENABLE=true -e REMOTE_MODE=false -e BOOKKEEPER_ADVERTISED_ADDRESS="localhost" -e BOOKIE_ROOT_LOGGER="DEBUG,INFO,CONSOLE" ttbb/bookkeeper:mate
```
### simple cluster
```bash
docker run --rm -e CLUSTER_INIT=true -e CLUSTER_ENABLE=true -e ZK_ADDR="localhost:2181" -e META_DATA_SERVICE_URI="localhost:2181" ttbb/bookkeeper:mate
docker run --rm -d -e CLUSTER_ENABLE=true -e REMOTE_MODE=false -e ZK_ADDR="localhost:2181" -e META_DATA_SERVICE_URI="localhost:2181" ttbb/bookkeeper:mate
```
### simple cluster tls
```bash
docker run --rm -d -e CLUSTER_ENABLE=true -e ZOOKEEPER_TLS_ENABLE=true -e BOOKKEEPER_TLS_ENABLE=true -e REMOTE_MODE=false -e ZK_ADDR="localhost:2181" -e META_DATA_SERVICE_URI="localhost:2181" ttbb/bookkeeper:mate
```
### cluster
```bash
docker run --rm -e REMOTE_MODE=false -e ZK_ADDR=172.16.0.56:2181,172.16.0.55:2181,172.16.0.54:2181 -e HOSTNAME=test-bookkeeper-0 -d ttbb/bookkeeper:cluster
```