import { getPost } from "./api.js";

export function postEvents() {

    const container = document.getElementById("posts-container");

    container.addEventListener("click", async (e) => {

        const post = e.target.closest(".post");

        if (!post) return;

        const id = post.dataset.id;

        const response = await getPost(id);

        console.log(response);

    });

}