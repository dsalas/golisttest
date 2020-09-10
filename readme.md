
# Pasos para crear este paquete

## Creación de modulo
go mod init github.com/dsalas/golisttest

## Inicializar el repositorio y subir a github.
git init 

git add -A

git commit -m "Commit message"

git remote add origin https://github.com/dsalas/golisttest.git

git push -u origin master

## Crear etiqueta para versión
git tag v1.0.0
git push --tags 

## Para importar paquete
import "github.com/dsalas/golisttest"

