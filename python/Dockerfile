# Use an official Python image
FROM python:3.10-slim

# Set the working directory in the container
WORKDIR /app

# Install dependencies
COPY requirements.txt .
RUN pip install -r requirements.txt

# Install debugpy for debugging
RUN pip install debugpy

# Copy the FastAPI app code into the container
COPY . .

# Expose the app and debug ports
EXPOSE 8888 5678

# Command to run FastAPI with Uvicorn and debugpy
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8888", "--reload"]
