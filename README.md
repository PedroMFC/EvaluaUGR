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

## Documentación
* [Problema a resolver][problema].
* [Puesta a punto del repositorio][configGitHub].
* [La documentación sobre la selección de herramientas][herramientas].

[configGitHub]: https://pedromfc.github.io/EvaluaUGR/docs/configuracion_github
[herramientas]: https://pedromfc.github.io/EvaluaUGR/docs/seleccion_herramientas
[problema]: https://pedromfc.github.io/EvaluaUGR/docs/problema
[issues]: https://github.com/PedroMFC/EvaluaUGR/issues
