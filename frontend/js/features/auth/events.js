import { showLogin, showRegister } from "./page.js";
import { registerAPI, loginAPI } from "./api.js";
import { getLoginData, getRegisterData } from "../../shared/helpers.js";
import { showHome } from "../home/page.js";
    const app = document.getElementById("app");
function handleAuthNavigation(){
    const loginLink = document.getElementById("login-link");
    if(loginLink){
        loginLink.addEventListener("click", ()=>{
            showLogin(app);
        });
    }
    const registerLink = document.getElementById("register-link");

    if(registerLink){

        registerLink.addEventListener("click", ()=>{

            showRegister(app);

        });
    }
}
async function handleRegisterSubmit(e){
    e.preventDefault();
    const data = getRegisterData();
    try{
        const response = await registerAPI(data);
        console.log(response);
    }
    catch(error){
        console.error(error);
    }
}
async function handleLoginSubmit(e){
    e.preventDefault();
    const data = getLoginData();
    try{
        const response = await loginAPI(data);
        console.log(response);
    }
    catch(error){
        console.error(error);
    }
}
function handleForms(){
    const registerForm =
    document.getElementById("register-form");
    if(registerForm){

        registerForm.addEventListener(
            "submit",
            handleRegisterSubmit
        );
        
    }
    const loginForm =
    document.getElementById("login-form");
    if(loginForm){

        loginForm.addEventListener(
            "submit",
            handleLoginSubmit
        );
    }
}
export function authEvents(){

    handleAuthNavigation();

    handleForms();

}