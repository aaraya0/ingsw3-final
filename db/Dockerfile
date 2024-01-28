# Dockerfile para MySQL
FROM mysql

# Variables de entorno para configurar la base de datos
ENV MYSQL_ROOT_PASSWORD=aaraya0
ENV MYSQL_DATABASE=fastfood

# Copiar el archivo SQL a la carpeta de inicialización del MySQL
COPY fastfood.sql /docker-entrypoint-initdb.d/fastfood.sql