import numpy as np
import pandas as pd

from keras.preprocessing.text import Tokenizer
from keras.preprocessing.sequence import pad_sequences

from services.load_model import load_model

# Load the trained DL model
loaded_model = load_model()

# Initialize the Tokenizer
sms_messages = pd.read_csv('../../analytics/data/spam.csv', encoding='latin1')
sms_messages = sms_messages.iloc[:, [1]]
sms_messages.columns = ["message"]
X = sms_messages["message"]
tokenizer = Tokenizer(num_words = 400, oov_token = "<OOV>")
tokenizer.fit_on_texts(X)

# Method for making prediction on SMS message
def get_predictions(txts):
    txts = tokenizer.texts_to_sequences(txts)
    txts = pad_sequences(txts, maxlen=250)
    preds = loaded_model.predict(txts)
    confidence = f"{float(np.array(preds, dtype=np.float32)[0]) * 100:.2f}%"

    if(preds[0] > 0.10):
        print("SPAM MESSAGE")
        return True, confidence
    else:
        print('NOT SPAM')
        return False, confidence
