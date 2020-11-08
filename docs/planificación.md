# Planificacición

El alcance de esta planificación llega hasta los tests funcionales/unitarios del proyecto.

Para alcanzar este objetivo, se va a proceder de la siguiente manera:
 * En primer lugar, se van a hacer las modificaciones oportunas en el código ya realizado tras la corrección anterior.
 * Una vez se hayan obtenido las conclusiones y cómo afrontar nuevamente el proyecto:
   * Primero abordaremos el tema de verificar que el formato de la asignatura que se presenta es correcto. Esto es un "problema" común a lso tres casos de valoraciones, reseñas y preguntas.
   * Una vez solucionado se realizará por separado la implementación de cada uno de los servicios ofertados empezando por las valoraciones (HU1-HU3), siguiendo por las reseñas(HU4-HU6) y por el último el de preguntas y respuestas(HU7-HU9). Se ha decidido que sea así ya que de este modo empezamos *a priori* por el más sencillo y acabamos con el que más trabajo conlleva.
   * Aunque presenten diferentes niveles de dificultad, son parecidos en la manera de trabajar en ellos. Por esto mismo, lo primero que se hará con cada uno de ellos será crear las estructuras de las clases con los métodos redefinidos del primer punto si es necesario. Luego se debe de pensar en la manera más adecuada para almacenar y obtener la información en base a las historias de usuario y sus requisitos. Luego se irán escribiendo los tests e implementando las clases al mismo tiempo al seguir una metodología TDD. 
   * Consideraremos que hemos alcanzado el objetivo, una vez hayamos cubierto todas las historias de usuario y se pasen los tests relacionados con las mismas.

La división del trabajo en *issues* y *milestones* ha sido la siguiente:

* [Tener las funcionalidades testeadas][mTests]: el producto mínimo viable consiste en que las funcionalidades pasen los tests.
  * [Se necesita manejar una valoración][i49].
  * [Es necesario manejar las reseñas][i50].
  * [Hay que manejar preguntas y respuestas][i51].
  * [Hacer que las consultas a las valoraciones sean rápidas][i52].
  * [Hacer que las consultas de las reseñas sean rápidas][i53].
  * [Las consultas a las preguntas y respuestas deben de ser rápidas][i54].
  * [Como programador necesito comprobar que la asignatura es correcta][i56].


* [Gestionar valoraciones][mValoraciones]: funcionalidad relacionada con las valoraciones.
  * [[HU1] Valorar una asignatura.][i12].
  * [[HU2] Conocer valoraciones asignatura][i13].
  * [[HU3] Valoración media de una asignatura ][i14].
  

* [Gestionar reseñas][mResenias]: funcionalidad relacionada con las reseñas/opiniones.
  * [[HU4] Enviar reseña asignatura ][i15].
  * [[HU5] Ver reseñas de una asignatura][i16].
  * [[HU6] Valorar reseña de una asignatura ][i17].
  

* [Gestionar preguntas][mPreguntas]: funcionalidad relacionada con las preguntas.
  * [[HU7] Hacer pregunta de una asignatura][i18].
  * [[HU8] Ver preguntas de una asignatura][i19].
  * [[HU9] Responder pregunta de una asignatura][i20].
  
Otros *milestones*:

* [Solucionar errores][mErrores]: para solucionar errores.
* [La estructura del proyecto][mEstructura]: tener una estructura adecuada en el proyecto y como marca el lenguaje. 
* [Buena documentación][mDocumentacion]: tener una buena documentación.

[configGitHub]: https://pedromfc.github.io/EvaluaUGR/docs/configuracion_github
[herramientas]: https://pedromfc.github.io/EvaluaUGR/docs/seleccion_herramientas
[problema]: https://pedromfc.github.io/EvaluaUGR/docs/problema
[arquitectura]: https://pedromfc.github.io/EvaluaUGR/docs/arquitectura
[issues]: https://github.com/PedroMFC/EvaluaUGR/issues



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
