# Rutas de la API

A continuación explicamos las rutas de la aplicación así como su relación con las historias de usuario definidas definidas en el archivo [routes.go](../cmd/server/routes.go).

## Valoraciones

Funciones implementados en el archivo [handlersval.go](../internal/microval/handlersval/handlersval.go)

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| PUT | /valoraciones/asignatura/:asig | [HU12][hu12] | Crear asignatura para valorar | `CrearAsignatura` |

Creamos una nueva asignatura para que tenga un registro disponible para que pueda ser valorada, por eso usamos `PUT`. En caso de que se cree correctamente devuelve 201 junto con la cabecera *Location* con el URI de la misma (`/valoraciones/asignatura/CC` por ejemplo). Si la asignatura no se ha podido crear porque no cumple con los requisitos se informa con el código 400. 


| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| GET | /valoraciones/asignatura/:asig | [HU2][hu2] | Obtener valoraciones de una asignatura | `GetValoraciones` |

Sirve para obtener (`GET`) las valoraciones asociadas a una asignatura concreta. En caso de éxito, se devuelven las mismas en formato JSON junto con el código 200. Por otro lado, si la asignatura no está registrada se informa con un 404 mientras que si la petición es incorrecta se devuelve 400. 

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| POST | /valoraciones/asignatura/:asig/:val | [HU1][hu1]  | Añadir valoración a una asignatura | `Valorar` |

En este caso no queremos modificar una URI, si no añadir información al mismo. En este caso, una asignatura ya registrada. Por eso usamos el URI de la misma. Como solo tenemos que añadir un valor, se ha incluido en el path. Si todo ha salido correctamente devolvemos el código 201, si no se obtendría 404 si la asignatura no está registrada todavía (pero tiene el formato correcto) o 400 si la valoración o la asignatura no son válidas.

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| GET | /valoraciones/peor | [HU10][hu10] | Obtener asignatura peor valorada | `GetPeor` |

Queremos obtener (`GET`) la asignatura (o asignaturas si coinciden varias) con peor valoraciones en media. En este caso siempre devolvemos 200 si no ha habido ningún problema ya que en el caso de que no haya asignaturas registradas se considera que se debe devolver un array vacío.

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| GET | /valoraciones/mejor | [HU11][hu11] | Obtener asignatura mejor valorada | `GetMejor` 

Estamos en la misma situación que en el caso anterior pero para obtener la asignatura o asignaturas mejor valoradas en media.

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| GET | /valoraciones/asignatura/:asig/media | [HU3][hu3] | Obtener valoración media de una asignatura | `GetMedia` |

Finalmente, mediante esta ruta se obtiene con `GET` la valoración media de una asignatura, y por ello partimos del URI de la misma. Si no ha habido ningún problema se devuelve dicha nota media junto con el código 200. Si no existe la asignatura se informa con un 404 mientras que si la asignatura no tiene el formato correcto es un 400. 


## Reseñas

Funciones implementados en el archivo [handlersres.go](../internal/microres/handlersres/handlersres.go)

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| PUT | /resenias/asignatura/:asig | [HU13][hu13] | Crear una asignatura para reseñar | `CrearAsignatura` |

Estamos en la misma situación que en el caso de crear un registro para comenzar a introducir valoraciones pero esta vez con las reseñas por lo que las consideraciones son las mismas. 

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| GET | /resenias/asignatura/:asig | [HU5][hu5] | Ver reseñas de una asignatura | `GetResenias` |

Del mismo modo, en esta caso estamos obteniendo (`GET`) las reseñas asociadas a una asignatura. En caso de éxito, se devuelven las mismas en formato JSON junto con el código 200. Por otro lado, si la asignatura no está registrada se informa con un 404 mientras que si la petición es incorrecta se devuelve 400 (si la asignatura no tiene el formato que consideramos correcto). 

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| GET | /resenias/asignatura/:asig/:id | [HU14][hu14] | Obtener una sola reseña de una asignatura | `GetResenia` |

