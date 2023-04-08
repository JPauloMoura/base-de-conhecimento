## Fiz um PR e subiu alterações que eu não fiz e agora?
  1. pegue a Hash do commit anterior aos seus
    > $ git log

  2. Volte para o momento anterior ao seus commits
    >	$ git reset --sof [hash]

  3. Vá para branch que seria alvo do PR
    >	$ git checkout [rc/]

  4. Delete sua branch anterior e crie ela novamente
    >	$ git branch -D [branch]
    >	$ git checkout -b [branch]

  5. faça um push forçado para sua branch remota
    >	$ git push origin [branch] -f
