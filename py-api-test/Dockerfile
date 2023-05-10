# Imagen base de Python 3.9
FROM python:3.9

# Copiar los archivos necesarios al contenedor
COPY requirements.txt /app/requirements.txt
COPY main.py /app/main.py

# Establecer el directorio de trabajo
WORKDIR /app

# Instalar las dependencias
RUN pip install --no-cache-dir -r requirements.txt

# Exponer el puerto en el que se ejecuta la API
EXPOSE 8080

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8081"]

