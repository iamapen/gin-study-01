

# golangのインストール (WSL)

## インストール
`$HOME/go` を GOPATH とする

```bash
mkdir ~/go
```

`/usr/local/go` を GOROOT とする。

```bash
cd ~
mkdir -p ~/local/src && cd ~/local/src
wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
```

```bash
/usr/local/go/bin/go version
```

以降はIntelliJのターミナルでwslを使って操作するのが楽。
環境変数の類を自動で設定してくれるから。
