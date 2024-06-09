# About

This repository contains useful "godoc" example, some of which I used while writing my articles.

More information about this application can be found in this article: [Como gerar documentos do seu código em Go?](https://willsena.dev/como-gerar-documentos-do-seu-codigo-em-go/).


# Installing dependencies

```shell
go install -v golang.org/x/tools/cmd/godoc@latest
go install golang.org/x/pkgsite/cmd/pkgsite@latest
```

# Running

```shell
(cd doc-app-example && make dev)
```

## Requests
```shell
curl http://localhost:4000/cards 

# {"holder_name":"Ephraim Hand","type":"American Express" "number":"6062825121549507","cvv":"857","exp":"08/25"}
```

# Docs
## pkgsite

```shell
make docs-pkgsite
```

## godoc
```shell
make docs-godoc
```