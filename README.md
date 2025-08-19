# Mood Tracker

A full-stack mood tracking application built with React, Go, and SQLite.

## 🏗️ Architecture

- **Frontend**: React with Vite (TypeScript) - served by Nginx
- **Backend**: Go server with Gorilla Mux
- **Database**: SQLite
- **Containerization**: Docker Compose

## 🚀 Quick Start

### Prerequisites

- Docker
- Docker Compose

### 1. Clone and Setup

```bash
git clone <your-repo>
cd mood-tracker
```

### 2. Environment Variables

Create a `.env` file in the root directory:

```env
HOST=localhost
PORT=8080
```

### 3. Run the Complete Stack

```bash
# Build and start all services
docker-compose up --build

# Or run in detached mode
docker-compose up -d --build
```

### 4. Access Your Application

- **Frontend**: http://localhost (port 80)
- **Backend API**: http://localhost:8080
- **API Endpoint**: http://localhost:8080/api/moods/{user_id}

## 🔧 Development

### Frontend Development

The React app is built and served by Nginx:

- **Build**: `npm run build` creates optimized production files
- **Serving**: Nginx serves static files efficiently
- **Routing**: React Router works with Nginx configuration

### Backend Development

For Go code changes:

```bash
# Rebuild and restart the backend
docker-compose up --build backend

# Or restart just the backend service
docker-compose restart backend
```

### Database

The SQLite database is automatically initialized with the schema from `dbo/mood.sql` and stored in the `./data` directory.

### Adding New Dependencies

```bash
# Frontend
docker-compose exec frontend sh
npm install <package-name>

# Backend
docker-compose exec backend sh
go get <package-name>
```

## 📁 Project Structure

```
mood-tracker/
├── frontend/          # React application (built and served by Nginx)
├── server/            # Go backend server
├── dbo/              # Database schema
├── data/             # SQLite database files (created by Docker)
├── docker-compose.yml # Docker services configuration
└── README.md
```

## 🛠️ Useful Commands

```bash
# View logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f backend
docker-compose logs -f frontend

# Stop all services
docker-compose down

# Rebuild and restart
docker-compose up --build

# Access container shell
docker-compose exec backend sh
docker-compose exec frontend sh
```

## 🔍 Troubleshooting

### Port Already in Use

If ports 80 or 8080 are already in use, modify the `docker-compose.yml` file to use different ports.

### Database Issues

If the database isn't initializing properly:

```bash
# Remove existing data and restart
docker-compose down -v
docker-compose up --build
```

### Frontend Build Issues

If the React app isn't building properly:

```bash
# Check the build logs
docker-compose logs frontend

# Rebuild the frontend
docker-compose up --build frontend
```

## 📝 API Endpoints

- `GET /api/moods/{user_id}` - Get all moods for a user
- `POST /api/moods` - Create a new mood (to be implemented)
- `PUT /api/moods/{id}` - Update a mood (to be implemented)
- `DELETE /api/moods/{id}` - Delete a mood (to be implemented)

## 🚀 Production Benefits

- **Nginx**: Efficient static file serving with gzip compression
- **Multi-stage builds**: Smaller production images
- **Security headers**: XSS protection, content type validation
- **Caching**: Static assets cached for 1 year
- **SPA routing**: Proper React Router support

## 🤝 Contributing

1. Make your changes
2. Test with Docker Compose
3. Commit and push your changes

## 📄 License

[Your License Here]
