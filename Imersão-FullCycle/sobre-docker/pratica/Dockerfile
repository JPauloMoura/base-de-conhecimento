#Descrição:
# servidor web simples que retorna um hello world
# excute: docker run -p 3333:3333 jp/server-go
FROM golang:1.17

# diretorio onde o projeto é iniciado
WORKDIR /src/app

# baixe os pacotes
COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o bin/main .
# execute o binario criado
CMD [ "./bin/main" ]