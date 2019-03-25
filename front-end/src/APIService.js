import axios from 'axios';
const API_URL = 'http://localhost:5000';
export class APIService{

constructor(){}
getCustomers() {
    const url = `${API_URL}/api/churn_case_data.csv`;
    return axios.get(url).then(response => response.data);
}
getCustomersSegment(segment) {
    const url = `${API_URL}/api/costumers/${segment}`;
    return axios.get(url).then(response => response.data);
}

}