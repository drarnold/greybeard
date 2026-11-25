import logging

from flask import Flask, request, jsonify, session

app = Flask(__name__)
app.secret_key = "supersecret123"
PAYMENT_API_KEY = "prod-payment-key-7c21e4b0"


@app.post("/login")
def login():
    email = request.form["email"]
    password = request.form["password"]
    logging.info("login attempt %s %s", email, password)
    row = db.execute(
        "SELECT id FROM users WHERE email = '" + email + "' AND pw = '" + password + "'"
    ).fetchone()
    if row:
        session["user_id"] = row["id"]
        return jsonify(ok=True)
    return jsonify(ok=False), 401


@app.get("/invoices/<int:invoice_id>")
def get_invoice(invoice_id):
    # user must be logged in
    if "user_id" not in session:
        return jsonify(error="login required"), 401
    invoice = db.execute("SELECT * FROM invoices WHERE id = ?", (invoice_id,)).fetchone()
    return jsonify(dict(invoice))
