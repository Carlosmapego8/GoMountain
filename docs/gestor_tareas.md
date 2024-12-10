# Gestor tareas

## Requisitos para la elección del gestor de tareas
- Con el principal objetivo de reducir la deuda téncica en esta parte, la herramienta debe de ser una que esté actualizada.
- Queremos continuar con las buenas prácticas que recomienda Go, así que elegiremos una herramienta que nos permita mantenerlas.

## Opciones consideradas

- **Make**: Es una herramienta que es compatible con cualquier lenguaje, sin embargo, debido a esto no es un gestor específico de Go. En un principio, alguien acostumbrado a trabajar en Go no tiene por qué conocer cómo se utiliza, lo que hace que pueda conllevar bastante tiempo su aprendizaje.
- **Sage**: Herramienta inspirada en Mage (un gestor de tareas ya existente escrito en Go)
- **Task**: Es una herramienta que utiliza archivos en formato YAML y está pensada tanto para Go como para otros lenguajes.
- **Just**: Es un gestor similar a Make pero con un enfoque en simplicidad y legibilidad. Está escrita en Go y se utiliza creando *Makefiles* que son mucho más similares a scripts.

## Elección
Valorando los pros y contras de las opciones anteriormente expuestas, considero que la elección más sensata será el uso de Just. Su sencillez facilitará la automatización de las tareas repetitivas (compilación y ejecuciones) sumado a la facilidad futura de mantenimiento. Enlazando con la gestión de dependencias, al estar escrito en el propio lenguaje Go no necesitaremos tener en cuenta dependencias extras.

## Tareas a automatizar:
- Compilación del proyecto y generación del binario en una carpeta `/bin`
- Instalar las dependencias necesarias
- Verificar el correcto formateado del código fuente 
