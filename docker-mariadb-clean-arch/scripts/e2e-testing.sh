#!/bin/bash
# Declare variables.
API_URL=http://localhost:8080

# Introduction to the script.
echo "Welcome to 'docker-mariadb-clean-arch' application!"
echo "Before running the end-to-end tests, please ensure that you have run 'make start'!"; echo

# Testing '/api/v1'.
echo
echo "Running end-to-end testing..."
echo "Testing GET route '/api/v1'..."
curl $API_URL/api/v1; echo

# Testing '/api/v1/users'.
echo
echo "Testing GET route '/api/v1/users'..."
curl $API_URL/api/v1/users; echo
echo
echo "Testing POST route '/api/v1/users'..."
curl -X POST -H 'Content-Type: application/json' -d '{"name":"Lucy Heartfilia","address":"Shinhotaka, Japan"}' $API_URL/api/v1/users; echo

# Testing '/api/v1/users/:userID'.
echo
echo "Using 'userID' with value of 11 (the one created beforehand)."
echo "Testing GET route '/api/v1/users/:userID'..."
curl $API_URL/api/v1/users/11; echo
echo
echo "Testing PUT route '/api/v1/users/:userID'..."
curl -X PUT -H 'Content-Type: application/json' -d '{"name":"Mirajane Strauss","address":"Osaka, Japan"}' $API_URL/api/v1/users/11; echo
echo
echo "Testing GET route after PUT '/api/v1/users/:userID'..."
curl $API_URL/api/v1/users/11; echo
echo
echo "Testing DELETE route '/api/v1/users/:userID'..."
curl -X DELETE $API_URL/api/v1/users/11; echo
echo
echo "Testing GET route after DELETE '/api/v1/users/:userID'..."
curl $API_URL/api/v1/users/11; echo

# Testing '/api/v1/auth/login'.
echo
echo "Testing POST route '/api/v1/auth/login'..."
curl -X POST -H 'Content-Type: application/json' -d '{"username":"fiber","password":"fiber"}' -c cookie.txt $API_URL/api/v1/auth/login; echo

# Testing '/api/v1/auth/private'.
echo
echo "Testing GET route '/api/v1/auth/private'..."
curl -b cookie.txt $API_URL/api/v1/auth/private; echo

# Testing '/api/v1/cities'.
echo
echo "Testing GET route '/api/v1/cities'..."
curl -b cookie.txt $API_URL/api/v1/cities; echo
echo
echo "Testing POST route '/api/v1/cities'..."
curl -b cookie.txt -X POST -H 'Content-Type: application/json' -d '{"name":"Kyoto"}' $API_URL/api/v1/cities; echo

# Testing '/api/v1/cities/:cityID'.
echo
echo "Using 'cityID' with value of 6 (the one created beforehand)."
echo "Testing GET route '/api/v1/cities/:cityID'..."
curl -b cookie.txt $API_URL/api/v1/cities/6; echo
echo
echo "Testing PUT route '/api/v1/cities/:cityID'..."
curl -b cookie.txt -X PUT -H 'Content-Type: application/json' -d '{"name":"Osaka"}' $API_URL/api/v1/cities/6; echo
echo
echo "Testing GET route after PUT '/api/v1/cities/:cityID'..."
curl -b cookie.txt $API_URL/api/v1/cities/6; echo
echo
echo "Testing DELETE route '/api/v1/cities/:cityID'..."
curl -b cookie.txt -X DELETE $API_URL/api/v1/cities/6; echo
echo
echo "Testing GET route after DELETE '/api/v1/cities/:cityID'..."
curl -b cookie.txt $API_URL/api/v1/cities/6; echo

# Testing '/api/v1/auth/logout'.
echo
echo "Testing POST route '/api/v1/auth/logout'..."
curl -X POST $API_URL/api/v1/auth/logout; echo

# Finish end-to-end testing.
rm cookie.txt
echo "Finished testing the application!"
