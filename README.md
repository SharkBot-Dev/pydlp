# pydlp
yt-dlpの簡単なバージョン切り替えツール

# ヘルプ
`pydlp --help` ヘルプを表示します。

# バージョンをインストールする方法
`pydlp install`で最新バージョンをインストールできます。<br>
`pydlp install <version>`でバージョンを指定してインストールできます。

# インストールしたバージョンを切り替える方法
`pydlp versions`でインストールしたバージョン一覧を確認し、<br>
`pydlp version <version>`でバージョンを切り替えられます。

# インストールしたyt-dlpを実行する方法
`pydlp run <yt-dlpのオプション>`で実行できます。<br>

# ビルドする
`go build`で実行した端末向けにビルドされます。<br>
WindowsからLinux向けにビルドする際は、
```bat
set GOOS=linux 
set GOARCH=amd64 
go build -o pydlp```でビルドできます。