package httpapi

import (
	"log"
	"pick-up/backend/internal/db"
	"pick-up/backend/internal/models"
	"time"
)

func simulateOrderProgress(order *models.Order) {
	statusRoadmap := []models.OrderStatus{
		models.OrderStatus(models.OrderStatusCreated),
		models.OrderStatus(models.OrderStatusAccepted),
		models.OrderStatus(models.OrderStatusPreparing),
		models.OrderStatus(models.OrderStatusReady),
		models.OrderStatus(models.OrderStatusPicked),
		models.OrderStatus(models.OrderStatusFinished),
	}
	intervals := []time.Duration{
		10 * time.Second,
		10 * time.Second,
		10 * time.Second,
		10 * time.Second,
		10 * time.Second,
		10 * time.Second,
	}
	for i := 0; i < len(statusRoadmap); i++ {
		if statusRoadmap[i] != order.Status {
			continue
		}
		time.Sleep(intervals[i])

		// Update order's status to next one
		log.Printf("Order %d status -> %s", order.ID, statusRoadmap[i+1])
		order.Status = statusRoadmap[i+1]
		if err := db.Conn().Save(order).Error; err != nil {
			log.Print(err)
		}
		if order.Status == models.OrderStatus(models.OrderStatusFinished) {
			break // order is finished, stop the loop
		}
	}
}
