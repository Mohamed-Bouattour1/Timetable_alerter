package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zhashkevych/scheduler"
)

func main() {
	logrus.Info("Scheduler démarré...")

	ctx := context.Background()
	sc := scheduler.NewScheduler()
	sc.Add(ctx, recoverAndSendEvents, time.Second*30)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	sc.Stop()
	logrus.Info("⏹️ Scheduler arrêté.")
}
func recoverAndSendEvents(ctx context.Context) {
	logrus.Info("🔁 Lancement de recoverAndSendEvents()...")
	// Ici on fera : appel API Config + récupération + parsing + envoi NATS
}
