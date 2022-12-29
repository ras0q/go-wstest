# go-wstest

## Usage

### Build

```sh
make
```

`bin/`以下に以下のコマンドを実行するためのバイナリを生成します。

- `app` : Websocketサーバーを起動するコマンド
- `wscat` : Websocketサーバーと対話形式で通信を行うコマンド
  - `app`を実行後、別のターミナルで実行する必要があります。

### Test

```sh
go test -v ./...
```

テスト内容は以下を参照してください。

- [ws/ws_test.go](./ws/ws_test.go)
