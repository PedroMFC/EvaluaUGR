# Contenedor para ejecutar los tests tomando com base Alpine
# Es necesario instalar Task (y entre eso curl para descargarlo) y el lenguaje GO

FROM alpine:latest 
LABEL maintainer="Pedro Flores <pedro_23_96@hotmail.com>"
LABEL version="0.0.1"

# Descargamos Go, Task y curl (luego lo borramos porque no hace falta) 
# La instalación es necesaria hacerla como root
RUN apk --no-cache add curl go \
  && sh -c "$(curl -ssL https://taskfile.dev/install.sh)" -- -d \
  && apk del curl

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