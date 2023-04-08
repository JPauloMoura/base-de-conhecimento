# ===== GIT GITHUB =====
> CREDENCIAIS
  $ git config  --global user.name
  $ git config  --global user.email
  $ git config --list

# LOGS
> mostra detalhes dos ultimos commits
  $ git log
> mostra detalhes dos ultimos commits de forma resumida 
  $ git log --oneline
> mostra detalhes dos ultimos commits em um arquivo especifico
  $ git log file.txt
> mostra detalhes de modificações dos ultimos commits em um arquivo especifico
  $ git log -p file.txt
> mostra detalhes dos ultimos commits de forma grafica
  $ git log --graph


#  DESFAZENDO COISAS
  > removendo de modified para estado original
    $ git restore --staged site.txt //opcao 2
    $ git checkout file.txt //opcao 2
  > removendo da staging-area para modified
    $ git add file.txt
    $ git reset HEAD site.txt
  > restaurando arquivo com status de Deletado
    $ git restore file.txt
  > removendo arquivo e comitando ele
    - apague o arquivo no projeto
    $ git rm file.txt //também serve para deletar o arquivo
    $ git commit -m "deletando file"
  > removendo commits
    ex:------------------------------
	6be7924 (HEAD -> branch-1) criando site
	92dc3b5 deletando texto
	051bf23 criando texto
	2e2439a deletando readme
	68da1b0 primeiro
       ------------------------------
     $ git reset --hard 92dc3b5
     $ git log --oneline
       ------------------------------
	92dc3b5 (HEAD -> branch-1) deletando texto
	051bf23 criando texto
	2e2439a deletando readme
	68da1b0 primeiro
       ------------------------------
   > retornando um commit para a stage area (apos git add)
     $ git reset --soft cod_commit_anterior
   > retornando um commit para modified (antes do git add)
     $ git reset --mixed cod_commit_anterior
   > retorna totalmente para o commit anterior
     $ git reset --hard cod_commit_anterior
   > reverte o commite atual criando um novo
     $ git revert cod_commit_atual 

# BRANCHS
 > listar as branch
   $ git branch
 > criar branch
   $ git branch minha-branch
 > deletar branch
   $ git branch -D minha-branch
 > renomear branch local
   $ git branch -m atual nova
 > realizando merge de uma branch atual
   (main)$ git merge nome-branch
 > deletando branch local
   $ git branch -D nome-branch

# Modificando Seu Último Commit
  imagine que  vc esqueceude add a img de logo no commit
  dessa forma você pode adiciona o arquivo esquecido ao commit
 (o commit não pode estar no github)
  $ git commit -m 'initial-Page'
  $ git add logo.png
  $ git commit --amend --no-edit

