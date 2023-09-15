
kubectl delete svc mongodb-0-service -n soda-crsytal
kubectl delete svc mongodb-1-service -n soda-crsytal
kubectl delete svc mongodb-2-service -n soda-crsytal

kubectl delete deployment  mongodb-0 -n soda-crsytal
kubectl delete deployment  mongodb-1 -n soda-crsytal
kubectl delete deployment  mongodb-2 -n soda-crsytal

# Delete pv, pvc
kubectl patch pvc mongo-0-pv-claim -p '{"metadata":{"finalizers": []}}' --type=merge -n soda-crsytal
kubectl patch pvc mongo-1-pv-claim -p '{"metadata":{"finalizers": []}}' --type=merge -n soda-crsytal
kubectl patch pvc mongo-2-pv-claim -p '{"metadata":{"finalizers": []}}' --type=merge -n soda-crsytal

kubectl delete pvc mongo-0-pv-claim -n soda-crsytal
kubectl delete pvc mongo-1-pv-claim -n soda-crsytal
kubectl delete pvc mongo-2-pv-claim -n soda-crsytal

kubectl delete pv mongo-0-pv-volume -n soda-crsytal
kubectl delete pv mongo-1-pv-volume -n soda-crsytal
kubectl delete pv mongo-2-pv-volume -n soda-crsytal
