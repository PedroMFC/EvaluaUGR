# Contenedor para ejecutar los tests tomando como base Alpine
# Es necesario instalar Task y curl (este lo eliminamos después)
# La hacemos en varias etapas

# PRIMERA ETAPA
FROM golang:alpine AS BUILD_IMAGE

# Nos vamos al directorio de trabajo
WORKDIR ../

# Nos descargamos El gestor de tareas y curl para el mismo
RUN apk --no-cache add curl \
  && sh -c "$(curl -ssL https://taskfile.dev/install.sh)" -- -d \
  && apk del curl


# SEGUNDA ETAPA
FROM alpine:latest
LABEL maintainer="Pedro Flores <pedro_23_96@hotmail.com>"
LABEL version="0.2.1"

# Copiamos la carpeta en la que se ha descargado el gestor de tareas y donde está el lenguaje
COPY --from=BUILD_IMAGE /bin/task ./bin
COPY --from=BUILD_IMAGE /usr/local/go ./usr/local/go

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

# Comando a ejecutar 
CMD ["task", "test"]