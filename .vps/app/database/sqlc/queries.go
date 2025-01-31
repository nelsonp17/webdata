package sqlc

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nelsonp17/webdata/app/database/sqlc/schemas"
)

// Table History
func (r Repo) GetLastHistory(money, source, source_web string) (schemas.History, error) {
	arg := schemas.GetLastHistoryParams{
		Money:     money,
		Source:    source,
		SourceWeb: source_web,
	}

	return r.Queries.GetLastHistory(context.Background(), arg)
}
func (r Repo) FindHistory(id int64) (schemas.History, error) {
	return r.Queries.FindHistory(context.Background(), id)
}
func (r Repo) CreateHistory(money, source, source_web string, change float64) (schemas.History, error) {
	arg := schemas.CreateHistoryParams{
		Money:     money,
		Source:    source,
		SourceWeb: source_web,
		Change:    change,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	return r.Queries.CreateHistory(context.Background(), arg)
}
func (r Repo) ListHistory(sourceWeb string, createdAt time.Time, limit int32) ([]schemas.History, error) {
	arg := schemas.ListHistoryParams{
		SourceWeb: sourceWeb,
		Limit:     limit,
		Column2:   pgtype.Timestamp{Time: createdAt, Valid: true},
	}

	return r.Queries.ListHistory(context.Background(), arg)
}

// Table users
func (r Repo) CreateUser(email, password string, createdAt pgtype.Timestamp) (schemas.User, error) {
	arg := schemas.CreateUserParams{
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
	}
	return r.Queries.CreateUser(context.Background(), arg)
}
func (r Repo) GetUser(email string) (schemas.User, error) {
	return r.Queries.GetUser(context.Background(), email)
}
func (r Repo) UpdateUser(email, password string) (schemas.User, error) {
	arg := schemas.UpdateUserParams{
		Email:    email,
		Password: password,
	}
	return r.Queries.UpdateUser(context.Background(), arg)
}
func (r Repo) DeleteUser(email string) error {
	return r.Queries.DeleteUser(context.Background(), email)
}

// Table subscriptions
func (r Repo) CreateSubscription(userId int64, status string, startDate, endDate, createdAt pgtype.Timestamp) (schemas.Subscription, error) {
	arg := schemas.CreateSubscriptionParams{
		CreatedAt: createdAt,
		StartDate: startDate,
		UserID:    userId,
		EndDate:   endDate,
		Status:    status,
	}
	return r.Queries.CreateSubscription(context.Background(), arg)
}
func (r Repo) GetSubscription(userId int64) (schemas.Subscription, error) {
	return r.Queries.GetSubscription(context.Background(), userId)
}
func (r Repo) UpdateSubscription(userId int64, status string, endDate pgtype.Timestamp) (schemas.Subscription, error) {
	arg := schemas.UpdateSubscriptionParams{
		UserID:  userId,
		Status:  status,
		EndDate: endDate,
	}
	return r.Queries.UpdateSubscription(context.Background(), arg)
}
func (r Repo) DeleteSubscription(userId int64) error {
	return r.Queries.DeleteSubscription(context.Background(), userId)
}

// Table payments
func (r Repo) CreatePayment(subscriptionId int64, amount pgtype.Numeric, paymentDate pgtype.Timestamp, paymentMethod string, createdAt pgtype.Timestamp) (schemas.Payment, error) {
	arg := schemas.CreatePaymentParams{
		Amount:         amount,
		CreatedAt:      createdAt,
		PaymentDate:    paymentDate,
		PaymentMethod:  paymentMethod,
		SubscriptionID: subscriptionId,
	}
	return r.Queries.CreatePayment(context.Background(), arg)
}
func (r Repo) GetPayment(subscriptionId int64) ([]schemas.Payment, error) {
	return r.Queries.GetPayment(context.Background(), subscriptionId)
}
func (r Repo) UpdatePayment(subscriptionId int64, amount pgtype.Numeric, paymentDate pgtype.Timestamp, paymentMethod string) (schemas.Payment, error) {
	arg := schemas.UpdatePaymentParams{
		ID:            subscriptionId,
		Amount:        amount,
		PaymentDate:   paymentDate,
		PaymentMethod: paymentMethod,
	}
	return r.Queries.UpdatePayment(context.Background(), arg)
}
func (r Repo) DeletePayment(subscriptionId int64) error {
	return r.Queries.DeletePayment(context.Background(), subscriptionId)
}
