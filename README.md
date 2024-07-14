# kotatuneko-rest

Hack-U 2024 福岡 こたつねこ ばっくえんど の Rest API

# cmd

```bash:github.com/rayfiyo/kotatuneko-rest
protoc \
	--go_out=gen/user/ --proto_path=proto/user/ \
	--go_opt=paths=source_relative \
	proto/user/user.proto proto/user/resources/user.proto proto/user/rpc/user.proto
```
