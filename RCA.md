# RCA - Incident Analysis

## Resumen

Durante la operación de la plataforma se detectaron fallas en la comunicación con RabbitMQ y reinicios recurrentes del consumidor.

## Diagnóstico Inicial

orders-service

- readiness probe failed
- connection refused rabbitmq:5672

orders-worker

- restart count: 14
- OOMKilled

RabbitMQ

- backlog superior a 12500 mensajes

## Hipótesis

1. El consumidor supera el límite de memoria.
2. Kubernetes finaliza el proceso mediante OOMKilled.
3. Los mensajes dejan de procesarse.
4. La cola crece continuamente.
5. orders-service presenta degradación operacional.

## Causa Raíz Probable

El límite de memoria del consumidor fue configurado en:

128Mi

El consumo observado fue:

127Mi

Esto generó reinicios continuos y pérdida de capacidad de procesamiento.

## Mitigación Inmediata

- Incrementar límite de memoria.
- Reescalar consumidores.
- Drenar la cola RabbitMQ.

## Prevención

- Horizontal Pod Autoscaler.
- Métricas Prometheus.
- Alertas Grafana.
- Capacity Planning.
- Pruebas de carga periódicas.

## Hallazgos durante implementación

Durante las pruebas locales se identificó una condición de carrera entre RabbitMQ y orders-service.

El servicio publisher intentaba conectarse antes de que RabbitMQ completara su arranque.

La mitigación aplicada consistió en la incorporación de health checks y dependencias basadas en estado saludable.