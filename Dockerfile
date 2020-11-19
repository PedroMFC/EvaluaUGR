# Contenedor para ejecutar los tests tomando com base una imagen no oficial del lenguaje de GO
# Es necesario instalar Task 

FROM webhippie/golang:latest
LABEL maintainer="Pedro Flores <pedro_23_96@hotmail.com>"
LABEL version="0.0.2"

# Move to working directory 
WORKDIR ../../

# Nos descargamos El gestor de tareas. Go y curl ya vienen instalados
RUN sh -c "$(curl -ssL https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

# Añadimos el nuevo usuario que será el que ejecute los tests
RUN adduser -D evaluaugr \
    && addgroup -S evaluaugr evaluaugr 

# Variables de entorno necearias para el lenguaje
ENV PATH=$PATH:/go/bin:/usr/local/go/bin \
    GOPATH=/home/evaluaugr/go \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Cambiamos al nuevo usuario e incluimos el directorio de trabajo donde vamos a empezar
USER evaluaugr
WORKDIR /app/test

# Comando a ejecutar cuando se ejecute el contenedor
CMD ["task", "test"]