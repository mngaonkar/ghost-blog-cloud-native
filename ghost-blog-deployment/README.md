## Ghost blog application deployment
```
kubectl create -f ghost-deployment.yaml
```

## Create blog service
```
kubectl create -f ghost-service.yaml
```

## Create Ingress namespace
```
kubectl create -f ingress-namespace.yaml
```

## Create Ingress deployment
```
kubectl create -f ingress-deployment.yaml
```

## Create Ingress default backend service
```
kubectl create -f ingress-service.yaml 
```

## Create Ingress rules for Ghost blog application
```
kubectl create -f ghost-ingress.yaml
```
