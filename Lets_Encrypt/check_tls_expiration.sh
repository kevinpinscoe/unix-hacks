#!/bin/bash

site="kevininscoe.com"

date_str=$(curl https://${site} -vI --stderr - | grep "expire date" | cut -d":" -f 2- | awk '{ print $1, $2, $4 }')

echo "TLS cert for ${site} expires on ${date_str}"

# Convert the given date string to a timestamp
given_date=$(date -d "$date_str" +%s)

# Get the timestamp for one week from today
one_week_from_now=$(date -d "7 days" +%s)

echo "given_date=${given_date}"
echo "one_week_from_now=${one_week_from_now}"

# Get today's timestamp
today=$(date +%s)

# Check if the given date is exactly one week from today
if [[ $given_date -eq $one_week_from_now ]]; then
	echo "The date $date_str is exactly one week from today."
	echo "I will perform a Let's Encrypt renewal"
	#Invoke lets_encrypt_renewal
else
	echo "The date $date_str is NOT one week from today will go back to sleep now."
fi
