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

Nombre del archivo fuente: [hola.go](https://github.com/pamelaAHeredia/go/blob/main/seminario-de-lenguajes/p1/soluciones/hola.go)
Compile y ejecute todo en un paso, y compile generando yn ejecutable con nmbre hola.exe o directamente hola. 

- *go run hola.go* > imprime "Hello world"
- *go build-o hola hola.go* > genera un binario llamado [hola](https://github.com/pamelaAHeredia/go/blob/main/seminario-de-lenguajes/p1/soluciones/hola)
- *./hola* > ejecuta el binario > imprime Hello World

## Ejercicio 3

- Dada la siguiente declaración de programa  GO, indicar si es correcta. Usar compilador. 

```
func main () {
    /* integers */
    var zz int = 0A
    var z int := x; 
    x := 10; 
    var y int8 = x+1; 
    const n := 5001
    /* float //
    var e float32 := 6
    f float32 = e
}
```

- No es correcta. La correción está en el archivo [e3.go](https://github.com/pamelaAHeredia/go/blob/main/seminario-de-lenguajes/p1/soluciones/e3.go)

l2) `var zz int = 0A` -> La variable zz fue declarada como int, que representa enteros, al ser un lenguaje fuertemente tipado, una variable declarada como int no puede tener strings. 

l3-l4) 
```
    var z int := x; 
    x := 10;
```
go es un lenguaje compilado, no debe tener errores semánticos ni sintácticos para poder traducirlo a lenguaje máquina y generar el ejecutable. La primera línea está asignando a la variable z el valor de x, pero x no tiene ningún valor asignado (undefined: x). 
En la línea 2, a x se le asigna el valor 10 y go infiere que x es una variable de tipo entero, lo cual es correcto. 
Para que este bloque tenga sentido, debemos mover la línea 2 arriba de la línea 1, de esta forma a z se le podrá asignar correctamente el valor de x. 

l5) `var y int8 = x+1` -> no existe int 8. Go tiene int, uint, y uintptr, suelen tener 32 bits de ancho en sistemas de 32-bits y 64 bits en sistemas de 64-bits.

l6) `const n := 5001` -> Las constantes no se pueden declarar con := 

l8) `var e float32 := 6` -> Si una variable es declarada con var, no se debe usar := 

l9) `f float32 = e` -> Puede ser *f := e* o var *f float32 = e* 

## [Ejercicio 4](https://github.com/pamelaAHeredia/go/blob/main/seminario-de-lenguajes/p1/soluciones/e4.go)

- Escriba en la salida estándar la suma de los primeros números positivos pares menores o iguales a 250. Cambiar el programa para que itere en el sentido contrario pero pobtener el mismo resultado. Cambiar el programa para que, en lugaa de usar un literal como tope, se use una constrante. Si lo desea, investigue la herramienta gofmt y pruebe sobre el ćodigo. 

## [Ejercicio 5](https://github.com/pamelaAHeredia/go/blob/main/seminario-de-lenguajes/p1/soluciones/e5.go)

- Realizar un programa que lea un número y muestre el valor correspondiente aplicando la siguiente función: 

si  x pertenece al conjunto (-inf, -18), entonces devolver x
si  x pertenece al conjunto [-18, -1], entonces devolver x mod 4
si  x pertenece al conjunto [1, 20), entonces devolver x al cuadrado
si  x pertenece al conmunto [20, inf), entonces devolver -x

## Ejercicio 6

- 