# Ações dentro do kubernets
	informações do cluster
		$ kubectl cluster-info

		$ kubectl get pods
		$ kubectl get namespaces 
		$ kubectl get pods -n kube-system
		$ kubectl describe pod kube-apiserver-minikube -n kube-system
		$ kubectl get pods --all-namespaces

	Criar um novo namespace
		$ kubectl create namespace teste
	Cria pod no namespace teste
		$ kubectl run nginx --generator=run-pod/v1 --image=nginx -n teste
	Verificar os pods em tempo real usando -w
		$ kubectl get pods -n teste -w
	Mais informaçoes sobre o pod
		$ kubectl get pods -n teste -o wide
