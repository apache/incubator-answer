# How to build and install

Before installing Answer, you need to install the base environment first.
 - database
     - [MySQL](http://dev.mysql.com) Version >= 5.7

You can then install Answer in several ways:

 - Deploy with Docker
 - binary installation
 - Source installation

## Docker-compose for Answer
```bash
$ mkdir answer && cd answer
$ wget https://raw.githubusercontent.com/answerdev/answer/main/docker-compose.yaml
$ docker-compose up
```

browser open URL [http://127.0.0.1:9080/](http://127.0.0.1:9080/).

You can log in with the default administrator username( **`admin@admin.com`** ) and password( **`admin`** ).

## Docker for Answer
Visit Docker Hub or GitHub Container registry to see all available images and tags.

### Usage
To keep your data out of Docker container, we do a volume (/var/data -> /data) here, and you can change it based on your situation.

```
# Pull image from Docker Hub.
$ docker pull answerdev/answer:latest

# Create local directory for volume.
$ mkdir -p /var/data

# Run the image first
$ docker run --name=answer -p 9080:80 -v /var/data:/data answer/answer

# After the first startup, a configuration file will be generated in the /var/data directory
# /var/data/conf/config.yaml
# Need to modify the Mysql database address in the configuration file
vim /var/data/conf/config.yaml

# Modify database connection
# connection: [username]:[password]@tcp([host]:[port])/[DbName]
...

# After configuring the configuration file, you can start the mirror again to start the service
$ docker start answer
```

## Binary for Answer
## Install Answer using binary

  1. Unzip the compressed package
  2. Use the command cd to enter the directory you just created
  3. Execute the command ./answer init
  4. Answer will generate a ./data directory in the current directory
  5. Enter the data directory and modify the config.yaml file
  6. Modify the database connection address to your database connection address

     connection: [username]:[password]@tcp([host]:[port])/[DbName]
  7. Exit the data directory and execute ./answer run -c ./data/conf/config.yaml

## Available Commands
Usage: answer [command]

- help: Help about any command
- init: Init answer application
- run: Run answer application
- check: Check answer required environment
- dump: Backup answer data

## config.yaml Description

```
server:
  http:
    addr: 0.0.0.0:80 #Project access port number
data:
  database:
    connection: root:root@tcp(127.0.0.1:3306)/answer #MySQL database connection address
  cache:
    file_path: "/tmp/cache/cache.db" #Cache file storage path
i18n:
  bundle_dir: "/data/i18n" #Internationalized file storage directory
swaggerui:
  show: true #Whether to display the swaggerapi documentation, address /swagger/index.html
  protocol: http #swagger protocol header
  host: 127.0.0.1 #An accessible IP address or domain name
  address: ':80'  #accessible port number
service_config:
  secret_key: "answer" #encryption key
  web_host: "http://127.0.0.1" #Page access using domain name address
  upload_path: "./upfiles" #upload directory
```

## Compile the image
If you have modified the source files and want to repackage the image, you can use the following statement to repackage the image
```
docker build -t  answer:v1.0.0 .
```
## common problem
 1. The project cannot be started, answer the main program startup depends on the configuration file config.yaml, the internationalization translation directory/i18n, the upload file storage directory/upfiles, you need to ensure that the configuration file is loaded when the project starts, answer run -c config.yaml and the correct config.yaml The configuration items that specify the i18n and upfiles directories
