# Use the official Node.js image as a base
FROM node:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY package*.json ./

# Install dependencies
RUN npm install

# Install typescript-related dev dependencies globally if necessary
RUN npm install -g nodemon ts-node typescript

# Copy the rest of the application source code
COPY . .

# Expose the port your application runs on and the debug port
EXPOSE 3000 9229

# Start the application with nodemon in debug mode
CMD ["npm", "run", "start"]
