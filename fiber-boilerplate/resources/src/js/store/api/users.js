import Axios from "axios";

export const GetUserList = () => Axios.get("/users").then(r => r.data)