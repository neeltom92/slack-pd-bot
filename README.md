# slack-pd-bot

A Slack bot to interact with PagerDuty. This is a work-in-progress, side project I created over the Independence day week ![Screenshot 4](https://raw.githubusercontent.com/stevenrskelton/flag-icon/master/png/16/country-4x3/us.png) off in July 23, and I'm sharing the first iteration for feedback loop only.(I use [semantic_version](https://github.com/neeltom92/slack-pd-bot/tags) to keep track of release)

## Introduction

This project was created out of a requirement to use [Chatops](https://response.pagerduty.com/resources/chatops/) to manage incident and interact with incident response management tool like PagerDuty. The functions and usecase of this bot is similar to [Hubot](https://hubot.github.com/) or [Pagerly](https://www.pagerly.io/) or [Errbot](https://errbot.readthedocs.io/en/latest/), but in a lightweight model for anyone to get started.

## Installation

1. Currently, the code only supports fetching the person who is on-call for a team from PagerDuty. More PagerDuty features will be added to the bot in the future. If you would like to add any features, please mention them in the GitHub issues.
2. The Helm chart for deploying this bot as a service in Kubernetes is currently in progress and will be shared eventually, currently this needs to be run in local.
3. To get started, obtain Slack Tokens and PagerDuty token by following the steps below.
4. Thanks to [this blog](https://www.bacancytechnology.com/blog/develop-slack-bot-using-golang) from Sourabh Chakravarty for providing insights on creating Slack tokens. Please follow the instructions in the blog to create a Slack app and obtain the required auth tokens.
5. Create the PagerDuty token by referring to the PagerDuty documentation on [API access keys](https://support.pagerduty.com/docs/api-access-keys).
6. Add the Slack tokens and PagerDuty token to the .env file. We use the [godotenv](https://github.com/joho/godotenv) module to fetch the tokens as environment variables. If the app is deployed to Kubernetes, expose the tokens in secrets. (Steps to deploy this as a service in Kubernetes will be added in the next iteration of this code.)

## Running the Code in local

1. To run the code, simply call the main.go file, after updating the tokens
2. In Slack, call "@botname devops oncall". This will **return the person who's on call for that team**. Please refer to the attached snapshots for reference.

### Attached Snapshots

#### Running the code
<img width="418" alt="Screenshot 2023-07-08 at 424 38 PM" src="https://github.com/neeltom92/slack-pd-bot/assets/135661004/c3f396ce-e3d7-42a5-82a1-df05d51feee0">

#### using bot to give Slack input
<img width="418" alt="Screenshot 2023-07-08 at 44 38 PM" src="https://github.com/neeltom92/slack-pd-bot/assets/135661004/5b737c66-f3ef-4e03-b6f2-f4148cd482fc">

#### oncall schedule in PagerDuty
![Screenshot 3](https://github.com/neeltom92/slack-pd-bot/assets/135661004/3444198b-5d7c-4b64-8809-ec41b470a6c8)

**As you can see on making the query in slack the bot returned the user who's on call for that team**.

## TODO

1. Helmify the app so it can be deployed in Kubernetes.
2. Add more features like schedule on-call, schedule overrides, and add/create teams using Slack commands.
