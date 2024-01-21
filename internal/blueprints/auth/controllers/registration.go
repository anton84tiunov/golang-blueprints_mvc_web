package controllers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	crud_user "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/database/crud/user"
	models "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/models"
	services "github.com/anton84tiunov/golang-blueprints_mvc_web/internal/services"
)

func ParserUser(data map[string]interface{}) (models.User_json, error) {
	var err error
	var jsonString []byte
	jsonString, err = json.Marshal(data)
	if err != nil {
		services.L.Err(err)
		return models.User_json{}, err
	}
	fmt.Println("string(jsonString)", string(jsonString))
	s := models.User_json{}
	err = json.Unmarshal(jsonString, &s)
	if err != nil {
		services.L.Err(err)
		return models.User_json{}, err
	}
	return s, err
}

// функция определяет соответствие длинны строки заданным значениям
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func MinMaxLength(data map[string]interface{}, errs map[string]interface{}, value string, min uint8, max uint8, ok *bool) {
	str_data_val, status_data_val := data[value].(string)
	if !status_data_val {
		*ok = false
		services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
	}
	if str_data_val == "" {
		str_errs_val, status_err_val := errs[value].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value] = append(str_errs_val, "Поле обязательно для заполнения")
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false
	}

	if len(str_data_val) < int(min) {
		str_errs_val, status_err_val := errs[value].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value] = append(str_errs_val, fmt.Sprintf("должна быть не менее %d символов", min))
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false
	}
	if len(str_data_val) > int(max) {
		str_errs_val, status_err_val := errs[value].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value] = append(str_errs_val, fmt.Sprintf("должна быть не больше %d символов", max))
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false
	}
}

// функция определяет соответствие строки на содержание только латинмких букв
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func IsLatin(data map[string]interface{}, errs map[string]interface{}, value string, ok *bool) {
	// bool_err := false
	str_data_val, status_data_val := data[value].(string)
	if !status_data_val {
		*ok = false
		services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
	}
	match, match_err := regexp.MatchString("^[a-zA-Z]*$", str_data_val)
	if match_err != nil {
		*ok = false
		services.L.Warn("сравнение строки на соответствие регулярного выражения ^[a-zA-Z]*$ завершилось неудачно")
	}
	if !match {
		str_errs_val, status_err_val := errs[value].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value] = append(str_errs_val, "Только латинсковые символы")
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false
	}

}

// функция определяет соответствие строки на содержание только цифры
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func IsNumber(data map[string]interface{}, errs map[string]interface{}, value string, ok *bool) {
	str_data_val, status_data_val := data[value].(string)
	if !status_data_val {
		*ok = false
		services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
	}
	match, match_err := regexp.MatchString("^[0-9]*$", str_data_val)
	if match_err != nil {
		*ok = false
		services.L.Warn("сравнение строки на соответствие регулярного выражения ^[a-zA-Z]*$ завершилось неудачно")
	}
	if !match {
		str_errs_val, status_err_val := errs[value].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value] = append(str_errs_val, "Только цифры")
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false
	}

}

// функция определяет соответствие строки ЭЛЕКТРОННОЙ ПОЧТЕ
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func IsEmail(data map[string]interface{}, errs map[string]interface{}, value string, ok *bool) {
	str_data_val, status_data_val := data[value].(string)
	if !status_data_val {
		*ok = false
		services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
	}
	match, match_err := regexp.MatchString(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, str_data_val)
	if match_err != nil {
		*ok = false
		services.L.Warn("сравнение строки на соответствие регулярного выражения ^[a-zA-Z]*$ завершилось неудачно")
	}
	if !match {
		str_errs_val, status_err_val := errs[value].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value] = append(str_errs_val, "неправильная почта")
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false
	}

}

// функция определяет соответствие строки паролю
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func IsPassword(data map[string]interface{}, errs map[string]interface{}, value string, ok *bool) {
	str_data_val, status_data_val := data[value].(string)
	if !status_data_val {
		*ok = false
		services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
	}
	match, match_err := regexp.MatchString("[0-9]+", str_data_val)
	if match_err != nil {
		*ok = false
		services.L.Warn("сравнение строки на соответствие регулярного выражения ^[a-zA-Z]*$ завершилось неудачно")
	}
	if !match {
		str_errs_val, status_err_val := errs[value].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value] = append(str_errs_val, "Строка должна содержать хоть одну цифру")
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false
	}
	match, match_err = regexp.MatchString("[a-z]+", str_data_val)
	if match_err != nil {
		*ok = false
		services.L.Warn("сравнение строки на соответствие регулярного выражения ^[a-zA-Z]*$ завершилось неудачно")
	}
	if !match {
		str_errs_val, status_err_val := errs[value].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value] = append(str_errs_val, "Строка должна содержать хоть одну строчную букву")
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false
	}
	match, match_err = regexp.MatchString("[A-Z]+", str_data_val)
	if match_err != nil {
		*ok = false
		services.L.Warn("сравнение строки на соответствие регулярного выражения ^[a-zA-Z]*$ завершилось неудачно")
	}
	if !match {
		str_errs_val, status_err_val := errs[value].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value] = append(str_errs_val, "Строка должна содержать хоть одну заглавную букву")
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false
	}
	match, match_err = regexp.MatchString(`[!@#$&*]+`, str_data_val)
	if match_err != nil {
		*ok = false
		services.L.Warn("сравнение строки на соответствие регулярного выражения ^[a-zA-Z]*$ завершилось неудачно")
	}
	if !match {
		str_errs_val, status_err_val := errs[value].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value] = append(str_errs_val, "Строка должна содержать хоть один спецсимвол @$!%*?&")
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false
	}

}

