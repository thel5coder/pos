package domain

import (
	"database/sql"
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/go-redis/redis/v7"
	jwtFiber "github.com/gofiber/jwt/v2"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"majoo-test/pkg/functioncaller"
	"majoo-test/pkg/jwe"
	"majoo-test/pkg/jwt"
	"majoo-test/pkg/logruslogger"
	db "majoo-test/pkg/postgresql"
	redisPkg "majoo-test/pkg/redis"
	"majoo-test/pkg/str"
	"os"
)

type Config struct {
	DB                       *sql.DB
	JweCredential            jwe.Credential
	JwtCredential            jwt.JwtCredential
	JwtConfig                jwtFiber.Config
	Validator                *validator.Validate
	GrpcClientSvcProductConn *grpc.ClientConn
	Redis                    redisPkg.RedisClient
}

var (
	ValidatorDriver *validator.Validate
	Uni             *ut.UniversalTranslator
	Translator      ut.Translator
)

func LoadConfig() (res Config, err error) {
	err = godotenv.Load("../../.env")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "load-env")
	}

	//setup db connection
	dbInfo := db.Connection{
		Host:                    os.Getenv("DB_HOST"),
		DbName:                  os.Getenv("DB_NAME"),
		User:                    os.Getenv("DB_USER"),
		Password:                os.Getenv("DB_PASSWORD"),
		Port:                    os.Getenv("DB_PORT"),
		SslMode:                 os.Getenv("DB_SSL_MODE"),
		DBMaxConnection:         str.StringToInt(os.Getenv("DB_MAX_CONNECTION")),
		DBMAxIdleConnection:     str.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		DBMaxLifeTimeConnection: str.StringToInt(os.Getenv("DB_MAX_LIFETIME_CONNECTION")),
	}
	res.DB, err = dbInfo.DbConnect()
	if err != nil {
		panic(err)
	}

	//redis
	redisOption := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	res.Redis = redisPkg.RedisClient{Client: redis.NewClient(redisOption)}
	pong, err := res.Redis.Client.Ping().Result()
	fmt.Println("Redis ping status: "+pong, err)

	//jwe credential
	res.JweCredential = jwe.Credential{
		KeyLocation: os.Getenv("JWE_PRIVATE_KEY"),
		Passphrase:  os.Getenv("JWE_PRIVATE_KEY_PASSPHRASE"),
	}

	//jwt credential
	res.JwtCredential = jwt.JwtCredential{
		TokenSecret:         os.Getenv("SECRET"),
		ExpiredToken:        str.StringToInt(os.Getenv("TOKEN_EXP_TIME")),
		RefreshTokenSecret:  os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiredRefreshToken: str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")),
	}

	//jwt config
	res.JwtConfig = jwtFiber.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		Claims:     &jwt.CustomClaims{},
	}
	ValidatorInit()
	res.Validator = ValidatorDriver

	return res, err
}

func ValidatorInit() {
	en := en.New()
	id := id.New()
	Uni = ut.New(en, id)

	transEN, _ := Uni.GetTranslator("en")
	transID, _ := Uni.GetTranslator("id")

	ValidatorDriver = validator.New()

	err := enTranslations.RegisterDefaultTranslations(ValidatorDriver, transEN)
	if err != nil {
		fmt.Println(err)
	}

	err = idTranslations.RegisterDefaultTranslations(ValidatorDriver, transID)
	if err != nil {
		fmt.Println(err)
	}
	switch os.Getenv("APP_LOCALE") {
	case "en":
		Translator = transEN
	case "id":
		Translator = transID
	}
}
