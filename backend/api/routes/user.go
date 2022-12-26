package routes

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/m4tthewde/happening/backend/api/internal"
	"github.com/nicklaw5/helix"
)

func GetUser(request events.APIGatewayProxyRequest) (string, int) {
	var params helix.UsersParams
	name, ok := request.QueryStringParameters["name"]
	if ok {
		params = helix.UsersParams{Logins: []string{name}}
	} else {
		id, ok := request.QueryStringParameters["id"]
		if !ok {
			return "", 404
		}

		params = helix.UsersParams{IDs: []string{id}}
	}

	resp, err := internal.GetUsers(&params)
	if err != nil {
		log.Println(err)
		return "", 500
	}

	if len(resp.Data.Users) == 0 {
		return "", 404
	}

	bodyBytes, err := json.Marshal(resp.Data.Users[0])
	if err != nil {
		log.Println(err)
		return "", 500
	}

	return string(bodyBytes), 200
}
