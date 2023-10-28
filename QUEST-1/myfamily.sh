#!/bin/bash
json_url='https://01.alem.school/assets/superhero/all.json'

curl -s "$json_url" -o superhero.json


myFamily=$(jq --arg HERO_ID "$HERO_ID"  '.[] | (select(.id == ($HERO_ID | tonumber)) | .connections.relatives )' superhero.json | tr -d '"')

echo "$myFamily"