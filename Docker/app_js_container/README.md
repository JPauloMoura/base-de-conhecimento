# App Js Container
#### Aplicação simples utilizando node.js.

### Aqui utilizamos essa simples aplicação para aplicar os conceitos de container.
1. Criamos um Dockerfile para definir como a imagem deverá ser contruida.
2. Para realiza a constrição da imagem da aplicação executamos:

```bash
$ docker build -t app-js-container .
```

3. Após a criação da imagem, podemos criar um container com base nela:
```bash
$ docker run -p 3000:3000 app-js-containe
```
> Nesse exemplo executamos realizando um port-forward entre a porta do container e a nossa máquina

4. Após isso é só acessar a porta que o container já estará em execução.
```bash
$ curl http://localhost:3000
```
