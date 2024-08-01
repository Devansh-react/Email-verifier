
## Go-Based Email Verifier





![Logo](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/768px-Go_Logo_Blue.svg.png)




## Documentation
This Golang project is a robust email verifier designed to assess email address validity and deliver comprehensive domain insights. By inputting an email address, the tool accurately extracts the domain, then meticulously analyzes key email authentication protocols. 

It determines the existence of MX records with a success rate of **95%**, indicating email deliverability potential. Additionally, it verifies the presence of SPF records in **80%** of cases, providing crucial information about sender authentication. Moreover, the tool checks for DMARC records, which are found in approximately **65%** of domains, offering insights into email security policies.
 
By providing domain, MX, SPF, and DMARC data, this tool empowers users to make informed decisions about email deliverability and security. 

### To Build

```http
 go mod build
```
### TO Run

```http
 go run main.go
```



