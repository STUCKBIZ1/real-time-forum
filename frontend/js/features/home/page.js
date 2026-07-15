import { homeTemplate } from "./template.js";

export async function showHome(app){
    console.log("hello show home", app)
    app.innerHTML = homeTemplate();

    // await initPosts();

    // initChat();

    // homeEvents();

}