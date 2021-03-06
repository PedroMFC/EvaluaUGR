## Gestor de tareas. Ejecución de los tests

Como se indicó anteriormente en el apartado [selección de herramientas][herramientas], se ha escogido como gestor de tareas [Task](https://taskfile.dev/#/). En su momento otras alternativas a usar eran por ejemplo [Tusk](https://github.com/rliebz/tusk). Sin embargo, se optó por esta herramienta ya que presenta una actualización reciente así como gran cantidad de estrellas y colaboradores en GitHub. Además, en su [documentación](https://taskfile.dev/#/usage) vemos que es una herramienta flexible que permite incluir resúmenes a las órdenes para saber qué hacen, dependencias entre tareas ... Por ejemplo, si ejecutamos la orden `task -l` vemos las tres tareas incluidas actualmente en el archivo [Taskfile](./Taskfile.yml) junto con una pequeña descripción de cada una. Más concretamente, las tareas incluidas son:
* format_code:  formatear código
* syntax_check: comprobar sintaxis
* test:         ejecutar tests

Si quisiéramos ejecutar alguna de ellas solo tendríamos que escribir `task nombre-tarea`. Finalmente, la herramienta `Task` se puede instalar fácilmente mediante el script  [install-task.sh](https://github.com/PedroMFC/EvaluaUGR/blob/main/install-task.sh).

## Biblioteca de aserciones

Del mismo modo que con el gestor de tareas, se decidió que usar [Testify](https://github.com/stretchr/testify) como biblioteca de aserciones. De este modo, se ha optado por un enfoque TDD. En [este enlace](https://bmuschko.com/blog/go-testing-frameworks/) aparecen una gran cantidad de bibliotecas el lenguaje escogido. A parte de `Testify`, encontramos otras como [Ginkgo](https://github.com/onsi/ginkgo) y [Goblin](https://github.com/franela/goblin) que presentan el sistema *Behaviour Driven Development* o incluso la propia librería estándar de `Go`. Las razones para el uso de `Testify` en el proyecto son por una lado, su valoración en GitHub. Al igual que el gestor de tareas, presenta una actualización reciente y gran número de estrellas y contribuidores. Por otro lado, la manera de trabajar con ella es sencilla, las funciones para aserciones que presenta son adecuadas para el proyecto y se integra bien con la librería estándar del lenguaje. Por lo que respecta a la escritura de los tests, aparte de la propia documentación de la herramienta se ha usado [este enlace](http://www.inanzzz.com/index.php/post/2t08/using-setup-and-teardown-in-golang-unit-tests). 

## Sistema de prueba de código

Como el lenguaje de desarrollo del proyecto es `Go`, se ha decidido usar el mecanismo estándar para ejecutar los tests en el mismo. Se ejecutan los tests escritos mediante [go test](https://golang.org/pkg/cmd/go/internal/test/) que ya viene proporcionado en el paquete de pruebas propio del lenguaje. Con este mecanismo, podemos ejecutar todos los tests contenidos en archivos del tipo `*_test.go`. También se ha decidido usar este método ya que es posible obtener un porcentaje de cobertura de los tests o indicar archivos o funciones específicos que ejecutar entre otros. 


## Clases y tests

Se ha avanzado código en los siguientes archivos:

- [Valoracion](./internal/microval/modelsval/valoracion.go)
- [ValoracionRepositorio](./internal/microval/modelsval/valoracionrepositorio.go)
- [Resenia](./internal/microres/modelsres/resenia.go)
- [ReseniaRepositorio](./internal/microres/modelsres/reseniasrepositorio.go)
- [Pregunta](./internal/micropre/modelspre/pregunta.go)
- [PreguntaRepositorio](./internal/micropre/modelspre/preguntasrepositorio.go)
- [Respuesta](./internal/micropre/modelspre/respuesta.go)
- [Asignatura](./internal/asignatura/asig/asignatura.go)

Y los archivos de test correspondientes son:

- [SetUp/TearDown](./tests/testmain_test.go)
- [valoracion_test](./tests/valoracion_test.go)
- [resenias_test](./tests/resenias_test.go)
- [preguntas_test](./tests/preguntas_test.go)
- [asignatura_test](./tests/asignatura_test.go)


[configGitHub]: https://pedromfc.github.io/EvaluaUGR/docs/configuracion_github
[herramientas]: https://pedromfc.github.io/EvaluaUGR/docs/seleccion_herramientas
[problema]: https://pedromfc.github.io/EvaluaUGR/docs/problema
[arquitectura]: https://pedromfc.github.io/EvaluaUGR/docs/arquitectura
[issues]: https://github.com/PedroMFC/EvaluaUGR/issues
[planificacion]: https://pedromfc.github.io/EvaluaUGR/docs/planificación


[mAuxiliar]: https://github.com/PedroMFC/EvaluaUGR/milestone/2
[mPreguntas]: https://github.com/PedroMFC/EvaluaUGR/milestone/5
[mErrores]: https://github.com/PedroMFC/EvaluaUGR/milestone/3
[mEstructura]: https://github.com/PedroMFC/EvaluaUGR/milestone/7
[mResenias]: https://github.com/PedroMFC/EvaluaUGR/milestone/6
[mDocumentacion]: https://github.com/PedroMFC/EvaluaUGR/milestone/1
[mValoraciones]: https://github.com/PedroMFC/EvaluaUGR/milestone/4
[mTests]: https://github.com/PedroMFC/EvaluaUGR/milestone/8

[i1]: https://github.com/PedroMFC/EvaluaUGR/issues/1
[i2]: https://github.com/PedroMFC/EvaluaUGR/issues/2
[i3]: https://github.com/PedroMFC/EvaluaUGR/issues/3
[i4]: https://github.com/PedroMFC/EvaluaUGR/issues/4
[i5]: https://github.com/PedroMFC/EvaluaUGR/issues/5
[i6]: https://github.com/PedroMFC/EvaluaUGR/issues/6
[i7]: https://github.com/PedroMFC/EvaluaUGR/issues/7
[i8]: https://github.com/PedroMFC/EvaluaUGR/issues/8
[i9]: https://github.com/PedroMFC/EvaluaUGR/issues/9
[i10]: https://github.com/PedroMFC/EvaluaUGR/issues/10
[i11]: https://github.com/PedroMFC/EvaluaUGR/issues/11
[i12]: https://github.com/PedroMFC/EvaluaUGR/issues/12
[i13]: https://github.com/PedroMFC/EvaluaUGR/issues/13
[i14]: https://github.com/PedroMFC/EvaluaUGR/issues/14
[i15]: https://github.com/PedroMFC/EvaluaUGR/issues/15
[i16]: https://github.com/PedroMFC/EvaluaUGR/issues/16
[i17]: https://github.com/PedroMFC/EvaluaUGR/issues/17
[i18]: https://github.com/PedroMFC/EvaluaUGR/issues/18
[i19]: https://github.com/PedroMFC/EvaluaUGR/issues/19
[i20]: https://github.com/PedroMFC/EvaluaUGR/issues/20
[i21]: https://github.com/PedroMFC/EvaluaUGR/issues/21
[i22]: https://github.com/PedroMFC/EvaluaUGR/issues/22
[i23]: https://github.com/PedroMFC/EvaluaUGR/issues/23
[i24]: https://github.com/PedroMFC/EvaluaUGR/issues/24
[i25]: https://github.com/PedroMFC/EvaluaUGR/issues/25
[i26]: https://github.com/PedroMFC/EvaluaUGR/issues/26
[i27]: https://github.com/PedroMFC/EvaluaUGR/issues/27
[i28]: https://github.com/PedroMFC/EvaluaUGR/issues/28
[i29]: https://github.com/PedroMFC/EvaluaUGR/issues/29
[i30]: https://github.com/PedroMFC/EvaluaUGR/issues/30
[i31]: https://github.com/PedroMFC/EvaluaUGR/issues/31
[i32]: https://github.com/PedroMFC/EvaluaUGR/issues/32
[i33]: https://github.com/PedroMFC/EvaluaUGR/issues/33
[i34]: https://github.com/PedroMFC/EvaluaUGR/issues/34
[i35]: https://github.com/PedroMFC/EvaluaUGR/issues/35
[i36]: https://github.com/PedroMFC/EvaluaUGR/issues/36
[i37]: https://github.com/PedroMFC/EvaluaUGR/issues/37
[i38]: https://github.com/PedroMFC/EvaluaUGR/issues/38
[i39]: https://github.com/PedroMFC/EvaluaUGR/issues/39
[i40]: https://github.com/PedroMFC/EvaluaUGR/issues/40
[i41]: https://github.com/PedroMFC/EvaluaUGR/issues/41
[i42]: https://github.com/PedroMFC/EvaluaUGR/issues/42
[i43]: https://github.com/PedroMFC/EvaluaUGR/issues/43
[i44]: https://github.com/PedroMFC/EvaluaUGR/issues/44
[i45]: https://github.com/PedroMFC/EvaluaUGR/issues/45
[i46]: https://github.com/PedroMFC/EvaluaUGR/issues/46
[i47]: https://github.com/PedroMFC/EvaluaUGR/issues/47
[i48]: https://github.com/PedroMFC/EvaluaUGR/issues/48
[i49]: https://github.com/PedroMFC/EvaluaUGR/issues/49
[i50]: https://github.com/PedroMFC/EvaluaUGR/issues/50
[i51]: https://github.com/PedroMFC/EvaluaUGR/issues/51
[i52]: https://github.com/PedroMFC/EvaluaUGR/issues/52
[i53]: https://github.com/PedroMFC/EvaluaUGR/issues/53
[i54]: https://github.com/PedroMFC/EvaluaUGR/issues/54
[i55]: https://github.com/PedroMFC/EvaluaUGR/issues/55
[i56]: https://github.com/PedroMFC/EvaluaUGR/issues/56
[i57]: https://github.com/PedroMFC/EvaluaUGR/issues/57
[i58]: https://github.com/PedroMFC/EvaluaUGR/issues/58
[i59]: https://github.com/PedroMFC/EvaluaUGR/issues/59
[i60]: https://github.com/PedroMFC/EvaluaUGR/issues/60
