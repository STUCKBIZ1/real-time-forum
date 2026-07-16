export function getRegisterData(){
    return {

        firstname:
        document.getElementById("firstname").value.trim(),


        lastname:
        document.getElementById("lastname").value.trim(),


        nickname:
        document.getElementById("nickname").value.trim(),


        email:
        document.getElementById("email").value.trim(),


        password:
        document.getElementById("password").value,

        age:
        Number(document.getElementById("age").value),


        gender:
        document.getElementById("gender").value

    };

}
export function getLoginData(){
    return {
        email:
        document.getElementById("login-email").value,
        password:
        document.getElementById("login-password").value
    }
}
export function showError(id, message) {
    const error = document.getElementById(id);
    error.textContent = message;
    error.style.display = "block";
}

export function hideError(id) {
    document.getElementById(id).style.display = "none";
}
export function showSuccess(id, message) {
    const success = document.getElementById(id);
    success.textContent = message;
    success.style.display = "block";
}

export function hideSuccess(id) {
    document.getElementById(id).style.display = "none";
}