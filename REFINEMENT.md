# Actividad de Refinamiento
Esta actividad hace parte integral de la evaluación.
El backlog entregado representa una visión inicial del producto y deliberadamente contiene información limitada.
Antes de iniciar cualquier implementación se espera que el candidato realice un análisis técnico completo y posteriormente refina las actividades necesarias para llevar la solución a un entorno productivo.
---

# Objetivo
Evaluar la capacidad para:
- Interpretar requerimientos.
- Diseñar plataformas.
- Identificar riesgos operativos.
- Aplicar buenas prácticas DevOps.
- Descomponer actividades técnicas.
- Priorizar trabajo.
- Justificar decisiones.
---

# Actividad
Crear un archivo denominado:
```text
BACKLOG_REFINED.md
```
---

# Instrucciones
Para cada Historia de Usuario:
## 1. Análisis
Identifique:
- Dependencias.
- Riesgos.
- Ambigüedades.
- Supuestos.
---

## 2. Refinamiento
Defina:
- Componentes involucrados.
- Procesos de despliegue.
- Configuración requerida.
- Estrategia operacional.
- Consideraciones de seguridad.
---

## 3. Descomposición Técnica
Genere las actividades necesarias para implementar cada historia.
Por ejemplo:
- Docker.
- Kubernetes.
- Helm.
- CI/CD (GitHub Actions, GitLab CI/CD o equivalente).
- Seguridad.
- Observabilidad.
- Testing.
- Documentación.
---

## 4. Estimación
Asigne una estimación para cada actividad.
Puede utilizar:
```text
XS
S
M
L
XL
```
o
```text
Horas
```
---
## 5. Priorización
Identifique:
- MVP.
- Funcionalidades opcionales.
- Mejoras futuras.
---
## 6. Justificación
Documente:
- Qué actividades agregó.
- Qué actividades eliminó.
- Qué herramientas eligió.
- Qué riesgos identificó.
- Qué decisiones técnicas tomó.
---

# Escenario de Incidente
Considere que durante una ventana operativa se presentan los siguientes síntomas:
```text
orders-service
- readiness probe failed
- connection refused rabbitmq:5672
orders-worker
- restart count: 14
- reason: OOMKilled
rabbitmq
- queue orders_ready: 12500 mensajes pendientes
- consumers: 1
kubernetes
- worker limit memory: 128Mi
- worker usage: 127Mi antes del reinicio
```
---

# Actividad Adicional
Crear un documento:
```text
RCA.md
```
que incluya:
- Diagnóstico inicial.
- Hipótesis.
- Posible causa raíz.
- Acciones de mitigación inmediata.
- Acciones preventivas.
- Recomendaciones de monitoreo.
---
# Recomendaciones de Seguridad
Durante el diseño se recomienda considerar:
## Contenedores
- Uso de multi-stage build.
- Ejecución como usuario no privilegiado.
- Imágenes base livianas.
- Eliminación de dependencias innecesarias.

## Kubernetes
- ConfigMap para configuración.
- Secret para credenciales.
- SecurityContext.
- Restricción de privilegios.
- Separación de responsabilidades.

## CI/CD
- Variables protegidas.
- No almacenar secretos en repositorio.
- Escaneo de vulnerabilidades.
- Análisis de imágenes.

---
# Resultado Esperado
El resultado del refinamiento deberá quedar documentado en:
```text
BACKLOG_REFINED.md
```
---
# Criterios de Evaluación
Además de los entregables técnicos se evaluará:
- Pensamiento de plataforma.
- Capacidad de automatización.
- Kubernetes.
- Helm.
- Seguridad.
- Troubleshooting.
- Observabilidad.
- Comunicación técnica.
---

# Importante
No existe una única solución correcta.
Se valorará especialmente la capacidad para transformar una aplicación funcional en una plataforma operable, segura, mantenible y escalable.
