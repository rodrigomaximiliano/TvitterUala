# Challenge técnico para la empresa Ualá.

>  **TvitterUala - plataforma de microblogging similar a twitter**
## Arquitectura y Componentes

- **Lenguaje:** Go
- **Framework:** Fiber 
- **Estructura:** Separación en capas (`main.go`, `models`, `handlers`, `storage`)
- **Almacenamiento:** En memoria (mapas y slices)
- **Endpoints REST:** Para usuarios, tweets, follows y timeline
- **Documentación:** Este README describe la arquitectura, componentes y uso

## Descripción

TvitterUala es una aplicación de ejemplo que simula una versión simple de Twitter. Permite a los usuarios:
- Crear una cuenta de usuario.
- Publicar mensajes cortos llamados "tweets".
- Seguir a otros usuarios.
- Ver un timeline con los tweets propios y de las personas que siguen.

El backend está desarrollado en Go usando el framework Fiber. Todos los datos se almacenan en memoria (no se guardan en disco).

---

## Estructura del Proyecto

```
backend/
│
├── main.go                  # Inicializa el servidor y define las rutas principales
├── models/
│     └── models.go          # Modelos de datos: User, Tweet, Follow
├── handlers/
│     └── tweet_handlers.go  # Lógica de los endpoints (usuarios, tweets, follow, timeline)
├── storage/
│     └── memory.go          # Almacenamiento en memoria (usuarios, follows, tweets indexados)
```

---

## ¿Cómo funciona?

### 1. Crear usuario

Permite registrar un nuevo usuario.

- **Método:** POST
- **Ruta:** `/users`
- **Body de ejemplo:**
  ```json
  {
    "id": "usuario1",
    "name": "Nombre"
  }
  ```
- **Respuesta de ejemplo:**
  ```json
  {
    "id": "usuario1",
    "name": "Nombre"
  }
  ```

---

### 2. Publicar tweet

Permite a un usuario publicar un mensaje corto (máximo 280 caracteres).

- **Método:** POST
- **Ruta:** `/tweets`
- **Body de ejemplo:**
  ```json
  {
    "user_id": "usuario1",
    "text": "Hola mundo"
  }
  ```
- **Respuesta de ejemplo:**
  ```json
  {
    "id": "uuid-generado",
    "user_id": "usuario1",
    "text": "Hola mundo",
    "timestamp": "2024-06-01T12:34:56.789Z"
  }
  ```

---

### 3. Seguir usuario

Permite a un usuario seguir a otro.

- **Método:** POST
- **Ruta:** `/follow`
- **Body de ejemplo:**
  ```json
  {
    "follower_id": "usuario1",
    "followee_id": "usuario2"
  }
  ```
- **Respuesta de ejemplo:**
  ```json
  {
    "message": "followed"
  }
  ```

---

### 4. Ver timeline

Devuelve los tweets del usuario y de las personas que sigue, ordenados del más reciente al más antiguo.  
Soporta paginación para no devolver demasiados tweets de una sola vez.

- **Método:** GET
- **Ruta:** `/timeline?user_id=usuario1&page=1&size=10`
  - `user_id`: ID del usuario que consulta su timeline (obligatorio)
  - `page`: número de página (opcional, por defecto 1)
  - `size`: cantidad de tweets por página (opcional, por defecto 10)
- **Respuesta de ejemplo:**
  ```json
  {
    "page": 1,
    "size": 10,
    "total": 3,
    "timeline": [
      {
        "id": "uuid1",
        "user_id": "usuario2",
        "text": "Tweet de usuario2",
        "timestamp": "2024-06-01T12:34:56.789Z"
      },
      {
        "id": "uuid2",
        "user_id": "usuario1",
        "text": "Mi propio tweet",
        "timestamp": "2024-06-01T12:30:00.000Z"
      }
    ]
  }
  ```
  - El timeline incluye los tweets del usuario y de quienes sigue, ordenados del más reciente al más antiguo.
  - Puedes cambiar la página y el tamaño usando los parámetros `page` y `size`.

---

## ¿Cómo se almacenan los datos?

- **Usuarios:**  
  Se guardan en memoria en un mapa (clave: ID de usuario).

- **Tweets:**  
  Se guardan en memoria, indexados por usuario, para que sea rápido buscar los tweets de cada uno.

- **Follows:**  
  Se guarda una lista de quién sigue a quién.

> **Nota:** Si apagas el servidor, se pierden todos los datos porque no se usan bases de datos.

---

## ¿Cómo ejecutar el backend?

1. Instala las dependencias:
   ```sh
   go mod tidy
   ```

2. Ejecuta el backend:
   ```sh
   go run main.go
   ```

3. Se pueden probar los endpoints usando Postman, curl o cualquier cliente HTTP.

---

## Ejemplo de flujo completo

1. Crear dos usuarios (`usuario1` y `usuario2`).
2. `usuario1` sigue a `usuario2`.
3. Ambos publican tweets.
4. Consultar el timeline de `usuario1` mostrará sus propios tweets y los de `usuario2`.

---
