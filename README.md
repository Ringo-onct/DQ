# DQProject

> バグとかあったら教えてほしい。

## 注意事項
かなり前に作成して飽きたやつだから
コードが悲惨です、あんまりジロジロみないで////

## プレイ方法

1. DQ/test/DQ1.exeとDQ/Documentフォルダをダウンロードする  
2. DQ1.exeを実行する  

ディレクトリ構成は、このようにしないと動かないので注意
```
├ DQ1.exe
└ Document
  ├ lvup_status
  ├ monster_list
  └ player_list
```
## v3.0.0実装

- 敵のランダム出現を実装
- データセーブ機能の実装
- 行動選択を矢印キーで操作できるようにした
- ゲーム開始前にエンターキー入力待機をするようにした
- 冒険の書作成・削除機能を実装
- ダメージ計算式を本家と同じにした
- 文字入力を、GUI的にできるようにした
- 経験値システム・レベル処理を実装
- 魔法を実装
- 体力に応じたウィンドウ枠・文字色変更を実装
