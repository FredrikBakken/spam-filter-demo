import json

from flask import Flask, request

from services.predict import get_predictions


# Initialize Flask
app = Flask(__name__)

# Setup incoming route path
@app.route("/sms", methods = ["GET", "POST"])
def sms():
    data = request.json
    messages = data["messages"]
    #print(messages) # Debugging

    predictions = []
    for msg in messages:
        message = [msg["message"]]
        #print(message) # Debugging

        prediction, confidence = get_predictions(message)
        msg["Spam"] = prediction
        msg["Confidence (Spam)"] = confidence
        predictions.append(msg)

    return json.dumps(predictions)
