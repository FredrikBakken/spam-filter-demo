import json

from flask import Flask, request

from services.predict import get_predictions


# Initialize Flask
app = Flask(__name__)

# Setup incoming route path
@app.route("/sms", methods = ["GET", "POST"])
def sms():
    data = request.json
    message = data["message"]
    prediction, confidence = get_predictions([message])

    msg = {
        "Message": message,
        "Spam": prediction,
        "Confidence": confidence,
    }

    return json.dumps(msg)

# Setup incoming route for bulk of sms messages
@app.route("/bulk-sms", methods = ["GET", "POST"])
def bulk_sms():
    data = request.json
    messages = data["messages"]

    predictions = []
    for msg in messages:
        message = [msg["message"]]

        prediction, confidence = get_predictions(message)
        msg["Spam"] = prediction
        msg["Confidence (Spam)"] = confidence
        predictions.append(msg)

    return json.dumps(predictions)
