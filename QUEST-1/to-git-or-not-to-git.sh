#!bin/bash

json_url='https://01.alem.school/assets/superhero/all.json'

curl -s "$json_url" -o superhero.json

name=$(jq -r '.[] | (select(.id == 170) | .name)' superhero.json)
power=$(jq -r '.[] | (select(.id == 170) | .powerstats.power)' superhero.json)
gender=$(jq -r '.[] | (select(.id == 170) | .appearance.gender)' superhero.json)

printf "$name\n"
printf "$power\n"
printf "$gender\n"