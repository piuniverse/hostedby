
## About The hostedby

This is a simple API for finding and returning data on specific IPs.

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

* [Fast API](https://fastapi.tiangolo.com/)
* [MongoDB](https://https://www.mongodb.com/)

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
   cd <DIRECTORY WHERE YOU CLONED THE REPO>
   docker-compose up
   ```
4. Test the API - curl example shown your browser or postman works too.
   ```sh
   curl http://http://localhost:9000/ip/52.93.153.170
   ```

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

