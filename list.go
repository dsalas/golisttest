package mlist

import "fmt"

// Unexported struct
type mNodoLista struct {
  sig *mNodoLista
  ant *mNodoLista
  valor  interface{}
}

// Unexported method
func fNodoLista(elem interface{}) *mNodoLista {
  nl := new(mNodoLista)
  nl.sig = nil
  nl.ant = nil
  nl.valor = elem
  return nl
}

// Unexported struct
type mLista struct {
  izquierdo *mNodoLista
  derecho *mNodoLista
  longitud int
}

// Constructra []
func ListaVacia() *mLista {
  l := new(mLista)
  l.derecho = nil
  l.izquierdo = nil
  l.longitud = 0
  return l
}

// Operacion observadora es-lista-vacia?
func (l *mLista) EsListaVacia() bool {
    return l.izquierdo == nil
}

// Operacion observadora longitud
func (l *mLista) Longitud() int {
    return l.longitud
}

// Operacion observadora derecho
func (l *mLista) Derecho() interface{} {
    return l.derecho.valor
}

// Constructra _:_
func (l *mLista) AnhadeIzquierda(elem interface{})  {
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
func ListaUnitaria(elem interface{}) *mLista {
    l := ListaVacia()
    l.AnhadeIzquierda(elem)
    return l
}

func (l *mLista) Traverse() {
  rec := l.izquierdo
  for ( rec != nil) {
    fmt.Printf("%v ", rec.valor)
    rec = rec.sig
  }
  fmt.Print("\n")
}