# slack-pd-bot
A slackbot to interact with PagerDuty, this is WIP project, just sharing the first iteration(we follow semantic versioning) only here just for feedback loop.

# Introduction

This project was created out of a requirement to use [Chatops](https://response.pagerduty.com/resources/chatops/) to manage incident and interact with incident response management tool like PagerDuty. The functions and usecase of this bot is similar to [Hubot](https://hubot.github.com/) or [Pagerly](https://www.pagerly.io/).


#Installation.

1. Currently the code supports only fetching the person whos oncall for a team from [PagerDuty](https://www.pagerduty.com/), rest of the PagerDuty features will be added to the bot and if you would like to add any feature please mention it in Github issues so can add those feature
2. Helm chart for deploying this as service in k8s is in progress, will share it eventually.
3. Get Slack Tokens and Pager Duty token, steps mentioned below.
4. Thanks to [this blog](https://www.bacancytechnology.com/blog/develop-slack-bot-using-golang) from Sourabh Chakravarty to get an idea of create the slack tokens, please follow and create the slack app and auth tokens.
5. Create the PD token , [refer](https://support.pagerduty.com/docs/api-access-keys)
6. The tokens are added to the .env file as we use (this module)[https://github.com/joho/godotenv] to fetch the tokens as env variables, if the app is deployed to k8s, expose it in secrets(I will add steps to deploy this as a service in k8s, do this in next iteration of this code)

#Run the code

1. to run the code just call the main.go file
2. in slack call "@botname devops oncall", this will return whoever is devops oncall, attached snapshots for reference.
3. Attached snapshots for reference.


# TODO

1. Helmify the app so it can be deployed in k8s
2. Add more features like schedule oncall, schedule overrides, add and create teams etc using Slack commands
