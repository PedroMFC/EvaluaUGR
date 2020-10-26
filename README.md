[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[![Build Status](https://travis-ci.org/PedroMFC/EvaluaUGR.svg?branch=main)](https://travis-ci.org/PedroMFC/EvaluaUGR)

# EvaluaUGR
Proyecto para la asignatura de Cloud Computing del Máster en Ingeniería Informática.

## Arquitectura

Tras consultar la información sobre las diferentes arquitecturas disponibles, se ha decidio que la arquitectura utilizada sea una basada en **microservicios**. Si consultamos en páginas como [Medium](https://medium.com/@goodrebels/microservicios-ventajas-y-contras-de-la-arquitectura-descentralizada-a3b7fc814422) o [RedHat](https://www.redhat.com/es/topics/microservices/what-are-microservices), vemos que esta arquitectura presenta una serie de ventajas como puede ser la versatilidad, facilidad de integración y escalado o la capacidad de recuperación ya que el fallo de un microservicio no afecta a los demás. Sin embargo, también presenta una serie de inconvenientes que debemos de tener en cuenta como la dificultad de gestionar gran cantidad de microservicios en una aplicación. 

En nuestro proyecto, se cuenta con tres microservicios obtenidos a partir de las historias de usuario que se pueden consultar en el apartado [issues]:

* **Valoraciones**: este microservicio será el encargado de llevar el control acerca de las valoraciones "numéricas" de las asignaturas y aportar los resultados tal y como desean los usuarios.
* **Reseñas**: por su parte, este se encargará de gestionar las opiniones/reseñas que los usuarios quieran incluir para cada una de las asignaturas.
* **Preguntas**: para finalizar, este microservicio se encargará de controlar las preguntas y repuestas que formulen los distintos usuarios.

Para el acceso a estos microservicios se creará una API REST. 

## Issues y milestones

* [Arcvhivos auxiliares][mAuxiliar]: para mantener los archivos auxiliares completos.
  * [Añadir traivs][i10].
  * [Añadir cc.yaml][i42].
  * [Incluir archivo de tareas][i36].
  * [Actualizar travis][i10].
  * [Añadir badge travis][i9].
  * [Añadir descripción][i3]

* [La estructura del proyecto][mEstructura]: tener una estructura adecuada en el proyecto y como marca el lenguaje. 
  * [Renombrar archivos][i38].
  * [Arreglar nombre de funciones][i41].
  * [Renombrar módulos][i40].
  * [Actualizar dependencias][i28].
  * [Crear módulo principal][i23].

* [Buena documetación][mDocumentacion]: tener una buena documentación.
  * [Añadir carpeta de documentación][i1]
  * [Documentar trabajo semana 1][i2].
  * [Documentar trabajo semana 2][i8].
  * [Describir arquitectura][i11].
  * [Incluir issues y milestones][i44].
  * [Explicar herramienta gofmt][i35].
  * [Arreglar fallo enlaces documentación][i22].
  * [Renombrar archivos documentación][i21].
  * [Cambiar nombre del proyecto][i7].
  * [Mover problema del REDAME][i43].


* [Gestionar valoraciones][mValoraciones]: gestión de toda la funcionalidad relacionada con las valoraciones.
  * [[HU1] Valorar una asignatura.][i12].
  * [[HU2] Conocer valoraciones asignatura][i13].
  * [[HU3] Valoración media de una asignatura ][i14].
  * [Conexión a la base de datos][i26].
  * [Interfaz con las operaciones][i27].
  * [Crear modelos de valoración][i24].
  

* [Gestionar reseñas][mResenias]: gestión de toda la funcionalidad relacionada con las reseñas/opiniones.
  * [[HU4] Enviar reseña asignatura ][i15].
  * [[HU5] Ver reseñas de una asignatura][i16].
  * [[HU6] Valorar reseña de una asignatura ][i17].
  * [Conexión a la base de datos reseñas][i33].
  * [Crear modelo de reseña][i29].
  * [Actualizar modelo][i39].
  * [Interfaz con las operaciones][i30].
  

* [Gestionar preguntas][mPreguntas]: gestión de toda la funcionalidad relacionada con las preguntas.
  * [[HU7] Hacer pregunta de una asignatura][i18].
  * [[HU8] Ver preguntas de una asignatura][i19].
  * [[HU9] Responder pregunta de una asignatura][i20].
  * [Conexión a la base de datos][i34].
  * [Crear modelo de pregunta][i31].
  * [Crear interfaz de operaciones][i32].
  * [Solucionar errores][mErrores]: para solucionar errores.
  * [Arreglar travis][i37].



## Documentación
* [Problema a resolver][problema].
* [Puesta a punto del repositorio][configGitHub].
* [La documentación sobre la selección de herramientas][herramientas].

[configGitHub]: https://pedromfc.github.io/EvaluaUGR/docs/configuracion_github
[herramientas]: https://pedromfc.github.io/EvaluaUGR/docs/seleccion_herramientas
[problema]: https://pedromfc.github.io/EvaluaUGR/docs/problema
[issues]: https://github.com/PedroMFC/EvaluaUGR/issues



[mAuxiliar]: https://github.com/PedroMFC/EvaluaUGR/milestone/2
[mPreguntas]: https://github.com/PedroMFC/EvaluaUGR/milestone/5
[mErrores]: https://github.com/PedroMFC/EvaluaUGR/milestone/3
[mEstructura]: https://github.com/PedroMFC/EvaluaUGR/milestone/7
[mResenias]: https://github.com/PedroMFC/EvaluaUGR/milestone/6
[mDocumentacion]: https://github.com/PedroMFC/EvaluaUGR/milestone/1
[mValoraciones]: https://github.com/PedroMFC/EvaluaUGR/milestone/4

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
