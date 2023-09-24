package bootstrap

import (
	"log"
	AppEvent "skeleton/internal/event"
	"time"

	"skeleton/app/event"
	"skeleton/app/task"
	"skeleton/internal/config"
	"skeleton/internal/config/driver"
	"skeleton/internal/crontab"
	"skeleton/internal/logger"
	"skeleton/internal/mq"
	"skeleton/internal/redis"
	"skeleton/internal/variable"
	"skeleton/internal/variable/consts"
)

func init() {
	var err error
	if variable.Config, err = config.New(driver.New(), config.Options{
		BasePath: variable.BasePath,
	}); err != nil {
		log.Fatal(consts.ErrorInitConfig)
	}
	if variable.Log, err = logger.New(
		logger.WithDebug(true),
		logger.WithEncode("json"),
		logger.WithFilename(variable.BasePath+"/storage/logs/system.log"),
	); err != nil {
		log.Fatal(consts.ErrorInitLogger)
	}
	if variable.DB, err = InitMysql(); err != nil {
		log.Fatal(consts.ErrorInitDb)
	}
	if err = InitMongo(); err != nil {
		log.Fatal(consts.ErrorInitMongoDb)
	}
	redisConfig := variable.Config.Get("Redis").(map[string]any)
	if redisConfig != nil && !redisConfig["disabled"].(bool) {
		variable.Redis = redis.New(
			redis.WithAddr(redisConfig["addr"].(string)),
			redis.WithPwd(redisConfig["pwd"].(string)),
			redis.WithDb(redisConfig["db"].(int)),
			redis.WithPoolSize(redisConfig["poolsize"].(int)),
			redis.WithMaxIdleConn(redisConfig["maxidleconn"].(int)),
			redis.WithMinIdleConn(redisConfig["minidleconn"].(int)),
			redis.WithMaxLifetime(time.Duration(redisConfig["maxlifetime"].(int))),
			redis.WithMaxIdleTime(time.Duration(redisConfig["maxidletime"].(int))),
		)
	}
	if variable.Config.GetBool("Crontab.Enable") {
		variable.Crontab = crontab.New()
		variable.Crontab.AddFunc(task.New().Tasks()...)
		variable.Crontab.Start()
	}
	if variable.Config.GetBool("MQ.Enable") {
		if variable.MQ, err = mq.New(
			mq.WithNameServers(variable.Config.GetStringSlice("MQ.Servers")),
			mq.WithGroupId(variable.Config.GetString("MQ.GroupId")),
			mq.WithConsumptionSize(variable.Config.GetInt("MQ.ConsumptionSize")),
			mq.WithRetries(variable.Config.GetInt("MQ.Retries")),
		); err != nil {
			log.Fatal(consts.ErrorInitMQ)
		}
	}
	variable.Event = AppEvent.New()
	if err = (&event.Event{}).Init(); err != nil {
		log.Fatal(consts.ErrorInitEvent)
	}
}
