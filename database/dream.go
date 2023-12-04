package database

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/rajesh818/BookMyShow/config"
)

func Connect() (*gorm.DB, error) {
    dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:%v)/%v?parseTime=true", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_PORT, config.MYSQL_DATABASE)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}