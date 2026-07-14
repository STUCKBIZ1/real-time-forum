const API_URL = "http://localhost:8080";

export async function request(url, options = {}){

    const response = await fetch(
        API_URL + url,
        options
    );
    const data = await response.json();

    return data;

}