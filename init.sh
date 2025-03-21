#!/bin/bash

REPO_NAME=$(basename $(git rev-parse --show-toplevel) 2>/dev/null || basename $(pwd))

echo "🔧 Inicializando Go Module con el nombre: github.com/usuario/$REPO_NAME"

go mod init "github.com/skolldire/$REPO_NAME"
go mod tidy

echo "Go Module inicializado correctamente como: github.com/usuario/$REPO_NAME"