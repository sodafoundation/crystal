#Delete the services
kubectl delete service  zookeeper -n soda-crsytal
kubectl delete service  tidb -n soda-crsytal
kubectl delete service  redis -n soda-crsytal
kubectl delete service  kafka -n soda-crsytal
kubectl delete service  api -n soda-crsytal
kubectl delete service  s3api -n soda-crsytal
kubectl delete service  s3 -n soda-crsytal
kubectl delete service  backend -n soda-crsytal
kubectl delete service  block -n soda-crsytal
kubectl delete service  file -n soda-crsytal
kubectl delete service  datamover -n soda-crsytal
kubectl delete service  dataflow -n soda-crsytal

# Delete the deployments
kubectl delete deployment zookeeper -n soda-crsytal
kubectl delete deployment redis -n soda-crsytal
kubectl delete deployment tidb -n soda-crsytal
kubectl delete deployment kafka -n soda-crsytal
kubectl delete deployment block -n soda-crsytal
kubectl delete deployment file -n soda-crsytal
kubectl delete deployment backend -n soda-crsytal
kubectl delete deployment s3api -n soda-crsytal
kubectl delete deployment s3 -n soda-crsytal
kubectl delete deployment api -n soda-crsytal
kubectl delete deployment datamover -n soda-crsytal
kubectl delete deployment dataflow -n soda-crsytal

# Delete config maps
kubectl delete configmap multicloud-config -n soda-crsytal
kubectl delete configmap tidb-config -n soda-crsytal
kubectl delete configmap s3-config -n soda-crsytal
kubectl delete configmap tidb-sql -n soda-crsytal
kubectl delete configmap s3-sql -n soda-crsytal
kubectl delete configmap yig-sql -n soda-crsytal


