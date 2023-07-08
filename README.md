# slack-pd-bot

A Slack bot to interact with PagerDuty. This is a work-in-progress project, and we are sharing the first iteration for feedback.

## Introduction

This project was created to fulfill the requirement of using ChatOps to manage incidents and interact with incident response management tools like PagerDuty. The functionality and use cases of this bot are similar to Hubot or Pagerly.

## Installation

1. Currently, the code only supports fetching the person who is on-call for a team from PagerDuty. More PagerDuty features will be added to the bot in the future. If you would like to add any features, please mention them in the GitHub issues.
2. The Helm chart for deploying this bot as a service in Kubernetes is currently in progress and will be shared eventually.
3. To get started, obtain Slack Tokens and PagerDuty token by following the steps below.
4. Thanks to [this blog](https://www.bacancytechnology.com/blog/develop-slack-bot-using-golang) from Sourabh Chakravarty for providing insights on creating Slack tokens. Please follow the instructions in the blog to create a Slack app and obtain the required auth tokens.
5. Create the PagerDuty token by referring to the PagerDuty documentation on [API access keys](https://support.pagerduty.com/docs/api-access-keys).
6. Add the Slack tokens and PagerDuty token to the .env file. We use the [godotenv](https://github.com/joho/godotenv) module to fetch the tokens as environment variables. If the app is deployed to Kubernetes, expose the tokens in secrets. (Steps to deploy this as a service in Kubernetes will be added in the next iteration of this code.)

## Running the Code

1. To run the code, simply call the main.go file.
2. In Slack, call "@botname devops oncall". This will return the person who is on-call for the DevOps team. Please refer to the attached snapshots for reference.

### Attached Snapshots

1. Running the code
![Screenshot 2023-07-08 at 4 24 50 PM](https://github.com/neeltom92/slack-pd-bot/assets/135661004/c3f396ce-e3d7-42a5-82a1-df05d51feee0)

2. Slack input
![Screenshot 2023-07-08 at 4 24 38 PM](https://github.com/neeltom92/slack-pd-bot/assets/135661004/97854c37-870e-4fba-a51e-249274e7a8c1)

3. Output
![Screenshot 2023-07-08 at 4 25 21 PM](https://github.com/neeltom92/slack-pd-bot/assets/135661004/3444198b-5d7c-4b64-8809-ec41b470a6c8)

## TODO

1. Helmify the app so it can be deployed in Kubernetes.
2. Add more features like schedule on-call, schedule overrides, and add/create teams using Slack commands.
