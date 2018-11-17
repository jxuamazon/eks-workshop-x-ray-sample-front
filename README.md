# eks-workshop-x-ray-sample-front

The example [AWS X-Ray](https://aws.amazon.com/xray/) instrumented front-end service in the [EKS Workshop](https://eksworkshop.com/)

**Command reference**

Deploy
```
kubectl apply -f x-ray-sample-front-k8s.yml
```

Delete
```
kubectl delete deployment x-ray-sample-front-k8s

kubectl delete service x-ray-sample-front-k8s
```

