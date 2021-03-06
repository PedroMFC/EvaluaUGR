## Framework para el microservicio

Estamos usando el lenguaje de programación `Golang`. Para conocer los *frameworks* disponibles para el mismo se han consultado varios enlaces entre los que destacan:
* [Comparison of Popular Web Frameworks in Go](https://dzone.com/articles/comparison-of-popular-web-frameworks-in-go).
* [Top 5 Golang Frameworks in 2020](https://www.geeksforgeeks.org/top-5-golang-frameworks-in-2020/).

Vemos que existen una gran cantidad de opciones disponibles. Para facilitar la elección del *framework* nos hemos centrado en tres que nos han parecido interesantes. Brevemente explicadas:

* [Mux](https://github.com/gorilla/mux): implementa la interfaz `http.Handler` por lo que es compatible con el estándar `http.ServerMux`. Además, permite anidar rutas para la definición de grupos. Su desventaja es la rapidez ya que no es adecuado si se necesita un alto rendimiento.
* [Echo](https://github.com/labstack/echo): se considera un *framework* minimalista. Está muy optimizado para llevar a cabo el enrutamiento y se usa para construir aplicaciones escalables y robustas. Soporta tipo de datos como JSON y XML entre otros. Su principal desventaja es la frecuencia de actualización.
* [Gin](https://github.com/gin-gonic/gin): también se considera minimalista como `Echo` y está pensado para un alto rendimiento. También acepta formatos como JSON y da soporte para middleware personalizado. Sin embargo, no está pensado par desarrollar aplicaciones muy complejas.

Para llevar a cabo una mejor elección, se ha decido hacer pruebas reales para trabajar con las tres opciones y luego ver cuál se adapta mejor al problema. Para ello, en el archivo [main.go](./cmd/prueba/main.go) se ha creado un pequeño simulador de la aplicación. Como decidimos, su finalidad no es la de ser el programa de la aplicación sino tener una primera aproximación para poder trabajar con los diferentes marcos de trabajo. Por lo tanto, solo se han implementado dos rutas para obtener las valoraciones de una asignatura (`GET`) y otra para añadir valoraciones (`POST`). Para los datos usamos un *single source of truth*. Más adelante se hará una explicación de todas las rutas así como su relación con las historias de usuario. Por ahora, se han elegido estas ya que se consideran que las dos tareas son igual de importantes y se llevan a cabo con la misma frecuencia. Para ver la velocidad de cada uno de los opciones de las opciones se han creado unos *benchmarks* en el archivo [benchmark_test.go](./tests/benchmark_test.go) aprovechando la funcionalidad para medir el rendimiento  del paquete [testing](https://golang.org/pkg/testing/) que también usamos para la automatización de tests. En estas pruebas estamos realizando el mismo número de peticiones `GET` y `PUT`. Los resultados han sido los siguientes, en los que se ha lanzado cada una de las pruebas tres veces.

| *Framework* | Prueba 1 | Prueba 2 | Prueba 3 |
| -- | -- | -- | --- |
| Echo | 2.159s | 2.163s | 2.240s |
| Gorilla/Mux | 2.027s | 3.212s | 3.279s |
| Gin |  2.219s | 2.262s | 2.214s |

Vemos que el más lento es `Gorilla/Mux` mientras que los otros dos, presentan un rendimiento similar. Así, vistos los resultados y tras el uso de ambos se ha decidido usar `Gin` como *framework*. Destacar además que `Gin` se basa en el paquete `net/http` por lo qye sirve directamente cada petición en una `goroutine` diferente, tal y como se indica en [esta pregunta](https://github.com/gin-gonic/gin/issues/1335).

## Diseño general de la API

El diseño general de la API se puede ver en el archivo [routes.go](cmd/server/routes.go), más concretamente en la función `NewAppGin()`. **La explicación de las mismas así como su relación de las historias de usuario se encuentra en el archivo** [rutas.md][rutas] **de la documentación**. Se han intentado seguir las buenas prácticas en los códigos devueltos en las mismas: 201 (Created) cuando se crean las asignaturas o se envían valoraciones o reseñas por ejemplo, 404 (NotFound) si no existe el recurso, 400 (BadRequest) cuando la petición es incorrecta o 200 (StatusOk) cuando se obtienen correctamente los datos. Los tipos devueltos son todos tipo JSON. 

## Configuración distribuida y logs

Para la configuración distribuida se ha usado la herramienta `etcd`. Para ello, usamos el [cliente](https://github.com/etcd-io/etcd/tree/master/client/v3) específico para nuestro lenguaje de programación. Un uso básico del mismo lo hemos utilizado anteriormente en [este ejercicio de autoevaluación](https://github.com/PedroMFC/Autoevaluacion-CC/blob/main/semana%208-10/Ejercicio%201.md) donde también se usa [godotenv](https://github.com/joho/godotenv) para obtener variables de entorno del archivo `.env` (ignorado en el `.gitignore`). Su uso se puede ver en la función `Start()` del archivo [routes.go](cmd/server/routes.go). Esta función sería la que arrancaría el servicio en un futuro. Notamos que por ahora no hacemos uso de la misma, ya que no tenemos una "aplicación" lanzable del microservicio. La estructura del proceso de almacenamiento es la siguiente: primero miraría si hay un archivo `.env` disponible y comprobaría si tenemos definida la variable deseada. Si falla, el valor del puerto tomaría uno por defecto. En todos los casos se almacenaría en `etcd` el puerto en el que se ha lanzado la aplicación finalmente. Lo mismo hacemos con la dirección de conexión.

Para los logs, está la opción de utilizar la biblioteca [log](https://golang.org/pkg/log/) que es la que viene por defecto en el lenguaje, [logrus](https://github.com/sirupsen/logrus) que permite logs estructurados o [logo](https://github.com/mbndr/logo) entre [otros muchos](https://github.com/avelino/awesome-go#logging). En nuestro caso, hemos optado por usar `logrus`. Una de las razones de su uso es que gracias a [gin-logrus](https://github.com/toorop/gin-logrus), podemos usar este logger  como el sistema de logs que utiliza el *framework* `Gin` por defecto para las rutas (realmente usamos una leve modificación en [customlog.go](internal/customlog/customlog.go) para hacer más clara la salida). Además, `Gin` ya realiza el log para cada una de las peticiones realizadas. La información que incluimos en los logs es la petición realizada, el path y el resultado entre otros. Indicamos el uso de este log en la función `NewAppGin()` del archivo [routes.go](cmd/server/routes.go). 
 
## Tests

Para llevar a cabo los tests de la API se ha utilizado la herramienta [apitest](https://github.com/steinfletcher/apitest). Los tests escritos para las rutas se encuentran en el archivo [api_test.go](tests/api_test.go). Para ello, indicamos el *handler* que contiene las rutas a la biblioteca anterior, que se encarga de comprobar que funcionan como se espera. En cuanto al acceso a datos, se ha decidido mockear el mismo en archivo [mock_test.go](tests/mock_test.go) donde se indica en cada caso el comportamiento que queremos que tenga. 

## Avances
Dadas la nuevos requerimientos del proyecto, se ha trabajo en nuevas historias de usuario:
* [HU12][hu12]: crear una asignatura para valorar.
* [HU13][hu13]: crear una asignatura para reseñar.
* [HU14][hu14]: ver una reseña de una asignatura mediante su identificador (hasta ahora solo se podían obtener todas las reseñas a la vez).
* [HU15][hu15]: crear una asignatura para enviar preguntas.
* [HU16][hu16]: ver una pregunta de una asignatura mediante su identificador (como las reseñas, solo se podían obtener todas las preguntas asociadas)

Estas nuevas historias han hecho que sea necesario incluir nuevas comprobaciones como la de ver que una asignatura ya está registrada y que se recoge en [este issue][i76]. 

También se ha configurado el gestor de tareas para `build` e `install`. En `Go` tenemos las órdenes `go build` y `go install`. Para conocer el funcionamiento se ha consultado el enlace [Cómo crear e instalar programas de Go](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs-es). Básicamente, `go build` crea los binarios ejecutables de la aplicación mientras que `go install`, en vez de dejar el ejecutable en el directorio actual, también lo mueve directamente al directorio `$GOPATH/bin`. Por ello, en la construcción no se va a llevar a cabo ninguna acción ya que podemos hacer lo mismo usando directamente `install`. Estas nuevas tareas se encuentran en el archivo [Taskfile.yml](Taskfile.yml).