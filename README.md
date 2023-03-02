
## hostedby

Hostedby is a simple (<a href="https://hostedby-hwphkwyqoq-uc.a.run.app/findip?ip=52.93.153.170">REST API</a>) written in GoLang that takes an IP address in string format and searches for it in a SQLite database.  If it finds the IP Address in the database it will return a json response. The example response shows an AWS IP address having been found in the database.

```
{
   "Net":"52.93.153.170/32",
   "Start_ip":878549418,
   "End_ip":878549418,
   "Url":"https://ip-ranges.amazonaws.com/ip-ranges.json",
   "Cloudplatform":"aws",
   "Iptype":"IPv4",
   "Error":"None"
 }
```

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With
- Golang
- SQLite
<p align="right">(<a href="#top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started Locally
The following instructions list how to install the API locally using docker and docker-compose.

### Installation

1. Install [Docker] (https://docs.docker.com/get-docker/)
2. Install [Docker-compose] (https://docs.docker.com/compose/install/)
2. Using a shell cli clone the repo
   ```sh
   git clone https://github.com/stclaird/hostedby.git
   ```
3. Run Docker-compose
   ```sh
   cd hostedby
   docker-compose up
   ```
4. Test the API - curl example shown, although a web browser or postman should work also.
   ```sh
   curl http://localhost:8080/findip?ip=52.93.153.170
   ```

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

