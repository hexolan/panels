FROM node:18 as build  
WORKDIR /app

# Build options
ARG VITE_API_URL
ENV VITE_API_URL $VITE_API_URL

# Install requirements
COPY package.json yarn.lock ./
RUN yarn install --pure-lockfile

# Copy files and build
COPY . .
RUN yarn run build

# Serve the frontend
FROM nginx:alpine-slim
EXPOSE 80
COPY --from=build /app/dist /usr/share/nginx/html
COPY --from=build /app/nginx.conf /etc/nginx/conf.d/default.conf