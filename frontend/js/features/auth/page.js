import { registerTemplate, loginTemplate } from "./template.js";
import { authEvents } from "./events.js";


export function showRegister(app){

    app.innerHTML = registerTemplate();

    authEvents();

}
export function showLogin(app){

    app.innerHTML = loginTemplate();

    authEvents();

}