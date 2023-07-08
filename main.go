package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type Team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OnCallResponse struct {
	Oncalls []struct {
		User struct {
			Summary string `json:"summary"`
		} `json:"user"`
	} `json:"oncalls"`
}

func main() {
	// Load Env variables from .env file
	godotenv.Load(".env")

	token := os.Getenv("SLACK_AUTH_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	// Create a new client to Slack by giving the token
	client := slack.New(token, slack.OptionAppLevelToken(appToken))

	// Create a new socket mode client
	socket := socketmode.New(client)

	// Create a context that can be used to cancel the goroutine
	ctx, cancel := context.WithCancel(context.Background())
	// Make sure to call this cancel properly in a real program (e.g., graceful shutdown)
	defer cancel()

	go func(ctx context.Context, socket *socketmode.Client) {
		// Create a for loop that selects either the context cancellation or the incoming events
		for {
			select {
			// In case context cancellation is called, exit the goroutine
			case <-ctx.Done():
				log.Println("Shutting down socketmode listener")
				return
			case event := <-socket.Events:
				// We have a new event, let's type switch on the event
				switch event.Type {
				// Handle Events API events
				case socketmode.EventTypeEventsAPI:
					// Type cast the event to the EventsAPIEvent
					eventsAPI, ok := event.Data.(slackevents.EventsAPIEvent)
					if !ok {
						log.Printf("Could not type cast the event to the EventsAPIEvent: %v\n", event)
						continue
					}
					// Send an acknowledgement to the Slack server
					socket.Ack(*event.Request)
					// Handle the Events API event
					err := HandleEventMessage(eventsAPI, client)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}(ctx, socket)

	socket.Run()
}

// HandleEventMessage handles the Events API events
func HandleEventMessage(event slackevents.EventsAPIEvent, client *slack.Client) error {
	switch event.Type {
	case slackevents.CallbackEvent:
		innerEvent := event.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			err := HandleAppMentionEventToBot(ev, client)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("unsupported event type")
	}
	return nil
}

// HandleAppMentionEventToBot handles the AppMentionEvent when the bot is mentioned
func HandleAppMentionEventToBot(event *slackevents.AppMentionEvent, client *slack.Client) error {
	// Extract the user's message after mentioning the bot's name
	text := strings.TrimSpace(strings.TrimPrefix(event.Text, fmt.Sprintf("<@%s> ", event.BotID)))
	pattern := fmt.Sprintf(`%s(.*?)oncall`, event.BotID)
	regex := regexp.MustCompile(pattern)
	matches := regex.FindStringSubmatch(text)

	if len(matches) > 1 {
		extractedText := strings.TrimSpace(matches[1])
		fmt.Println(extractedText)





		//#######PAGER DUTY

		// Load environment variables from .env file
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Error loading .env file:", err)
			return err
		}

		// Retrieve the PagerDuty token from environment variables
		pagerDutyToken := os.Getenv("PAGERDUTY_TOKEN")
		if pagerDutyToken == "" {
			fmt.Println("PagerDuty token not found in .env file.")
			return errors.New("PagerDuty token not found in .env file")
		}

		// Set the PagerDuty API endpoint URLs
		teamsURL := "https://api.pagerduty.com/teams"
		onCallsURL := "https://api.pagerduty.com/oncalls"

		// Set the team name

		mentionRegex := regexp.MustCompile("<@\\w+>")
		extractedTextval := mentionRegex.ReplaceAllString(extractedText, "")

		teamName := strings.TrimSpace(extractedTextval)


		// Create a new HTTP GET request to fetch the team ID based on the team name
		teamsReq, err := http.NewRequest("GET", teamsURL, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}
		teamsReq.Header.Set("Authorization", "Token token="+pagerDutyToken)

		// Send the HTTP request to fetch teams
		httpClient := &http.Client{}
		teamsResp, err := httpClient.Do(teamsReq)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return err
		}
		defer teamsResp.Body.Close()

		// Read the teams response body
		teamsBody, err := ioutil.ReadAll(teamsResp.Body)
		if err != nil {
			fmt.Println("Error reading teams response:", err)
			return err
		}

		// Parse the teams JSON response
		var teamsData struct {
			Teams []Team `json:"teams"`
		}
		err = json.Unmarshal(teamsBody, &teamsData)
		if err != nil {
			fmt.Println("Error parsing teams JSON:", err)
			return err
		}

		// Find the team ID based on the team name
		var teamID string
		for _, team := range teamsData.Teams {
			if strings.EqualFold(team.Name, teamName) {
				teamID = team.ID
				break
			}
		}

		if teamID == "" {
			fmt.Println("Team not found.")
			return errors.New("Team not found")
		}

		// Create a new HTTP GET request to fetch on-call information using the team ID
		onCallsReq, err := http.NewRequest("GET", onCallsURL, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return err
		}
		onCallsReq.Header.Set("Authorization", "Token token="+pagerDutyToken)
		onCallsReq.URL.RawQuery = "team_ids[]=" + teamID

		// Send the HTTP request to fetch on-call information
		onCallsResp, err := httpClient.Do(onCallsReq)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return err
		}
		defer onCallsResp.Body.Close()

		// Read the on-call response body
		onCallsBody, err := ioutil.ReadAll(onCallsResp.Body)
		if err != nil {
			fmt.Println("Error reading on-call response:", err)
			return err
		}

		// Parse the on-call JSON response
		var onCallResp OnCallResponse
		err = json.Unmarshal(onCallsBody, &onCallResp)
		if err != nil {
			fmt.Println("Error parsing on-call JSON:", err)
			return err
		}

		// Check if on-call data exists
		if len(onCallResp.Oncalls) > 0 {
			finalOnCallUsername := onCallResp.Oncalls[len(onCallResp.Oncalls)-1].User.Summary
			fmt.Println("Final on-call username:", finalOnCallUsername)

			// Send the extracted text as a response
			response := fmt.Sprintf("%s", finalOnCallUsername)
			_, _, err := client.PostMessageContext(
				context.Background(),
				event.Channel,
				slack.MsgOptionText(response, false),
			)
			if err != nil {
				return fmt.Errorf("failed to send message: %w", err)
			}
		} else {
			fmt.Println("No on-call data found.")
		}
	}

	return nil
}
