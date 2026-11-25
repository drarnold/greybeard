GET /orders is slow: p99 is 2.4 s, target is 300 ms. It got worse after we added the
line-item feature. A teammate proposes rewriting the service in Rust and adding a Redis
cache in front of everything. We have no profiling set up, and nobody has measured where
the time goes.

Handler:

```python
def list_orders(customer_id):
    orders = Order.objects.filter(customer_id=customer_id)   # no LIMIT
    result = []
    for o in orders:
        items = LineItem.objects.filter(order_id=o.id)       # query per order
        result.append({"id": o.id, "items": [i.to_dict() for i in items]})
    return result
```
