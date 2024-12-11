# Gestor tareas

## Requisitos para la elección del gestor de tareas
- Con el principal objetivo de reducir la deuda téncica en esta parte, la herramienta debe de ser una que esté actualizada.
- Queremos continuar con las buenas prácticas que recomienda Go, así que elegiremos una herramienta que nos permita mantenerlas.
- Buscaremos una herramienta con una sintáxis sencilla de aprender y que no requiera de gran experiencia previa, de esta forma no gastaremos mucho tiempo en aprender a utilizarla si no estamos experimentados.
- Intentaremos que su uso genere la menor cantidad de dependencias posibles, de esta forma conseguiremos que la herramienta no sea un problema por ella misma.
- Adecuarnos a la complejidad necesaria para el proyecto. Algunas herramientas pueden ser más sofisticadas pero requirir una configuración complicada desde el principio. Sabiendo que el proyecto que estamos tratando no es de gran tamaño, buscaremos alguna herramienta que pueda ser capaz de funcionar con la mínima configuración.

## Opciones consideradas

- **Make**: Es una herramienta que es compatible con cualquier lenguaje, sin embargo, debido a esto no es un gestor específico de Go, lo que hace que no fomente las mejores prácticas. En un principio, alguien acostumbrado a trabajar en Go no tiene por qué conocer cómo se utiliza, lo que hace que pueda conllevar bastante tiempo su aprendizaje. Al ser una herramienta altamente disponible no genera dependencias significativas.
- **Sage**: Herramienta inspirada en Mage (un gestor de tareas ya existente escrito en Go). Al estar inspirado en Mage ayudará a seguir las buenas prácticas de Go y será conocida por un desarrollador experimentado en Go. Sin embargo, Mage no recibe actualizaciones desde hace más de un año, lo que hace pensar que Sage pueda acabar por el mismo camino.
- **Task**: Es una herramienta que utiliza archivos en formato YAML y está pensada tanto para Go como para otros lenguajes. Requiere un binario de la herramienta, lo que es una dependencia sumado al archivo YAML. A pesar de no estar escrito en lenguaje Go como otras herramientas, la sintáxis YAML es sencilla. 
- **Just**: Es un gestor similar a Make pero con un enfoque en simplicidad y legibilidad. Está escrita en Go y se utiliza creando *Makefiles* que son mucho más similares a scripts. Al ser parecida a Make, su aprendizaje será rápido, pues junta su lógica con una sintáxis mucho más sencilla. Al estar escrita en Go permitirá continuar con las buenas prácticas y no sumar dependencias como otras herramientas. Es en general una herramienta moderna con un soporte activo y actualizaciones frecuentes. Permite elaborar una rápida configuración para proyectos pequeños y medianos como lo es este.

## Elección
Valorando los pros y contras de las opciones anteriormente expuestas, considero que la elección más sensata será el uso de Just. Su sencillez facilitará la automatización de las tareas repetitivas (compilación y ejecuciones) sumado a la facilidad futura de mantenimiento. Enlazando con la gestión de dependencias, al estar escrito en el propio lenguaje Go no necesitaremos tener en cuenta dependencias extras.