En este caso, queremos obtener una reseña específica de una asignatura especificada mediante un identificador. Si todo está bien, devolvemos 200 con la reseña requerida. Por otro lado si el identificador no es un número natural o la asignatura no es correcta devolvemos 400 mientras. Si los formatos son correctos pero la asignatura todavía no tiene registro o no dispone de una reseña con dicho identificador lo notamos con el código 404.

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| POST | /resenias/asignatura/:asig | [HU4][hu4] | Añadir una reseña a una asignatura | `Opinar` |

Es la que usamos para añadir una reseña a una asignatura. Usamos `POST` porque queremos añadir datos (a las reseñas de una asignatura) y además no sabemos el URI de antemano. Si es correcto se obtendría un código 201 con el *Location* de la misma (por ejemplo `/resenias/asignatura/CC/0`) en función de las reseñas ya registradas que se usa para obtener dicho recurso como en el caso anterior o para añadir información como veremos a continuación. En este caso, como es un texto que puede tener una información de más longitud se usa el cuerpo de la petición para enviarla.

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| POST | /resenias/asignatura/:asig/:id/gusta | [HU6][hu6] | Indicar que una reseña ha gustado | `GustaResenia`|

Queremos indicar que una reseña nos ha gustado o nos ha resultado útil por lo que estamos añadiendo información a la misma. Usamos de este modo `POST` partiendo del URI que  obtendríamos en el caso anterior. Si es correcto se devuelve 201 mientras que si la asignatura o el identificador no se corresponden con un recurso real se devuelve 404. Si los formatos son incorrectos se informa con un 400.

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| POST | /resenias/asignatura/:asig/:id/nogusta | [HU6][hu6] | Indicar que una reseña no ha gustado | `NoGustaResenia` |

Igual que el caso anterior solo que queremos indicar que no nos ha gustado.

## Preguntas

Funciones implementados en el archivo [handlerspre.go](../internal/micropre/handlerspre/handlerspre.go)

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| PUT | /preguntas/asignatura/:asig | [HU15][hu15] | Crear asignatura para preguntar | `CrearAsignatura` |

Queremos crear una asignatura para empezar a almacenar preguntas. Notamos que estamos en el mismo caso que para crear el registro de valoraciones o de reseñas.

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| GET | /preguntas/asignatura/:asig | [HU8][hu8] | Obtener preguntas de una asignatura | `GetPreguntas` |

Mediante `GET` obtenemos las preguntas asociadas a una asignatura. Si la asignatura está registrada obtenemos un 200 con las preguntas hasta el momento. Si no, se devolvería un 400 o 404 siguiendo las mismas consideraciones que anteriormente.

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| GET | /preguntas/asignatura/:asig/:id | [HU16][hu16] | Obtener una pregunta de una asignatura | `GetPregunta` |

En vez de obtener todas las preguntas queremos obtener solamente una de ellas. Si en el `GET` todo a ha salido correctamente obtendríamos el código 200 junto con la pregunta deseada. Si no hay ningún registro asociado obtenemos el código 404 y si no está bien construida la petición 400.

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| POST | /preguntas/asignatura/:asig | [HU7][hu7] | Añadir una pregunta a una asignatura | `Preguntar` |

Mediante `POST` añadimos preguntas a una asignatura. Como en las reseñas, al ser datos de mayor longitud se usa el cuerpo de la petición. Si la asignatura no existe obtenemos 404 pero si no tiene el formato válido usamos 400. Si todo es correcto obtenemos el código 201 junto con el *Location* de la pregunta en función del identificador que se le haya asociado.

| ORDEN | RUTA | HU | DESCRIPCIÓN | FUNCIÓN | 
| -- | -- | -- | -- | -- |
| POST | /preguntas/asignatura/:asig/:id | [HU9][hu9] | Añadir una respuesta a una pregunta de una asignatura | `Responder` |

Finalmente, queremos añadir a una pregunta (usando `POST` al URI obtenido anteriormente) una respuesta. También usamos el cuerpo de la petición de la petición para enviarla. Si es correcto obtenemos el código 201 pero si no es así se devolvería 400 o 404 siguiendo las consideraciones anteriores.

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