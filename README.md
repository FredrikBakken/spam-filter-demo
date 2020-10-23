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

# Project Overview
In this project we will take a closer look at the process of developing a data streaming service application for filtering SMS messages as either spam or ham messages. [Apache Kafka](https://kafka.apache.org/) will be used as the dedicated streaming platform and [Keras](https://keras.io/) is used as the deep learning API for developing the prediction model.

The project consists of the following four modules:
1. **DL Analytics:** Design and develop a deep learning prediction model.
2. **Model Service:** Build the deep learning prediction API.
3. **Kafka Producer:** Publish incoming SMS-messages to a Kafka topic.
4. **Kafka Filter Stream:** Develop the streaming application for filtering incoming SMS-messages.

## Architecture
<img src="https://github.com/FredrikBakken/spam-filter-demo/raw/main/docs/assets/architecture.png"/>

# Applications

## DL Analytics
We start by examining the deep learning module as it will function as a service for the streaming application to filter the incoming SMS-messages.

### Launch the Application
Run the following command within the `/analytics`-directory to access the notebook:

```
>> jupyter notebook
```

A new page should now open in your browser on http://localhost:8888. Go to the `/notebooks` directory and open the `Deep Learning - Spam Message Classification.ipynb` notebook.

...
