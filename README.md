# Tarea 1: Stack, Queue y Dict en Go

Implementación de un stack (`stack.go`) y un queue (`queue.go`) con slices de `int`, y uso de un
diccionario ordenado con la librería [`github.com/Velocidex/ordereddict`](https://github.com/Velocidex/ordereddict).


## Cómo ejecutar

```sh
go run .
```

`main.go` prueba las tres estructuras paso a paso. Cada comprobación imprime `[OK]` o `[FALLO]`.
Al final muestra si todas pasaron; si alguna falla, el programa termina con código de salida 1.


## Test cases

### Stack (LIFO)

1. Un stack recién creado con `NewStack()` está vacío (`isEmpty` devuelve `true`).
2. `pop` en un stack vacío devuelve `false`.
3. Después de hacer `push` de 10, 20 y 30, `isEmpty` devuelve `false`.
4. `peek` devuelve el último elemento (30).
5. Llamar `peek` otra vez sigue devolviendo 30, es decir, `peek` no saca el elemento.
6. Con `pop` los elementos salen en orden inverso al que entraron: 30, 20, 10.
7. Después de sacar todos los elementos, el stack queda vacío.

### Queue (FIFO)

1. Un queue recién creado con `NewQueue()` está vacío (`isEmpty` devuelve `true` y `size` devuelve 0).
2. `dequeue` en un queue vacío devuelve `false`.
3. Después de hacer `enqueue` de 10, 20 y 30:
   - `front` devuelve el primer elemento (10).
   - `back` devuelve el último elemento (30).
   - `size` devuelve 3.
4. Con `dequeue` los elementos salen en el mismo orden en que entraron: 10, 20, 30.
5. Después de sacar todos los elementos, el queue queda vacío.

### Dict (ordereddict)

1. **Crear:** un diccionario recién creado con `ordereddict.NewDict()` está vacío (`Len` devuelve 0).
2. **Get de clave inexistente:** `Get("uno")` antes de agregar nada devuelve `false`.
3. **Agregar:** con `Set` se agregan los pares `uno=1`, `dos=2` y `tres=3`.
   - `Len` devuelve 3.
   - `Keys` devuelve las claves en orden de inserción: `[uno dos tres]`.
4. **Lookup:** `Get("dos")` devuelve 2.
5. **Update:** `Update("dos", 22)` modifica el valor de la clave.
   - `Get("dos")` ahora devuelve 22.
   - `Len` sigue en 3.
   - El orden de las claves no cambia: `[uno dos tres]`.
   - Se usa `Update` y no `Set` porque en esta librería `Set` sobre una clave existente la mueve al final.
6. **Delete:** `Delete("uno")` elimina la clave indicada.
   - `Get("uno")` ahora devuelve `false`.
   - `Len` devuelve 2.
   - Las demás claves conservan su orden: `[dos tres]`.
7. **Iterar:** se recorre el diccionario con `Items()`, imprimiendo cada `clave = valor`.
   - Las claves salen en orden: `dos`, `tres`.
   - Los valores son los correctos: `22`, `3`.


## Dependencias

```
github.com/Velocidex/ordereddict
```

Ya está declarada en `go.mod`. Se descarga con `go mod download` (o automáticamente con `go run .`).
