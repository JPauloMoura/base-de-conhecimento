# pods
  - pode executar varios containers dentro dele
  - o indicado é apenas um, para que assim possamos escalar os nossos pods
    serviços individualmente.

# replicaset
  - permite que possamos escalar os nossos pods informando sua quantidade
  - é caso algum pod morra, ele garante que sempre vai ter uma quantidade X em execução
# deployment
  - permite o controle de versão dos nossos pods.
  - garante que sempre que for realizada uma alteração na especificação, os pods
    antigos vão morrer e novos atualizados com a alteração vão nascer