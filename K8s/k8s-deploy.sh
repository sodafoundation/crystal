# Run this file from crsytal folder.
kubectl apply -f rbac.yaml
kubectl apply -f soda-secret.yaml

#1. Create all the config Maps from files.
kubectl create configmap multicloud-config --from-file=../examples/crsytal.conf -n soda-crsytal
kubectl create configmap tidb-config --from-file=../s3/tidbconf/tidb.toml -n soda-crsytal
kubectl create configmap s3-config --from-file=../s3/pkg/conf/s3.toml -n soda-crsytal
kubectl create configmap tidb-sql --from-file=../s3/pkg/conf/tidb.sql -n soda-crsytal
kubectl create configmap s3-sql --from-file=../s3/pkg/conf/tidb.sql -n soda-crsytal
kubectl create configmap yig-sql --from-file=../s3/pkg/datastore/yig/conf/yig.sql -n soda-crsytal


#2. Create all pods with no dependencies.
kubectl apply -f zookeeper-deployment.yaml
kubectl apply -f redis-deployment.yaml
kubectl apply -f tidb-deployment.yaml
kubectl apply -f kafka-deployment.yaml

kubectl apply -f block-deployment.yaml
kubectl apply -f file-deployment.yaml
kubectl apply -f backend-deployment.yaml
kubectl apply -f s3api-deployment.yaml
kubectl apply -f api-deployment.yaml
kubectl apply -f datamover-deployment.yaml
kubectl apply -f dataflow-deployment.yaml
kubectl apply -f s3-deployment.yaml

#4 Expose the services
kubectl apply -f zookeeper-service.yaml
kubectl apply -f tidb-service.yaml
kubectl apply -f redis-service.yaml
kubectl apply -f kafka-service.yaml
kubectl apply -f api-service.yaml
kubectl apply -f s3api-service.yaml
kubectl apply -f s3-service.yaml
kubectl apply -f block-service.yaml
kubectl apply -f backend-service.yaml
kubectl apply -f file-service.yaml
kubectl apply -f datamover-service.yaml
kubectl apply -f dataflow-service.yaml

#5. Get all the deployed objects
kubectl get all

