```
 ian@wintermute  ~/g/s/g/g/tls-trunk   master 
 ❯ curl https://secret.theknown.net/\~ian/                                                                                                                                                                                        [12:35:37]
<html>
<head><title>400 No required SSL certificate was sent</title></head>
<body bgcolor="white">
<center><h1>400 Bad Request</h1></center>
<center>No required SSL certificate was sent</center>
<hr><center>nginx/1.14.0 (Ubuntu)</center>
</body>
</html>

 ian@wintermute  ~/g/s/g/g/tls-trunk   master 
 ❯ curl -E ~/ian.crt.pem --key ~/ian.key.pem https://secret.theknown.net/\~ian/                                                                                                                                                   [12:35:50]
cant see this

 ian@wintermute  ~/g/s/g/g/tls-trunk   master 
 ❯ curl -H 'Host: secret.theknown.net' localhost:8080/~ian/                                                                                                                                                                       [12:36:06]
curl: (7) Failed to connect to localhost port 8080: Connection refused

 ✘ ian@wintermute  ~/g/s/g/g/tls-trunk   master 
 ❯ ./tls-trunk                                                                                                                                                                                                                    [12:36:10]
Listening on :8080
^Z
[1]  + 12192 suspended  ./tls-trunk

 ✘ ⚙ ian@wintermute  ~/g/s/g/g/tls-trunk   master 
 ❯ bg %1                                                                                                                                                                                                                          [12:36:15]
[1]  + 12192 continued  ./tls-trunk

 ⚙ ian@wintermute  ~/g/s/g/g/tls-trunk   master 
 ❯ curl -H 'Host: secret.theknown.net' localhost:8080/~ian/                                                                                                                                                                       [12:36:19]
cant see this
```
