# Comparación de API en Python y Golang

Este repositorio contiene una comparación entre una API desarrollada en Python (py-api-test) y una API desarrollada en Golang (api-test). El objetivo es evaluar y comparar el rendimiento y la respuesta de ambas API utilizando Docker Compose para levantar ambos servicios y un Jupyter Notebook para analizar los resultados.

## Estructura del repositorio

El repositorio se organiza de la siguiente manera:

- `py-api-test/`: Carpeta que contiene el código fuente de la API en Python.
- `api-test/`: Carpeta que contiene el código fuente de la API en Golang.
- `docker-compose.yml`: Archivo de configuración de Docker Compose para levantar los servicios.
- `results.ipynb`: Jupyter Notebook para analizar y comparar los resultados.



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
