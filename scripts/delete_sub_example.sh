#!/bin/bash
curl -X DELETE "https://api.twitch.tv/helix/eventsub/subscriptions?id=$1" \
-H 'Authorization: Bearer <TOKEN>' \
-H 'Client-Id: <CLIENT_ID>'