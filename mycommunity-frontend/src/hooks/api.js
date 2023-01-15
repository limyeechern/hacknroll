import axios from "axios";

const AxiosConstants = {
  BASEURL: "http://localhost:8080",
};

export default axios.create({
  baseURL: AxiosConstants.BASEURL,
});
