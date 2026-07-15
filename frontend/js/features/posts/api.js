export async function getPost(id) {

    return await request(`/posts/${id}`);

}