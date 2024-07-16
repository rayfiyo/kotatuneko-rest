# kotatuneko-rest

Hack-U 2024 福岡 こたつねこ ばっくえんど の Rest API

# 構造

## Enterprise Business Rules

proto/user/
gen/user/
internal/domain/

## Application Business Rules

internal/usecase/service

## Interface Adapters

## Frameworks & Drivers

# cmd

```bash:github.com/rayfiyo/kotatuneko-rest
buf lint && buf format -w && buf generate
```
