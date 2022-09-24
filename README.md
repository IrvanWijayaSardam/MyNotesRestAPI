# MyNotesRESTAPI

Requirements :

- Docker 
- Go
- Mysql

## INSTALLATION PROCCES ##

1. Run the Docker-Compose using
```
docker-compose up
```
2. Make the databases inside your image docker , use MyNotes.sql , default user and password is root 
3. To access your mysql inside docker image use localhost and with port 3308 
4. Run the app
```
go run main.go
```
#API ENDPOINT 
```   

POST : notes/ Insert Notes
{
    "title" : "Your Notes Title",
    "description" : "Your Notes Description",
    "userid" : "69"
}
GET : notes/ Grab All Notes
{
   
}
GET : notes/{id} Grab Notes By ID
{
    "title" : "Your Notes Title",
    "description" : "Your Notes Description",
    "userid" : "69"
}

PUT : notes/{id} Update Notes
{
    "title" : "Your new Notes Title",
    "description" : "Your New Notes Description",
    "userid" : "69"
}
```

```
GET : user/ Get All User

POST : user/ Insert User
{
    "username" : "john doe",
    "email" : "aminivan@gmail.com",
    "Password" : "urpassword"
}
PUT : user/{userID} update user
{
    "username" : "New Updated doe",
    "email" : "Updated@gmail.com",
    "Password" : "urnewpassword"
}


```
