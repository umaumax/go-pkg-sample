# how to run

	go run *.go

# how to test
```
# Unauthorized
curl http://localhost:5555/
curl http://user:pass@localhost:5555/
curl http://user:pass@localhost:5555/index.html
# not found
curl http://user:pass@localhost:5555/assets
curl http://user:pass@localhost:5555/assets/
curl http://user:pass@localhost:5555/assets/css/app.css
curl http://user:pass@localhost:5555/assets/js/app.js
# empty name error
curl http://user:pass@localhost:5555/api/hello/
curl http://user:pass@localhost:5555/api/hello/nanoha
```
