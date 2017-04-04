# api.somesim
API for MoE-Somesim

## API開発者用

### はじめに
まず最初に`make init`してディレクトリと、ローカルでHTTPS通信するためのオレオレ証明書を作成します。

オレオレ証明書の作成中にパスフレーズの入力が計4回、各種情報の入力が何回かあります。
パスフレーズは適当に入力して、各種情報はデフォルト値のままEnter押して進めて大丈夫です。

もしすでにオレオレ証明書がある場合、`.crt`ディレクトリの中に
`testserver.crt`と`testserver.key`
という名前でシンボリックリンクを張って下さい。

次に`make container/develop`でGoのバイナリや依存ライブラリ、
`go-swagger`の入ったイメージ(`pocka/api.somesim-dev`)をビルドして下さい。
(もしかしたら`docker pull pocka/api.somesim-dev`でもできるかも)

### コンパイル
`make`(デフォルト)で`pocka/api.somesim`というイメージが作成されます。

ローカル開発環境でのテストは`make run/server`で
`8443`ポートをリッスンするサーバが立ち上がります。


実際にデプロイした場合は以下の様な感じで起動して下さい。

```sh
docker run \
	--rm \
	-p <リッスンするポート>:443 \
	-v <証明書のパス>:/work/.crt/tls.crt \
	-v <秘密鍵のパス>:/work/.crt/tls.key \
	pocka/api.somesim:v1 \
		--tls-certificate=/work/.crt/tls.crt \
		--tls-key=/work/.crt/tls.key
```
