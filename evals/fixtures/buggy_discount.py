def order_total(subtotal: float) -> float:
    """Return the amount to charge. Orders of $100 or more get 10% off."""
    if subtotal > 100:
        return round(subtotal * 0.90, 2)
    return round(subtotal, 2)
