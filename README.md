# MyNotesRESTAPI

Requirements :

- Docker 
- Go
- Mysql

## INSTALLATION PROCCES ##

1. Run the Docker-Compose using
        docker-compose up
2. Make the databases inside your image docker , use MyNotes.sql , default user and password is root 
3. To access your mysql inside docker image use localhost and with port 3308 
4. Run the app
        go run main.go

#API ENDPOINT 
UPDATED SOON : 
        
        /notes
        POST :
        {
       
              "title" : "Your Notes Title",
              "description" : "Your Notes Description",
              "userid" : "69"
                  
        }
