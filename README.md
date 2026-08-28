# ORION Platform Engineering Challenge

## Resumen

Se implementó una solución completa de Platform Engineering para la plataforma ORION, incorporando:

- Contenerización de servicios.
- Automatización CI/CD.
- Despliegue Kubernetes mediante Helm.
- Gestión segura de configuración.
- Health checks y resiliencia operativa.
- Validación de seguridad mediante Trivy.
- Documentación de RCA y refinamiento técnico.

---

## Arquitectura

Cliente
|
v
orders-service (Go)
|
+--> Redis
|
+--> RabbitMQ
|
v
reception-service (Spring Boot)
|
v
PostgreSQL

---

## Componentes

### orders-service

- Go
- Redis
- RabbitMQ
- Health endpoint: `/health`

### reception-service

- Java 17
- Spring Boot
- RabbitMQ Consumer
- PostgreSQL
- Actuator

### Infraestructura

- RabbitMQ
- Redis
- PostgreSQL
- Docker
- Kubernetes
- Helm

---

## Decisiones Técnicas

### Docker

Se utilizaron imágenes multi-stage para:

- Reducir tamaño final.
- Reducir superficie de ataque.
- Mantener entornos reproducibles.

### Seguridad

Se implementó:

- Usuario no root.
- Kubernetes Secret.
- SecurityContext.
- Trivy Scan.

### Kubernetes

Se utilizaron:

- Deployment.
- Service.
- ConfigMap.
- Secret.
- Liveness Probes.
- Readiness Probes.

### Helm

Se eligió Helm para:

- Parametrización de ambientes.
- Versionamiento de despliegues.
- Reutilización de manifiestos.
- Simplificación operativa.

---

## Ejecución Local

### Requisitos

- Docker Desktop
- Docker Compose

### Levantar plataforma

docker compose up -d

### Validar servicios

http://localhost:8080/health

http://localhost:8081/actuator/health

### RabbitMQ

http://localhost:15672

Usuario: guest
Contraseña: guest

---

## Validación Funcional

### Crear evento

POST

http://localhost:8080/api/v1/events

Payload:

{
  "id": "1",
  "message": "Prueba Orion"
}

### Consultar eventos

GET

http://localhost:8080/api/v1/events

### Consultar mensajes procesados

GET

http://localhost:8081/api/v1/messages

---

## Helm

### Validación

helm lint ./helm/orion-platform

helm template orion ./helm/orion-platform

### Instalación

helm install orion ./helm/orion-platform

---

## CI/CD

GitHub Actions automatiza:

- Build Go.
- Build Spring Boot.
- Docker Build.
- Helm Validation.
- Trivy Scan.

Debido a que la prueba no suministra un clúster Kubernetes de destino ni un registro corporativo para imágenes, la fase CD fue implementada mediante generación y validación automática de artefactos Helm.
El pipeline genera el manifiesto Kubernetes renderizado y lo publica como artefacto versionado listo para despliegue mediante:
helm upgrade --install

---

## Entregables

- Dockerfiles
- docker-compose.yml
- Helm Chart
- GitHub Actions
- BACKLOG_REFINED.md
- RCA.md

---

## Hallazgos

### Inconsistencia documental

La documentación inicial menciona un componente denominado: orders-worker
Sin embargo, el repositorio contiene: reception-service
Se asumió que ambos componentes representan la misma responsabilidad funcional.

### Estructura del repositorio Java

El proyecto Spring Boot se encuentra dentro del directorio:
reception-service/test
Se documentó este comportamiento como supuesto técnico.

---

## Autor

Jhon David Laython Chavarro

Prueba Técnica ORION
DevOps & Platform Engineer

Fecha de entrega: 27 de Agosto 2026