# DQProjectについて書いてみた。

> markdown記法苦手だから逃げてtextファイルで描いてたけど、やってみないといけないなって、思った☆

## やりたいこと[^1]

[^1]:これに関しては自分が適当に追加しただけだから、これにこだわる必要はない。勝手に追加しやがれ
 - [ ] ANSIエスケープシーケンスを使って綺麗な画面を作る。[ANSIについての詳細な説明](#ansiの書式について)
 - [ ] モンスターを複数作って、ランダム出現化
 - [ ] キャラクター複数作成
 - [ ] キャラクター別のステータス割り振り
 - [ ] プレイヤーの行動を増やす
 - [ ] モンスターの攻撃方法を増やす
 - [x] 関数を分ける

## 関数について

***分けるコード***[^2]
[^2]:これに関してもやっていくうちにどんどん増えていくと思うし、勝手に決めただけだから勝手に追加しやがれ
- `main`当たり前だな。
- `file`ファイル読み込み系列はこれにまとめる予定
- `action`モンスターとかプレイヤーの攻撃とかのアクションを処理する。`m_atk`/`m_1`みたいな感じで引数渡して処理すれば結構いい感じになりそう
- `console`コンソール画面に関する処理を行う。
- `prompt`これは前作ったやつでいう`menu`と同じ。「促す、要請する」という意味


## 作業について[^3]
[^3]:これに関しても、勝手に決m...以下同文

***俺が勝手に作ろうと思っているところ***

- [x] 関数を分ける
- [ ] プレイヤーの攻撃方法を増やす



## ANSIの書式について


#### 便利なANSIエスケープシーケンス！

<details><summary>多いから折りたたむわ。クリックして全部表示してや。</summary>

- カーソルの移動
```Go
fmt.Print("\033[<行>;<列>H")
```
- 画面クリア
```Go
fmt.Print("\033[2J")
```

- 行クリア
```Go
fmt.Print("\033[K")
```
- 文字色変更
```Go
fmt.Print("\033[<色番号>m")
```
- フォントの太さ変更
```Go
fmt.Print("\033[<番号>m")
```
- 背景色変更
```Go
fmt.Print("\033[<背景色番号>m")
```
- 画面バッファの保存(これ便利やぞ！！)
```Go
fmt.Print("\033[s")
```
- 画面バッファの復元(これもセットやけどな！)
```Go
fmt.Print("\033[v")
```
</details>

[ANSIエスケープシーケンスのI詳細](https://www.mm2d.net/main/prog/c/console-02.html)
