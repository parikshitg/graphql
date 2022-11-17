# GraphQL and Go

This project consist of a simple example that demonstrates how to write a graphql server in Go. It consists of a graphql server with minimal go code and is a written using "github.com/99designs/gqlgen" package.

### To run the server

```
make build or go run server.go
```

### Schema Update
You can also update the schema according to your needs. You just have to update the schema in `graph/schema.graphqls` file.
Once the schema is updated you can generate the updated resolvers using:

```
go run github.com/99designs/gqlgen generate
```


### Example Queries
You can copy/paste them at __http://localhost:8080/__ and try & test the graphql server on browser.

#### 1. Get a list of blogs.

```
query{
  blogList{
    id
    title
    body
    author{
      id
      firstname
      lastname
      nickname
    }
    date
  }
}
```

#### 2. Create a Blog

```
mutation CreateBlog($input: BlogInput!){
  createBlog(input:$input){
    id
    title
    date
    status
    author{
      firstname
      lastname
    }
  }
}
```
You can pass the below variable input for this mutation.

```
{
  "input": {
    "title": "Hello World",
    "body": "Some long string",
    "date": "2022-01-09T12:22:20Z",
    "author": {
      "id": 1,
      "firstname": "Jon",
      "lastname": "Doe",
      "nickname": "jondoe"
    }
  }
}
```