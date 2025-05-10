package helpers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
)

func IsEmailValid(email string) bool {
	// Must be email format with @ symbol
	emailRegex := regexp.MustCompile(`^[\w\.-]+@[\w\.-]+\.\w{2,}$`)
	return emailRegex.MatchString(email)
}

func IsPhoneValid(phone string) bool {
	// Must all be numbers and start with non zero digits (assumed as country codes)
	phoneRegex := regexp.MustCompile(`^[1-9]\d{5,19}$`)
	return phoneRegex.MatchString(phone)
}

func IsDOBValid(dob string) bool {
	// Must be in YYYY-MM-DD format
	_, err := time.Parse("2006-01-02", dob)
	return err == nil
}

func IsUsernameValid(username string) bool {
	// Must be alphanumeric and underscore only
	// Username must be between 3 to 30 characters
	usernameRegex := regexp.MustCompile(`^[A-Za-z0-9_]{3,30}$`)
	return usernameRegex.MatchString(username)
}

func IsPasswordValid(password string) bool {
	// Must be alphanumeric and special characters with no whitespace
	// Password must be between 8 to 30 characters
	passwordRegex := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_\-+=\[\]{}|\\:;"'<>,.?/~` + "`" + `]{8,30}$`)

	isValid := passwordRegex.MatchString(password)

	// Can't be commonly used passwords
	// * COMMON HASHES LIBRARY DITCHED BECAUSE OF SLOW PERFORMANCE: TESTED ~ >30 MINUTES TO CHECK ALL HASHES IN LIBRARY
	// * Compared with hashed library in assets/common_password_libs_hashed.txt
	// * commonHashes := make(map[string]struct{})
	// Use plain common pass lib instead
	commonPasswords := make(map[string]struct{})
	filePath := os.Getenv("COMMON_PASS_LIB_PATH")
	file, err := os.Open(filePath)
	if err != nil {
		isValid = false
		LogWithSeverity("ERR", fmt.Sprintf("Failed to open file %s: %v", filePath, err))
		return isValid
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password := strings.TrimSpace(scanner.Text())
		if password != "" {
			commonPasswords[password] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		isValid = false
		LogWithSeverity("ERR", fmt.Errorf("error reading file: %w", err))
		return isValid
	}

	// * COMMON HASHES LIBRARY DITCHED BECAUSE OF SLOW PERFORMANCE: TESTED ~ >30 MINUTES TO CHECK ALL HASHES IN LIBRARY
	// * Compared with hashed library in assets/common_password_libs_hashed.txt
	// * for hashed := range commonHashes {
	// *   if CheckPasswordHash(password, hashed) {
	// * 	   // Password is common
	// * 	   isValid = false
	// * 	   return isValid
	// *   }
	// * }

	// Use plain common pass lib instead
	if _, exists := commonPasswords[password]; exists {
		isValid = false
		return isValid
	}

	return isValid
}

func ValidateRegisterReq(req dto.RegisterReq) (bool, string) {
	// VALIDATE EMAIL
	isEmailValid := IsEmailValid(req.Email)
	if !isEmailValid {
		return false, "Invalid email"
	}

	// VALIDATE PHONE
	isPhoneValid := IsPhoneValid(req.Phone)
	if !isPhoneValid {
		return false, "Invalid phone"
	}

	// VALIDATE USERNAME
	isUsernameValid := IsUsernameValid(req.Username)
	if !isUsernameValid {
		return false, "Invalid username"
	}

	// VALIDATE DOB
	isDOBValid := IsDOBValid(req.DOB)
	if !isDOBValid {
		return false, "Invalid DOB"
	}

	// VALIDATE PASSWORD
	isPasswordValid := IsPasswordValid(req.Password)
	if !isPasswordValid {
		return false, "Invalid password"
	}

	return true, ""
}

func ValidateLoginReq(req dto.LoginReq) (bool, string) {
	// VALIDATE EMAIL AND USERNAME
	isEmailValid := IsEmailValid(req.UsernameOrEmail)
	isUsernameValid := IsUsernameValid(req.UsernameOrEmail)

	if !isEmailValid && !isUsernameValid {
		return false, "Invalid username or email"
	}

	return true, ""
}

func ValidateTrxType(trxType string) (bool, string) {
	if trxType != constants.TRANSACTION_TYPE_INCOME && trxType != constants.TRANSACTION_TYPE_EXPENSE {
		return false, "Invalid transaction type"
	}

	return true, ""
}

func ValidateTransactionQueryParam(query dto.TransactionQueryParam) (bool, string) {
	// VALIDATE TRX TYPE
	isTrxTypeValid, message := ValidateTrxType(query.Type)
	if !isTrxTypeValid {
		return false, message
	}

	// START DATE AND END DATE MUST BE YYYY-MM-DD
	if query.StartDate != "" {
		_, err := time.Parse("2006-01-02", query.StartDate)
		if err != nil {
			return false, "Invalid date"
		}
	}

	if query.EndDate != "" {
		_, err := time.Parse("2006-01-02", query.EndDate)
		if err != nil {
			return false, "Invalid date"
		}
	}

	// VALIDATE ORDER
	if query.Order != "" {
		if query.Order != constants.TRANSACTION_ORDER_NEWEST && query.Order != constants.TRANSACTION_ORDER_OLDEST {
			return false, "Invalid order"
		}
	}

	return true, ""
}

func ValidateTransactionSummaryQueryParam(query dto.TransactionSummaryQueryParam) (bool, string) {
	// PERIOD MUST BE YYYY-MM
	if query.Period != "" {
		_, err := time.Parse("2006-01", query.Period)
		if err != nil {
			return false, "Invalid period"
		}
	}

	return true, ""
}

func ValidateCreateTransactionRequest(req dto.CreateTransactionRequest) (bool, string) {
	// TYPE MUST BE "income" OR "expense"
	if req.Type != constants.TRANSACTION_TYPE_INCOME && req.Type != constants.TRANSACTION_TYPE_EXPENSE {
		return false, "Invalid transaction type"
	}

	// TRX DATE MUST BE YYYY-MM-DD
	if req.TrxDate != "" {
		_, err := time.Parse("2006-01-02", req.TrxDate)
		if err != nil {
			return false, "Invalid transaction date"
		}
	}

	return true, ""
}

func ValidateUpdateTransactionRequest(req dto.UpdateTransactionRequest) (bool, string) {
	// TYPE MUST BE "income" OR "expense"
	if req.Type != constants.TRANSACTION_TYPE_INCOME && req.Type != constants.TRANSACTION_TYPE_EXPENSE {
		return false, "Invalid transaction type"
	}

	// TRX DATE MUST BE YYYY-MM-DD
	if req.TrxDate != "" {
		_, err := time.Parse("2006-01-02", req.TrxDate)
		if err != nil {
			return false, "Invalid transaction date"
		}
	}

	return true, ""
}
