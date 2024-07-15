# kotatuneko-rest

Hack-U 2024 福岡 こたつねこ ばっくえんど の Rest API

# 構造

## Enterprise Business Rules

### Entity

proto/user
gen/user

### Repository

internal/domain/repository
internal/domain/service

## Application Business Rules

internal/app/service // usecase

## Interface Adapters

## Frameworks & Drivers

# cmd

```bash:github.com/rayfiyo/kotatuneko-rest
buf lint && buf format -w && buf generate
```
