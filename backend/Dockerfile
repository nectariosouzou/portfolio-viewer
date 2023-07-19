FROM redis:latest

# Set the desired port
EXPOSE 6379

# Set the IP address to listen on
ENV REDIS_BIND 0.0.0.0

# Start Redis with custom configuration
CMD ["redis-server", "--port", "6379", "--bind", "0.0.0.0"]

