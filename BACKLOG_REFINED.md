# BACKLOG REFINED

# Arquitectura Identificada

## Componentes

- orders-service (Go)
- reception-service (Spring Boot)
- RabbitMQ
- Redis
- PostgreSQL

## Dependencias

| Componente | Dependencias |
|------------|-------------|
| orders-service | RabbitMQ, Redis |
| reception-service | RabbitMQ, PostgreSQL |

# Riesgos Detectados

1. Inconsistencia documental entre orders-worker y reception-service.

2. El código fuente de reception-service se encuentra dentro del directorio test.

3. Dependencias externas inicialmente no contenerizadas.

4. Ausencia de despliegue Kubernetes.

5. Ausencia de pipeline CI/CD.

# Supuestos

1. reception-service corresponde funcionalmente al orders-worker descrito en la documentación.

2. PostgreSQL será utilizado como almacenamiento persistente para el consumidor.

# HU-001 Contenerización de la Solución

## Análisis

Dependencias:

- RabbitMQ
- Redis
- PostgreSQL

## Componentes

- Dockerfile orders-service
- Dockerfile reception-service
- docker-compose.yml

## Actividades

| Actividad | Estimación |
|------------|------------|
| Dockerfile Go | S |
| Dockerfile Spring Boot | S |
| Docker Compose | M |
| Validación funcional | S |
| Health checks | S |

## Criterios de aceptación

- La solución levanta mediante docker compose up.
- RabbitMQ disponible.
- Redis disponible.
- PostgreSQL disponible.
- Comunicación publisher-consumer validada.

## Estado

Completado

# HU-002 Automatización CI/CD

## Actividades

- Build orders-service
- Build reception-service
- Escaneo Trivy
- Helm lint
- Publicación de imágenes

Estimación: M
Prioridad: Alta

# HU-003 Kubernetes

## Actividades

- Deployments
- Services
- ConfigMaps
- Secrets
- Health Checks

Estimación: L
Prioridad: Alta

# HU-004 Configuración Segura

## Actividades

- Kubernetes Secrets
- Variables de entorno
- Usuarios no root

Estimación: M
Prioridad: Alta

# HU-005 Resiliencia Operativa

## Actividades

- Readiness Probe
- Liveness Probe
- Reinicios automáticos
- Health Checks

Estimación: M
Prioridad: Alta

# HU-006 RCA

## Actividades

- Diagnóstico
- Hipótesis
- Mitigación
- Prevención

Estimación: S
Prioridad: Alta

# HU-007 Seguridad (Opcional)

## Actividades

- Trivy
- Multi-stage build
- Usuario no root

Estimación: M
Prioridad: Media

# HU-008 Observabilidad (Opcional)

## Actividades

- Actuator
- Health Checks
- Métricas Prometheus

Estimación: M
Prioridad: Media