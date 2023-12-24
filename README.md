# certifizer

Certifizer is super handy when you've a bunch of applications running on different ports yet you want all of them behind SSL certs. This has been motivated by a scenario where one of the customer has multiple services - FE and BE services in a single VM but all the traffic needs to be encrypted. 


```
go get github.com/sivaramsajeev/certifizer
./certifizer
```



# Configuration
Config file: ~/certifizer.yml

```
cat ~/certifizer.yml

domain: ssajeev.hopto.org
email: certifizer@gmail.com
# offset: 7

ports:
- 3000
- 5000
```





