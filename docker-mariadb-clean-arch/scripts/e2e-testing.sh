#!/bin/bash
# Declare variables.
API_URL=http://localhost:8080

# Introduction to the script.
echo "Welcome to 'docker-mariadb-clean-arch' application!"
echo "Before running the end-to-end tests, please ensure that you have run 'docker-compose up -d'!"
read -n 1 -s -r -p "If you have done so, please press any key to continue!"; echo

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
echo "Testing GET route '/api/v1/users/:userID'..."
curl $API_URL/api/v1/users/11; echo
echo
echo "Testing PUT route '/api/v1/users/:userID'..."
curl -X PUT -H 'Content-Type: application/json' -d '{"name":"Mirajane Strauss","address":"Osaka, Japan"}' $API_URL/api/v1/users/11; echo
echo
echo "Testing DELETE route '/api/v1/users/:userID'..."
curl -X DELETE $API_URL/api/v1/users/11; echo

# Testing '/api/v1/auth/login'.
echo
echo "Testing POST route '/api/v1/auth/login'..."
curl -X POST -H 'Content-Type: application/json' -d '{"username":"fiber","password":"fiber"}' -c cookie.txt $API_URL/api/v1/auth/login; echo

# Testing '/api/v1/auth/private'.
echo
echo "Testing GET route '/api/v1/auth/private'..."
curl -b cookie.txt $API_URL/api/v1/auth/private; echo

# Testing '/api/v1/auth/logout'.
echo
echo "Testing POST route '/api/v1/auth/logout'..."
curl -X POST $API_URL/api/v1/auth/logout; echo

# Finish end-to-end testing.
rm -rf cookie.txt
echo "Finished testing the application!"
