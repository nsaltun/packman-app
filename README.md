# packman-app
Package Manager UI

## Overview
A web-based UI for the Packman package manager, built with Go and deployed on Heroku.

## Ready to Use Link (Heroku)
- Heroku: https://packman-app-e1f16a2cdc2f.herokuapp.com/

## Prerequisites
- Go 1.25
- Docker (for local testing)
- Heroku CLI
- Git

## Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/nsaltun/packman-app.git
   cd packman-app
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env and set PACKMAN_API_BASE_URL to your API endpoint
   ```

3. **Run locally**
   ```bash
   make run
   ```
   The app will be available at http://localhost:8080

4. **Test with Docker**
   ```bash
   # run
   make docker-up
   # stop and remove container
   make docker-stop
   ```

## Deployment to Heroku

### Initial Setup

1. **Create a Heroku app**
   ```bash
   heroku create your-app-name
   ```

2. **Set the stack to container**
   ```bash
   heroku stack:set container -a your-app-name
   ```

3. **Configure environment variables**
   ```bash
   # Set your API URL (replace with your actual Heroku API URL)
   heroku config:set PACKMAN_API_BASE_URL=https://your-api-name.herokuapp.com -a your-app-name
   ```

### Deploy via GitHub Integration

1. **Connect your GitHub repository to Heroku**
   - Go to your app in the Heroku Dashboard (https://dashboard.heroku.com)
   - Click on the "Deploy" tab
   - Under "Deployment method", select "GitHub"
   - Search for and connect to your `packman-app` repository
   
2. **Enable Automatic Deploys (optional)**
   - Under "Automatic deploys", choose your branch (usually `main`)
   - Click "Enable Automatic Deploys"
   
3. **Manual Deploy**
   - Under "Manual deploy", select your branch
   - Click "Deploy Branch"

### Deploy via Git Push

Alternatively, you can deploy directly via Git:

```bash
# Add Heroku remote (if not already added)
heroku git:remote -a your-app-name

# Push to deploy
git push heroku main
```

## Environment Variables

Configure these in Heroku (Dashboard → Settings → Config Vars):

- `PACKMAN_API_BASE_URL`: The URL of your deployed API (e.g., `https://your-api-name.herokuapp.com`)
- `PORT`: Automatically set by Heroku (default: 8080)


**Environment-Based Configuration**
The app automatically detects the environment:
- Local development: Uses `http://localhost:8081` by default
- Production: Uses the `PACKMAN_API_BASE_URL` environment variable

## Monitoring

View logs in real-time:
```bash
heroku logs --tail -a your-app-name
```

## Troubleshooting

- **Build fails**: Check `heroku logs` for build errors
- **API calls fail**: Verify `PACKMAN_API_BASE_URL` is set correctly with `heroku config -a your-app-name`
- **App crashes**: Check logs with `heroku logs --tail -a your-app-name`

## Architecture

```
┌─────────────────┐         ┌─────────────────┐
│   packman-app   │ ─────→  │   packman-api   │
│  (Frontend UI)  │  HTTPS  │   (Backend)     │
│  Heroku App     │         │   Heroku App    │
└─────────────────┘         └─────────────────┘
```
