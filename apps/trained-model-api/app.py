import tensorflow as tf

from flask import Flask

from keras.preprocessing.text import Tokenizer
from keras.preprocessing.sequence import pad_sequences


loaded_model = tf.keras.models.load_model(
    "model/",
    custom_objects = None,
    compile = True,
)

def get_predictions(txts):
    tokenizer = Tokenizer(num_words = 400, oov_token = "<OOV>")
    tokenizer.fit_on_texts(txts)
    
    txts = tokenizer.texts_to_sequences(txts)
    txts = pad_sequences(txts, maxlen=250)
    preds = loaded_model.predict(txts)
    if(preds[0] > 0.5):
        print("SPAM MESSAGE")
    else:
        print('NOT SPAM')

app = Flask(__name__)

@app.route("/")
def home_route():
    return "Hello world!"

@app.route("/sms/<message>", methods = ["GET", "POST"])
def sms(message):
    print(message)

    #txts = ["Hi man, I was wondering if we can meet tomorrow."]
    #txts = ["Free entry in 2 a weekly competition to win FA Cup final tkts 21st May 2005"]

    get_predictions(message)

    predicted = loaded_model.predict_classes([[0, 23, 65, 34, 12, 89]])
    return str(predicted)
