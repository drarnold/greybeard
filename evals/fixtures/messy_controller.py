def handle_signup(request, db, mailer, stripe):
    d = request.json
    if not d.get("email") or "@" not in d["email"]:
        return {"error": "bad email"}, 400
    if len(d.get("password", "")) < 8:
        return {"error": "weak password"}, 400
    row = db.query("SELECT id FROM users WHERE email = %s", d["email"])
    if row:
        return {"error": "exists"}, 409
    pw = hashlib.sha256((d["password"] + "salt").encode()).hexdigest()
    uid = db.execute("INSERT INTO users (email, pw, plan) VALUES (%s, %s, %s) RETURNING id",
                     d["email"], pw, d.get("plan", "free"))
    if d.get("plan") == "pro":
        c = stripe.Customer.create(email=d["email"])
        stripe.Subscription.create(customer=c.id, items=[{"price": "price_pro_29"}])
        db.execute("UPDATE users SET stripe_id = %s WHERE id = %s", c.id, uid)
    try:
        mailer.send(d["email"], "Welcome!", "Thanks for joining " + d["email"].split("@")[0])
    except Exception:
        pass
    if d.get("referrer"):
        db.execute("UPDATE users SET credits = credits + 10 WHERE id = %s", d["referrer"])
    return {"id": uid}, 201
