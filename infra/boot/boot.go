package boot

import (
	"cart-manager/infra/dbrepo"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var env_dir_flag string

func Boot() error {
	if err := parseFlags(); err != nil {
		return err
	}

	if err := initEnvFile(".env"); err != nil {
		return err
	}

	if err := InitLog(env_dir_flag); err != nil {
		return err
	}

	if err := dbrepo.NewDBConn(); err != nil {
		return err
	}
	return nil
}

func parseFlags() error {
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	flags.StringVar(&env_dir_flag, "env_dir", ".", "")

	if err := flags.Parse(os.Args[1:]); err != nil {
		return err
	}

	return nil
}

func initEnvFile(fileName string) error {
	fullPath := filepath.Join(env_dir_flag, fileName)
	err := godotenv.Load(fullPath)
	if err != nil {
		log.Print(err)
		return fmt.Errorf("error loading .env file")
	}

	return nil
}

func InitLog(baseDir string) (err error) {
	logDirPath := os.Getenv("LOG_DIR_PATH")
	if logDirPath == "" {
		logrus.SetOutput(os.Stdout)
		return nil
	}

	if !filepath.IsAbs(logDirPath) {
		logDirPath = filepath.Join(baseDir, logDirPath)
	}

	logDirPath, err = filepath.Abs(logDirPath)
	if err != nil {
		return err
	}

	logFileName := fmt.Sprintf("%d", os.Getpid())

	logFilePath, err := filepath.Abs(filepath.Join(logDirPath, logFileName))
	if err != nil {
		return err
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0666)
	if err != nil {
		return err
	}

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	logrus.SetFormatter(formatter)
	logrus.SetOutput(io.MultiWriter(os.Stdout, file))
	logrus.SetLevel(logrus.InfoLevel)

	return nil
}
