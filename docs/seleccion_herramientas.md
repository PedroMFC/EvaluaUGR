# Selección de herramientas: trabajo autónomo segunda semana 

Explicaremos en conjunto tanto el trabajo realizado tanto para el avance de proyecto y los ejercicios de autoevaluación. 

## Selección del lenguaje

Para la elaboración del proyecto se han considerado una serie de lenguajes iniciales siendo los más destacados JavaScript, Go y Rust. Las razones para elegir entre estos tres lenguajes ha sido aprender un lenguaje nuevo, a pesar que en JavaScript ya tenía una pequeña base. Tras consultar varias páginas sobre el tema:

* [Lenguajes de programación que debes aprender en 2020](https://www.chakray.com/es/lenguajes-programacion-debes-aprender-2019/).
* [Lenguajes de programación 2020](https://computerhoy.com/noticias/tecnologia/estos-son-lenguajes-programacion-deberas-manejar-2020-607013).
* [Rust vs Go](https://medium.com/@devathon_/rust-vs-go-in-2020-1d472b5ee15).
* [Rust vs Go 2](https://stackshare.io/stackups/go-vs-rust).
* [Go vs JavaScript](https://www.slant.co/versus/111/126/~javascript_vs_go).
* [JavaScript vs Rust](https://www.slant.co/versus/111/5522/~javascript_vs_rust).
* [Cuáles son los lenguajes de programación con los sueldos más altos](https://www.businessinsider.es/cuales-son-lenguajes-programacion-sueldos-altos-737907).

se ha decido seleccionar Go como lenguaje de programación del proyecto. Las principales razones para elegirlo frente a los otros ha sido que en la bibliografía consultada se ha resaltado la facilidad de aprender el lenguaje desde cero así como la rapidez de programación en él. También se ha tenido en cuenta que es un lengiaje que actualemte está en auge y se prevé que en los próximos años hagan falta informáticos con conocimiento en el mismo. Como nunca se ha trabajado con él, han sido necesarios consultar tutoriales para conocer la sintaxis del programa y también ha sido necesaria (y será) una investigación para conocer las herramientas que dispone y que nos harán falta durante el desarrollo. Estos aspectos se irán comentando según se crea necesario. En primer lugar, ponemos los enlaces de los tutoriales de Go utilizados para conocer la sintaxis del mismo y realizar el primer acercamineto al mismo:

* [Tutorial GO I](https://tour.golang.org/list).
* [Tutorial GO II](https://golang.org/doc/tutorial/getting-started).

Se ha seguido [este tutorial](https://golang.org/doc/install) para instalar Go. También se ha seguido [esta guía](https://kgolding.co.uk/blog/2020/02/19/golang-application-directory-structure/) sobre cómo estructurar correctamente un proyecto en Go.

## Dependiencias
Pasamos ahora ha explicar cómo se van a manejar las dependencias del proyecto. Este apartado tiene relación con el [ejercicio 2](https://github.com/PedroMFC/Autoevaluacion-CC/blob/main/semana%202/Ejercicio2.md). En el trabajo de investigación se ha visto que tienes la opción de usar `vendoring` ([más información](https://riptutorial.com/go/topic/978/vendoring)) o `Go modules`. Se ha decidido usar esta última herramienta y su funcionamiento se ha visto en [este enlace](https://blog.friendsofgo.tech/posts/go-modules-en-tres-pasos/) y [este](https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31). Se ha priorizado su uso frente a `vendoring` ya que es una herramienta más moderna y que permite manejar las dependencias de una manera más sencilla.

## Tests
Este apartado guarda relación con el [ejercicio4](https://github.com/PedroMFC/Autoevaluacion-CC/blob/main/semana%202/Ejercicio4.md).  La información se ha obtenido de [este enlace](https://blog.golang.org/using-go-modules) con un uso básico de la librería estandar de GO y [este otro](https://bmuschko.com/blog/go-testing-frameworks/) donde explica algunos *frameworks* para los tests, donde destacamos [Testify](https://github.com/stretchr/testify) que sirve para las aserciones y *mocking* y [Ginkgo](https://github.com/onsi/ginkgo) que presenta el sistema *Behaviour Driven Development* que permiten describir los tests. En principio, se ha decidido usar `Ginkgo` que es el que más información proporciona a la hora de testear. Si accedemos a los enlaces correspondientes a `Ginkgo` y `Testify` vemos que la última *release* del primero es bastante reciente mientras que la otra fue a primeros de años. Sin embargo, `Testify` presenta mayor número de estrellas. Por lo tanto, se considera que ambas herramientas serían adecuadas para el desarrollo del proyecto.

## Integración Continua
Este apartado guarda relación con el [ejercicio5](https://github.com/PedroMFC/Autoevaluacion-CC/blob/main/semana%202/Ejercicio5.md). Para llevar a cabo la integración continua se va a usar `Travis` que es una herramienta que ya ha usado, por ejemplo en [DatosDemograficos-curso-tdd](https://github.com/tdd-organization-afp/DatosDemograficos). También se ha consultado [este enlace]() para añadir el *badge* al `README.md`.
 
## Cobertura
Para la cobertura de los test se va a usar `Codecov` que es una herramientos que también se usó en [DatosDemograficos-curso-tdd](https://github.com/tdd-organization-afp/DatosDemograficos) y presenta un muy buen funcionamiento junta con la herramienta de integración continua. Para añadir el *badge* al `README.md` se ha consultado [este enlace](https://stackoverflow.com/questions/54010651/codecov-io-badge-in-github-readme-md).
 

## Automatización de tareas

En caso de ser necearia la automatización de tareas se usará o [Tusk](https://github.com/rliebz/tusk) o [Task](https://taskfile.dev/#/). Se ha decidio usar el segundo ya que en GitHub aparece que se ha actualizado más recientemente y tiene mayor valoración y contribuidores.

## Base de datos
La herramientas de bases de datos barajadas han sido las usuales: MySQL, PostgreSQL, MongoDB, SQLite, MariaDB, etc. Todas las opciones serían adecuadas para el desarrollo del proyecto. Si comparamos PostgreSQL y MongoDB:
* [MongoDB vs PostgreSQL I](https://www.educative.io/blog/mongodb-versus-postgresql-databases).
* [MongoDB vs PostgreSQL II](https://www.educative.io/blog/mongodb-versus-postgresql-databases) 
  
vemos que ambas herramientas son robustas, adecuadas y cada una de ellas tiene sus ventajas e inconvenientes. Sin embargo, se ha decidio usar PostgeSQL, ya que, aunque MongoDB es adecuada para desarrollos ágiles, PostgeSQL tiene mejor itegración con otras herramientas que usaremos después como [Gorm](https://github.com/go-gorm/gorm) (una ORM que nos facilitará las operaciones en la base de datos). Además, presenta opciones como la facilidad de identificadores con autoincremento lo que será muy útil para poder identificar sin problemas cada una de las entradas de la base de datos. Los enlaces consultados sobre su utilización en el lenguaje y con otras herramientas han sido los siguientes:

* [Instalar PostgreSQL](https://www.digitalocean.com/community/tutorials/como-instalar-y-utilizar-postgresql-en-ubuntu-18-04-es).
* [PostgreSQL and Golang Tutorial](https://www.enterprisedb.com/postgres-tutorials/postgresql-and-golang-tutorial).
* [PostgreSQL en Travis](https://docs.travis-ci.com/user/database-setup/#postgresql).

## Comunicación y tareas

La aplicación que se va adesarrollar, va a utilizar microservicios como mostramos en al arquitectura de la misma. Para ello, es necesario un sistema para la comunicación entre los microservicios. Sobre este tema se han consultado los siguientes enlaces con distinas herrramientas:

* [Message Broker In Microservices](https://medium.com/@usha.dewasi/message-broker-in-microservices-c3c9dce003ef).
* [Redis, Kafka or RabbitMQ: Which MicroServices Message Broker To Choose?](https://otonomo.io/blog/redis-kafka-or-rabbitmq-which-microservices-message-broker-to-choose/)
* [AMQP, RabbitMQ and Celery - A Visual Guide For Dummies](https://www.abhishek-tiwari.com/amqp-rabbitmq-and-celery-a-visual-guide-for-dummies/).
* [Empezando con Apache Kafka en Golang](https://blog.friendsofgo.tech/posts/empezando-con-apache-kafka-en-golang/).
* [RabbitMQ](https://www.rabbitmq.com/tutorials/tutorial-one-go.html).
* [RabbitMQ GitHub](https://github.com/streadway/amqp).
* [Go RabbitMQ Beginners Tutorial](https://tutorialedge.net/golang/go-rabbitmq-tutorial/).
* [GoCelery I](https://godoc.org/github.com/taoh/gocelery).
* [GoCelery II](https://github.com/gocelery/gocelery).
* [Taskqueue](https://pkg.go.dev/google.golang.org/appengine/taskqueue).
* [Redis QuickStar](https://redis.io/topics/quickstart).
* [Redis GitHub](https://github.com/gomodule/redigo).

Una vez consultados estos enlaces y obtener una idea de cómo funcionan, se ha decido usar como cola de tareas para le ejecución asíncrona GoCelery. Este puede funcionar sobre Redis o AMQP(RAbbitMQ) como bróker de mensajes. Vemos que en GitHub Redis está mejor posicionado. Es cierto que presenta ciertas desventajas ya que, por ejemplo, no tiene persistencia y es más básico. Sin embargo, para nuestro aplicación puede funcionar adecuadamente. 

Otros enlaces para conocer mejor estos mecanismos han sido:
* [What is the difference between a message queue and a task queue?](https://www.quora.com/What-is-the-difference-between-a-message-queue-and-a-task-queue-Why-would-a-task-queue-require-a-message-broker-like-RabbitMQ-Redis-Celery-or-IronMQ-to-function)
* [What is the relationship between Celery and RabbitMQ?](https://stackoverflow.com/questions/43379554/what-is-the-relationship-between-celery-and-rabbitmq)

## Configuración remota

Los servicios en la nuben necesitan una configuración remota. Entre las [posibilidades a usar](https://www.g2.com/products/etcd/competitors/alternatives) tenemos `etcd`, `zookeper` o `consul`. Si vemos los repositorios de estas herramientas:

* [etcd](https://github.com/etcd-io/etcd).
* [Zookeper](https://github.com/go-zookeeper/zk).
* [Consul](https://github.com/hashicorp/consul).

y actuamos de manera similar a las herramientas anteriores vemos que las que mejor opiniones tienen son `Consul` y `etcd`. `etcd` tiene mayor valoración que `Consul` (aunque también tiene muchas estrellas) por lo que se va a usar esta. Su uso en Go se puede ver en [este tutorial](https://programmer.help/blogs/a-concise-tutorial-of-golang-etcd.html).

## *Logging*

También es necesara una herrmienta de *logging*. La información consultada para este tema se encuentra en
las siguientes páginas:
* [Logging in Go: Choosing a System and Using it](https://www.honeybadger.io/blog/golang-logging/).
* [Golang Logging: ship GoLang application logs to logstash](https://logit.io/sources/configure/golang).
* [How to integrate your Go Service with ELK](https://pmihaylov.com/go-service-with-elk/).

Por lo tanto se va usar `logrus` y `Logstash`. Dentro del lenguaje de programación no se han econtrado muchas alternativas a `logrus` pero aún vemos que está bien valorada en [su repositorio](https://github.com/sirupsen/logrus). [Otras alternativas](https://sematext.com/blog/logstash-alternatives/) a `Logstash` serían las mostradas pero la integración junto con `logrus` es mejor y además es [gratuito](https://www.elastic.co/es/blog/elasticsearch-free-open-limitless).

## API REST Framework
El acceso a los microservicios será mediante una API REST. Dentro del leguaje Go, tenemos [framworks](https://nordicapis.com/7-frameworks-to-build-a-rest-api-in-go/) para construirlo como Revel, Martini o Gorilla. En nuestro caso, hemos optado por usar [Gin](https://github.com/gin-gonic/gin), muy valorado en GitHub con 42.6K estrellas. Algunos ejemplos de su uso los podemos ver en [este tutorial](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/) o [en este](https://medium.com/wesionary-team/create-your-first-rest-api-with-golang-using-gin-gorm-and-mysql-d439bcc6f987) donde se observa una buena división del proyecto en carpetas.


[arquitectura]: COMPLETAR