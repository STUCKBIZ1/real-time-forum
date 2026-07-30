import { request } from "../../shared/fetch.js";
    export async function getPost(id) {

        return await request(`/posts/${id}`);

    }
    export async function getPosts(){
        return await request(`/getPosts?cursor=${cursor}&limit=${limit}`)
    }