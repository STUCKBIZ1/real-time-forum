import { request } from "../../shared/fetch.js";

export async function registerAPI(data){
    const response = await request(
        "/register",
        {
            method:"POST",

            headers:{
                "Content-Type":"application/json"
            },

            body:JSON.stringify(data)
        }
    );

    return response;
}
export async function loginAPI(data){
    const response = await request(
        "/login",
        {
            method:"POST",

            headers:{
                "Content-Type":"application/json"
            },

            body:JSON.stringify(data)
        }
    );
    return response;

}