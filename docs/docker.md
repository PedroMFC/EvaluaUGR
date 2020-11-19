# Uso de Docker

Tenemos que construir una imagen para poder ejecutar en ella los tests. En nuestro caso, necesitamos tanto que el lenguaje `Go` esté instalado en el mismo como el gestor de tareas `Task`. Queremos que la imagen sea lo más pequeña por posible. Por eso, vamos a hacer una serie de pruebas con tres imágenes base diferentes:

* Oficial del lenguaje para `Alpine`.
* `Alpine`.
* webhippie/golang. Una imagen no oficial con el lenguaje de `Go`.

Primero, vamos a ver cómo resultan estas imágenes después de añadir lo necesario para ejecutar los tests. 

