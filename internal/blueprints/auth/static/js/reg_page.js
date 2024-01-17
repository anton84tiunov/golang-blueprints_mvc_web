var name_inp = document.getElementById("name");
var surname_inp = document.getElementById("surname");
var email_inp = document.getElementById("email");
var birthday_inp = document.getElementById("birthday");
var phone_inp = document.getElementById("phone");
var password1_inp = document.getElementById("password1");
var password2_inp = document.getElementById("password2");
var login_inp = document.getElementById("login");
var submit_btn = document.getElementById("btn_submit_reg");

var name_err = document.getElementById("name_err");
var surname_err = document.getElementById("surname_err");
var email_err = document.getElementById("email_err");
var birthday_err = document.getElementById("birthday_err");
var phone_err = document.getElementById("phone_err");
var password1_err = document.getElementById("password1_err");
var password2_err = document.getElementById("password2_err");
var login_err = document.getElementById("login_err");

var msg_err = document.querySelectorAll(".msg_err");

function GetValue() {
    var name_val = name_inp.value;
    var surname_val = surname_inp.value;
    var email_val = email_inp.value;
    var birthday_val = birthday_inp.value;
    var phone_val = phone_inp.value;
    var password1_val = password1_inp.value;
    var password2_val = password2_inp.value;
    var login_val = login_inp.value;

    var data = {
        name: name_val,
        surname: surname_val,
        email: email_val,
        birthday: birthday_val,
        phone: phone_val,
        password1: password1_val,
        password2: password2_val,
        login: login_val,
    }
    return data;
}

function minMaxLength(data, errors, value, min, max) {
    if (data[value].length == 0) {
        errors[value].push(`Поле обязательно для заполнения`);
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    if (data[value].length < min) {
        errors[value].push(`Строка должна быть не менее ${min} символов`);
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    if (data[value].length > max) {
        errors[value].push(`Cтрока должна быть не больше ${max} символов`);
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    return errors
}

function isLatin(data, errors, value){
    var pattern = /^[a-zA-Z]*$/;
    if (!pattern.test(data[value])) {
        errors[value].push("Только латинсковые символы");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    return errors
}

function isNumber(data, errors, value){
    var pattern = /^[0-9]*$/;
    if (!pattern.test(data[value])) {
        errors[value].push("Только цифры");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    return errors
}

function isPassword(data, errors, value){
    // var pattern = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]*$/;
    var pattern = /[0-9]+/;
    if (!pattern.test(data[value])) {
        errors[value].push("Строка должна содержать хоть одну цифру");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    pattern = /[a-z]+/;
    if (!pattern.test(data[value])) {
        errors[value].push("Строка должна содержать хоть одну строчную букву");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    pattern = /[A-Z]+/;
    if (!pattern.test(data[value])) {
        errors[value].push("Строка должна содержать хоть одну заглавную букву");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    pattern = /(?=.*[@$!%*?&])/;
    if (!pattern.test(data[value])) {
        errors[value].push("Строка должна содержать хоть один спецсимвол @$!%*?&");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    return errors
}
function isEmail(data, errors, value){
    // /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    var pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    if (!pattern.test(data[value])) {
        errors[value].push("неправильная почта");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    return errors
}

function isEquality(data, errors, value1, value2) {
    if (["value1"] != ["value2"]) {
        errors[value2].push("пароли не совпадают");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    return errors
}





function isValid(data) {
    var errors = {
        name: [],
        surname: [],
        email: [],
        birthday: [],
        phone: [],
        password1: [],
        password2: [],
        login: [],
        counter: 0,
        isEmpty: true,
    }
    errors = minMaxLength(data, errors, "name", 2, 20);
    errors = isLatin(data, errors, "name");
    errors = minMaxLength(data, errors, "surname", 2, 20);
    errors = isLatin(data, errors, "surname");
    errors = minMaxLength(data, errors, "email", 11, 50);
    errors = isEmail(data, errors, "email");
    errors = minMaxLength(data, errors, "birthday", 10, 10);
    errors = minMaxLength(data, errors, "phone", 11, 11);
    errors = isNumber(data, errors, "phone");
    errors = minMaxLength(data, errors, "password1", 8, 50);
    errors = isPassword(data, errors, "password1");
    errors = minMaxLength(data, errors, "password2", 8, 50);
    errors = isPassword(data, errors, "password2");
    errors = minMaxLength(data, errors, "login", 2, 20);
    errors = isEquality(data, errors, "password1", "password2");
    errors = isLatin(data, errors, "login");
  
    // console.log(errors);
    if (!errors["isEmpty"]){ 
        for(const element of msg_err) { 
            element.innerHTML = "";
            // while (element.firstChild) {
            //     element.removeChild(element.firstChild);
            // }
        }
        
        for (const [key, value] of Object.entries(errors)) {
            console.log(`${key}`);
            var teg = document.getElementById(key + "_err");
            // var str_error_teg = "";
            if (Array.isArray(value)){
                value.forEach(function(elem) {
                    var p = document.createElement("p");
                    p.textContent = elem;
                    teg.appendChild(p)
                    // str_error_teg += elem + "\n";
                });
            }
            // teg.textContent = str_error_teg;
            // console.log(str_error_teg);
           
          }

        return false
    }else{
        for(const element of msg_err) { 
            element.innerHTML = "";
            // while (element.firstChild) {
            //     element.removeChild(element.firstChild);
            // }
        }

        return true
    }
}


function sendData(data) {
    fetch('/auth/reg', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data), // Convert the object to a JSON string
    })
        .then((response) => response.json())
        .then((data) => {
            console.log('data:', data);
        })
        .catch((error) => {
            console.log('Error:', error);
    });
}

function SubmitReg() {
    var data = GetValue();
    if (isValid(data)){
        sendData(data);
    }
    
}

submit_btn.addEventListener('click', SubmitReg);

