# CLEAN ARCHTECTURE
	Dominio
		O coração do software é sua regra de negocio 

# NÃO MISTURE !!
Complexidade do negocio:
		Calcular juros
Complecidade acidental:
		api rest, go, mongo DB ....

# CLEAN ARCH :
	FOCAR NO CORAÇÃO DO NEGOCIO E
	DEIXAR AS COMPLEXIDADES EXTERNAS POR ULTIMO

# ENTITY
	- regras de negocio da empresa
	- não mudam com frequencia
	- são universais em qualquer lugar do sistema
	- protegem o negocio da empresa
	- possui ESTADO e COMPORTAMENTOS
	- não é so GETs e SETs
	- um entidate possui comportamentos 
#USER CASES
	- usa as ntidades e acessa intefaces externas para realizar seu fluxo
	- fluxos da aplicação e suas regras
	- realizar venda
	- gerar nota fiscal
# ADAPTERS (input) -> output
	output = usecase(input)
	- a forma como nossa aplicação se comunica com o mundo externo
	- o seu input de converter a entrada e se adaptar os dados
	- DTO, data transfer object | camada de controler
	- portas para I/O  do nosso use case
	- só atende quem implementa a interface

#GATEWAY
	- qualquer coisa que tenha acesso ao mundo externo
	DB, API, file....
