# api.somesim
API for MoE-Somesim

## API開発者用

### はじめに
まず最初に`make init`してディレクトリを作成します。

次に`make container/develop`でGoのバイナリや依存ライブラリ、
`go-swagger`の入ったイメージ(`pocka/api.somesim-dev`)をビルドして下さい。
(もしかしたら`docker pull pocka/api.somesim-dev`でもできるかも)

### コンパイル
`make`(デフォルト)で`pocka/api.somesim`というイメージが作成されます。

ローカル開発環境でのテストは`make run/server`で
`80`ポートをリッスンするサーバが立ち上がります。


実際にデプロイした場合は以下の様な感じで起動して下さい。

```sh
docker run -p <リッスンするポート>:80 pocka/api.somesim:v1
```
