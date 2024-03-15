package configs

import "github.com/lpernett/godotenv"


func LoadEnv() error {
    err := godotenv.Load()
    if err != nil {
        return err
    }
    return nil
}

