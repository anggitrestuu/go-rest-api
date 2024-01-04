package paginate

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// FilterOperator type for defining filter operators
type FilterOperator string

// Define constants for FilterOperator
const (
	EQ   FilterOperator = "$eq"
	LIKE FilterOperator = "$like"
	IN   FilterOperator = "$in"
	BTW  FilterOperator = "$btw"
)

// Set of valid operators
var validOperators = map[FilterOperator]bool{
	EQ: true, LIKE: true, IN: true, BTW: true,
}

type FilterComparator string

const (
	AND FilterComparator = "AND"
	OR  FilterComparator = "OR"
)

// Set of valid comparators
var validComparators = map[FilterComparator]bool{
	AND: true, OR: true,
}

// Filter represents a filter
type Filter struct {
	Value      interface{}      `json:"value"`
	Comparator FilterComparator `json:"comparator"`
}

// ApplyFilters applies filters to a query
func ApplyFilters(db *gorm.DB, filters []Filter) *gorm.DB {
	if len(filters) == 0 {
		return db
	}

	for _, filter := range filters {
		if filter.Comparator == AND {
			db = db.Where(filter.Value)
		} else {
			db = db.Or(filter.Value)
		}
	}

	return db
}

func ParseFilterFromString(filterString string) ([]Filter, error) {
	if filterString == "" {
		return nil, nil
	}

	var filters []Filter
	var currentComparator FilterComparator = AND // Default comparator

	filterParts := strings.Split(filterString, ";")
	for _, part := range filterParts {
		if part == "" {
			continue
		}

		// Handle comparator
		if strings.HasPrefix(part, "$or:") || strings.HasPrefix(part, "$and:") {
			currentComparator = parseComparator(part)
			strComparator := ""
			if currentComparator == OR {
				strComparator = "$or:"
			} else {
				strComparator = "$and:"
			}

			part = strings.TrimPrefix(part, strComparator)
			if part == "" {
				continue
			}
		}

		// Split into operator and value
		operatorValue := strings.SplitN(part, ":", 2)
		if len(operatorValue) != 2 {
			return nil, fmt.Errorf("invalid filter format: %s", part)
		}

		operator := operatorValue[0]
		value := operatorValue[1]

		// Validate operator
		if _, exists := validOperators[FilterOperator(operator)]; !exists {
			return nil, fmt.Errorf("invalid operator: %s", operator)
		}

		// Create SQL-like value string
		filterValue, err := createSQLLikeValue(operator, value)
		if err != nil {
			return nil, err
		}

		// Skip adding filter if value is empty
		if filterValue == "" {
			continue
		}

		filters = append(filters, Filter{
			Value:      filterValue,
			Comparator: currentComparator,
		})
	}

	return filters, nil
}

func parseComparator(part string) FilterComparator {
	if strings.HasPrefix(part, "$or:") {
		return OR
	}
	return AND
}
func createSQLLikeValue(operator string, value string) (string, error) {
	// Skip creating a filter if the value is empty
	if value == "" {
		return "", nil
	}

	switch operator {
	case "$eq":
		parts := strings.SplitN(value, "=", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid format for $eq: %s", value)
		}
		return fmt.Sprintf("%s = '%s'", parts[0], parts[1]), nil
	case "$like":
		parts := strings.SplitN(value, "=", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid format for $like: %s", value)
		}
		return fmt.Sprintf("%s LIKE '%%%s%%'", parts[0], parts[1]), nil
	case "$in":
		parts := strings.SplitN(value, "=", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid format for $in: %s", value)
		}
		return fmt.Sprintf("%s IN (%s)", parts[0], parts[1]), nil
	case "$btw":
		parts := strings.SplitN(value, "=", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid format for $btw: %s", value)
		}
		rangeParts := strings.Split(parts[1], ",")
		if len(rangeParts) != 2 {
			return "", fmt.Errorf("invalid range for $btw: %s", parts[1])
		}
		return fmt.Sprintf("%s BETWEEN %s AND %s", parts[0], rangeParts[0], rangeParts[1]), nil
	default:
		return "", fmt.Errorf("unsupported operator: %s", operator)
	}
}
