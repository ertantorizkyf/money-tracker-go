package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/initializers"

	"github.com/redis/go-redis/v9"
)

type TransactionRedisRepository struct {
	RedisClient *redis.Client
}

func NewTransactionRedisRepository() *TransactionRedisRepository {
	return &TransactionRedisRepository{
		RedisClient: initializers.RedisClient,
	}
}

func (r *TransactionRedisRepository) GetSummaryByUserAndPeriod(ctx context.Context, userID uint, period string) (bool, dto.TransactionSummaryData, error) {
	response := dto.TransactionSummaryData{
		Period: period,
	}

	key := fmt.Sprintf("user:%d:trx_summary:%s", userID, period)
	values, err := r.RedisClient.HMGet(ctx, key, constants.TRANSACTION_TYPE_INCOME, constants.TRANSACTION_TYPE_EXPENSE).Result()
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return false, response, err
	}

	// CHECK KEY EXISTENCE
	if values[0] == nil || values[1] == nil {
		return false, response, nil
	}

	// 1ST IDX IS INCOME, 2ND IDX IS EXPENSE
	if val, ok := values[0].(string); ok && val != "" {
		income, err := strconv.ParseFloat(val, 64)
		if err != nil {
			helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, fmt.Errorf("invalid income value: %v", err))
			return false, response, fmt.Errorf("invalid income value: %v", err)
		}

		response.IncomeAmount = income
	}

	if val, ok := values[1].(string); ok && val != "" {
		expense, err := strconv.ParseFloat(val, 64)
		if err != nil {
			helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, fmt.Errorf("invalid expense value: %v", err))
			return false, response, fmt.Errorf("invalid expense value: %v", err)
		}

		response.ExpenseAmount = expense
	}

	return true, response, nil
}

func (r *TransactionRedisRepository) SetSummaryByUserAndPeriod(ctx context.Context, userID uint, period string, summary dto.TransactionSummaryData) error {
	key := fmt.Sprintf("user:%d:trx_summary:%s", userID, period)

	fields := map[string]interface{}{
		constants.TRANSACTION_TYPE_INCOME:  fmt.Sprintf("%.2f", summary.IncomeAmount),
		constants.TRANSACTION_TYPE_EXPENSE: fmt.Sprintf("%.2f", summary.ExpenseAmount),
	}

	if err := r.RedisClient.HMSet(ctx, key, fields).Err(); err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	return nil
}
