## Docker network
Essa é a forma com que podemos fazer os nossos containers se comunicarem.

### Criando uma network bredg e conectando containers
Vamos criar uma network chamada `minha-rede`
```bash
$ docker network create minha-rede
$ docker network ls
```
Agora vamos subir um container do ngnix.
```bash
$ docker run nginx
```

E em um terminal separado vamos subir um container do ubuntu com o curl
```bash
$ docker run ubuntu --it /bin/bash
$~# apt-get update && apt-get install curl
```

Se tentarmos acessar o nginx de dentro do container do ubuntu via curl não vamos conseguir
pos eles estão em redes diferentes.

Para que isso seja possível, temos que conectalos a mesma rede, essa que criamos:
```bash
$ docker network connect minha-rede <id-container-ubuntu>    
$ docker network connect minha-rede <id-container-nginx>7
```

Agora dentro do container do ubuntu vai ser possível interagir com o container do nginx
```bash
$~# curl http://<id-container-nginx>
```

# Adicionado um container a uma rede no momento da sua criação
Dado que a rede já esteja criada, podemos inicializar nossos containers dentro dessa rede usando
```bash
$ docker run -d --network minha-rede nginx 
```

# Criando container na network host
Quando criamos um container dessa forma, ele usar a rede da nossa máquina e não
precisamos fazer um port-forward para vincular as portas do container, apenas acessando direto no navegador.
```bash
$ docker run -d --network host nginx 
$ curl http://localhost:80
```