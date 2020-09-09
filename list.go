package mlist

import "fmt"

type NodoLista struct {
  sig *NodoLista
  ant *NodoLista
  valor  interface{}
}

func fNodoLista(elem interface{}) *NodoLista {
  nl := new(NodoLista)
  nl.sig = nil
  nl.ant = nil
  nl.valor = elem
  return nl
}

type Lista struct {
  izquierdo *NodoLista
  derecho *NodoLista
  longitud int
}

// Constructra []
func ListaVacia() *Lista {
  l := new(Lista)
  l.derecho = nil
  l.izquierdo = nil
  l.longitud = 0
  return l
}

// Operacion observadora es-lista-vacia?
func (l *Lista) EsListaVacia() bool {
    return l.izquierdo == nil
}

// Operacion observadora longitud
func (l *Lista) Longitud() int {
    return l.longitud
}


// Operacion observadora derecho
func (l *Lista) Derecho() interface{} {
    return l.derecho.valor
}


// Constructra _:_
func (l *Lista) AnhadeIzquierda(elem interface{})  {
  nl := fNodoLista(elem)
  if (l.EsListaVacia()) {
    l.derecho = nl
  } else {
    nl.sig = l.izquierdo
    l.izquierdo.ant = nl
  }
  l.izquierdo = nl
  l.longitud += 1
}

// Operacion lista unitaria [e]
func ListaUnitaria(elem interface{}) *Lista {
    l := ListaVacia()
    l.AnhadeIzquierda(elem)
    return l
}

// Anhade derecha

// Operacion izquierdo

// Operacion derecho

// Operacion elimina izquierdo

// Operacion elimina derecho

// Operacion concatenar

// Operacion recorrer

func (l *Lista) Traverse() {
  rec := l.izquierdo
  for ( rec != nil) {
    fmt.Printf("%v ", rec.valor)
    rec = rec.sig
  }
  fmt.Print("\n")
}