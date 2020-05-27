# Data Academy Serverless

## Introduction

* Cloud Adoption presentation
* Create a Lambda function in the AWS Console
  * Trigger it in the console
  * Trigger it via the command line
  * Apply an external trigger

# Serverless Framework

* Install Node.js
* Install Serverless Framework
* Create a Serverless app using Serverless framework
  * Deploy Hello World
  * Handle a HTTP post request
  * Write the post request content to the console as JSON
  * Write a Lambda that subscribes to Event Bridge
  * Post data to Event Bridge via the CLI
  * Post data to Event Bridge using the Go program
* Other events
  * Write a Lambda that subscribes to a file arriving in an S3 bucket.
    * Create a test file with the generator.
    * It should then read the file and send each JSON line to Event Bridge.
  * Use AWS Log Insights to calculate the best selling product by analysing the logs.
  * Use AWS Athena to calculate the total amount of products sold directly by querying the data in the S3 bucket.
* Stock calculation
  * Create a DynamoDB table by using CloudFormation
  * Rick Houlian talk https://www.youtube.com/watch?v=6yqfmXiZTlM
  * Create a "delivery" event schema to handle receipt of stock. When stock is received, update the database.
  * Update the "transaction" EventBridge handler to remove items from the Stock table as items are sold.
  * Update the Lambda to post a warning to Event Bridge if a stock record exists and the stock is lower than 10.

# Stretch goal

* Using the console, subscribe Kinesis to all EventBridge events and push them to S3 for querying via AWS Athena.

# AWS Step Functions for batch processing

* Importing 4M rows into DynamoDB talk - Q&A

