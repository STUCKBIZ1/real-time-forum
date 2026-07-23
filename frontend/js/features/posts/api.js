    import { request } from "../../shared/fetch";

    export async function getPost(id) {

        return await request(`/posts/${id}`);

    }
    export async function getPosts(){
        return await request(`/posts?cursor=${cursor}&limit=${limit}`)
    }