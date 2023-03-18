# Commits Convencionales

Adoptar estas convenciones facilita la lectura de la historia del proyecto. 

No es obligatorio, pero para mantener buenas prácticas se prioriza el uso del inglés al escribir el commit.

## Formato para estas convenciones
```
<Tipo>(Scope Opcional): Descripción

[Cuerpo (Opcional)]

[Pie (Opcional)]

```

### Encabezado

No debe exceder los 100 caracteres. Esto lo vuelve más fácil de leer en Github. Usar el imperativo y presente simple impersonal: "change" no "changes" o "changed", "arregla" en vez de "arreglado". Evitar las mayúsculas y el punto final.



Los tipos esenciales son:


|Tag     |Descripción|
|--------|---|
|``feat``    |Nueva característica agregada, excluyendo los archivos de build (Makefile, go.mod, etc.).|
|``test``    |Agregar o refactorizar tests, no cambia el código de producción.|
|`fix`     |Reparación de bugs.|
|`refactor`|Refactorización de alguna parte del código, como renombrar variables o corregir typos.|
|`docs`    |Cambios a la documentación, ya sea dentro del código o aparte.|
|`style`   |Cambios que no afectan el funcionamiento del código (linting en general).|
|`build`   |Cambios que afectan al sistema de compilación, e.g. Makefile.|

Si es un cambio importante se puede usar (!) para señalarlo (experimental, debe ser consultado).

El **_scope_** apunta a la parte del proyecto modificada, eg. API, web-server, config, etc. Revisar el CHANGELOG para conocer los scopes usados. Puede estar vacío si es un cambio global, o es difícil asignarlo a un componente (en ese caso se omiten los paréntesis).

### Cuerpo

Usar el imperativo y presente simple. Si se incluye, debe presentar el motivo de ese cambio y los cambios efectuados, siempre lo mas breve posible.

### Pie

Sólo se usa si el commit concluye algún issue, es necesario referenciarlo:

```
Close #3142
```







## Ejemplos

```
fix: correct minor typos in code

see the issue for details.

Reviewed-by: Z
Refs #133

```
```
refactor!: change API entirely, incompatible with last version

Refactor to use the new Go features not available on version 1.20
```
```
docs: correct spelling of CHANGELOG
```
```
test: hash MD5
``` 

Estos cambios mejoran la interacción de los usuarios con el proyecto, buscando facilitar la comprensión de su historia y la manera en que la comunidad trabaja. Mantenga en mente que:

> “Any fool can write code that a computer can understand. Good programmers write code that humans can understand.” - Martin Fowler


## Bibliografía

[https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#commit](https://github.com/angular/angular/blob/22b96b9/CONTRIBUTING.md#commit)

[https://karma-runner.github.io/0.10/dev/git-commit-msg.html](https://karma-runner.github.io/0.10/dev/git-commit-msg.html)

[https://www.conventionalcommits.org/en/v1.0.0/](https://www.conventionalcommits.org/en/v1.0.0/)
