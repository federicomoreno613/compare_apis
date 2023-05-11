# Comparación de API en Python y Golang

Este repositorio contiene una comparación entre una API desarrollada en Python (py-api-test) y una API desarrollada en Golang (api-test). El objetivo es evaluar y comparar el rendimiento y la respuesta de ambas API utilizando Docker Compose para levantar ambos servicios y un Jupyter Notebook para analizar los resultados.

## Estructura del repositorio

El repositorio se organiza de la siguiente manera:

- `py-api-test/`: Carpeta que contiene el código fuente de la API en Python.
- `api-test/`: Carpeta que contiene el código fuente de la API en Golang.
- `docker-compose.yml`: Archivo de configuración de Docker Compose para levantar los servicios.
- `EDA_requests.ipynb`: Jupyter Notebook para analizar y comparar los resultados.



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

Cómo esta demostrado en el código la api en FASTAPI sin atomics no mantiene consistencia en el conteo en cambio la de golang si.
Por otro lado el tiempo de delay es en promedio decil a decile menor en Golang que en python. Sin embargo encontramos que golang tarda más en responder en 1246 muestras sobre 10.000 esto es un 12%
de las veces pero en promedio tarda menos.

### Displot de tiempos de respuestas
![Distplot](/tester_api/distplot.png)

### Media del response tiempo por deciles
![response time by deciles](/tester_api/mean_response_time.png)

### Describe de las respuestas de golang con delay

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
