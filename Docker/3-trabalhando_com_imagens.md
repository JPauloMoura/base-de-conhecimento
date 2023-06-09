## Trabalhando com imagens Docker 
#### Lista com alguns comandos úteis:
  _ | O que faz                           |     Comando    
  --| -------------------------------     |  ---------------     
  1 | Verificar os contêineres em execução |`$ docker ps`
  2 | Verificar todos os contêineres       |`$ docker ps -a`
  3 | Acessar o terminal de um container|`$ docker run -it <container>`
  4 | Inicia um container                 |`$ docker start <id_container>`
  5 | Para um container em execução       |`$ docker stop <id_container>  //demora alguns segundos`<br> `$ docker stop -t 0 <id_container> // para imediatamente`
  6 | Inicia um container já com um terminal integrado|`$ docker start -a -i <id_container>`
  7 | Entra no terminal de um container quer esteja rodando|`$ docker exec -it <nome_container> bash`
  8 | Remover um container                |`$ docker rm <id_container>`
  9 | Remover todos os contêineres parados |`$ docker container prune`
  10 | Remover uma imagem                  |  `$ docker rmi <nome_img>`
  11 | Baixando a imagem de um servidor estático.<br> Nesse caso ele trava o terminal pois o conteiner fica em execução|`$ docker run dockersamples/static-site`
  12 | Baixar e executar um container em background|`$ docker run -d <nome_imagem>`
  13 | Executar um container linkando as portas da máquina</br> a do container de forma aleatória: use a flag "-P"|`$ docker run -d -P dockersamples/static-site`
  14 | Definindo uma porta especifica para acessar o container: `[local]:[container]`| `$ docker run -d -p 1233:80 dockersamples/static-site`
  15 | Verifica as portas usadas por um container|`$ docker port <id_container>`
  16 | Renomeia um container no momento da sua crição</br> usando a flag "--name"|`$ docker run -d -P --name meu-site dockersamples/static-site`
  16 | `--rm` deleta um container após sua execução | `$ docker run --rm dockersamples/static-site`
  17 | Permite passar variáveis de ambiente com a flag "-e"|`$ docker run -d -p 3001:80 -e AUTHOR="jpaulo" dockersamples/static-site`
  18 | Verifica apenas os IDSs dos contêineres em execução|`$ docker ps -q`
  18 | Verifica os logs de um container. `-n` define o numero de linhas |`$ docker container logs -n 10 <id/nome_do_containet>`
  19 | Para todos os contêineres em execução </br> concatenando os comandos| `$ docker stop -t 0 $(docker ps -q) `

---
### Criando um container do mongo db
Para isso podemos da uma olhada na documentação da imagem no [docker hub](https://hub.docker.com/_/mongo)

Vamos então executar o seguinte comando:
```sh
$ docker run -d --name container-mongo -p 27018:27017 --rm -e MONGO_INITDB_ROOT_USERNAME=jp -e MONGO_INITDB_ROOT_PASSWORD=<senha> mongo
```

Passo a passo do comando!
- `docker run` é usado para criação de um container.
- ` -d` permite executar o container em backgroud.
- ` --name container-mongo` com isso damos um nome para o nosso container.
- ` -p 27018:27017` vincula nossa porta local "27018" a porta do container "27017".
- ` --rm` faz com que o container seja deletado quando ele para a sua execução.
- `e MONGO_INITDB_ROOT_USERNAME=jp -e MONGO_INITDB_ROOT_PASSWORD=<senha>` essa é a forma como passamos variaveis de ambiente para o container.
- `mongo` por ultimo colocamos o nome da imagem que o container vai usar para ser criado.

Você pode usar o robo3t para se conectar e manipular o banco :).

# Como as imagens  docker são criadas?
Elas são criadas atravez de camadas reaproveitaveis.
ex: Se criamos duas imagens e elas compartilham algo em comum, por exemplo o SO
elas não teram que ter uma comada individual pra cada, mas apenas referencia-la

 | camada A |       | camada T |
 | camada B |   ->  | camada B |
 | camada C |       | camada F |

Para criar imagens docker utilizamos aquivos `Dockerfile` onde descrevermos passo a passo como ela deve ser criada.
```yaml
FROM ubuntu
RUN apt-get update
RUN apt-get install curl --yes
```
` $ docker build -t ubuntu-com-curl ./Dockerfile`
Essa imagem tem 3 camadas:
- usa como base a imagem do unbuntu
- roda o comando de update
- roda o comanda para instalação do curl
Quando realizamos a construção dessa imagem cada step é colocado em cache, assim
caso executemos o build novamente,vamos aproveitar a camada que já existe.

O problema disso nesse com base nesse exemplo é que se alterarmos a instalação do curl para outra lib
é ele não existir nos repositorios do apt-get atualizados que estão no cache, teremos problemas
ela deve ser criada.
```yaml
FROM ubuntu # em cache
RUN apt-get update # em cache
RUN apt-get install vim --yes # latest
```

Para resolver esse problema basta colocarmos esses passos na mesma camada
ela deve ser criada.
```yaml
FROM ubuntu # em cache
RUN apt-get update && apt-get install curl --yes # latest
```
Assim cada vez que essa camada tiver alguma alteração, tudo nela sera atualizado.

> Como cada caso é um caso, termos que identifificar quando faz ou não sentido usar o cache