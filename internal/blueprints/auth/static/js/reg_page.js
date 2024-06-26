var name_inp = document.getElementById("name");
var surname_inp = document.getElementById("surname");
var email_inp = document.getElementById("email");
var birthday_inp = document.getElementById("birthday");
var phone_inp = document.getElementById("phone");
var password1_inp = document.getElementById("password1");
var password2_inp = document.getElementById("password2");
var login_inp = document.getElementById("login");
var check_btn = document.getElementById("btn_check_reg");
var submit_btn = document.getElementById("btn_submit_reg");
var togglePassword1 = document.querySelector('#togglePassword1');
var togglePassword2 = document.querySelector('#togglePassword2');


var name_err = document.getElementById("name_err");
var surname_err = document.getElementById("surname_err");
var email_err = document.getElementById("email_err");
var birthday_err = document.getElementById("birthday_err");
var phone_err = document.getElementById("phone_err");
var password1_err = document.getElementById("password1_err");
var password2_err = document.getElementById("password2_err");
var login_err = document.getElementById("login_err");

var name_result = document.getElementById("name_result");

var msg_err = document.querySelectorAll(".msg_err");


submit_btn.setAttribute('disabled', '');

// функция получающая значения всех полей
// возвращает обьект в виде  json
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
    pattern = /[!@#$&*]+/;
    if (!pattern.test(data[value])) {
        errors[value].push("Строка должна содержать хоть один спецсимвол @$!%*?&");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    return errors
}
function isEmail(data, errors, value){
    var pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    if (!pattern.test(data[value])) {
        errors[value].push("неправильная почта");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    return errors
}

function isEquality(data, errors, value1, value2) {
    if (data[value1] != data[value2]) {
        errors[value2].push("пароли не совпадают");
        errors["counter"] += 1;
        errors["isEmpty"] = false;
    }
    return errors
}

function isErrors(errors) {
    if (!errors["isEmpty"]){ 
        for(const element of msg_err) { 
            element.innerHTML = "";
        }
        for (const [key, value] of Object.entries(errors)) {
            var teg = document.getElementById(key + "_err");
            if (Array.isArray(value)){
                value.forEach(function(elem) {
                    var p = document.createElement("p");
                    p.textContent = elem;
                    teg.appendChild(p)
                });
            }
          }
        return false
    }else{
        for(const element of msg_err) { 
            element.innerHTML = "";
        }
        return true
    }
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
        isInserted: false,
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
    errors = minMaxLength(data, errors, "login", 2, 20);
    errors = isLatin(data, errors, "login");
    errors = minMaxLength(data, errors, "password1", 8, 50);
    errors = isPassword(data, errors, "password1");
    errors = minMaxLength(data, errors, "password2", 8, 50);
    errors = isPassword(data, errors, "password2");
    errors = isEquality(data, errors, "password1", "password2");
  
  
    return isErrors(errors)
  
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
            isErrors(data)
            console.log('data:', data);
        })
        .catch((error) => {
            console.log('Error:', error);
    });
}

function checkData(data) {
    fetch('/auth/reg_check', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data), // Convert the object to a JSON string
    })
        .then((response) => response.json())
        .then((data) => {
            isErrors(data)
            console.log('data:', data);
        })
        .catch((error) => {
            console.log('Error:', error);
    });
}

function CheckReg() {
    var data = GetValue();
    if (isValid(data)){
        checkData(data);
    }
    check_btn.setAttribute('disabled', '');
    submit_btn.removeAttribute('disabled', '');
}
check_btn.addEventListener('click', CheckReg);

function SubmitReg() {
    var data = GetValue();
    if (isValid(data)){
        sendData(data);
    }
    submit_btn.setAttribute('disabled', '');
    check_btn.removeAttribute('disabled', '');

}


submit_btn.addEventListener('click', SubmitReg);

togglePassword1.addEventListener('click', () => {
    const type = password1_inp
        .getAttribute('type') === 'password' ?
        'text' : 'password';
        password1_inp.setAttribute('type', type);
});

togglePassword2.addEventListener('click', () => {
    const type = password2_inp
        .getAttribute('type') === 'password' ?
        'text' : 'password';
        password2_inp.setAttribute('type', type);
});


name_inp.addEventListener('input', () => {
    var data = {
        name: name_inp.value,
    }
    var errors = {
        name: [],
        counter: 0,
        isEmpty: true,
    }
    errors = minMaxLength(data, errors, "name", 2, 20);
    errors = isLatin(data, errors, "name");
    // var name_err = document.getElementById("name_err");

    if(!errors["isEmpty"]){
        name_err.innerHTML = "";
        errors["name"].forEach(function(value) {
            var p = document.createElement("p");
            p.textContent = value;
            name_err.appendChild(p)
        });
    }else{
        name_err.innerHTML = "";
    }
});

