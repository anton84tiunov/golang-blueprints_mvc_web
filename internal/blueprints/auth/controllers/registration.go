package controllers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	services "../../../services"

	crud_user "../../../database/crud/user"

	models "../../../models"
)

func ParserUser(data map[string]interface{}) models.User_json {
	// convert map to json
	jsonString, _ := json.Marshal(data)
	fmt.Println("string(jsonString)", string(jsonString))
	// convert json to struct
	s := models.User_json{}
	err_json := json.Unmarshal(jsonString, &s)
	if err_json != nil {
		fmt.Println(err_json)
	}
	return s
}

// функция определяет соответствие длинны строки заданным значениям
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func MinMaxLength(data map[string]interface{}, errs map[string]interface{}, value string, min uint8, max uint8) map[string]interface{} {
	bool_err := false
	str, ok1 := data[value].(string)
	if !ok1 {
		bool_err = true
		fmt.Println("str, ok1 := data[value].(string) !ok", str)
	}
	if str == "" {
		val_str, ok_str := errs[value].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value] = append(val_str, "Поле обязательно для заполнения")
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}

	if len(str) < int(min) {
		val_str, ok_str := errs[value].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value] = append(val_str, fmt.Sprintf("должна быть не менее %d символов", min))
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}
	if len(str) > int(max) {
		val_str, ok_str := errs[value].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value] = append(val_str, fmt.Sprintf("должна быть не больше %d символов", max))
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}
	fmt.Println(bool_err)
	return errs
}

// функция определяет соответствие строки на содержание только латинмких букв
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func IsLatin(data map[string]interface{}, errs map[string]interface{}, value string) map[string]interface{} {
	bool_err := false
	str, ok := data[value].(string)
	fmt.Println(ok)
	match, err := regexp.MatchString("^[a-zA-Z]*$", str)
	fmt.Println(err)
	if !match {
		val_str, ok_str := errs[value].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value] = append(val_str, "Только латинсковые символы")
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}
	fmt.Println(bool_err)
	return errs
}

// функция определяет соответствие строки на содержание только цифры
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func IsNumber(data map[string]interface{}, errs map[string]interface{}, value string) map[string]interface{} {
	bool_err := false
	str, ok := data[value].(string)
	fmt.Println(ok)
	match, err := regexp.MatchString("^[0-9]*$", str)
	fmt.Println(err)
	if !match {
		val_str, ok_str := errs[value].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value] = append(val_str, "Только цифры")
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}
	fmt.Println(bool_err)
	return errs
}

// функция определяет соответствие строки ЭЛЕКТРОННОЙ ПОЧТЕ
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func IsEmail(data map[string]interface{}, errs map[string]interface{}, value string) map[string]interface{} {
	bool_err := false
	str, ok := data[value].(string)
	fmt.Println(ok)
	match, err := regexp.MatchString(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, str)
	fmt.Println(err)
	if !match {
		val_str, ok_str := errs[value].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value] = append(val_str, "неправильная почта")
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}
	fmt.Println(bool_err)
	return errs
}

// функция определяет соответствие строки паролю
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func IsPassword(data map[string]interface{}, errs map[string]interface{}, value string) map[string]interface{} {
	bool_err := false
	str, ok := data[value].(string)
	fmt.Println(ok)
	match, err := regexp.MatchString("[0-9]+", str)
	fmt.Println(err)
	if !match {
		val_str, ok_str := errs[value].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value] = append(val_str, "Строка должна содержать хоть одну цифру")
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}
	match, err = regexp.MatchString("[a-z]+", str)
	fmt.Println(err)
	if !match {
		val_str, ok_str := errs[value].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value] = append(val_str, "Строка должна содержать хоть одну строчную букву")
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}
	match, err = regexp.MatchString("[A-Z]+", str)
	fmt.Println(err)
	if !match {
		val_str, ok_str := errs[value].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value] = append(val_str, "Строка должна содержать хоть одну заглавную букву")
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}

	match, err = regexp.MatchString(`[!@#$&*]+`, str)
	fmt.Println(err)
	if !match {
		val_str, ok_str := errs[value].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value] = append(val_str, "Строка должна содержать хоть один спецсимвол @$!%*?&")
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}
	fmt.Println(bool_err)
	return errs
}

// функция определяет равенство двух строк
// и возвращает map[string]interface{} в котором находится ошибки (в случае их обнаружения)
// и булевое значение в случае ошибки
func IsEquality(data map[string]interface{}, errs map[string]interface{}, value1 string, value2 string) map[string]interface{} {
	bool_err := false
	str1, ok1 := data[value1].(string)
	str2, ok2 := data[value1].(string)
	fmt.Println(ok1, " ", ok2)
	if str1 != str2 {
		val_str, ok_str := errs[value2].([]string)
		if !ok_str {
			bool_err = true
			fmt.Println("val_str, ok_str := errs[value].([]string)", ok_str)
		}
		errs[value2] = append(val_str, "пароли не совпадают")
		val_counter, ok_counter := errs["counter"].(int)
		if !ok_counter {
			bool_err = true
			fmt.Println("val_counter, ok_counter := errs['counter'].(int)", ok_counter)
		}
		errs["counter"] = val_counter + 1
		errs["isEmpty"] = false
	}
	fmt.Println(bool_err)
	return errs
}

func ValidUser(data map[string]interface{}) map[string]interface{} {
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
		"isEmpty":    false,
		"isInserted": false,
	}

	errs = MinMaxLength(data, errs, "name", 2, 20)
	errs = IsLatin(data, errs, "name")
	errs = MinMaxLength(data, errs, "surname", 2, 20)
	errs = IsLatin(data, errs, "surname")
	errs = MinMaxLength(data, errs, "birthday", 10, 10)
	errs = MinMaxLength(data, errs, "email", 11, 50)
	errs = IsEmail(data, errs, "email")
	fmt.Println("errs", errs)
	errs = MinMaxLength(data, errs, "phone", 11, 11)
	errs = IsNumber(data, errs, "phone")
	errs = MinMaxLength(data, errs, "login", 2, 20)
	errs = IsLatin(data, errs, "login")
	errs = MinMaxLength(data, errs, "password1", 8, 50)
	errs = IsPassword(data, errs, "password1")
	errs = MinMaxLength(data, errs, "password2", 8, 50)
	errs = IsPassword(data, errs, "password2")
	errs = IsEquality(data, errs, "password1", "password2")

	fmt.Println(errs)
	return errs
}

func AddUser(data map[string]interface{}) map[string]interface{} {

	data_valid := ValidUser(data)
	fmt.Println("data", data)
	str1, _ := data_valid["isEmpty"].(bool)
	user_struct := models.User_json{}
	if !str1 {
		user_struct = ParserUser(data)
		hash_pass, err := services.Hashing("passs123@")
		if err != nil {
			panic(err)
		}
		date, _ := time.Parse("2006-01-02", user_struct.Birthday)

		err_ins := crud_user.Insert_user(user_struct.Name, user_struct.Surname, date, user_struct.Email, user_struct.Phone, user_struct.Login, hash_pass)
		if err_ins == nil {
			data_valid["isInserted"] = true

		}
		fmt.Println("err_ins", err_ins)

	}
	return data_valid
}
