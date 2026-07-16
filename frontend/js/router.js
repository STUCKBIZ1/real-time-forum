import { showLogin } from "./features/auth/page.js";
export function router(){

    const app = document.getElementById("app");

    showLogin(app);

}