# parallel-poke-request
## 概要 - Overview  
Goの並列処理を用いてPokemonAPIにリクエスト  
PokeAPIにリクエストを送信し、151匹のポケモンデータを取得する

PokeAPI: https://pokeapi.co/

## 実行方法 - How to run  
pフラグを付けて並行処理で実行
```
Let's get started to capture all pokemons!
-------------------
id: 56, name: mankey
id: 129, name: magikarp
id: 117, name: seadra
:
:
id: 102, name: exeggcute
-------------------
It took 334.734625ms.
```

sフラグを付けて逐次処理で実行
```
Let's get started to capture all pokemons!
-------------------
id: 1, name: bulbasaur
id: 2, name: ivysaur
id: 3, name: venusaur
:
:
id: 151, name: mew
-------------------
It took 2.865047209s.
```
