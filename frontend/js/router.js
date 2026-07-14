import { showRegister } from "./features/auth/page.js";
export function router(){

    const app = document.getElementById("app");

    showRegister(app);

}