package middle

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"golang.org/x/net/context"
)

func getDBConnectionAddress() string {
	config := GetConnectionInfo()
	return config.User + ":" + config.Pass + "@tcp(" + config.Addr + ":" + strconv.Itoa(config.Port) + ")/" + config.DbName + "?charset=utf8"
}

func GetDBContext() *xorm.Engine {
	db, err := xorm.NewEngine("mysql", getDBConnectionAddress())
	if err != nil {
		panic(err)
	}
	// db.ShowSQL(true)
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(150)

	return db
}

func ConnectionHandler(db *xorm.Engine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session := db.NewSession()
			defer session.Close()

			req := c.Request()
			c.SetRequest(req.WithContext(
				context.WithValue(req.Context(), "DB", session),
			))

			if err := session.Begin(); err != nil {
				log.Println(err)
			}
			if err := next(c); err != nil {
				session.Rollback()
				return err
			}
			if c.Response().Status >= 500 {
				session.Rollback()
				return nil
			}
			if err := session.Commit(); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			return nil
		}
	}
}
