# UTILIZANDO O KIND PARA TESTAR O KUBERNETS LOCAL

# Criar um cluter utilizando o kind
	$ kind create cluster --name cluster-kind-k8s

# Listar os clusters
	$ kind get clusters

# Listagem
	$ kubectl get nodes
	$ kubectl get pods
	$ kubectl get pods --watch // acompanha o status do pod
	$ kubectl get namespaces
	$ kubectl get services
		aqui podemos usar os mesmos comandos do k8s

# Lista os contextos que temos na nossa maquina
	$ kubectl config get-contexts

# Mostra qual o nosso contexto atual
	$ kubectl config current-context

# Alterar entre contextos
	$ kubectl config use-context <nome-do-contexto>

# Criando um pod
	$ kubectl kubectl run <nome-do-pod> --image=<nome>:<versão>

# Detalhes de um pod
	$ kubectl describe pod <nome-do-pod>

# Alterar a imagem  de um pod
	$ kubectl edit pod <nome-do-pod>

# Entrar no bash de um pod 
	$ kubectl exec -it <nome-do-pod> -- bash 

# Serviçes
	- atua como uma camada na frente dos pods
	- como os pods não tem ip fixo, os services criam um ip para atuar como proxy que encaminham as requisições para os pods cuja label tenha sido setada.
	- ao definir um service do tipo ClusterIP para um pod, só conseguimos bater nesse ip caso estejamos dentro do mesmo cluster.
	- para que nosso seviço seja acessado fora no cluste, temos que definir seu type como NodePort, assim podemos acessar esse serviço atravez do node .

