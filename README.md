# Comparación de API en Python y Golang

Este repositorio contiene una comparación entre una API desarrollada en Python (py-api-test) y una API desarrollada en Golang (api-test). El objetivo es evaluar y comparar el rendimiento y la respuesta de ambas API utilizando Docker Compose para levantar ambos servicios y un Jupyter Notebook para analizar los resultados.

## Estructura del repositorio

El repositorio se organiza de la siguiente manera:

- `py-api-test/`: Carpeta que contiene el código fuente de la API en Python.
- `api-test/`: Carpeta que contiene el código fuente de la API en Golang.
- `docker-compose.yml`: Archivo de configuración de Docker Compose para levantar los servicios.
- `EDA_requests.ipynb`: Jupyter Notebook para analizar y comparar los resultados.
- `request_apis.ipynb`: Jupyter Notebook para generar request concurrentes.



clonar y levantar el servicio

```bash
git clone https://github.com/tu-usuario/repo.git

cd repo

docker-compose up -d
```

## EDA de resultados

```
/tester_api/EDA_requests.ipynb
```

Como se puede comprobar en el código, la API en FASTAPI, sin el uso de atomics, no logra mantener la consistencia en el conteo, mientras que en el caso de Golang sí lo hace de manera confiable.

Por otro lado, se observa que el tiempo de retraso promedio es considerablemente menor en Golang en comparación con Python. No obstante, se ha encontrado que Golang tarda más en responder en 1246 muestras de un total de 10.000, lo cual representa un 12% de las veces. Sin embargo, en promedio, Golang presenta tiempos de respuesta más rápidos.

### Displot de tiempos de respuestas

Un displot es una herramienta visual que permite observar la distribución de los datos en un conjunto de tiempos de respuesta. Proporciona una representación gráfica de la frecuencia con la que ocurren diferentes valores de tiempo de respuesta.

![Distplot](/tester_api/distplot.png)

### Media del response tiempo por deciles

El motivo de realizar un gráfico que compare la media de respuesta por deciles entre las APIs de Golang y Python es analizar y visualizar de manera más clara cómo se distribuyen los tiempos de respuesta en ambos sistemas.

El título "Media del response tiempo por deciles" indica que se está calculando la media (promedio) de los tiempos de respuesta en diferentes puntos de la distribución, dividiendo los datos en diez partes iguales, conocidas como deciles.

![response time by deciles](/tester_api/mean_response_time.png)

### Describe de las respuestas de golang con delay

El resumen estadístico (describe) del delay en las respuestas de la API de Golang nos proporciona información importante sobre los tiempos de respuesta en un escenario específico. Este resumen nos ayuda a comprender mejor el comportamiento de la API cuando se introduce un retraso intencional en el procesamiento de las solicitudes.

|           | delay_golang |
|-----------|--------------|
| 25%       | 2.149170     |
| 50%       | 5.360718     |
| 75%       | 13.298706    |
| count     | 1246.000000  |
| max       | 587.762939   |
| mean      | 28.635691    |
| min       | 0.000977     |
| std       | 73.729192    |