// функция определяет равенство двух строк
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func IsEquality(data map[string]interface{}, errs map[string]interface{}, value1 string, value2 string, ok *bool) {
	str_data_val1, status_data_val1 := data[value1].(string)
	if !status_data_val1 {
		*ok = false
		services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
	}
	str_data_val2, status_data_val2 := data[value2].(string)
	if !status_data_val2 {
		*ok = false
		services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
	}

	if str_data_val1 != str_data_val2 {
		str_errs_val, status_err_val := errs[value2].([]string)
		if !status_err_val {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(string) в string завершилось неудачно")
		}
		errs[value2] = append(str_errs_val, "пароли не совпадают")
		int_errs_counter, status_errs_counter := errs["counter"].(int)
		if !status_errs_counter {
			*ok = false
			services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
		}
		errs["counter"] = int_errs_counter + 1
		errs["isEmpty"] = false

	}
}

func ValidUser(data map[string]interface{}) (map[string]interface{}, bool) {
	ok := true
	errs := map[string]interface{}{
		"name":       []string{},
		"surname":    []string{},
		"birthday":   []string{},
		"email":      []string{},
		"phone":      []string{},
		"login":      []string{},
		"password1":  []string{},
		"password2":  []string{},
		"counter":    0,
		"isEmpty":    true,
		"isInserted": true,
	}

	MinMaxLength(data, errs, "name", 2, 20, &ok)
	IsLatin(data, errs, "name", &ok)
	MinMaxLength(data, errs, "surname", 2, 20, &ok)
	IsLatin(data, errs, "surname", &ok)
	MinMaxLength(data, errs, "birthday", 10, 10, &ok)
	MinMaxLength(data, errs, "email", 11, 50, &ok)
	IsEmail(data, errs, "email", &ok)
	MinMaxLength(data, errs, "phone", 11, 11, &ok)
	IsNumber(data, errs, "phone", &ok)
	MinMaxLength(data, errs, "login", 2, 20, &ok)
	IsLatin(data, errs, "login", &ok)
	MinMaxLength(data, errs, "password1", 8, 50, &ok)
	IsPassword(data, errs, "password1", &ok)
	MinMaxLength(data, errs, "password2", 8, 50, &ok)
	IsPassword(data, errs, "password2", &ok)
	IsEquality(data, errs, "password1", "password2", &ok)

	// fmt.Println(errs, ok)
	return errs, ok
}

func AddUser(data map[string]interface{}) map[string]interface{} {
	// var bool_err bool
	var err error
	// var data_valid map[string]interface{}
	errs_valid, bool_err := ValidUser(data)
	if !bool_err {
		services.L.Warn("При валидации данных произошла ошибка")
	}
	isEmpty, _ := errs_valid["isEmpty"].(bool)
	user_struct := models.User_json{}
	fmt.Println("isEmpty", isEmpty)
	if isEmpty {
		user_struct, err = ParserUser(data)
		if err != nil {
			services.L.Err(err)

		}
		var hash_pass []byte
		hash_pass, err = services.Hashing("passs123@")

		if err != nil {
			services.L.Err(err)

		}
		date, _ := time.Parse("2006-01-02", user_struct.Birthday)

		err_ins := crud_user.Insert_user(user_struct.Name, user_struct.Surname, date, user_struct.Email, user_struct.Phone, user_struct.Login, hash_pass)
		// fmt.Println("errs_valid", errs_valid)
		if err_ins != nil {
			errs_valid["isInserted"] = false
			services.L.Err(err_ins)
			// fmt.Println("err_ins.Error()", err_ins.Error())
			if strings.Contains(err_ins.Error(), " for key 'users.Login'") {
				str_errs_val, status_err_val := errs_valid["login"].([]string)
				if !status_err_val {
					services.L.Warn("преобразование значения map[string]interface{}['login'].(string) в string завершилось неудачно")
				}
				errs_valid["login"] = append(str_errs_val, "такой логин уже сущуствуе")
			}
			if strings.Contains(err_ins.Error(), " for key 'users.Email'") {
				str_errs_val, status_err_val := errs_valid["email"].([]string)
				if !status_err_val {
					services.L.Warn("преобразование значения map[string]interface{}['email'].(string) в string завершилось неудачно")
				}
				errs_valid["email"] = append(str_errs_val, "к этой почте уже привязан аккаунт")
			}
			int_errs_counter, status_errs_counter := errs_valid["counter"].(int)
			if !status_errs_counter {
				services.L.Warn("преобразование значения map[string]interface{}[value].(int) в int завершилось неудачно")
			}
			errs_valid["counter"] = int_errs_counter + 1
			errs_valid["isEmpty"] = false

		}

	}
	// services.L.Warn(data_valid)
	return errs_valid
}

func CheckUser(data map[string]interface{}) map[string]interface{} {
	// var bool_err bool
	var err error
	// var data_valid map[string]interface{}
	errs_valid, bool_err := ValidUser(data)
	if !bool_err {
		services.L.Warn("При валидации данных произошла ошибка")
	}
	isEmpty, _ := errs_valid["isEmpty"].(bool)
	user_struct := models.User_json{}
	fmt.Println("isEmpty", isEmpty)

	user_struct, err = ParserUser(data)
	if err != nil {
		services.L.Err(err)

	}

	crud_user.Existes_col_user("Login", user_struct.Login)
	// fmt.Println("errs_valid", errs_valid)

	// services.L.Warn(data_valid)
	return errs_valid
}
