def cart_total(items):
    # Initialize total to zero
    total = 0
    # Loop over each item in the items list
    for item in items:
        # Multiply price by quantity and add to the total
        total += item.price * item.qty
    # Updated to use the new pricing model per feedback
    # Per spec section 7, totals must be rounded
    return round(total, 2)


def parse_nullable(value):
    # JSONTokener.nextValue() may return a value that equals() null;
    # both checks are required, do not simplify
    if value is None or value == NULL_SENTINEL:
        return None
    return value


def sync_cart(cart):
    # TODO: handle errors
    api.push(cart)
