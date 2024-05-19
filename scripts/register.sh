#!/bin/bash

source .env

curl -XPOST --header "Content-Type: application/json"  \
  --header "Authorization: Bot $TOKEN" \
  -d '{"name":"sheet","description":"show charatersheet"}' \
  "https://discord.com/api/applications/1241384650039889931/commands"


curl -XPOST --header "Content-Type: application/json"  \
  --header "Authorization: Bot $TOKEN" \
  -d '{"name":"me","description":"show informations"}' \
  "https://discord.com/api/applications/1241384650039889931/commands"


curl -XPOST --header "Content-Type: application/json"  \
  --header "Authorization: Bot $TOKEN" \
  -d '{"name":"statistic","description":"make a statistic check","options":[{"name":"statistic","description":"the stat to check","type":3}]}' \
  "https://discord.com/api/applications/1241384650039889931/commands"