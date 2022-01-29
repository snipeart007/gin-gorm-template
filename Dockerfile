FROM golang

ENV POSTGRES_USER=book-store
ENV POSTGRES_PASSWORD=meow.meow@meow
ENV POSTGRES_DB=book-store

WORKDIR /app
COPY . .

CMD [ "go", "run", "./main.go" ]
