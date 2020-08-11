import axios from 'axios';

export const getDiaRate = () => {
    return new Promise(async (resolve, reject) => {
        const URL  = `https://api.diadata.org/v1/quotation/DIA`;
        axios.get(URL)
        .then( (response) => {
            // handle success
            resolve(response.data);
        })
        .catch((error) => {
            // handle error
            reject(error);
        });
    })

}