import { homeTemplate } from "./template.js";
import { initPosts } from "../posts/page.js";
export async function showHome(app){
    console.log("hello show home", app)
    app.innerHTML = homeTemplate();

    await initPosts();

    initChat();

    homeEvents();

}