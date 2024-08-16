package domainError

import "errors"

var (
	Err_NAME_NOT_NIL_OR_EMPTY                  = errors.New("name must not be nil or empty")
	Err_CONTENT_NOT_NIL_OR_EMPTY               = errors.New("content must not be nil or empty")
	Err_CONTACTS_NOT_EMPTY                     = errors.New("contacts must not be empty")
	Err_PROPERTIES_NOT_LESS_THAN_OR_EQUAL_ZERO = errors.New("properties %s can't be less than or equal zero")
	Err_PROPERTIES_PERPAGE_NOT_EMPTY           = errors.New("properties perPage can't be zero")
	Err_PARAMETER_ID_NOT_EMPTY                 = errors.New("parameter id must not be nil or empty")
	Err_PARAMETER_NOT_EMPTY                    = errors.New("parameter %s must not be nil or empty")
	Err_CAMPAIGN_NOT_FOUND                     = errors.New("campaign %s not found")
)
