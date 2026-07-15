import { getPosts } from "./api.js";
import { postTemplate } from "./template.js";

export async function initPosts() {

    const response = await getPosts();

    const container =
        document.getElementById("posts-container");

    container.innerHTML = response.posts
        .map(postTemplate)
        .join("");

}