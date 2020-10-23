# Spam Filter Demo
This is a simple demonstration project to showcase some of the interesting technologies and projects you might get to work on as a developer at [Telenor](https://www.telenor.no/privat/).

<img src="https://github.com/FredrikBakken/spam-filter-demo/raw/main/docs/assets/telenor-logo.png" width="250"/>

# Introduction
The purpose of this demonstration project is to illustrate how deep learning (DL) can be used to filter incoming SMS messages as either `spam`- or `ham`-messages. It uses an open-source dataset from [Kaggle](https://www.kaggle.com/) - [SMS Spam Collection Dataset](https://www.kaggle.com/uciml/sms-spam-collection-dataset).

## Requirements
The following set of requirements has to be fulfilled in order to run the example project:
- [Docker](https://www.docker.com/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [Python v3.8.6+](https://www.python.org/)
- [Jupyter Notebook](https://jupyter.org/)
- [Golang v1.15.2+](https://golang.org/)
- [Postman](https://www.postman.com/)
- [Graphviz](https://graphviz.org/download/) (Only required if you want to make the model image)

# Project Overview
In this project we will take a closer look at the process of developing a data streaming service application for filtering SMS messages as either spam or ham messages. [Apache Kafka](https://kafka.apache.org/) will be used as the dedicated streaming platform and [Keras](https://keras.io/) is used as the deep learning API for developing the prediction model.

The project consists of the following four modules:
1. **DL Analytics:** Design and develop a deep learning prediction model.
2. **Model Service:** Build the deep learning prediction API.
3. **Kafka Producer:** Publish incoming SMS-messages to a Kafka topic.
4. **Kafka Filter Stream:** Develop the streaming application for filtering incoming SMS-messages.

## Architecture
<img src="https://github.com/FredrikBakken/spam-filter-demo/raw/main/docs/assets/architecture.png"/>

- **Green boxes:** Kafka applications developed in Golang.
- **Blue boxes:** Kafka topics.
- **Orange box:** External prediction service API.

# Applications

## DL Analytics
We start by examining the deep learning module as it will function as a service for the streaming application to filter the incoming SMS-messages.

### Launch the Application
Run the following command within the `/analytics`-directory to access the notebook:

```
>> jupyter notebook
```

A new page should now open in your browser on http://localhost:8888. Go to the `/notebooks` directory and open the `Deep Learning - Spam Message Classification.ipynb` notebook.

### Exploratory Data Analysis
When performing data analysis it is important to inspect the type of data that one is working with, from our dataset we can find the following:

```
v1,v2,,,
ham,"Go until jurong point, crazy.. Available only in bugis n great world la e buffet... Cine there got amore wat...",,,
ham,Ok lar... Joking wif u oni...,,,
spam,Free entry in 2 a wkly comp to win FA Cup final tkts 21st May 2005. Text FA to 87121 to receive entry question(std txt rate)T&C's apply 08452810075over18's,,,
ham,U dun say so early hor... U c already then say...,,,
...
```

From this data we find that it is structured into five columns, where only the first two columns includes information that is interesting for us.


Further examination of the data distribution shows us the difference between the number of ham and spam messages in the dataset:

<img src="https://github.com/FredrikBakken/spam-filter-demo/raw/main/docs/assets/ham-vs-spam-count.png" width="600"/>

It is also possible to investigate the message length of the incoming SMS-messages in the dataset and compare the differences in spam vs. not-spam:

<img src="https://github.com/FredrikBakken/spam-filter-demo/raw/main/docs/assets/ham-vs-spam-length.png" width="600"/>

Further more - an illustration of the word distribution for ham and spam messages is also performed:

<img src="https://github.com/FredrikBakken/spam-filter-demo/raw/main/docs/assets/ham-vs-spam-most-frequent.png" width="600"/>

### Define the Model Architecture
For building the deep learning (DL) model the [Keras](https://keras.io/) API is used. A deep neural network (DNN) is an artificial neural network (ANN) with multiple hidden layers between the input layer and the output layer. These layers are fully connected and includes a weighted parameter which is adjusted during the training phase of the model. Once a deep learning model is trained it's architecture and trained weights can used to make predictions on new messages. Our deep learning model architecture is illustrated below:

<img src="https://github.com/FredrikBakken/spam-filter-demo/raw/main/docs/assets/model_plot.png" width="600"/>

## Model Service
The model service application is designed to make our prediction model production ready by loading it into a REST-API and make predictions on incoming SMS-messages.

### Launch the Application
Run the following commands within the `/apps/trained-model-api`-directory to download the dependencies and launch the application:

```
>> python3 -m pip install -r requirements.txt
>> flask run
```

The application will now be running locally on your machine (http://localhost:5000) and can be used by sending a POST-request to one of the endpoints (http://localhost:5000/sms or http://localhost:5000/bulk-sms).

### Message Example
Sending POST-requests to this model API can easily be done by using `curl` or Postman.

#### Example 1: Singular SMS-Message
POST-Request:
```
{
    "message": "Hi man, I was wondering if we can meet tomorrow."
}
```

Returned:
```
{
    "Message": [
        "Hi man, I was wondering if we can meet tomorrow."
    ],
    "Spam": false,
    "Confidence": "0.00%"
}
```

#### Example 2: Bulk of SMS-Messages
POST-Request:
```
{
    "messages": [
        {
            "message": "Free entry in 2 a wkly comp to win FA Cup final tkts 21st May 2005"
        },
        {
            "message": "Hi man, I was wondering if we can meet tomorrow."
        }
    ]
}
```

Returned:
```
[
    {
        "message": "Free entry in 2 a wkly comp to win FA Cup final tkts 21st May 2005",
        "Spam": true,
        "Confidence (Spam)": "56.01%"
    },
    {
        "message": "Hi man, I was wondering if we can meet tomorrow.",
        "Spam": false,
        "Confidence (Spam)": "0.00%"
    }
]
```

## Kafka Producer
...

## Kafka Filter Stream
...
