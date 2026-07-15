export function postTemplate(post) {
    return `
        <article class="post card" data-id="${post.id}">

            <h2>${post.title}</h2>

            <p>${post.content}</p>

        </article>
    `;
}