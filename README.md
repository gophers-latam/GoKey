# GoKey 🔑

## Concepto:

<pre>
Sistema de base de datos clave valor, distribuido. En forma de cache en memoria.
</pre>

## Especificaciones:

Para concentrar información documentada y contar con canal de comunicación para indicaciones y especificaciones sobre el desarrollo del proyecto se encuentra el [**discord**](https://discord.com/invite/AEarh2kSvn) de la comunidad con el canal de texto **#🔑gokey** en su descripción se fijó el acceso al [**Google Doc**](https://docs.google.com/document/d/1qwuyUnS0YtBLRmz4qXaywxWSTIyc023iMe9aPQXHHZk/edit) para documentar sobre realización del proyecto.

De ser necesario para reuniones puede hacerse uso de los canales **#👥sala** y **#stage**

<br>

### 🔗 [Tablero tareas](https://github.com/gophers-latam/GoKey/projects/1)

### 🔗 [Guia estilos de código](./docs/styles-go.md)

### 🔗 [Convenciones commits](./docs/conv-commits.md)

### 🔗 [Código de conducta](./docs/CODE_OF_CONDUCT.md)

<br>


## Avance etapas

- [] V1. ```Core funcionalidad base``` (WIP)
* Agregar, leer, eliminar elementos.
* Darle un tamaño determinado de entradas.
* Opciones de configuración a la hora de crear la instancia del cache.
* Funcionalidades de Set (agregar/eliminar uno/muchos valor(es), leer todos).
* Cache eviction LIFO para que no exceda la capacidad máxima.
* Estadisticas.
<br><br>
- [] V2. ``` Cache Distribuido ```
* Opcion de guardado en disco.
* Poder correr independientemente, no solo embebido.
* Conexión TCP, GRPC y HTTP.
* Shards/nodos.
### [Licencia MIT](./LICENSE)
