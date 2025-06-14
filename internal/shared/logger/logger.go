package logger

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

func Init(cfg config.LogConfig) {
	fileFormatter := &logrus.TextFormatter{}

	fileWriter, _ := rotatelogs.New(
		filepath.Join(cfg.Dir, "%Y-%m-%d.log"),
		rotatelogs.WithMaxAge(time.Duration(cfg.MaxAge)*24*time.Hour),
	)

	logrus.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  fileWriter,
			logrus.WarnLevel:  fileWriter,
			logrus.ErrorLevel: fileWriter,
			logrus.FatalLevel: fileWriter,
			logrus.PanicLevel: fileWriter,
		},
		fileFormatter,
	))

	if environment.IsDev {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})
		logrus.SetOutput(os.Stdout)
	}
}
