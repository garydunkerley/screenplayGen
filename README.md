# screenplayGen
A simple implementation of colly to scrape screenplays from television shows and a fork of GPT-2 with modified dockerfiles. 

After cloning the repository, you can construct the docker image by navigating into the gpt-2 folder and running 

docker build -f Dockerfile.cpu . -t 'your_cool_title:latest'

With this, you can access and run gpt-2 by running 

docker run -it --rm your_cool_title


