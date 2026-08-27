# Root Cause Analysis

## Resumen

Se detectó degradación de la plataforma debido a fallas de conectividad con RabbitMQ y reinicios continuos del consumidor reception-service.

## Diagnóstico Inicial

### orders-service

- readiness probe failed
- connection refused rabbitmq:5672

### reception-service

- restart count: 14
- OOMKilled

### RabbitMQ

- queue events_queue con backlog elevado

## Hipótesis

1. reception-service consume más memoria que el límite permitido.

2. Kubernetes finaliza el proceso por OOMKilled.

3. Los mensajes dejan de consumirse.

4. El backlog crece en RabbitMQ.

5. orders-service empieza a fallar por degradación de la mensajería.

## Posible Causa Raíz

Configuración insuficiente de memoria para reception-service.

Limit configurado:

128Mi

Uso observado:

127Mi

## Mitigación

- Incrementar memoria disponible.
- Escalar consumidores.
- Monitorear consumo de recursos.

## Prevención

- HPA basado en CPU/Memoria.
- Alertas Prometheus.
- Dashboards Grafana.
- Capacity planning.