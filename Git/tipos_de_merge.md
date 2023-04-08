## Tipos de Merge
### Squash:
Usado quando queremos mergear uma branch em uma release candidate, ou uma alteração direto na master
```js
	branch_1 <--  branch_fix
	master <-- branch_fix
	historico de commits:
		[ master: c1 - c3 ] <-- [ fix:	c2 - c4 - c5  ]
		master: c1 - c3 - ( c2 + c4 + c5)
		master: c1 - c3 - squash4
```
### Rebase:
Usado quando queremos mergear uma  release candidate na master
```js
	master <-- branch_1
	historico de commits: é mergiado no ponto onde a branch foi criada
		[ master: c1 - c3-c6 ] <-- [ rc:	c2 - c4 - c5  ]
		master: c1 - c3 - C2 - C4 - C5- c6
```

### Merge:
Historico de commits: são mergiados de forma cronologica na linha do tempo
```js
	[ master: c1 - c3 ] <-- [ fix:	f2 - f4 - f5  ]
	master: c1 - f2 - c3 - f4 - f5 -merge6 -f
```

### Cherry pick
Pega um commit especifico de outra branch e trás pra branch atual
```js
	[ master: c1 - c3 ]
	[ branch2:	f2 - f4 - f5  ]

	(master)$ git cherry-pick <hash-commit-f4>
	
	[ master: c1 - c3 - f4 ]
```