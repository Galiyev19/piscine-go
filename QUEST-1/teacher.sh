#!/bin/bash
interview_number=$(head  -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo "$interview_number"
file=$(cat interviews/interview-${interview_number})
echo "${file}"
echo "$MAIN_SUSPECT"