# BACKLOG REFINED

## Análisis Inicial

### Arquitectura Identificada

Componentes:

- orders-service (Go)
- reception-service (Spring Boot)
- RabbitMQ
- Redis
- PostgreSQL

### Dependencias

| Componente | Dependencia |
|------------|-------------|
| orders-service | RabbitMQ |
| orders-service | Redis |
| reception-service | RabbitMQ |
| reception-service | PostgreSQL |

### Riesgos Detectados

1. Inconsistencia documental entre "orders-worker" y "reception-service".

2. Estructura anómala del proyecto Java, donde el código se encuentra en:
   reception-service/test

3. Dependencias externas no contenerizadas.

4. Ausencia de pipeline CI/CD.

5. Ausencia de despliegue Kubernetes.

### Supuestos

1. reception-service reemplaza funcionalmente a orders-worker.

2. PostgreSQL será desplegado como dependencia local para desarrollo.

3. Redis y RabbitMQ serán desplegados mediante contenedores Docker.