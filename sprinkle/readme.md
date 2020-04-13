# Sprinkle
Web向きな語句(sprinkle word)をいくつか生成し、空いているドメイン名が見つかる確率を高める。

## Usage
```bash
# 標準出力
$ ./sprinkle
chat
chat
chat
chatapp
chat
gochat

# ファイル
$ cat ./data/words.txt
java
javascript
c
c##
c++
scala
go
rust
php
python
perl

$ ./sprinkle ./data/words.txt
javasite
javascript
getc
c##app
c++
letsscala
gogo
rust
gophp
pythonsite
perl
```

