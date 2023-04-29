package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

// Product struct

type Product struct{
	ID int64 `json:"id"`
	Name string `json:"name"`
	Info string `json:"info,omitempty"`
	Price float64 `json:"price"`
}

// all products
var products = [] Product{
	{
		ID: 1,
		Name: "Fanta",
		Info: " soft drink ",
		Price: 15.00 ,
	},
	{
		ID: 2,
		Name: "coco cola",
		Info : "soft drink",
		Price : 20.00,
	},
	{
		ID: 3,
		Name : "seven- up",
		Info : "soft drink",
		Price: 25.00,
	},
}


// defining schema

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id" : &graphql.Field{
				Type: graphql.Int,
			},
			"name" : &graphql.Field{
                 Type: graphql.String, 
			},
			"info" : &graphql.Field{
                  Type: graphql.String,
			}, 
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)


// query

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"product" : &graphql.Field{
				Type: productType,
				Description: "Get product by Id",
				Args: graphql.FieldConfigArgument{
					"id" : &graphql.ArgumentConfig{
                          Type: graphql.Int,
					},

				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok{
						// Find product
						for _,product := range products{
							if int(product.ID) == id{
								return product,nil
							}
						}
					}
					return nil, nil
				},

			},
			"list" : &graphql.Field{
				Type: graphql.NewList(productType),
				Description: "Get Product List",
				Resolve : func (params graphql.ResolveParams)(interface{}, error)  {
					return products, nil
				},
			},

		},
	})

// for create update delete
var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			// create new product item
			"create": &graphql.Field{
				Type : productType,
				Description: "create new product",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
                        Type: graphql.NewNonNull(graphql.String),
					},
					"info" : &graphql.ArgumentConfig{
						Type: graphql.String,
					}, 
					"price" : &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Float),
					},

				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					rand.Seed(time.Now().UnixNano())
					product := Product{
						ID : int64(rand.Intn(100000)),
						Name : params.Args["name"].(string),
						Info : params.Args["info"].(string),
						Price : params.Args["price"].(float64),
					}
					products = append(products, product)
					return product,nil
				},

			},
			// update product by Id
			"update" : &graphql.Field{
				Type : productType,
				Description: "update by product id",
				Args: graphql.FieldConfigArgument{
					"id" : &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),

					},
					"name": &graphql.ArgumentConfig{
                        Type : graphql.String,
					},
					"info": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},

				},
				Resolve: func (params graphql.ResolveParams) (interface{},error) {
					id , _ := params.Args["id"].(int)
					name , nameOk := params.Args["name"].(string)
					info, infoOk := params.Args["info"].(string)
					price , priceOk := params.Args["price"].(float64)
					product := Product{}
					for i ,p := range products{
						if int64(id) == p.ID{
							if nameOk{
								products[i].Name = name
							}
							if infoOk{
								products[i].Info = info
							
							}
							if priceOk{
								products[i].Price = price 
							}
							product = products[i]
							break
						}
					}
					return product , nil
					
				},

			},

			// Delete product by Id

			"delete": &graphql.Field{
				Type: productType,
				Description: "Delete product by Id",
				Args : graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
                        Type : graphql.NewNonNull(graphql.Int),                               
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id , _ := params.Args["id"].(int)
					product := Product{}
					for i , p := range products{
						if int64(id) == p.ID{
							product = products[i]
							products = append(products[:i],products[i+1:]...)
						}
					}
					return product , nil
				},
			},
		},

})



var schema , _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
		Mutation: mutationType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result{
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,
	})
	if len(result.Errors) > 0{
		fmt.Printf("errors: %v",result.Errors)
	}
	return result
}

func main(){
	http.HandleFunc("/product",func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"),schema)
		json.NewEncoder(w).Encode(result)

	})
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080",nil)
}





















