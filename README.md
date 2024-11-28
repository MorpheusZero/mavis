# Mavis
A simple and light-weight Web Server that can also be used as a reverse proxy, load balancer, and handle caching. 
It is named after the character  [Mavis Vermillion](https://fairytail.fandom.com/wiki/Mavis_Vermilion) from the anime [Fairy Tail](https://en.wikipedia.org/wiki/Fairy_Tail) that 
excels in strategy and tactics while using illusion magic.

<p align="center">
  <img width="150" src="https://i.pinimg.com/736x/98/83/2b/98832bc44c2fa4137a98cd7542184865.jpg" alt="Description of the image">
</p>

## WARNING
Mavis was built using Golang as a hobby and to learn more about how web servers work. It is not meant to be used in production environments and should be used for educational purposes only.
**Mavis IS NOT Production ready!**

## Features
- [ ] HTTP/2
- [ ] HTTPS support (Let's Encrypt)
- [ ] Reverse Proxy
- [ ] Load Balancer
- [ ] Caching
- [ ] Logging
- [ ] Simple Configuration file
- [ ] Custom error pages
- [ ] Web Admin Interface
- [ ] Rate Limiting
- [ ] Security Headers

## Running Locally
You will need a `mavis.json` file in the root directory of the project. You can copy the example file for more references, but the minimal file should look like this:
```json
{
    "server-port": "8080",
    "proxy-hosts": [
        {
            "domain": "localhost",
            "protocol": "http",
            "host": "10.0.0.4",
            "port": "8090"
        }
    ]
}

```
To run the server locally, you can use the following command:
```bash
go run main.go
```