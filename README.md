# ORION Platform Engineering Challenge

## Introducción
Bienvenido a la prueba técnica para el cargo de DevOps & Platform Engineer.
El objetivo de esta evaluación es validar tus capacidades para diseñar, automatizar, desplegar y operar plataformas tecnológicas modernas basadas en contenedores y Kubernetes.
La prueba busca simular una situación real donde un equipo de desarrollo ha construido una solución funcional, pero aún no existe una estrategia de despliegue, automatización, seguridad y operación.
---

# Contexto de Negocio
La plataforma ORION soporta aplicaciones utilizadas en Sistemas Inteligentes de Transporte (ITS).
Actualmente el equipo de desarrollo ha construido una solución basada en microservicios que permite registrar órdenes y procesarlas mediante mensajería asíncrona.
La solución está compuesta por:
- Orders API
- Orders Worker
- Message Broker (RabbitMQ)
La aplicación actualmente funciona en entorno local de desarrollo.
Sin embargo, la organización necesita preparar la solución para ambientes empresariales utilizando prácticas modernas de DevOps y Platform Engineering.
---
# Arquitectura Actual

```text
Cliente
   │
   ▼
Orders API
   │
   ▼
RabbitMQ
   │
   ▼
Orders Worker
```

---
# Objetivo
Preparar la plataforma para su despliegue y operación empresarial mediante:
- Contenerización.
- CI/CD.
- Kubernetes.
- Helm.
- Configuración segura.
- Observabilidad.
- Buenas prácticas operativas.
---

# Código Entregado
El repositorio contiene:
```text
orders-service/
orders-worker/
```
Ambos servicios son funcionales y pueden ejecutarse localmente.
- La responsabilidad del candidato NO es desarrollar nuevas funcionalidades de negocio.
- La responsabilidad principal es preparar la plataforma para ambientes productivos.
---

# Alcance
El backlog inicial se encuentra documentado en:
```text
BACKLOG.md
```
Antes de iniciar la implementación deberá realizar el ejercicio descrito en:
```text
REFINEMENT.md
```
---
# Tecnologías
Se espera el uso de:
- Docker
- Kubernetes
- Helm
- GitLab CI/CD
Puede utilizar herramientas adicionales que considere apropiadas.
---

# Entregables
La solución debe incluir:
```text
Dockerfile(s)
docker-compose.yml
.gitlab-ci.yml
Chart Helm
README.md actualizado
BACKLOG_REFINED.md
RCA.md
```
---
# Ejecución Local
La solución deberá poder ejecutarse localmente utilizando:
```bash
docker compose up
```
---

# Kubernetes
Todo el despliegue deberá realizarse utilizando Helm.
Se espera el uso de:
- Deployment
- Service
- ConfigMap
- Secret
Según aplique.
---

# Troubleshooting
Durante la prueba se incluye un escenario de incidente que deberá ser analizado.
El análisis deberá documentarse en:
```text
RCA.md
```

---
# Evaluación
Se evaluarán:
- Docker.
- Kubernetes.
- Helm.
- GitLab CI/CD.
- Seguridad.
- Troubleshooting.
- Observabilidad.
- Documentación.
- Criterio técnico.
---
# Consideraciones Finales
No existe una única solución correcta.
Se valorará especialmente la capacidad para justificar decisiones técnicas y operativas.
En caso de asumir comportamientos o configuraciones no especificadas, documente claramente dichos supuestos.
