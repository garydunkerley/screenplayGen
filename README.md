![bitemyshinymetalass](https://user-images.githubusercontent.com/58280110/112587816-48410e80-8dcc-11eb-9c06-7ede87a9c49d.png)

Use GPT-2 to produce Futurama scripts from another dimension!

Scripts for our dataset were scraped from imdbs.com using the golang package Colly.

-----------------------

After cloning the repository, you can construct the docker image by navigating into the gpt-2 folder and running 

docker build -f Dockerfile.cpu . -t 'your_cool_title:latest'

With this, you can access and run gpt-2 by running 

docker run -it --rm your_cool_title


