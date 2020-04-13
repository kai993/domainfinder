# Domainify
受け取った文字列をドメイン名として利用できる正当な物に変換する。
* 不正な文字を削除
* 空白文字をハイフンに変換
* トップレベルドメイン(`.com`, `.net`など)を末尾に追加する

## Usage
```bash
$ ./domainify 
chat
chat.net
chat
chat.com

# top level domain指定
$ ./domainify -tld org
chat
chat.org
```

