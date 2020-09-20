
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

## Uso del paquete

Crear modulo de prueba: go mod init testmodule

Crear main.go con el ejemplo:

```
package main

import (
	"fmt"
	"github.com/dsalas/golisttest"
)

func main() {
    l1 := mlist.ListaVacia()
    fmt.Printf("L1 long: %v\n",l1.Longitud())
    l2 := mlist.ListaUnitaria(1)
    fmt.Printf("L2 long: %v value %v\n",l2.Longitud(), l2.Derecho())
    l1.AnhadeIzquierda(2)
    fmt.Printf("L1 long: %v value %v\n",l1.Longitud(), l1.Derecho())
    l1.AnhadeIzquierda(3)
    l1.Traverse()
}
```
Compilar: go build

Ejecutar: ./testmodule

Salida esperada:
```
L1 long: 0
L2 long: 1 value 1
L1 long: 1 value 2
3 2 
```

# Referencias
https://blog.francium.tech/go-modules-go-project-set-up-without-gopath-1ae601a4e868XS