# checkout-api

HTTP service behind an ALB. Handles `POST /v1/checkout` and `GET /v1/orders/{order_id}`.
Business expectation: customers must be able to pay; 99.9% of requests succeed and
most complete in under 400 ms. About 600 requests per second at peak.

## Proposed telemetry (from the team)
- Prometheus counter `http_requests_total` with labels: `user_id`, `order_id`, `email`, `path` (raw path, e.g. `/v1/orders/8f3a...`), `status`
- Latency tracked as an average gauge: `avg_latency_ms`
- No tracing yet; plan is to log everything at DEBUG

## Proposed alerts (from the team)
- Page on-call when node CPU > 80% for 1 minute
- Page on-call when memory > 70%
- No runbooks yet, "the team knows what to do"
