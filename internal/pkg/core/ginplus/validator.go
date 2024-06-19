package ginplus

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"sync"
)

const (
	fieldErrMsg = "%s=%s"
)

// Errors 多字段验证错误类型
type Errors []FieldError

// FieldError 字段验证错误封装
type FieldError struct {
	validator.FieldError
	tagMsg string
	errMsg string
}

func (ve Errors) Error() string {

	buff := bytes.NewBufferString("")

	var fe FieldError

	buff.WriteString("字段校验错误：[")
	for i := 0; i < len(ve); i++ {

		fe = ve[i]
		msg := fmt.Sprintf(fieldErrMsg, fe.Field(), fe.Tag())

		if i != 0 {
			buff.WriteString(",")
		}
		buff.WriteString(msg)
	}
	buff.WriteString("]")

	return strings.TrimSpace(buff.String())
}

// FieldErrorMsg 字段错误信息
type FieldErrorMsg struct {
	Field   string `json:"field"`
	Msg     string `json:"msg"`
	ErrType string `json:"err_type"`
}

// ErrorData 获取错误 ErrorData 对象
func (ve Errors) ErrorData(trans ut.Translator) map[string]*FieldErrorMsg {

	data := make(map[string]*FieldErrorMsg, len(ve))

	var fe FieldError

	for i := 0; i < len(ve); i++ {
		fe = ve[i]
		msg := fe.tagMsg

		if msg == "" {
			msg = fe.Translate(trans)
		}

		data[fe.Field()] = &FieldErrorMsg{
			Field:   fe.Field(),
			ErrType: fe.Tag(),
			Msg:     msg,
		}
	}

	return data
}

type customerValidator struct {
	once     sync.Once
	validate *validator.Validate
}

// NewValidator 新建验证器
func NewValidator() binding.StructValidator {
	return &customerValidator{}
}

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *customerValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
		value = value.Elem()
	}
	if valueType == reflect.Struct {
		v.lazyInit()
		if err := v.validate.Struct(obj); err != nil {
			var validatorErrors validator.ValidationErrors
			if errors.As(err, &validatorErrors) {
				l := len(validatorErrors)
				errs := make(Errors, l)
				for i := 0; i < l; i++ {
					fieldError := validatorErrors[i]
					fe := FieldError{
						FieldError: fieldError,
					}

					tag := fieldError.Tag()
					field := fieldError.Field()
					st, err := value.Type().FieldByName(field)
					if err {
						fe.tagMsg = st.Tag.Get(tag)
					}
					errs[i] = fe
				}
				return errs
			}
			return err
		}
	}
	return nil
}

func (v *customerValidator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

func (v *customerValidator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
	})
}
