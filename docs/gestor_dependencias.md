## Gestor de dependencias

En Go, tras la introducción del concepto de módulos, no existe un gestor de dependencias externo.
La gestión de las dependencias se realiza directamente mediante los archivos `go.mod` y `go.sum`.
En estos archivos se definen el path del módulo, se especifican dependencias necesarias (en el primero de ellos) y se gestionan la integridad de las dependencias y las versiones (en el segundo de ellos).

De esta forma, Go Module nos permite crear colecciones de paquetes que se versionan conjuntamente
