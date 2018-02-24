#Creating Simple Web Server
This is a simple web server, the main.go file is as simple as the following code. 
```
templates := populateTemplates()
controller.Startup(templates)
http.ListenAndServe(":8000", nil)
```