### Project Mini-Task with JSONB POSTGRESQL
1.  GET ALL DATA <br />
    GET `http://localhost:8080/people`
2.  GET DATA WHERE NAME <br />
    GET `http://localhost:8080/apiumur/Lala`
3.  CREATE NEW DATA <br />
    POST `http://localhost:8080/apiumur` <br />
    Content-Type = application/json <br />
    {<br />
        "name": "Fatin",<br />
        "age": 22<br />
    }<br />
4.  UPDATE DATA <br />
    PUT `http://localhost:8080/apiumur/Fatin` <br />
    Content-Type = application/json <br />
    {<br />
        "name": "Fatin",<br />
        "age": 33<br />
    }<br />
5.  DELETE DATA <br />
    DELETE `http://localhost:8080/apiumur/Fatin` 