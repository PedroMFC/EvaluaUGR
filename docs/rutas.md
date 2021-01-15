# Rutas de la API

A continuación explicamos las rutas de la aplicación así como su relación con las historias de usuario definidas definidas en el archivo [routes.go](../cmd/server/routes.go).

## Valoraciones

Funciones implementados en el archivo [handlersval.go](../internal/microval/handlersval/handlersval.go)

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| PUT | /valoraciones/asignatura/:asig | [HU12][hu12] | Crear asignatura para valorar | `CrearAsignatura` |
| GET | /valoraciones/asignatura/:asig | [HU2][hu2] | Obtener valoraciones de una asignatura | `GetValoraciones` |
| POST | /valoraciones/asignatura/:asig/:val | [HU1][hu1]  | Añadir valoración a una asignatura | `Valorar` |
| GET | /valoraciones/peor | [HU10][hu10] | Obtener asignatura peor valorada | `GetPeor` |
| GET | /valoraciones/mejor | [HU11][hu11] | Obtener asignatura mejor valorada | `GetMejor` |
| GET | /valoraciones/asignatura/:asig/media | [HU3][hu3] | Obtener valoración media de una asignatura | `GetMedia` |


## Reseñas

Funciones implementados en el archivo [handlersres.go](../internal/microres/handlersres/handlersres.go)

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| PUT | /resenias/asignatura/:asig | [HU13][hu13] | Crear una asignatura para reseñar | `CrearAsignatura` |
| GET | /resenias/asignatura/:asig | [HU5][hu5] | Ver reseñas de una asignatura | `GetResenias` |
| GET | /resenias/asignatura/:asig/:id | [HU14][hu14] | Obtener una sola reseña de una asignatura | `GetResenia` |
| POST | /resenias/asignatura/:asig | [HU4][hu4] | Añadir una reseña a una asignatura | `Opinar` |
| POST | /resenias/asignatura/:asig/gusta | [HU6][hu6] | Indicar que una reseña ha gustado | `GustaResenia`|
| POST | /resenias/asignatura/:asig/nogusta | [HU6][hu6] | Indicar que una reseña no ha gustado | `NoGustaResenia` |

## Preguntas

Funciones implementados en el archivo [handlerspre.go](../internal/micropre/handlerspre/handlerspre.go)

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| PUT | /preguntas/asignatura/:asig | [HU15][hu15] | Crear asignatura para preguntar | `CrearAsignatura` |
| GET | /preguntas/asignatura/:asig | [HU8][hu8] | Obtener preguntas de una asignatura | `GetPreguntas` |
| GET | /preguntas/asignatura/:asig/:id | [HU16][hu16] | Obtener una pregunta de una asignatura | `GetPregunta` |
| POST | /preguntas/asignatura/:asig | [HU7][hu7] | Añadir una pregunta a una asignatura | `Preguntar` |
| POST | /preguntas/asignatura/:asig/:id | [HU9][hu9] | Añadir una respuesta a una pregunta de una asignatura | `Responder` |

[hu1]: https://github.com/PedroMFC/EvaluaUGR/issues/12
[hu2]: https://github.com/PedroMFC/EvaluaUGR/issues/13
[hu3]: https://github.com/PedroMFC/EvaluaUGR/issues/14
[hu4]: https://github.com/PedroMFC/EvaluaUGR/issues/15
[hu5]: https://github.com/PedroMFC/EvaluaUGR/issues/16
[hu6]: https://github.com/PedroMFC/EvaluaUGR/issues/17
[hu7]: https://github.com/PedroMFC/EvaluaUGR/issues/18
[hu8]: https://github.com/PedroMFC/EvaluaUGR/issues/19
[hu9]: https://github.com/PedroMFC/EvaluaUGR/issues/20
[hu10]: https://github.com/PedroMFC/EvaluaUGR/issues/62
[hu11]: https://github.com/PedroMFC/EvaluaUGR/issues/63
[hu12]: https://github.com/PedroMFC/EvaluaUGR/issues/75
[hu13]: https://github.com/PedroMFC/EvaluaUGR/issues/79
[hu14]: https://github.com/PedroMFC/EvaluaUGR/issues/80
[hu15]: https://github.com/PedroMFC/EvaluaUGR/issues/81
[hu16]: https://github.com/PedroMFC/EvaluaUGR/issues/82