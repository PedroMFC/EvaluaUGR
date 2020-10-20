# Selección de herramientas: trabajo autónomo segunda semana 

Explicaremos en conjunto tanto el trabajo realizado tanto para el avance de proyecto y los ejercicios de autoevaluación. 

## Selección del lenguaje

Para la elaboración del proyecto se han considerado una serie de lenguajes iniciales siendo los más destacados JavaScript, Go y Rust. Tras consultar varias páginas sobre el tema:

* [Lenguajes de programación que debes aprender en 2020](https://www.chakray.com/es/lenguajes-programacion-debes-aprender-2019/).
* [Lenguajes de programación 2020](https://computerhoy.com/noticias/tecnologia/estos-son-lenguajes-programacion-deberas-manejar-2020-607013).
* [Rust vs Go](https://medium.com/@devathon_/rust-vs-go-in-2020-1d472b5ee15).
* [Rust vs Go 2](https://stackshare.io/stackups/go-vs-rust).
* [Go vs JavaScript](https://www.slant.co/versus/111/126/~javascript_vs_go).
* [JavaScript vs Rust](https://www.slant.co/versus/111/5522/~javascript_vs_rust).
* [Cuáles son los lenguajes de programación con los sueldos más altos](https://www.businessinsider.es/cuales-son-lenguajes-programacion-sueldos-altos-737907).

Se ha decido seleccionar Go como lenguaje de programación del proyecto. Como nunca se ha trabajado con él, han sido necesarios consultar tutoriales para conocer la sintaxis del programa y también ha sido necesaria (y será) una investigación para conocer las herramientas que dispone y que nos harán falta durante el desarrollo. Estos aspectos se irán comentando según se crea necesario. En primer lugar, ponemos los enlaces de los tutoriales de Go utilizados para conocer la sintaxis del mismo y realizar el primer acercamineto al mismo:

* [Tutorial GO I](https://tour.golang.org/list).
* [Tutorial GO II](https://golang.org/doc/tutorial/getting-started).

Se ha seguido [este tutorial](https://golang.org/doc/install) para instalar Go. También se ha seguido [esta guía](https://kgolding.co.uk/blog/2020/02/19/golang-application-directory-structure/) sobre cómo estructurar correctamente un proyecto en Go.

## Dependiencias
Pasamos ahora ha explicar cómo se van a manejar las dependencias del proyecto. Este apartado tiene relación con el [ejercicio 2](https://github.com/PedroMFC/Autoevaluacion-CC/blob/main/semana%202/Ejercicio2.md). En el trabajo de investigación se ha visto que tienes la opción de usar `vendoring` ([más información](https://riptutorial.com/go/topic/978/vendoring)) o `Go modules`. Se ha decidido usar esta úñtima herramienta y su funcionamiento se ha visto en [este enlace](https://blog.friendsofgo.tech/posts/go-modules-en-tres-pasos/) y [este](https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31).

## Tests
Este apartado guarda relación con el [ejercicio4](https://github.com/PedroMFC/Autoevaluacion-CC/blob/main/semana%202/Ejercicio4.md).  La información se ha obtenido de [este enlace](https://blog.golang.org/using-go-modules) con un uso básico de la librería estandar de GO y [este otro](https://bmuschko.com/blog/go-testing-frameworks/) donde explica algunos *frameworks* para los tests, donde destacamos [Testify](https://github.com/stretchr/testify) que sirve para las aserciones y *mocking* y [Ginkgo](https://github.com/onsi/ginkgo) que presenta el sistema *Behaviour Driven Development* que permiten describir los tests. En principio, se ha decidido usar `Ginkgo` que es el que más información proporciona.

## Integración Continua
Este apartado guarda relación con el [ejercicio5](https://github.com/PedroMFC/Autoevaluacion-CC/blob/main/semana%202/Ejercicio5.md). Para llevar a cabo la integración continua se va a usar `Travis` que es una herramienta que ya ha usado, por ejemplo en [DatosDemograficos-curso-tdd](https://github.com/tdd-organization-afp/DatosDemograficos). También se ha consultado [este enlace]() para añadir el *badge* al `README.md`.
 
## Cobertura
Para la cobertura de los test se va a usar `Codecov` que es una herramientos que también se usó en [DatosDemograficos-curso-tdd](https://github.com/tdd-organization-afp/DatosDemograficos). Para añadir el *badge* al `README.md` se ha consultado [este enlace](https://stackoverflow.com/questions/54010651/codecov-io-badge-in-github-readme-md).
 

## Automatización de tareas

En caso de ser necearia la automatización de tareas se usará o [Tusk](https://github.com/rliebz/tusk) o [Task](https://taskfile.dev/#/), preferblemente el segundo de ellos.

## Base de datos
La herramientas de bases de datos barajadas han sido las usuales: MySQL, PostgreSQL, MongoDB, etc. Se ha decidio usar MongoDB para almacenar la información. Los enlaces consultados sobre su utilización en el lenguaje y con otras herramientas han sido los siguientes:

* [Instalar MondgoDB](https://www.digitalocean.com/community/tutorials/how-to-install-mongodb-on-ubuntu-18-04).
* [¿Cómo usar MongoDB driver oficial en Go?](https://blog.friendsofgo.tech/posts/driver-oficial-mongodb-golang/).
* [MongoDB en Travis](https://docs.travis-ci.com/user/database-setup/#mongodb).

## Comunicación

La aplicación que se va adesarrollar, va a utilizar microservicios como mostramos en al [arquitectura][arquitectura] de la misma. Para ello, es necesario una cola de tareas o un bróker de mensajería para la comunicación entre los microservicios. Sobre este tema se han consultado los siguientes enlaces con distanas herrramientas:

* [Message Broker In Microservices](https://medium.com/@usha.dewasi/message-broker-in-microservices-c3c9dce003ef).
* [Empezando con Apache Kafka en Golang](https://blog.friendsofgo.tech/posts/empezando-con-apache-kafka-en-golang/).
* [RabbitMQ en Go](https://www.rabbitmq.com/tutorials/tutorial-one-go.html).
* [Go RabbitMQ Beginners Tutorial](https://tutorialedge.net/golang/go-rabbitmq-tutorial/).
* [GoCelery I](https://godoc.org/github.com/taoh/gocelery).
* [GoCelery II](https://github.com/gocelery/gocelery).
* [Taskqueue](https://pkg.go.dev/google.golang.org/appengine/taskqueue).

Como se observa, las opciones son GoCelery, Kafka, RabbitMQ y Taskqueue. En principio, se ha decidio usar RabbitMQ aunque no se descarta que se use GoCelery en función del avance del proyecto.

## Configuración remota

Para la configuración temota se va hacer uso de `etcd` cuyo uso en Go se puede ver en [este tutorial](https://programmer.help/blogs/a-concise-tutorial-of-golang-etcd.html).

## *Logging*

También es necesara una herrmienta de *logging*. La información consultada para este tema se encuentra en
las siguientes páginas:
* [Logging in Go: Choosing a System and Using it](https://www.honeybadger.io/blog/golang-logging/).
* [Golang Logging: ship GoLang application logs to logstash](https://logit.io/sources/configure/golang).
* [How to integrate your Go Service with ELK](https://pmihaylov.com/go-service-with-elk/).

Por lo tanto se va usar `logrus` y `Logstash`.

[arquitectura]: COMPLETARRRRRRRRRRRR