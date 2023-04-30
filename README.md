## What is golang ?

* Go is statically typed, compiled high-level programming language designed at Google .It is syntactically similar to C , But memory safety , garbage collection, structual typing and CSP- style concurrency. It is often referrred to as Golang. Its proper name is go 


## What is GraphQL ?

* GraphQL is a query language for APIs and a runtime for fulfilling those Queries with your existing data.GraphQL provides a complete and understandable description of the data in your API, gives clients the power to ask for exactly what they need and nothing more, makes it easier to evolve APIs over time, and enables powerful developer tools.

## What is graphql-go ?

* An implementation of GraphQL in Go. Follows the official reference implementation graphql-js.
*Supports: queries, mutations & subscriptions.

## What is CRUD ? 
*Create, Read, Update, and Delete (CRUD) are the four basic functions that models should be able to do, at most.

## What is CRUD.go ?

* CRUD is a graphql-api in golang . Where one can query by ID of products. Also create , update , and delete new products .

* Get single product by id 
` http://localhost:8080/product?query={product(id:1){name,info,price}}`

* Get product list 
` http://localhost:8080/product?query={list{id,name,info,price}}`

* Create New product Item

`http://localhost:8080/product?query=mutation+_{create(name:"fizz",info:"moja" ,price:2.99){id,name,info,price}`


* Update Item
`http://localhost:8080/product?query=mutation+_{update(id:1,price:9.95){id,name,info,price}}`

* Delete Item
`http://localhost:8080/product?query=mutation+_{delete(id:1){id,name,info,price}}`
