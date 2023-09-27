FROM node:20-slim

WORKDIR /app

# Install requirements
COPY package.json package-lock.json ./
RUN npm install

# Copy files
COPY . ./

EXPOSE 9090
CMD ["npm", "run", "start"]