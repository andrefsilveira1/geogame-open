FROM node:14-alpine as builder

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

ARG REACT_APP_API_URL
ENV REACT_APP_API_URL=$REACT_APP_API_URL

RUN npm run build

RUN npm install -g serve

FROM node:14-alpine

RUN npm install -g serve

COPY --from=builder /app/build /app

EXPOSE 3000

CMD ["serve", "-s", "app", "-l", "3000"]
