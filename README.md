## A simple rest api built with go

```go
  r.GET("/books", controllers.FindBooks)
  r.POST("/books", controllers.CreateBook)
  r.GET("/books/:id", controllers.FindBook)
  r.PATCH("/books/:id", controllers.UpdateBook)
  r.DELETE("/books/:id", controllers.DeleteBook)
```