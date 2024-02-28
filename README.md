# Github Oauth Api
An api that connects with github's api via github's Oauth service

## Installation Guide

   1.  [Install Chocolatey](https://docs.chocolatey.org/en-us/choco/setup)``windows``
   2.  Install Make ```choco install make``` ``windows``
   3.  [Install Scoop](https://scoop.sh/)
   4.  Install [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)```scoop install migrate```
   5.  [Install Docker](https://docs.docker.com/get-docker/)

## Setup Local Environment 

   1. Create in root directory ```.env``` and ```token.env``` files
   2. ```token.env``` has one parameter ``ACTIVE_TOKEN`` which shouldn't be filled by default   
   3. ```.env``` has the following parameters 
         1. ``DB_URL`` will contain the db url for the migration file it will have this format 
      
         ``postgres://username:password@host:port/database?host=/var/run/postgresql/data&sslmode=disable``
      
         2. ``CLIENT_ID`` will contain  client ID from the Oauth2.0 application of Github
         
         3. ``DB_URL_API`` will contain the database url that the dockerized api will access
         
            ``postgres://username:password@host:port/databasename``
         4. ``CLIENT_SECRET`` will contain  client secret from the Oauth2.0 application of Github
         
## Start the environment

Open a powershell terminal and run the following command in project's root directory to create the migrations
```make
make migrate
```
Run the following command to start the project
```make
make up
```
