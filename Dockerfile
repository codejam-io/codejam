FROM library/node:20-alpine3.20 as frontend_builder

WORKDIR /app

COPY ./html/package.json /app
COPY ./html/package-lock.json /app

RUN npm install

COPY ./html /app

RUN npm run build

RUN mkdir /app/dist/static/
RUN cp --recursive /app/static/. /app/dist/static/

FROM library/golang:1.22.3-alpine3.20

WORKDIR /app

COPY ./backend/app/go.mod /app
COPY ./backend/app/go.sum /app

RUN go mod download

COPY ./backend/app/ /app
COPY --from=frontend_builder /app/dist/ /app/server/static_files/

ENTRYPOINT [ "go", "run", "./main.go" ]
