# SSH COM DUAS CONTAS GIT 
> gerar chave
  $ ssh-keygen -t rsa -C "email_1@gmail.com"
> quando solicitado, salve o arquivo como "~/.ssh/id_rsa_email_1"

> Cria uma nova chave ssh no github colocando a string criada anteriomente
  $ vim ~/.ssh/id_rsa_email_1.pub

> Como salvamos nossa chave com um nome exclusivo, precisamos informar ao SSH sobre isso: o resultado deve ser "Identity Added."
  $ ssh-add ~/.ssh/id_rsa_email_1 
  
 > Agora precisamos especificar quando desejamos enviar para nosso email_1 e quando devemos enviar para o email_2
  $ touch ~/.ssh/config
  $ vim config
 
 > email_1
	#Default GitHub
	Host github.com-email-1
	  HostName github.com
	  User git
	  IdentityFile ~/.ssh/id_rsa_email_1

> email_2
	Host github.com
	  HostName github.com
	  User git
	  IdentityFile ~/.ssh/id_rsa_email_2
