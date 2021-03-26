![bitemyshinymetalGPT-2](https://user-images.githubusercontent.com/58280110/112588421-5b081300-8dcd-11eb-8a9b-eb0c979e1e91.png)

## BiteMyShinyGPT-2

Harness the power of GPT-2 to create Futurama scripts from another dimension!

Scripts for our dataset were scraped from imdbs.com using the golang package Colly.

-----------------------

After cloning the repository, you can construct the docker image by navigating into the gpt-2 folder and running 

docker build -f Dockerfile.cpu . -t 'your_cool_title:latest'

With this, you can access and run gpt-2 by running 

docker run -it --rm your_cool_title


