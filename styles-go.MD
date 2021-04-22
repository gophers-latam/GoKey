# Gopher-latam. Guia de estilos en Go.

## TOC


## Intro

En esta guía queremos mostrar algunos conceptos básicos que sean los que conduzcan a una calidad de código excelente.
Ademas de respetar las normas de Go ya tiene incorporadas en sus herramientas como fmt o imports.

El objetivo es reducir discusiones técnicas basandose en esta guia como también la complejidad del código.


- Correr `goimports`.
- Correr `golint` y `go vet` para el chequeo de errores.


## Guia

### Verificar interfaces

Verificar las interfaces en tiempo de compilación.

<table>
<thead><tr><th>Mal</th><th>Bien</th></tr></thead>
<tbody>
<tr><td>

```go
type Handler struct {
  // ...
}



func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  ...
}
```

</td><td>

```go
type Handler struct {
  // ...
}

var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}
```

</td></tr>
</tbody></table>

`var _ http.Handler = (*Handler)(nil)` debería fallar al compilar porque
`*Handler` no cumple como una interfaz de `http.Handler`.

El lado derecho de la asignacion debe ser el zero value de ese tipo. En este caso `nil` para los punteros (like `*Handler`), slices, mapas
y struct vacias para structs.

```go
type LogHandler struct {
  h   http.Handler
  log *zap.Logger
}

var _ http.Handler = LogHandler{}

func (h LogHandler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}
```
