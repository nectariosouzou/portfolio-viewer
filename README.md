# portfolio Viewer
Built a little dashboard to break down stocks by GICS sectors. Uses ChatGPT to determine stocks' classification. Check `portfolioTest.csv` for example input data.

<img width="2233" alt="Screenshot 2023-08-04 at 12 45 34 PM" src="https://github.com/nectariosouzou/portfolio-viewer/assets/88638503/493d7607-9320-45c9-b56a-3f2bf97e9f26">

# Set Up
## Backend
Set `.env` variables: 
```
API_KEY='ENTER OPENAI API KEY HERE'
REDIS_ADDR=redis:6379
```
Next from `/backend` directory run `make build-backend` which will build and run redis service and the golang server.

## Frontend
From the `\frontend\portfolio-viewer` directory run `make build-dev` to build and run the NextJs app.
