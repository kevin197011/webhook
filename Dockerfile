FROM gcr.io/distroless/base-debian10 as production

# Copy application binary from build/dev stage to the distroless container
COPY webhook /app/

# Application port (optional)
EXPOSE 8000

# Container start command for production
CMD ["/webhook"]
