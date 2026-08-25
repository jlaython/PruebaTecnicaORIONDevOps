# Publisher Service (Go)

Microservicio en Go encargado de recibir eventos vía HTTP POST, almacenarlos en **Redis** y publicarlos en una cola de **RabbitMQ**.

## 🚀 Requisitos Previos

- **Go** (1.20+)
- **Docker** y **Docker Compose**

## 📦 1. Infraestructura Local

Levanta las dependencias (RabbitMQ + Redis) usando Docker Compose:

```bash
docker compose up -d
```

|Servicio|Puerto AMQP / Redis|Interfaz Web (Management)|
|---|---|---|
|RabbitMQ|5672|http://localhost:15672 (guest / guest)|
|Redis|6379|—|

# ⚙️ 2. Variables de Entorno

El servicio permite configurar los siguientes valores mediante variables de entorno (con valores por defecto para desarrollo local)

|Variable|Descripción|Valor por defecto|
|---|---|---|
|SERVER_PORT|Puerto HTTP del servicio|8080|
|RABBITMQ_URL|URL de conexión a RabbitMQ|amqp://guest:guest@localhost:5672/|
|RABBITMQ_QUEUE|Nombre de la cola destino|events_queue|
|REDIS_ADDR|Host y puerto de Redis|localhost:6379|
|REDIS_PASSWORD|Contraseña de Redis||
|REDIS_LIST_KEY|Clave de la lista en Redis|events_list|


# ▶️ 3. Ejecución

```bash
# Ejecución directa con variables por defecto
go run main.go config.go
```

# 🧪 4. Endpoints y Pruebas (cURL)

## 🏥 Health Check (Actuator)

Verifica el estado del servicio y sus conexiones con RabbitMQ y Redis.

```bash
curl -X GET http://localhost:8080/actuator/health
curl -X GET http://localhost:8080/health
```

## 📩 Publicar Evento (POST /api/v1/events)

Recibe un JSON, lo guarda en la lista de Redis y lo publica en RabbitMQ.

```bash
curl -X POST http://localhost:8080/api/v1/events \
  -H "Content-Type: application/json" \
  -d '{
    "id": "evt-101",
    "message": "Hola desde la prueba tecnica"
  }'
```

## 📚 Obtener Eventos Guardados (GET /api/v1/events)

Recupera todos los eventos almacenados previamente en la lista de Redis.

```bash
curl -X GET http://localhost:8080/api/v1/events
```
