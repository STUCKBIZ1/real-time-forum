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