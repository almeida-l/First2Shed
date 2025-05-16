package core

type Notification struct {
	Name string
	Data any
}

type NotificationQueue []Notification

func (q *NotificationQueue) Push(item Notification) {
	*q = append((*q), item)
}

func (q *NotificationQueue) Pop() (Notification, bool) {
	if len(*q) == 0 {
		return Notification{}, false
	}
	notification := (*q)[0]
	*q = (*q)[1:]
	return notification, true
}
