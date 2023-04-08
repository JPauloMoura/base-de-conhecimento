## Evolução dos hosts de aplicações

#### 1) Vários servidores físicos
  - Cada servidor era responsavel por um serviço.
  - Cada serviço necessitava de um SO.
  - Para haver uma comunicação entre os serviços, deviamos colocar-los em rede.
  - Custo excessivo de hardware, manutenção e eletricidade.
  - Servidores parrudos com com uma unica aplicação rodando com capacidade mínima
    afim de suporta possíveis demadas de recursos.
  - Os servidores passavam muito tempo com recursos ociosos.

#### 2) Virtualização
  - Hypervision: criação de uma máquina virtual em cima do SO possibilitando
    virtualizar recursos físicos da máquina como: HD, CPU, RAM etc..
  - Assim podemos ter vários servidores dentro de uma unica máquina.
  - Diminuição de luz, rede, manutenção.
  - Aproveitamento melhor do hardware.
  - Para cada máquina virtual é necessario um SO, o que também consome recursos.
  - Cada SO também necessitar de configuração, manutenção e seguração.

#### 3) Containers
  - Eles funcionam junto ao nosso SO base, e é nele que nossa aplicação vai rodar.
  - Temos então uma máquina com um SO e vários containes que dividem os recursos.
  - São muito mais leves já que não possui um SO para cada um.
  - Podemos limitar o uso de recursos da nossa máquina para cada container.
  - Também podemos instalar programas em versões diferentes em containers diferentes.



