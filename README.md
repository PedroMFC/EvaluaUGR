[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

[![Build Status](https://travis-ci.org/PedroMFC/EvaluaUGR.svg?branch=main)](https://travis-ci.org/PedroMFC/EvaluaUGR)

[![Build status](https://ci.appveyor.com/api/projects/status/j0jnyv7lgm7mkjkn?svg=true)](https://ci.appveyor.com/project/PedroMFC/evaluaugr)

# EvaluaUGR
Proyecto para la asignatura de Cloud Computing del Máster en Ingeniería Informática.

## Travis

La primera herramienta que se ha utilizado para la integración continua ha sido Travis. Esta es una herramienta que ya se había usado anteriormente aunque no con el lenguaje `Golang`. Se ha decidido que en Travis se ejecute el **contenedor** que hemos creado para poder lanzar los test. En este caso, estamos trabajando con la [útlima versión](https://golang.org/doc/devel/release.html) del lenguaje, la 1.15.6. 

Para poder realizar la integración continua en Travis se ha creado el archivo [.travis.yml](.travis.yml). Para ello, hemos usado los siguientes enlaces de ejemplo:
* [Raku-aulas](https://github.com/JJ/raku-aulas/blob/master/.travis.yml): archivo de CI en el ejemplo visto en clase. Para conocer cómo trabajar con Docker y Travis.
* [Building a Go Project](https://docs.travis-ci.com/user/languages/go/): explicación de cómo usar Travis con Go.
* [Travis Lifecicle](https://docs.travis-ci.com/user/job-lifecycle/).

En primer lugar lo que hacemos es indicar que necesitamos el lenguaje `Go` en su última versión. En un principio no lo tendríamos que usar ya que estamos utilizando un contenedor. Sin embargo, lo necesitamos para poder instalar el gestor de tareas `Task` y lanzar el contenedor mediante una orden del mismo. Luego, indicamos la variable de entorno `GO111MODULE=on`, que es necesario a partir de la versión 1.11 del lenguaje para llevar el control de las dependencias mediante la herramienta Go Modules. En el apartado `services` indicamos que queremos usar `Docker`. Ahora, *sobreescrbimos* el comportamiento por defecto que tiene Travis para Go en los apartados `install` y `script` para adaptarlos a nuestras necesidades. En la primera de ellas nos descargamos el contenedor que hemos creado para ejecutar los tests (no necesitaríamos especificarlo ya que lo descargaría cuando se ejecute, pero se ha decidido hacerlo así por si hay algún problema en la descarga verlo antes) e instalamos el gestor de tareas mediante el archivo [install-task.sh](install-task.sh) creado para tal efecto. Finalmente, en `script` ejecutamos la tarea `task docker_travis` (especificada en [Taskfile.yml](Taskfile.yml)) pata ejecutar el contenedor. 

También se ha creado una tarea para poder ejecutar directamente el contenedor en local.

## Alternativas

Una vez seleccionado Travis como primera herramienta de integración continua, vemos otras alternativas:

* [Alternatives to Travis](https://alternativeto.net/software/travis-ci/?license=free).
* [What is Travis CI and what are its top alternatives?](https://stackshare.io/travis-ci/alternatives)

Algunas que nos han llamado la atención son Gitlab (aunque *solo* tiene [400 minutos al mes de manera gratuita](https://about.gitlab.com/pricing/)), CircleCI (con [2 500 créditos a la semana](https://circleci.com/pricing/?utm_source=gb&utm_medium=SEM&utm_campaign=SEM-gb-200-Eng-ni&utm_content=SEM-gb-200-Eng-ni-CirclePricing&utm_term=a2&gclid=Cj0KCQiAzZL-BRDnARIsAPCJs71IGpSUzcd8woVIuXc0MY6RC0ytEGOQMN6FLBNqN4qw-h55ijrqEuYaAtV1EALw_wcB) para la cuenta gratis), Jenkins y AppVeyor (ambos totalmente gratuitos). De este modo, para no tener que estar pendiente de si nos pasamos con los créditos que dan de manera gratuita, se ha escogido trabajar con **Appveyor**. La única restricción para la cuenta gratuita es que solamente se puede ejecutar un *job* a la vez pero en nuestro caso no es ningún impedimento. Los principales enlaces consultados han sido:
* [AppVeyor Build Pipeline](https://www.appveyor.com/docs/build-configuration/#build-pipeline).
* [Testing with Go | AppVeyor](https://www.appveyor.com/docs/lang/go/).

El archivo que se ha escrito es [appveyor.yml](appveyor.yml). Hemos observado que la sintaxis es similar a la de Travis. Lo primero que hacemos es indicar que vamos a utilizar una máquina Ubuntu y que no vamos a hacer *build*, solo vamos a ejecutar los tests. En la parte `stack` indicamos las versiones del lenguaje que vamos a testear. En [este enlace](https://www.appveyor.com/docs/linux-images-software/#golang) podemos ver las diferentes versiones disponibles. Como estamos usando GoModules para la gestión de dependencias, tenemos que usar una versión de Go mayor que la 1.11. Sin embargo, la biblioteca de aserciones da soporte como mínimo a la versión 1.13. Por ello, se van a testear las versiones 1.13.15 y 1.14.12 (son las versiones que tienen instaladas y [las últimas de las respectivas versiones](https://golang.org/doc/devel/release.html)). No probamos la última versión ya que esta se testea en Travis. Finalmente, en la parte `install` nos descargamos el gestor de tareas y en `test_script` ejecutamos los tests mediante la orden especificada en el archivo de tareas [Taskfile.yml](Taskfile.yml). 

## Avances

En esta ocasión lo que se ha hecho ha sido un *refactoring* del código ya existente según los issues [#70][i70], [#71][i71] y [#72][i72], correspondientes a cada una de las estructuras *map* que teníamos predefinidas. Hasta ahora, las clases que controlaban la lógica dependían de la estructura de almacenamiento escogida y rompíamos el principio de inversión de dependencias. .Así lo que se ha hecho es crear las interfaces [IValSaver](internal/microval/modelsval/valsaverinterface.go), [IResSaver](internal/microres/modelsres/ressaverinterface.go) e [IPreSaver](internal/micropre/modelspre/presaverinterface.go) que definen el comportamiento que deben tener las clases que se dediquen exclusivamente al almacenamiento de datos. Siguiendo estas interfaces se han creado las clases [ValoracionMap](internal/microval/modelsval/valoracionmap.go), [ReseniasMap](internal/microres/modelsres/reseniasmap.go) y [PreguntasMap](internal/micropre/modelspre/preguntasmap.go) que implementan las interfaces y usan una tabla hash para almacenar los datos. Ahora las clases [ValoracionRepositorio](internal/microval/modelsval/valoracionrepositorio.go), [ReseniasRepositorio](internal/microres/modelsres/reseniasrepositorio.go) y [PreguntasRespositorio](internal/micropre/modelspre/preguntasrepositorio.go) reciben un objeto de las interfaces correspondientes que es el que utilizan. 

Se han creado los archivos de tests correspondientes a estas nuevas clases: [valmapsaver_test.go](tests/valmapsaver_test.go), [resmapsaver_test.go](tests/resmapsaver_test.go) y [premapsaver_tes.go](tests/premapsaver_test.go). Los ya existentes ([valoracion_tes.go](tests/valoracion_test.go), [resenias_test.go](tests/resenias_test.go) y [preguntas_test.go](tests/preguntas_test.go)) se han modificado *falseando* el comportamiento de las interfaces creadas con los *mocks* situados en la carpeta [mocks](./mocks). Se ha usado [Testify](https://github.com/stretchr/testify) que ya veníamos usando para las aserciones y que tiene soporte para trabajar con *mocks*. Los mismos se han creado con la herramienta [Mockery](https://github.com/vektra/mockery) que crea los archivos de la nueva clase directamente a partir de la definición de la interfaz. Para este proceso se ha consultado [este enlace](https://blog.codecentric.de/2019/07/gomock-vs-testify/).


## Documentación
Puede consultar más información acerca del proyecto en los siguientes enlace:

* [Problema a resolver][problema].
* [Puesta a punto del repositorio][configGitHub].
* [La documentación sobre la selección de herramientas][herramientas].
* [Información sobre la arquitectura empleada][arquitectura].
* [Planificación][planificacion].
* [Test][tests].
* [Docker][dockerR].

[configGitHub]: https://pedromfc.github.io/EvaluaUGR/docs/configuracion_github
[herramientas]: https://pedromfc.github.io/EvaluaUGR/docs/seleccion_herramientas
[problema]: https://pedromfc.github.io/EvaluaUGR/docs/problema
[arquitectura]: https://pedromfc.github.io/EvaluaUGR/docs/arquitectura
[issues]: https://github.com/PedroMFC/EvaluaUGR/issues
[planificacion]: https://pedromfc.github.io/EvaluaUGR/docs/planificación
[docker]: https://pedromfc.github.io/EvaluaUGR/docs/docker
[tests]: https://pedromfc.github.io/EvaluaUGR/docs/tests
[dockerR]: https://pedromfc.github.io/EvaluaUGR/docs/docker_README

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

[i70]: https://github.com/PedroMFC/EvaluaUGR/issues/70
[i71]: https://github.com/PedroMFC/EvaluaUGR/issues/71
[i72]: https://github.com/PedroMFC/EvaluaUGR/issues/72
