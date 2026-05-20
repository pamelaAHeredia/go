# Practica 1 - Seminario de lenguajes - opción Go

**Objetivo**: Aprender la sintaxis y semántica de la estructura de un programa go. Conocer los tipos escalares básicos predefinidos por el lenguaje con operaciones y declaraciones. Aprender a usar las sentencias de control básicas del lenguaje: for, breake, continue, if , case. Aprender a usarel modelo de E/S de Go. Introducción de strings. 

## Ejercicio 1
- Indique las diferencias entre un lenguaje interpretado y uno compilado. Indique ejemplos de cada uno, si conoce. ¿en cuál categoría catalogaría a Go? 

### respuesta: 
- Un lenguaje compilado se traduce a lenguaje máquina antes de ser ejecutado, se "compila", produciendo un archivo binario que se puede ejecutar todas las veces que sea necesario. Un ejemplo de este tipo de lenguajes es Java. 
- Un lenguaje interpretado se va traduciendo en tiempo real a medida que se ejecutan las sentencias. Un ejemplo de este tipo de lenguajes es python.
- Una de las dirernecias más notables entre un lenguaje compilado y un lenguaje intepretado es la velocidad de ejecución, siendo más rápido el compilado. 

- Go es un lenguaje compilado. 

## Ejercicio 2
- Compile en el compilador Go el famoso Hello World. 

Nombre del archivo fuente: [hola.go](https://github.com/pamelaAHeredia/go/blob/main/seminario-de-lenguajes/p1/hola.go)
Compile y ejecute todo en un paso, y compile generando yn ejecutable con nmbre hola.exe o directamente hola. 

- *go run hola.go* > imprime "Hello world"
- *go build-o hola hola.go* > genera un binario llamado hola
- *./hola* > ejecuta el binario > imprime Hello World

