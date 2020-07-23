# JWT+Mongo+Docker+Nginx 🚀
<img align="right" width="274" height="237" src="https://raw.githubusercontent.com/solrac97gr/solrac97gr/master/carlos.png">

## Freatures

- **JWT** : Custom middleware for JWT auth
- **Mongo** : Connection with mongodb and configuration file with the possibility of changing the environment
- **Docker and Nginx** : Deploy in docker using 5 replicas and load balancer with Nginx
- **Logger** : Request logger

## Endpoints

| Name         | Rute      | Parameters | State     | Protected | Method |
| ------------ | --------- | ---------- | --------- | --------- | ------ |
| Register     | /register | No         | Completed | No        | POST   |
| Login        | /login    | No         | Completed | No        | POST   |
| Get Profile  | /profile  | id         | Completed | Yes       | GET    |
| Edit Profile | /profile  | No         | Completed | Yes       | PUT    |
