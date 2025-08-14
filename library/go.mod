module github.com/library

go 1.23.0

toolchain go1.23.8

require (
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
	github.com/caarlos0/env/v10 v10.0.0
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/microcosm-cc/bluemonday v1.0.27
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.26.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/go-sql-driver/mysql v1.9.2 // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/text v0.24.0 // indirect
)

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
