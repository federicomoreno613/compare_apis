@Author: Federico Moreno

# API

## Requisitos
* Docker
* Make

## Uso
Para ejecutar la API, se debe ejecutar el siguiente comando en la raíz del proyecto:

``` bash
make run
```
Este comando ejecutará la API en un contenedor de Docker y expondrá el puerto 8080. La API estará lista para recibir solicitudes en http://localhost:8080.

## Endpoints
La API expone los siguientes endpoints:

* ```POST /date?full_format=BOOLEAN```: Incrementa el contador y devuelve la fecha actual en formato. Tiene como parámetro obligatorio full_format
  * Ejemplo: ```POST /date?full_format=true``` Devuelve la fecha y hora actual en formato ISO 8601 **incluyendo** hora, minutos y segundos
  * Ejemplo: ```POST /date?full_format=false``` Devuelve la fecha actual en formato ISO 8601, **sin incluir** hora, minutos ni segundos.
* ```GET /counter```: incrementa el contador y devuelve su valor actual.


### Notas de implementación
#### Atomics vs locks

```go
// Definimos el contador como un entero sin signo de 64 bits
var counter uint64 = 0
...

// Incrementamos el contador atómicamente
atomic.AddUint64(&counter, 1)
```
En lugar de simplemente incrementar el valor de un contador de forma convencional (por ejemplo, con counter++), se utiliza esta función que incrementa el valor de la variable counter en 1 de manera atómica. Es decir, ningún otro hilo puede acceder a la variable counter mientras se está realizando la operación de incremento.

En contraposición, si usamos una operación de incremento no atómica en un recurso compartido, corremos el riesgo de obtener resultados inesperados, debido a que un hilo puede interrumpir a otro y el resultado final podría no ser el esperado.

Referencias:
* https://golangdocs.com/atomic-operations-in-golang-atomic-package
* https://pkg.go.dev/sync/atomic


#### Ejemplo de concurrencia
```go
var counter uint64 = 0

func increment() {
counter = counter + 1
}

func main() {
// Lanzamos dos goroutines para incrementar el contador.
go increment()
go increment()

    // Esperamos que ambas goroutines finalicen.
    time.Sleep(1 * time.Second)

    fmt.Println(counter) // El resultado puede ser inesperado
}
```

Este código incrementa el contador simple counter en dos goroutines diferentes, lo que puede generar una condición de carrera o race condition, donde múltiples goroutines acceden y modifican el mismo valor de manera concurrente, generando un resultado inesperado.

Para solucionar este problema, podemos utilizar atomic.AddUint64(&counter, 1) para asegurarnos de que el incremento del contador sea atómico y libre de race conditions.

La diferencia entre atomic.AddUint64(&counter, 1) y un contador simple es que el primero asegura que la operación de incremento del contador sea ejecutada de manera atómica, lo que significa que la operación es ejecutada en una sola instrucción de la CPU, sin interferencia de otras operaciones concurrentes.

En cambio, si utilizamos un contador simple sin protección para accederlo concurrentemente, múltiples goroutines pueden intentar acceder y modificar el valor del contador al mismo tiempo, lo que puede resultar en inconsistencias y errores en el valor final del contador.
