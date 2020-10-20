import tensorflow as tf

def load_model():
    loaded_model = tf.keras.models.load_model(
        "model/",
        custom_objects = None,
        compile = True,
    )

    return loaded_model
