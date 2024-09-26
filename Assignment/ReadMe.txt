Create Network (It will be used to communicate between containers) - 

1. docker network create -d bridge assignment_network


Service1:WebApp - 
Commands - 
1. In terminal go to directory 'Assignment'
2. docker build -t assignment:webappV0 ./src
3. docker run -p 3333:3333 --network assignment_network --name assignment_webapp -d assignment:webappV0


Service2:SQLServer -
Commands - 
1. In terminal, go to directory 'Assignment'
2. docker build -t assignment:SqlServerV0 ./db
3. docker run -e "MSSQL_SA_PASSWORD=admin@123" -p 1433:1433 --network assignment_network --name assignment_sqlserver -d assignment:SqlServerV0


// to check, if sql server container is started, try below commands from host machine's terminal:
1. docker exec -it assignment_sqlserver "bash"
2. /opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -C
    Enter password of sa user (admin@1234) //this has been altered through script
3. 1> SELECT Name from sys.databases;
   2> go
   Press Enter. Below output must be shown:

Name
--------------------------------------------------------------------------------------------------------------------------------
master
tempdb
model
msdb
STUDENTDATA


Web APIs - 
http://localhost:3333/check_db
http://localhost:3333/execSql?query=insert into Student(rollno,name,class) values(1, 'S1', 1)
http://localhost:3333/execSql?query=insert into Student(rollno,name,class) values(2, 'S2', 2)
http://localhost:3333/execSql?query=insert into Student(rollno,name,class) values(3, 'S3', 3)
http://localhost:3333/execSql
http://localhost:3333/updateSql
http://localhost:3333/execSql
http://localhost:3333/execSql?query=delete from Student where rollno = 3
http://localhost:3333/execSql


Using docker compose - 

1. On Terminal, inside "Assignment" dir, run below command - 
   > docker compose up -d
 This will create all the specified services and networks in docker-compose.yml.
 2. Now you may try if, mssql server container is started as specified above.
 3. Then webapp api can be tried. 


 Bonus Ques - 

 Setting up different environments (Development, Testing, Production) - (Commands to be executed)
 1. Development :
      > docker compose up -d
    This will pick the default compose yml file i.e. docker-compose.yml

 2. Testing : (first execute command - > docker compose down)
      > docker compose -f docker-compose.yml -f docker-compose-test.yml up -d
    This will execute test yml file on top of defualt yml file.
    
 3. Production : (first execute command - > docker compose down)
      > docker compose -f docker-compose.yml -f docker-compose-prod.yml up -d
    This will execute prod yml file on top of defualt yml file.      

