# Register
curl -X POST http://localhost:8080/auth/register -d "username=JhonDoe123&password=easyPasswordWillBe&email=johnDoe@web.com"

# Login
curl -X POST http://localhost:8080/auth/login -d "username=JhonDoe123&password=easyPasswordWillBe&email=johnDoe@web.com"

# Check User status
curl -X GET http://localhost:8080/auth/user -d "username=JhonDoe123&password=easyPasswordWillBe&email=johnDoe@web.com"
