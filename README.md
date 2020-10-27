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
- [Flutter v1.23.0-18.1.pre](https://flutter.dev/docs/get-started/install) (Only required for the client application)

# Project Overview
In this project we will take a closer look at the process of developing a data streaming service application for filtering SMS messages as either spam or ham messages. [Apache Kafka](https://kafka.apache.org/) will be used as the dedicated streaming platform and [Keras](https://keras.io/) is used as the deep learning API for developing the prediction model. A chat messaging client is also developed and deployed for interactive testing.

The project consists of the following five modules:
1. **DL Analytics:** Design and develop a deep learning prediction model.
2. **Model Service:** Build the deep learning prediction API.
3. **Kafka Producer:** Publish incoming SMS-messages to a Kafka topic.
4. **Kafka Filter Stream:** Develop the streaming application for filtering incoming SMS-messages.
5. **Messaging Client:** Develop and publish the messaging chat application for sending example messages.

## Architecture
<img src="https://github.com/FredrikBakken/spam-filter-demo/raw/main/docs/assets/architecture.png"/>

- **Green boxes:** Kafka applications developed in Golang.
- **Blue boxes:** Kafka topics.
- **Orange box:** External prediction service API.

# Applications
The applications in this project requires that the `cp-all-in-one` docker containers are running. These can be launched by the command `docker-compose up -d`.

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

The application will now be running locally on your machine (http://localhost:5000) and can be used by sending a POST-request to one of the following endpoints:
 - http://localhost:5000/sms
 - http://localhost:5000/bulk-sms

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
    "Message": "Hi man, I was wondering if we can meet tomorrow.",
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
        "Message": "Free entry in 2 a wkly comp to win FA Cup final tkts 21st May 2005",
        "Spam": true,
        "Confidence": "56.01%"
    },
    {
        "Message": "Hi man, I was wondering if we can meet tomorrow.",
        "Spam": false,
        "Confidence": "0.00%"
    }
]
```

## Kafka Producer
The incoming messages into the application will first arrive at the Kafka Producer application, called `sms-event-integration`. This application works as a REST-API with an endpoint for serving incoming messages. These messages are then structured into objects and published to a Kafka topic named `new-sms-json-v1`.

The application also pulls new messages from the Firebase - Cloud Firestore and pushes these messages to the `new-sms-json-v1` topic, which is used as a client-based example application.

### Launch the Application
Run the following command within the `/apps/sms-event-integration`-directory to launch the application:

```
>> go run app.go
```

## Kafka Filter Stream
The filter stream application is a streaming application that listens for new messages to arrive into the `new-sms-json-v1` topic, structures the messages into accepted JSON-format for the model API, sends POST-requests to the model API, and then pushes the messages to one of the following Kafka topics based on the predicted result:
- `safe-sms-json-v1`
- `spam-sms-json-v1`

In the client demonstration case, the results are also pushed to update the corresponding messages in Firebase - Cloud Firestore.

### Launch the Application
Run the following command within the `/apps/sms-filter-stream`-directory to launch the application:

```
>> go run app.go
```

## Messaging Client
The messaging client is a simple Flutter web chatroom application where clients can send messages. These messages are stored into the Firebase - Cloud Firestore and continously streamed by the web application. Any updates to the messages, such as categorization as spam or not-spam, will be accessible to the users using the chatroom.

### Launch the Application
To run this application on your own machine, it requires that you have created a new Firebase project and updated the `firebaseConfig` variable found within the `/client/web/index.html` file with your Firebase project config. After this is done, the client application can be launched by running the following commands within the `/client`-directory:

```
>> flutter run -d chrome
```

Important to note is that if you want to run the entire project locally, it also requires you to add the `serviceAccountKey.json` file into the `/config` directory of both the `sms-event-integration`- and the `sms-filter-stream` application.

# Have any Questions or in Need of Support?
Do you want a demonstration of the project or have any questions/issues, please feel free to create a `New Issue` on the [Issues](https://github.com/FredrikBakken/spam-filter-demo/issues) page and I will get back to you as soon as possible.