surname_inp.addEventListener('input', () => {
    var data = {
        surname: surname_inp.value,
    }
    var errors = {
        surname: [],
        counter: 0,
        isEmpty: true,
    }
    errors = minMaxLength(data, errors, "surname", 2, 20);
    errors = isLatin(data, errors, "surname");
    // var surname_err = document.getElementById("surname_err");

    if(!errors["isEmpty"]){
        surname_err.innerHTML = "";
        errors["surname"].forEach(function(value) {
            var p = document.createElement("p");
            p.textContent = value;
            surname_err.appendChild(p)
        });
    }else{
        surname_err.innerHTML = "";
    }
});

login_inp.addEventListener('input', () => {
    var data = {
        login: login_inp.value,
    }
    var errors = {
        login: [],
        counter: 0,
        isEmpty: true,
    }
    errors = minMaxLength(data, errors, "login", 2, 20);
    errors = isLatin(data, errors, "login");
    // var login_err = document.getElementById("login_err");

    if(!errors["isEmpty"]){
        login_err.innerHTML = "";
        errors["login"].forEach(function(value) {
            var p = document.createElement("p");
            p.textContent = value;
            login_err.appendChild(p)
        });
    }else{
        login_err.innerHTML = "";
    }
});


birthday_inp.addEventListener('input', () => {
    var data = {
        birthday: birthday_inp.value,
    }
    var errors = {
        lobirthdaygin: [],
        counter: 0,
        isEmpty: true,
    }
    errors = minMaxLength(data, errors, "birthday", 10, 10);

    if(!errors["isEmpty"]){
        birthday_err.innerHTML = "";
        errors["birthday"].forEach(function(value) {
            var p = document.createElement("p");
            p.textContent = value;
            birthday_err.appendChild(p)
        });
    }else{
        birthday_err.innerHTML = "";
    }
});
       
email_inp.addEventListener('input', () => {
    var data = {
        email: email_inp.value,
    }
    var errors = {
        email: [],
        counter: 0,
        isEmpty: true,
    }
    errors = minMaxLength(data, errors, "email", 11, 50);
    errors = isEmail(data, errors, "email");

    if(!errors["isEmpty"]){
        email_err.innerHTML = "";
        errors["email"].forEach(function(value) {
            var p = document.createElement("p");
            p.textContent = value;
            email_err.appendChild(p)
        });
    }else{
        email_err.innerHTML = "";
    }
});

phone_inp.addEventListener('input', () => {
    var data = {
        phone: phone_inp.value,
    }
    var errors = {
        phone: [],
        counter: 0,
        isEmpty: true,
    }
    errors = minMaxLength(data, errors, "phone", 11, 11);
    errors = isNumber(data, errors, "phone");

    if(!errors["isEmpty"]){
        phone_err.innerHTML = "";
        errors["phone"].forEach(function(value) {
            var p = document.createElement("p");
            p.textContent = value;
            phone_err.appendChild(p)
        });
    }else{
        phone_err.innerHTML = "";
    }
}); 

password1_inp.addEventListener('input', () => {
    var data = {
        password1: password1_inp.value,
    }
    var errors = {
        password1: [],
        counter: 0,
        isEmpty: true,
    }
    errors = minMaxLength(data, errors, "password1", 8, 50);
    errors = isPassword(data, errors, "password1");

    if(!errors["isEmpty"]){
        password1_err.innerHTML = "";
        errors["password1"].forEach(function(value) {
            var p = document.createElement("p");
            p.textContent = value;
            password1_err.appendChild(p)
        });
    }else{
        password1_err.innerHTML = "";
    }
});

password2_inp.addEventListener('input', () => {
    var data = {
        password1: password1_inp.value,
        password2: password2_inp.value,
    }
    var errors = {
        password1: [],
        password2: [],
        counter: 0,
        isEmpty: true,
    }
    errors = minMaxLength(data, errors, "password2", 8, 50);
    errors = isPassword(data, errors, "password2");
    errors = isEquality(data, errors, "password1", "password2");

    if(!errors["isEmpty"]){
        password2_err.innerHTML = "";
        errors["password2"].forEach(function(value) {
            var p = document.createElement("p");
            p.textContent = value;
            password2_err.appendChild(p)
        });
    }else{
        password2_err.innerHTML = "";
    }
}); 