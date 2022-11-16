# Sessions - SQLite3

A real-world example of how to use Fiber sessions with Storage package.\
Run localhost:3000 from multiple browsers to see active sessions for different users.

## Explanation

This example uses SQLite3 Storage package to persist users sessions.\
Storage package can create sessions table for you at init time but for the purpose of this example I created it manually expanding its structure with an "u" column to better query all user-related sessions.